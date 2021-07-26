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
		"billet_discount": 1,
		"card_discount": 1,
		"message": "teste",
		"expire_at": "2018-12-12",
		"request_delivery_address": false,
		"payment_method": "all",
	}

	res, err := gn.UpdateChargeLink(1, body) // no lugar do 1 coloque o charge_id certo

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}