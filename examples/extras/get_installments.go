package main

import (
	"fmt"
	"github.com/FilipeMata/gn-api-sdk-go/gerencianet"
	"github.com/FilipeMata/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	res, err := gn.GetInstallments(20000, "visa") // total e bandeira

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}