package main

import (
	"net/rpc"
	"testing"

	"github.com/konghui/gospider/proto"
)

func Connect() (client *rpc.Client, err error) {

	return
}

func SendTestRequest(t *testing.T, request *proto.Request) {
	t.Log(request)
	client, err := Connect()
	if err != nil {
		t.Fatal(err.Error())
	}
	response := new(proto.Response)
	err = client.Call("Server.Request", request, response)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
}

// send a keepalive request to the server
func Test_SendKeepAliveRequest(t *testing.T) {
	request := proto.NewKeepAliveRequest()
	SendTestRequest(t, request)
}

// send a getlist request to the server
func Test_SendGetListRequest(t *testing.T) {
	request := proto.NewGetListRequest()
	SendTestRequest(t, request)
}

// send a reportdata request to the server
func Test_SendReportDataRequest(t *testing.T) {
	request := proto.NewReportDataRequest()
	SendTestRequest(t, request)
}

// send a parserpage request to the server
func Test_SendParsePageRequest(t *testing.T) {
	Request := prot
}
