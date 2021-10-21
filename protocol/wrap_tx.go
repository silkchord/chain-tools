package protocol

// WrapTx represent the wrap of raw transaction, it provide consistent method to different coin transaction
type WrapTx interface {
	BaseTx
	SerializedSize() uint64
	InputControlProgram(int) []byte
	InputArguments(int) [][]byte
	SetInputArguments(int, [][]byte)
	SigHash(int) string
	SignHashBytes(int) []byte
	InputAssetAmount(int) (string, uint64)
	OutputAssetAmount(int) (string, uint64)
	SetTimeRange(uint64)
	New() WrapTx
	Build()
}
