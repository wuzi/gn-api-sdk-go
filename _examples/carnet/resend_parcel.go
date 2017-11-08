package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go/_examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	body := map[string]interface{} {
		"email": "oldbuck@gerencianet.com.br",
	}

	res, err := gn.ResendParcel(1, 1, body) // no lugar dos 1s coloque o carnet_id e o numero da parcela respectivamente

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}