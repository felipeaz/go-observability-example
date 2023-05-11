package domain

type FuelState string

var (
	FuelFull  FuelState = "FULL"
	FuelEmpty FuelState = "EMPTY"
)

type FuelTank struct {
	AcceptsGasoline bool      `json:"acceptsGasoline"`
	AcceptsAlcohol  bool      `json:"acceptsAlcohol"`
	State           FuelState `json:"state"`
}
