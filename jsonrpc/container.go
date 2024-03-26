package jsonrpc

import (
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/myuksal/updico/config"
)

var once sync.Once

type jsonRPCContainer struct {
	client *ethclient.Client
}

func newJsonRPCContainer(host string) *jsonRPCContainer {
	container := jsonRPCContainer{}
	client, err := ethclient.Dial(host)
	if err != nil {
		panic(err)
	}
	container.client = client
	return &container
}

var _config config.JsonRPCConfig
var _it *jsonRPCContainer

func BootStrap(config config.JsonRPCConfig) {
	// Singleton initialize
	if _it == nil {
		once.Do(
			func() {
				_config = config
				_it = newJsonRPCContainer(config.Host)
				log.Println("⚡ now json rpc is ready")
			},
		)
	}
}

func GetJsonRPC() *jsonRPCContainer {
	if _it == nil {
		panic("None-bootstrapped container (jsonRpcContainer)")
	}
	return _it
}
