package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet/pix"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := pix.NewGerencianet(credentials)

	
	const txid = "adssshdsjdsjeyccdyddsasdstxid01"

	body := map[string]interface{} {
		
		"calendario": map[string]interface{} {
			"dataDeVencimento": "2022-12-30",
			"validadeAposVencimento": 30,
		},
		"devedor": map[string]interface{}{
			"logradouro": "Alameda Souza, Numero 80, Bairro Braz",
    		"cidade": "Recife",
    		"uf": "PE",
    		"cep": "70011750",
    		"cpf": "12345678909",
    		"nome": "Francisco da Silva",
		},
		"valor": map[string]interface{} {
			"original": "123.45",
    		"multa": map[string]interface{} {
      			"modalidade": 2,
      			"valorPerc": "2.00",
    		},
    		"juros": map[string]interface{} {
      			"modalidade": 2,
      			"valorPerc": "2.00",
    		},
		},
		"chave": "",
		"solicitacaoPagador": "Teste.",
	}

	res, err := gn.CreateDueCharge(txid,body)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}