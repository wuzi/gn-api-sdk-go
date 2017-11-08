package gerencianet

type gerencianet struct {
	endpoints
}

func NewGerencianet(configs map[string]interface{}) *gerencianet {
	clientID := configs["client_id"].(string)
	clientSecret := configs["client_secret"].(string)
	sandbox := configs["sandbox"].(bool)
	timeout := configs["timeout"].(int)

	requester := newRequester(clientID, clientSecret, sandbox, timeout)
	gn := gerencianet{}
	gn.requester = *requester
	return &gn
}