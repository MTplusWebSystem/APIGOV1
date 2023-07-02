package main

import (
	"fmt"
	"github.com/MTplusWebSystem/APIGOV1/router"
	"github.com/MTplusWebSystem/APIGOV1/config"
)

var (
	logger *config.Logger
)

func main() {
	fmt.Println("\033[1;33mAPI INICIALIZADA")

	logger = config.ResponseLogger("main")
	err := config.StartDB()
	if err == nil {
		logger.Error("ERRO NA INICIALIZAÇÃO",err)
		return
	}
	router.Start()
}
