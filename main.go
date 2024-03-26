package main

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/myuksal/updico/config"
	"github.com/myuksal/updico/jsonrpc"
)

func main() {

	config := config.GetConfig()
	jsonrpc.BootStrap(config.JsonRpcConfig)
	rpc := jsonrpc.GetJsonRPC()
	receipt := rpc.TransactionReceipt(common.HexToHash("0x028507096f6d825d6612c5c045b430031bd7a473b9039687531a16060b22d0be"))
	result, _ := receipt.MarshalJSON()

	log.Println(string(result))

	// db := database.CreateConnection()
	// fmt.Println(db.Config)

	// http.Handle("/metrics", promhttp.Handler())
	// go dev.CurrentBlock()

	// http.ListenAndServe(":2112", nil)
}
