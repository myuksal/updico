package jsonrpc

import (
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/myuksal/updico/config"
)

var once sync.Once
var rpc *jsonRPCContainer

type jsonRPCContainer struct {
	client *ethclient.Client
}

func BootStrap(config config.JsonRPCConfig) {
	// Singleton initialize
	if rpc == nil {
		once.Do(
			func() {
				client, err := ethclient.Dial(config.Host)
				if err != nil {
					panic(err)
				}
				container := jsonRPCContainer{}
				container.client = client

				rpc = &container
				log.Println("⚡ now json rpc is ready")
			},
		)
	}
}

func GetJsonRpc() *jsonRPCContainer {
	if rpc == nil {
		panic("JsonRpc is not bootstrapped")
	}
	return rpc
}
