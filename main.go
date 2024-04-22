package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/myuksal/updico/config"
	"github.com/myuksal/updico/jsonrpc"
	"github.com/myuksal/updico/transform"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	config.BootStrap()
	configMap := config.GetConfig()

	jsonrpc.BootStrap(configMap.JsonRpcConfig)
	rpc := jsonrpc.GetJsonRpc()

	fmt.Println("json rpc => ", configMap.JsonRpcConfig.Host)

	http.HandleFunc("/block", func(w http.ResponseWriter, r *http.Request) {
		blockNumberString := r.URL.Query().Get("number")
		blockNumber, _ := strconv.Atoi(blockNumberString)

		hid := uuid.New()
		ctx := context.WithValue(context.Background(), "request-id", hid.String())

		block, err := rpc.Block(ctx, *big.NewInt(int64(blockNumber)))
		if err != nil {
			fmt.Println("block rpc error => ", err)
		}

		blockJson, err := json.Marshal(transform.BlockToResponse(block))
		if err != nil {
			fmt.Println("block parse error => ", err)
		}

		w.Write(blockJson)

	})

	http.Handle("/metrics", promhttp.Handler())

	http.ListenAndServe(":2112", nil)

}
