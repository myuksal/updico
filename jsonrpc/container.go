package jsonrpc

import (
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/myuksal/updico/config"
)

var once sync.Once

var _config config.JsonRPCConfig

var JsonRPC *jsonRPCContainer

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

func BootStrap(config config.JsonRPCConfig) {
	// Singleton initialize
	if JsonRPC == nil {
		once.Do(
			func() {
				_config = config
				JsonRPC = newJsonRPCContainer(config.Host)
				log.Println("⚡ now json rpc is ready")
			},
		)
	}
}
