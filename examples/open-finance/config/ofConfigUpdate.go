package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet/open_finance"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := open_finance.NewGerencianet(credentials)

	body := map[string]interface{} {
		"redirectURL" : "https://your-domain.com.br/redirect/",
		"webhookURL" : "https://your-domain.com.br/webhook/",
	}

	res, err := gn.OfConfigUpdate(body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}