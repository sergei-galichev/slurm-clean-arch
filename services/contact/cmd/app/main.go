package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"slurm-clean-arch/pkg/store/postgres"
	deliveryGrpc "slurm-clean-arch/services/contact/internal/delivery/grpc"
	deliveryHttp "slurm-clean-arch/services/contact/internal/delivery/http"
	repositoryStorage "slurm-clean-arch/services/contact/internal/repository/storage/postgres"
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

	var (
		repoStorage, _ = repositoryStorage.New(conn.Pool, repositoryStorage.Options{})
		ucContact      = useCaseContact.New(repoStorage, useCaseContact.Options{})
		ucGroup        = useCaseGroup.New(repoStorage, useCaseGroup.Options{})
		_              = deliveryGrpc.New(ucContact, ucGroup, deliveryGrpc.Options{})
		listenerHttp   = deliveryHttp.New(ucContact, ucGroup, deliveryHttp.Options{})
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
