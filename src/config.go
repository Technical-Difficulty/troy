package src

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
)

type (
	Config struct {
		RPCUrl  string
		Network string
		Colors  ColorsConfig
	}

	ColorsConfig struct {
		Instructions map[string]InstructionColors
	}

	InstructionColors struct {
		Opcode ColorTags
	}

	ColorTags struct {
		Prefix string
		Suffix string
	}
)

func InitConfig(args ParsedArgs) (config Config) {
	config = Config{
		Network: args.Network,
		Colors:  initColors(),
	}

	if args.Network != "" && args.ApiKey != "" {
		config.RPCUrl = fmt.Sprintf("https://%s.infura.io/v3/%s", args.Network, args.ApiKey)
	}

	return config
}

func initColors() (colors ColorsConfig) {
	dat, err := os.ReadFile("config/colors/instructions.json")
	if err != nil {
		log.WithField("Error", err).Fatal("Failed to initialise color config")
	}

	err = json.Unmarshal(dat, &colors)
	if err != nil {
		log.WithField("Error", err).Fatal("Failed to initialise color config")
	}

	return colors
}
