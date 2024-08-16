package authclient

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log"
	"net/http"
)

const clientApiKey string = "API-Key"

var client = &AuthClient{}

type AuthClient struct {
	*fasthttp.HostClient
	apiKey string
}

func Init(host, apiKey string) {
	client = &AuthClient{
		HostClient: &fasthttp.HostClient{
			Addr: host,
		},
		apiKey: apiKey,
	}
}

type Response struct {
	UserId string `json:"user_id"`
	Role   string `json:"role"`
}

func ValidateToken(token string) (Response, bool) {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI("http://" + client.Addr + "/service/get_user_info")
	req.Header.Set(fasthttp.HeaderAuthorization, token)
	req.Header.Set(clientApiKey, client.apiKey)
	req.Header.SetHost(client.Addr)
	req.Header.SetMethod(fasthttp.MethodGet)

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	err := client.Do(req, resp)
	if err != nil {
		return Response{}, false
	}
	log.Println(resp)
	if resp.StatusCode() != http.StatusOK {
		return Response{}, false
	}

	var respData Response
	err = json.Unmarshal(resp.Body(), &respData)
	if err != nil {
		return Response{}, false
	}

	return respData, true
}
