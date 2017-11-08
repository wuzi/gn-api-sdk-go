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
		"expire_at": "2020-12-12",
	}

	res, err := gn.UpdateBillet(1, body) // no lugar do 1 coloque o charge_id certo

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}