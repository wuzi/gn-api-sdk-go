package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet/pix"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := pix.NewGerencianet(credentials)

	body := map[string]interface{} {
		"dataMovimento": "2022-04-24",
    	"tipoRegistros": map[string]interface{} {
			"pixRecebido": true,
			"pixEnviadoChave": true,
			"pixEnviadoDadosBancarios": true,
			"estornoPixEnviado": true,
			"pixDevolucaoEnviada": true,
			"pixDevolucaoRecebida": true,
			"tarifaPixEnviado": true,
			"tarifaPixRecebido": true,
			"estornoTarifaPixEnviado": true,
			"saldoDiaAnterior": true,
			"saldoDia": true,
			"transferenciaEnviada": true,
			"transferenciaRecebida": true,
			"estornoTransferenciaEnviada": true,
			"tarifaTransferenciaEnviada": true,
			"estornoTarifaTransferenciaEnviada": true,
    	},
	}

	
	res, err := gn.CreateReport(body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}