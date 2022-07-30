package src

import (
	"errors"
	"fmt"
	"github.com/akamensky/argparse"
	"log"
)

type ParsedArgs struct {
	Address string
	ApiKey  string
	Code    string
	Network string
}

func ParseArgs(input []string) (args ParsedArgs) {
	parser := initParser()

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

	err := parser.Parse(input)
	if err != nil {
		log.Fatal(parser.Help(err))
	}

	args.Address = *address
	args.ApiKey = *apiKey
	args.Code = *code
	args.Network = *network

	err = validateArgs(args)
	if err != nil {
		log.Fatal(parser.Help(err))
	}

	return args
}

func initParser() (parser *argparse.Parser) {
	parser = argparse.NewParser("troy", "The EVM foot soldier")

	parser.HelpFunc = func(c *argparse.Command, msg interface{}) (result string) {
		result = parser.Usage(nil)

		if msg != nil {
			switch msg.(type) {
			case error:
				result += fmt.Sprintf("error: %s\n", msg.(error).Error())
			case string:
				result += fmt.Sprintf("error: %s\n", msg.(string))
			}
		}

		return result
	}

	return parser
}

func validateArgs(args ParsedArgs) error {
	addressIsSet := args.Address != ""

	if addressIsSet && args.ApiKey == "" {
		return errors.New("--api-key must be provided with --address")
	}

	if !addressIsSet && args.Code == "" {
		return errors.New("--address or --code are required")
	}

	return nil
}
