package metrics

const (
	MetricRpcCallCountName = "rpc_call_count"
	MetricRpcCallCountHelp = "json rpc call count gauge"

	MetricRpcCallLatencyName = "rpc_call_latency"
	MetricRpcCallLatencyHelp = "json rpc response time histogram"

	CollectMetricNamespace = "updico_collect"

	// 블록 높이
	CollectMetricBlockHeight        = "block_height"
	CollectMetricChainBlockHeight   = "chain_block_height"
	CollectMetricIndexedBlockHeight = "indexed_block_height"

	// 처리된 블록 수
	CollectMetricBlockCount               = "block_count"
	CollectMetricTotalIndexedBlocks       = "total_indexed_blocks"
	CollectMetricBlocksProcessedPerSecond = "blocks_processed_per_second"

	// 트랜잭션 처리량
	CollectMetricTxProcessed                  = "tx_processed"
	CollectMetricTps                          = "tps"
	CollectMetricTransactionsPerSecondByChain = "transactions_per_second_by_chain"
)
