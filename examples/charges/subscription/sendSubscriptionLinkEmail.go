package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	body := map[string]interface{} {
		"email": "yasmin.goncalves@gerencianet.com.br",
	}

	res, err := gn.SendSubscriptionLinkEmail(1, body) // no lugar do 1 coloque o charge_id correto

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}