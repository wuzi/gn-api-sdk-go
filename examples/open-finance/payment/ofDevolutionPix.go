package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet/open_finance"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := open_finance.NewGerencianet(credentials)

	body := map[string]interface{} {
		"identificadorPagamento" : "urn:gerencianet:ea807997-9c28-47a7-8ebc-eeb672ea59f0",
		"valorDevolucao" : "0.01",
	}

	res, err := gn.OfDevolutionPix(body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}