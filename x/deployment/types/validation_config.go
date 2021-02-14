package types

import "github.com/ovrclk/akash/types/unit"

// This is the validation configuration that acts as a hard limit
// on what the network accepts for deployments. This is never changed
// and is the same across all members of the network

var validationConfig = struct {

	// MaxUnitCPU is the maximum number of milli (1/1000) cpu units a unit can consume.
	MaxUnitCPU uint
	// MaxUnitMemory is the maximum number of bytes of memory that a unit can consume
	MaxUnitMemory uint
	// MaxUnitStorage is the maximum number of bytes of storage that a unit can consume
	MaxUnitStorage uint
	// MaxUnitCount is the maximum number of replias of a service
	MaxUnitCount uint
	// MaxUnitPrice is the maximum price that a unit can have
	MaxUnitPrice uint

	MinUnitCPU     uint
	MinUnitMemory  uint
	MinUnitStorage uint
	MinUnitCount   uint
	MinUnitPrice   uint

	// MaxGroupCount is the maximum number of groups allowed per deployment
	MaxGroupCount int
	// MaxGroupUnits is the maximum number services per group
	MaxGroupUnits int

	// MaxGroupCPU is the maximum total amount of CPU requested per group
	MaxGroupCPU int64
	// MaxGroupMemory is the maximum total amount of memory requested per group
	MaxGroupMemory int64
	// MaxGroupStorage is the maximum total amount of storage requested per group
	MaxGroupStorage int64
}{
	MaxUnitCPU:     10 * 1000,     // 10 CPUs
	MaxUnitMemory:  16 * unit.Gi,  // 16 Gi
	MaxUnitStorage: 100 * unit.Gi, // 100 Gi
	MaxUnitCount:   50,
	MaxUnitPrice:   10000000, // 10akt

	MinUnitCPU:     10,
	MinUnitMemory:  unit.Mi,
	MinUnitStorage: 256 * unit.Mi, // 256 Mi
	MinUnitCount:   1,
	MinUnitPrice:   1,

	MaxGroupCount: 20,
	MaxGroupUnits: 20,

	MaxGroupCPU:     20 * 1000,     // 10 CPUs
	MaxGroupMemory:  32 * unit.Gi,  // 32 Gi
	MaxGroupStorage: 256 * unit.Gi, // 256 Gi
}
