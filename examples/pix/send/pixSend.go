package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet/pix"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := pix.NewGerencianet(credentials)

	const idEnvio = "0S000000000000000000000000000000000"

	body := map[string]interface{} {
		"valor": "0.01",
		"pagador": map[string]interface{} {
			"chave": "",
			"infoPagador": "Segue o pagamento da conta",
		},
		"favorecido": map[string]interface{} {			
			"chave": "",
		},
	}

	res, err := gn.PixSend(idEnvio, body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}