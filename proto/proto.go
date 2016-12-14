package proto

import (
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"time"

	log "github.com/Sirupsen/logrus"
)

const (
	GetUrlList = iota
	ParsePage
	ReportData
	Ping
	Pong
)

var Command = [...]string{
	"GetUrlList",
	"ParsePage",
	"ReportData",
	"Ping",
	"Pong",
}

const Version = 1
const PROTO = "tcp"

type Request struct {
	Version   int
	Id        int
	Cmd       byte
	TimeStamp int64
}

func (this *Request) String() (s string) {
	s = fmt.Sprintf("REQUEST:\n\tVersion:%d, Cmd:%s, Id:%d, TimeStamp:%d\n", this.Version, Command[this.Cmd], this.Id, this.TimeStamp)
	return
}

func NewRequest(cmd byte) (req *Request) {
	req = &Request{Version: Version, Cmd: cmd, TimeStamp: time.Now().Unix(), Id: 1}
	return
}

func (req *Request) SendRequest(host string) (response *Response, err error) {
	log.Debugf("Send Request:%s", req)
	var client *rpc.Client

	// need refactoring. use connect pool instead
	client, err = rpc.DialHTTP(PROTO, host)

	if err != nil {
		return
	}

	response = new(Response)
	err = client.Call("Server.Request", req, response)
	if err != nil {
		return
	}
	log.Debugf("Get Response: %s", response)
	return
}

func NewKeepAliveRequest() (req *Request) {
	req = NewRequest(Ping)
	return
}

func SendKeepAliveRequest(host string) (response *Response, err error) {
	response, err = NewKeepAliveRequest().SendRequest(host)
	return
}

func NewGetListRequest() (req *Request) {
	req = NewRequest(GetUrlList)
	return
}

func NewReportDataRequest() (req *Request) {
	req = NewRequest(ReportData)
	return
}

func NewParsePageRequest() (req *Request) {
	req = NewRequest(ParsePage)
	return
}

type Response struct {
	Version   int
	Cmd       byte
	Data      interface{}
	Length    uint32
	Error     string
	TimeStamp int64
}

func (this *Response) String() (s string) {
	s = fmt.Sprintf("RESPONSE:\n\tVersion:%d, Type:%s, Data:%s, Length:%d, Error:%s, Time:%d\n", this.Version, Command[this.Cmd], this.Data, this.Length, this.Error, this.TimeStamp)
	return
}

func NewResponse(cmd byte, data interface{}, leng uint32, error string) (rep *Response) {
	rep = new(Response)
	rep.BuildResponse(cmd, data, leng, error)
	//rep = &Response{Cmd: cmd, Data: data, Length: leng, Version: Version, Error: Error, TimeStamp: time.Now().Unix()}
	return
}

func (this *Response) BuildResponse(cmd byte, data interface{}, leng uint32, Error string) {
	this.Cmd = cmd
	this.Data = data
	this.Length = leng
	this.Version = Version
	this.Error = Error
	this.TimeStamp = time.Now().Unix()
}

func (this *Response) IsPongResponse() (rv bool) {
	if this.Cmd == Pong {
		return true
	}
	return
}

func NewKeepAliveResponse() (rep *Response) {
	rep = NewResponse(Pong, nil, 0, "")
	return
}

func (this *Response) BuildKeepAliveResponse() {
	this.BuildResponse(Pong, nil, 0, "")
}

func NewGetListResponse(list []string) (rep *Response) {
	rep = NewResponse(GetUrlList, list, 0, "")
	return
}

func (this *Response) BuildGetListResponse(list []string) {
	this.BuildResponse(GetUrlList, list, 0, "")
}

func NewParsePageResponse(url string) (rep *Response) {
	rep = NewResponse(ParsePage, url, 0, "")
	return
}

func (this *Response) BuildParsePageResponse(url string) {
	this.BuildResponse(ParsePage, url, 0, "")
}

// rpc function
type Server struct {
	handler map[byte]func(*Request, *Response)
}

func NewServer() (server *Server, err error) {
	var l net.Listener

	server = new(Server)
	server.handler = make(map[byte]func(*Request, *Response))
	rpc.Register(server)
	rpc.HandleHTTP()
	l, err = net.Listen("tcp", ":1234")
	if err != nil {
		return
	}
	go http.Serve(l, nil)

	return
}

func (t *Server) Request(args *Request, rv *Response) (err error) {
	callback, yes := t.handler[args.Cmd]
	if !yes {
		rv = new(Response)
		rv.Error = fmt.Sprintf("unknown cmd code %d", args.Cmd)
		return
	}
	log.Debugf("new request arrive %s", args)
	callback(args, rv)
	return
}

func (t *Server) RegisterCallbackFunc(cmd byte, callback func(*Request, *Response)) {
	t.handler[cmd] = callback
}
