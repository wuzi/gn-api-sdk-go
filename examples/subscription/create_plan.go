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
		"name": "My plan",
		"interval": 2,
		"repeats": nil,
	}

	res, err := gn.CreatePlan(body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}