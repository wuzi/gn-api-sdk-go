package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet/pix"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials	
	gn := pix.NewGerencianet(credentials)

	params := map[string]string {
		"inicio" : "2022-01-01T00:00:00Z",
		"fim" : "2024-12-31T23:59:59Z",
		// "status" : "CONCLUIDA", // "ATIVA","CONCLUIDA", "REMOVIDA_PELO_USUARIO_RECEBEDOR", "REMOVIDA_PELO_PSP"
		// "cpf" : "12345678909", // Filter by payer's CPF. It cannot be used at the same time as the CNPJ.
		// "cnpj" : "12345678909", // Filter by payer's CNPJ. It cannot be used at the same time as the CPF.
		// "locationPresente" : true,
		// "loteCobVId" : "123456789",
		// "paginacao.paginaAtual" : 1,
		// "paginacao.itensPorPagina" : 10,
	}

	res, err := gn.PixListDueCharges(params) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}