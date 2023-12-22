package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"slurm-clean-arch/pkg/store/postgres"
	deliveryGrpc "slurm-clean-arch/services/contact/internal/delivery/grpc"
	deliveryHttp "slurm-clean-arch/services/contact/internal/delivery/http"
	repositoryContact "slurm-clean-arch/services/contact/internal/repository/contact/postgres"
	repositoryGroup "slurm-clean-arch/services/contact/internal/repository/group/postgres"
	useCaseContact "slurm-clean-arch/services/contact/internal/usecase/contact"
	useCaseGroup "slurm-clean-arch/services/contact/internal/usecase/group"
	"syscall"
)

func main() {
	conn, err := postgres.New(postgres.Settings{})
	if err != nil {
		panic(err)
	}

	defer conn.Pool.Close()

	repoContact, err := repositoryContact.New(conn.Pool, repositoryContact.Options{})
	if err != nil {
		panic(err)
	}
	repoGroup, err := repositoryGroup.New(conn.Pool, repoContact, repositoryGroup.Options{})
	if err != nil {
		panic(err)
	}

	//repoStorage, err := repositoryStorage.New(conn.Pool, repositoryStorage.Options{})
	//if err != nil {
	//	panic(err)
	//}

	var (
		ucContact = useCaseContact.New(repoContact, useCaseContact.Options{})
		ucGroup   = useCaseGroup.New(repoGroup, useCaseGroup.Options{})
		//ucGroup        = useCaseGroup.New(repoStorage, useCaseGroup.Options{})
		_            = deliveryGrpc.New(ucContact, ucGroup, deliveryGrpc.Options{})
		listenerHttp = deliveryHttp.New(ucContact, ucGroup, deliveryHttp.Options{})
	)

	go func() {
		fmt.Printf("service started successfully on htp port: %d", viper.GetUint("HTTP_PORT"))
		if err = listenerHttp.Run(); err != nil {
			panic(err)
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh
}
