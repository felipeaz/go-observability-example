package domain

type FuelState string

var (
	Full    = "FULL"
	Empty   = "EMPTY"
	Reserve = "RESERVE"
)

type FuelTank struct {
	AcceptsGasoline bool
	AccepsAlcohol   bool
	State           FuelState
}
