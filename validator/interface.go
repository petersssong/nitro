package validator

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/nitro/util/readymarker"
)

type ValidationSpawner interface {
	Launch(entry *ValidationInput, moduleRoot common.Hash) ValidationRun
	Start(context.Context) error
	Stop()
	Name() string
	Room() int
}

type ValidationRun interface {
	readymarker.ReadyMarkerInt
	WasmModuleRoot() common.Hash
	Result() (GoGlobalState, error)
	Close()
}

type ExecutionSpawner interface {
	ValidationSpawner
	CreateExecutionRun(wasmModuleRoot common.Hash, input *ValidationInput) (ExecutionRun, error)
	LatestWasmModuleRoot() (common.Hash, error)
	WriteToFile(input *ValidationInput, expOut GoGlobalState, moduleRoot common.Hash) error
}

type ExecutionRun interface {
	GetStepAt(uint64) MachineStep
	GetLastStep() MachineStep
	PrepareRange(uint64, uint64)
	Close()
}

type MachineStep interface {
	readymarker.ReadyMarkerInt
	Hash() (common.Hash, error)
	Proof() ([]byte, error)
	Position() (uint64, error)
	Status() (MachineStatus, error)
	GlobalState() (GoGlobalState, error)
	Close()
}
