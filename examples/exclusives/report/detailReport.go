package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet/pix"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := pix.NewGerencianet(credentials)

	const id = "cacf4fbf-09bc-44f5-bc32-143ffb2a37ff"

	res, err := gn.DetailReport(id) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}