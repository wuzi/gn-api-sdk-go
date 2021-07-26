package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	body := map[string]interface{} {}	
	const chave = "2c5c7441-a91e-4982-8c25-61e18ae"
	
	

	res, err := gn.DeleteWebhook(chave, body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}