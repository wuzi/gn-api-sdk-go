package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go/_examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	customer := map[string]interface{}{
		"name": "Gorbadoc Oldbuck",
		"cpf": "04267484171",
		"phone_number": "5144916523",
	}

	body := map[string]interface{} {
		"payment": map[string]interface{} {
			"banking_billet": map[string]interface{} {
				"expire_at": "2018-12-12",
				"customer": customer,
			},
		},
	}

	res, err := gn.PaySubscription(1, body) // no lugar do 1 coloque o subscription_id certo

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}