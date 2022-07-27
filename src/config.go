package src

import (
	"fmt"
	"log"
	"os"

	"github.com/akamensky/argparse"
)

type Config struct {
	Address string
	RPCUrl  string
	Network string
	Code    string
}

func InitConfig() Config {
	parser := argparse.NewParser("troy", "The EVM foot soldier")

	address := parser.String("a", "address", &argparse.Options{
		Required: false, Help: "The contract address to analyse",
	})

	apiKey := parser.String("k", "api-key", &argparse.Options{
		Required: false, Help: "Infura.io API key",
	})

	code := parser.String("c", "code", &argparse.Options{
		Required: false, Help: "Byte code hex string",
	})

	network := parser.String("n", "network", &argparse.Options{
		Help: "The Infura network to use", Default: "mainnet",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		log.Fatal(parser.Usage(err))
	}

	addressIsSet := len(*address) > 0

	if addressIsSet && len(*apiKey) <= 0 {
		log.Fatal(parser.Usage("--apiKey must be provided with --address"))
	}

	if !addressIsSet && len(*code) <= 0 {
		log.Fatal(parser.Usage("--address or --code must be provided"))
	}

	return Config{
		Address: *address,
		RPCUrl:  fmt.Sprintf("https://%s.infura.io/v3/%s", *network, *apiKey),
		Network: *network,
		Code:    *code,
	}
}
