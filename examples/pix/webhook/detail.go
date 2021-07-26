package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet/pix"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := gerencianet.NewGerencianet(credentials)

	
	const chave = "2c5c7441-a91e-4982-8c25-e18ae"

	

	res, err := gn.GetWebhook(chave) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}