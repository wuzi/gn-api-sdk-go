package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go/_examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	paymentToken := "payment_token";

	customer := map[string]interface{}{
		"name": "Gorbadoc Oldbuck",
		"cpf": "04267484171",
		"phone_number": "5144916523",
		"email": "oldbuck@gerencianet.com.br",
		"birth": "1977-01-15",
	}

	billingAddress := map[string]interface{} {
		"street": "Av JK",
		"number": 909,
		"neighborhood": "Bauxita",
		"zipcode": "35400000",
		"city": "Ouro Preto",
		"state": "MG",
	}

	body := map[string]interface{} {
		"payment": map[string]interface{} {
			"credit_card": map[string]interface{} {
				"installments": 1,
				"billing_address": billingAddress,
				"payment_token": paymentToken,
				"customer": customer,
			},
		},
	}

	res, err := gn.PayCharge(1, body) // no lugar do 1 coloque o charge_id certo

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}