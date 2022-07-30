package src

import "fmt"

type Config struct {
	RPCUrl  string
	Network string
}

func InitConfig(args ParsedArgs) (config Config) {
	config = Config{
		Network: args.Network,
	}

	if args.Network != "" && args.ApiKey != "" {
		config.RPCUrl = fmt.Sprintf("https://%s.infura.io/v3/%s", args.Network, args.ApiKey)
	}

	return config
}
