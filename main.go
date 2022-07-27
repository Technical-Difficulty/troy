package main

import (
	"os"
	"fmt"
	"time"
	"testing"
	"math/big"
	goruntime "runtime"

	"github.com/ethereum/go-ethereum/eth/tracers/logger"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/rawdb"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/core/state"
  "github.com/ethereum/go-ethereum/core/vm/runtime"
)

type execStats struct {
	time           time.Duration // The execution time.
	allocs         int64         // The number of heap allocations during execution.
	bytesAllocated int64         // The cumulative number of bytes allocated during execution.
}

func timedExec(bench bool, execFunc func() ([]byte, uint64, error)) (output []byte, gasLeft uint64, stats execStats, err error) {
	if bench {
		result := testing.Benchmark(func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				output, gasLeft, err = execFunc()
			}
		})

		// Get the average execution time from the benchmarking result.
		// There are other useful stats here that could be reported.
		stats.time = time.Duration(result.NsPerOp())
		stats.allocs = result.AllocsPerOp()
		stats.bytesAllocated = result.AllocedBytesPerOp()
	} else {
		var memStatsBefore, memStatsAfter goruntime.MemStats
		goruntime.ReadMemStats(&memStatsBefore)
		startTime := time.Now()
		output, gasLeft, err = execFunc()
		stats.time = time.Since(startTime)
		goruntime.ReadMemStats(&memStatsAfter)
		stats.allocs = int64(memStatsAfter.Mallocs - memStatsBefore.Mallocs)
		stats.bytesAllocated = int64(memStatsAfter.TotalAlloc - memStatsBefore.TotalAlloc)
	}

	return output, gasLeft, stats, err
}

func main() {
	fmt.Println("The EVM foot soldier")
  
	logconfig := &logger.Config{
		EnableMemory:     false, 
		DisableStack:     false, 
		DisableStorage:   false, 
		EnableReturnData: true, 
		Debug:            true,
	}

	var (
		tracer        vm.EVMLogger
		debugLogger   *logger.StructLogger
		statedb       *state.StateDB
		sender        = common.BytesToAddress([]byte("sender"))
		receiver      = common.BytesToAddress([]byte("receiver"))
		genesisConfig *core.Genesis
	)

  debugLogger = logger.NewStructLogger(logconfig)
  tracer = debugLogger

  statedb, _ = state.New(common.Hash{}, state.NewDatabase(rawdb.NewMemoryDatabase()), nil)
  genesisConfig = new(core.Genesis)

	statedb.CreateAccount(sender)

  code := common.FromHex("608060405234801561001057600080fd5b5060fb8061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c806342e2d4d914602d575b600080fd5b605860048036036040811015604157600080fd5b506001600160a01b0381358116916020013516605a565b005b6040805163095ea7b360e01b81526001600160a01b0383811660048301526003602483015291519184169163095ea7b39160448082019260009290919082900301818387803b15801560ab57600080fd5b505af115801560be573d6000803e3d6000fd5b50505050505056fea265627a7a7231582074d2a91f880b7effa21aa471ea48b6bcf75305a0c1e9d86c0e0fda548a127b2464736f6c63430005110032")

	runtimeConfig := runtime.Config{
		Origin:      sender,
		State:       statedb,
		GasLimit:    genesisConfig.GasLimit,
		GasPrice:    new(big.Int), 
		Value:       new(big.Int),
		Difficulty:  genesisConfig.Difficulty,
		Time:        new(big.Int).SetUint64(genesisConfig.Timestamp),
		Coinbase:    genesisConfig.Coinbase,
		BlockNumber: new(big.Int).SetUint64(genesisConfig.Number),
		EVMConfig: vm.Config{
			Tracer: tracer,
			Debug:  true,
		},
	}

  runtimeConfig.ChainConfig = params.AllEthashProtocolChanges

	input := common.FromHex("42e2d4d900000000000000000000000023bbc53e1904b589d685e811c5d3410146f0ab0200000000000000000000000023bbc53e1904b589d685e811c5d3410146f0ab02")
  statedb.SetCode(receiver, code)

	var execFunc func() ([]byte, uint64, error)
  execFunc = func() ([]byte, uint64, error) {
    return runtime.Call(receiver, input, &runtimeConfig)
  }

	output, _, _, _ := timedExec(false, execFunc)

  fmt.Fprintln(os.Stderr, "#### TRACE ####")
  logger.WriteTrace(os.Stderr, debugLogger.StructLogs())
		
  fmt.Fprintln(os.Stderr, "#### LOGS ####")
  logger.WriteLogs(os.Stderr, statedb.Logs())

		fmt.Printf("%#x\n", output)
}
