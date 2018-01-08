package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go/_examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	body := map[string]interface{} {
		"name": "My new plan",
	}

	res, err := gn.UpdatePlan(1, body) // no lugar do 1 coloque o plan_id certo

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}