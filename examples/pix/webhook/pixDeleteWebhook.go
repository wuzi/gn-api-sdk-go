package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet/pix"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := pix.NewGerencianet(credentials)

	body := map[string]interface{} {}	
	const chave = ""
	
	

	res, err := gn.PixDeleteWebhook(chave, body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}