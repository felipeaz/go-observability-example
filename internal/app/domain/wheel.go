package domain

type WheelState string

var (
	FlatState   = "FLAT"
	FullState   = "FULL"
	BrokenState = "BROKEN"
)

type Wheel struct {
	State WheelState
	Size  int
}
