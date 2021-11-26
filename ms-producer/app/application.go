package app

import (
	"github.com/DevNiklesh/kafka-producer-consumer/ms-producer/config"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
	conf   = config.New()
)

func StartApplication() {
	router.Run(conf.PORT)
}
