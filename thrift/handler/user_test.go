package handler

import (
	"testing"
	"thrift/proto/req"
	"thrift/proto/resp"
	"thrift/proto/userservice"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/stretchr/testify/assert"
)

const SERVER_ADDRESS = "127.0.0.1:56783"

func init() {
	transportFactory := thrift.NewTTransportFactory()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	// TODO : change test server
	go runServer(transportFactory, protocolFactory, SERVER_ADDRESS)
	time.Sleep(time.Second)
}

// runServer 启动服务器
func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, addr string) error {
	var transport thrift.TServerTransport
	var err error
	transport, err = thrift.NewTServerSocket(addr)
	if err != nil {
		return err
	}
	handler := NewUserHandler()
	processor := userservice.NewUserServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	return server.Serve()
}

func TestServer(t *testing.T) {
	var (
		friendResp *resp.AddFriendResp
	)
	// 建立连接
	var transport, err = thrift.NewTSocket(SERVER_ADDRESS)
	assert.NoError(t, err)
	assert.NotNil(t, transport)
	client := userservice.NewUserServiceClientFactory(transport, thrift.NewTBinaryProtocolFactoryDefault())
	assert.NotNil(t, client)
	err = transport.Open()
	assert.NoError(t, err)
	friendResp, err = client.AddFriend(&req.AddFriendReq{
		RequesterId: 100001,
		ApproverId:  100002,
	})
	assert.NotNil(t, friendResp)
	assert.NoError(t, err)
}
