package main

import (
	"fmt"
	"github.com/gerencianet/gn-api-sdk-go/gerencianet/pix"
	"github.com/gerencianet/gn-api-sdk-go/examples/configs"
)

func main(){
	
	credentials := configs.Credentials
	gn := pix.NewGerencianet(credentials)
	
	const e2eid = "E00416968202105211756Rh0iPsaJ1RK"

	res, err := gn.PixSendDetail(e2eid) 

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}