package dev

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/myuksal/updico/metrics"
)

var (
	ctx         = context.Background()
	url         = "https://polygon-rpc.com"
	client, err = ethclient.DialContext(ctx, url)
)

func transactionReceipt(hashChannel chan (common.Hash)) {
	for hash := range hashChannel {

		start := time.Now()
		receipt, _ := client.TransactionReceipt(ctx, hash)
		metrics.MetricRpcCallLatency.Observe(float64(time.Since(start).Milliseconds()))
		if receipt == nil {
			return
		}

		receiptResult, _ := json.Marshal(receipt)
		fmt.Println("[recript]")
		fmt.Println(string(receiptResult))

		fmt.Println("[logs]")
		for _, vLog := range receipt.Logs {
			logBytes, _ := vLog.MarshalJSON()
			fmt.Println(string(logBytes))
		}

		fmt.Println()
	}
}

func CurrentBlock() {

	for i := int64(52551917); i < 54051917; i++ {
		block, err := client.BlockByNumber(ctx, big.NewInt(i))
		if err != nil {
			log.Println(err)
		}

		block.Header()

		workerPoolSize := 3
		workerPool := make(chan common.Hash, workerPoolSize)

		for j := 0; j < (workerPoolSize); j++ {
			go transactionReceipt(workerPool)
		}

		// transaction hash publisher
		for _, transaction := range block.Transactions() {
			fmt.Println("insert ", transaction.Hash())
			workerPool <- transaction.Hash()
			transaction.Data()
		}

		// if block.Transactions().Len() > 0 {

		// 	tx, _, _ := client.TransactionByHash(ctx, block.Transactions()[0].Hash())
		// 	txResult, _ := json.Marshal(tx)
		// 	fmt.Println("[transaction]")
		// 	fmt.Print(string(txResult))
		// 	fmt.Println()

		// 	receipt, _ := client.TransactionReceipt(ctx, tx.Hash())

		// 	if receipt != nil {
		// 		receiptResult, _ := json.Marshal(tx)
		// 		fmt.Println("[recript]")
		// 		fmt.Print(string(receiptResult))
		// 		fmt.Println()
		// 	}

		// }
	}

}
