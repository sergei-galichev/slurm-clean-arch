package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"slurm-clean-arch/services/contact/internal/usecase"
)

func init() {
	viper.SetDefault("HTTP_PORT", 8080)
	viper.SetDefault("HTTP_HOST", "127.0.0.1")
	viper.SetDefault("IS_PRODUCTION", "false")
}

type Delivery struct {
	ucContact usecase.Contact
	ucGroup   usecase.Group
	router    *gin.Engine

	options Options
}

type Options struct{}

func New(ucContact usecase.Contact, ucGroup usecase.Group, options Options) *Delivery {
	var d = &Delivery{
		ucContact: ucContact,
		ucGroup:   ucGroup,
	}

	d.SetOptions(options)

	d.router = d.initRouter()

	return d
}

func (d *Delivery) SetOptions(options Options) {
	if d.options != options {
		d.options = options
	}
}

func (d *Delivery) Run() error {
	return d.router.Run(
		fmt.Sprintf(
			"%s:%d",
			viper.GetString("HTTP_HOST"),
			uint16(viper.GetUint("HTTP_PORT")),
		),
	)
}

func checkAuth(c *gin.Context) {
	c.Next()
}
