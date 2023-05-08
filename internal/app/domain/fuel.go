package domain

type FuelState string

var (
	FuelFull  FuelState = "FULL"
	FuelEmpty FuelState = "EMPTY"
)

type FuelTank struct {
	AcceptsGasoline bool
	AcceptsAlcohol  bool
	State           FuelState
}
