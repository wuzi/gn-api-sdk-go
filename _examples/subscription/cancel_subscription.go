package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go/_examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	res, err := gn.CancelSubscription(1) // no lugar do 1 coloque o subscription_id correto

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}