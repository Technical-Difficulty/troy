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
		EnableMemory:     true, 
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

  code := common.FromHex("60ff60ff")

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

	input := common.FromHex(string(""))
  statedb.SetCode(receiver, code)

	var execFunc func() ([]byte, uint64, error)
  execFunc = func() ([]byte, uint64, error) {
    return runtime.Call(receiver, input, &runtimeConfig)
  }

	output, _, _, _ := timedExec(false, execFunc)

  // f, err := os.Create(memProfilePath)
  // pprof.WriteHeapProfile(f)
  // f.Close()

  fmt.Fprintln(os.Stderr, "#### TRACE ####")
  logger.WriteTrace(os.Stderr, debugLogger.StructLogs())
		
  fmt.Fprintln(os.Stderr, "#### LOGS ####")
  logger.WriteLogs(os.Stderr, statedb.Logs())

		fmt.Printf("%#x\n", output)
}
