package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet/pix"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials	
	gn := pix.NewGerencianet(credentials)

	const txid = "adssshdsjdsjeyccdyddsasdstxid01"

	body := map[string]interface{} {
		"valor" : map[string]interface{}{
			"original" : "1.00",
		},
		"solicitacaoPagador" : "Enter the order number or identifier.",
	}

	res, err := gn.PixUpdateDueCharge(txid, body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}