package eth

import (
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

type (
	RPC struct {
		url     string
		network string
		client  *resty.Client
	}

	RPCBase struct {
		ID      int    `json:"id"`
		JsonRPC string `json:"jsonrpc"`
	}

	RPCRequest struct {
		RPCBase
		Method string   `json:"method"`
		Params []string `json:"params"`
	}

	RPCSuccessResponse struct {
		RPCBase
		Result string
	}

	RPCErrorResponse struct {
		RPCBase
		Error struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}
)

func NewRPC(url string, network string) RPC {
	return RPC{
		url:     url,
		network: network,
		client:  resty.New(),
	}
}

func newRPCRequest(method string, params []string) RPCRequest {
	r := RPCRequest{
		Method: method,
		Params: params,
	}
	r.ID = 1
	r.JsonRPC = "2.0"

	return r
}

func (r *RPC) GetCode(address string) string {
	body := newRPCRequest("eth_getCode", []string{
		address, "latest",
	})

	resp, err := r.client.R().
		SetBody(body).
		SetResult(RPCSuccessResponse{}).
		SetError(RPCErrorResponse{}).
		Post(r.url)

	if err != nil {
		log.WithField("Error", err).
			WithField("URL", r.url).
			Fatal("Failed to get contract byte code")
	}

	return resp.Result().(*RPCSuccessResponse).Result
}
