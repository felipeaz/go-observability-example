package domain

type WheelState string

var (
	FlatWheel WheelState = "FLAT"
	FullWheel WheelState = "FULL"
)

type Wheel struct {
	State WheelState
	Size  int
}
