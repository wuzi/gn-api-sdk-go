package main

import (
	"fmt"
	"github.com/FilipeMata/gn-api-sdk-go/gerencianet"
	"github.com/FilipeMata/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	body := map[string]interface{} {
		"description": "This charge was not fully paid",
	}

	res, err := gn.CreateCarnetHistory(1, body) // no lugar do 1 coloque o carnet_id certo

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}