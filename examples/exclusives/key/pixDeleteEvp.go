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
	const key = "38903a81-1c30-456f-a4f6-6b682165d08d"

	res, err := gn.PixDeleteEvp(key, body) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}