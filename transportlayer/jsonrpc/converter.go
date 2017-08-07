package jsonrpc

import jsonrpctransport "github.com/devimteam/go-kit/transport/jsonrpc"

type EndpointServerConverter struct {
	EncodeResp jsonrpctransport.EncodeResponseFunc
	DecodeReq  jsonrpctransport.DecodeRequestFunc
}

type EndpointClientConverter struct {
	EncodeReq  jsonrpctransport.EncodeRequestFunc
	DecodeResp jsonrpctransport.DecodeResponseFunc
	ReplyType  interface{}
}
