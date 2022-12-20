package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet/pix"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := pix.NewGerencianet(credentials)

	const id = "423"

	res, err := gn.PixGenerateQRCode(id) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}