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
		Table        TableColors
	}

	InstructionColors struct {
		Opcode string
	}

	TableColors struct {
		Default           ColorFlag
		Selected          ColorFlag
		ProgramCounter    string                  `json:"program_counter"`
		FunctionSignature FunctionSignatureColors `json:"function_signatures"`
	}

	FunctionSignatureColors struct {
		Divider            string
		Function           string
		PossibleSignatures string `json:"possible_signatures"`
	}

	ColorFlag struct {
		Foreground string
		Background string
		Flags      string
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
	instructions, err := os.ReadFile("config/colors/instructions.json")
	if err != nil {
		log.WithField("Error", err).Fatal("Failed to initialise instructions color config")
	}

	err = json.Unmarshal(instructions, &colors)
	if err != nil {
		log.WithField("Error", err).Fatal("Failed to unmarshal instructions color config")
	}

	table, err := os.ReadFile("config/colors/table.json")
	if err != nil {
		log.WithField("Error", err).Fatal("Failed to initialise table color config")
	}

	err = json.Unmarshal(table, &colors)
	if err != nil {
		log.WithField("Error", err).Fatal("Failed to unmarshal table color config")
	}

	return colors
}
