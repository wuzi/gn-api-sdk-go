package main

import (
	"fmt"
	"github.com/FilipeMata/gn-api-sdk-go/gerencianet"
	"github.com/FilipeMata/gn-api-sdk-go/examples/configs"
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

	res, err := gn.PayCharge(1, body) // no lugar do 1 coloque o charge_id certo

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}