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
		"items": []map[string]interface{}{
			{
				"name" : "Product One",
				"value": 600,
				"amount": 1,
			},
		},
		"settings": map[string]interface{}{
				"payment_method": "all",
				"expire_at": "2022-12-15",
				"request_delivery_address": false,
		},
	}

	res, err := gn.CreateOneStepLink(body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}