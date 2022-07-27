package src

import (
	"fmt"
	"github.com/akamensky/argparse"
	"log"
	"os"
)

type Config struct {
	Address string
	RPCUrl  string
	Network string
}

func InitConfig() Config {
	parser := argparse.NewParser("troy", "The EVM foot soldier")

	address := parser.String("a", "address", &argparse.Options{
		Required: true, Help: "The contract address to analyse",
	})

	apiKey := parser.String("k", "api-key", &argparse.Options{
		Required: true, Help: "Infura.io API key",
	})

	network := parser.String("n", "network", &argparse.Options{
		Help: "The Infura network to use", Default: "mainnet",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		log.Fatal(parser.Usage(err))
	}

	return Config{
		Address: *address,
		RPCUrl:  fmt.Sprintf("https://%s.infura.io/v3/%s", *network, *apiKey),
		Network: *network,
	}
}
