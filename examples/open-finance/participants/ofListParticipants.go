package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet/open_finance"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	credentials := configs.Credentials
	gn := open_finance.NewGerencianet(credentials)

	const nome = ""

	res, err := gn.OfListParticipants(nome)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}