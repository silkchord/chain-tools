package protocol

// BaseTx represent the base of raw transaction, it provide consistent method to different coin transaction
type BaseTx interface {
	ID() string
	MarshalText() ([]byte, error)
	UnmarshalText([]byte) error
	InputSize() int
	OutputSize() int
	SpentOutputIDs() []string
	RawTx() interface{}
	SetVersion(uint64)
	AddInput(interface{})
	AddOutput(interface{})
}
