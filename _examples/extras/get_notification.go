package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go/_examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	res, err := gn.GetNotification("token") // subsituir pelo token recebido

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}