package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet/payments"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials	
	gn := payments.NewGerencianet(credentials)

	params := map[string]string {
		"codBarras" : "36400000000000000000000000000000000000000000000",
	}

	res, err := gn.PayDetailBarCode(params) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}