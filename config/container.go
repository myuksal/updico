package config

import (
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

var once sync.Once

const CONFIG_PATH = "config.yaml"

type PrometheusConfig struct {
	Port int32 `yaml:"port"`
}

type JsonRPCConfig struct {
	Host    string `yaml:"host"`
	Timeout struct {
		TransactionReceipt int32 `yaml:"transactionReceipt" default:"10000"`
	} `yaml:"timeout"`
}

type config struct {
	JsonRpcConfig    JsonRPCConfig    `yaml:"jsonRpc"`
	PrometheusConfig PrometheusConfig `yaml:"prometheus"`
}

var _it *config

func GetConfig() *config {
	if _it == nil {
		once.Do(
			func() {
				data, err := os.ReadFile(CONFIG_PATH)
				if err != nil {
					panic(err)
				}

				if err := yaml.Unmarshal(data, &_it); err != nil {
					panic(err)
				}
			},
		)
	}

	return _it
}
