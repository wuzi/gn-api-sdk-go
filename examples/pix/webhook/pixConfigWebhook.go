package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet/pix"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := pix.NewGerencianet(credentials)

	
	const chave = ""

	body := map[string]interface{} {
		
		"webhookUrl": "https://seu_webhook",
	}

	res, err := gn.PixConfigWebhook(chave,body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}