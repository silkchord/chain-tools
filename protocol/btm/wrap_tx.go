package btm

import (
	"github.com/bytom/bytom/protocol/bc/types"

	"github.com/songqianba/chain-tools/protocol"
)

// WrapTx represent the wrap of bytomclassic transaction
type WrapTx struct {
	tx *types.Tx
}

// NewWrapTx return a new instance of the WrapTx
func NewWrapTx(tx *types.Tx) *WrapTx {
	return &WrapTx{tx: tx}
}

// ID return the hash of tx
func (b *WrapTx) ID() string {
	return b.tx.ID.String()
}

// SerializedSize return the SerializedSize of the tx
func (b *WrapTx) SerializedSize() uint64 {
	return b.tx.SerializedSize
}

// MarshalText return the marshaled text of the tx
func (b *WrapTx) MarshalText() ([]byte, error) {
	return b.tx.MarshalText()
}

// UnmarshalText used to convet byte data to bytomclassic tx
func (b *WrapTx) UnmarshalText(data []byte) error {
	b.tx = &types.Tx{}
	return b.tx.UnmarshalText(data)
}

// InputSize return size of tx's inputs
func (b *WrapTx) InputSize() int {
	return len(b.tx.Inputs)
}

// OutputSize return size of tx's Outputs
func (b *WrapTx) OutputSize() int {
	return len(b.tx.Outputs)
}

// SpentOutputIDs return tx's spent output id list
func (b *WrapTx) SpentOutputIDs() []string {
	var outputIDs []string
	for _, outputID := range b.tx.SpentOutputIDs {
		outputIDs = append(outputIDs, outputID.String())
	}
	return outputIDs
}

// InputAssetAmount return the asset and amount of inputs[i]
func (b *WrapTx) InputAssetAmount(index int) (string, uint64) {
	assetID := b.tx.Inputs[index].AssetID()
	return assetID.String(), b.tx.Inputs[index].Amount()
}

// OutputAssetAmount return the asset and amount of outputs[i]
func (b *WrapTx) OutputAssetAmount(index int) (string, uint64) {
	assetAmount := b.tx.Outputs[index].AssetAmount
	return assetAmount.AssetId.String(), assetAmount.Amount
}

// SigHash return tx's input hash for signature
func (b *WrapTx) SigHash(index int) string {
	hash := b.tx.SigHash(uint32(index))
	return hash.String()
}

// SignHashBytes return tx's input hash bytes for signature
func (b *WrapTx) SignHashBytes(index int) []byte {
	hash := b.tx.SigHash(uint32(index))
	return hash.Bytes()
}

// InputControlProgram return the inputs[i]'s ControlProgram
func (b *WrapTx) InputControlProgram(i int) []byte {
	return b.tx.Inputs[i].ControlProgram()
}

// InputArguments return the inputs[i]'s arguments
func (b *WrapTx) InputArguments(i int) [][]byte {
	return b.tx.Inputs[i].Arguments()
}

// SetInputArguments used to set the inputs[i]'s arguments
func (b *WrapTx) SetInputArguments(i int, witness [][]byte) {
	b.tx.SetInputArguments(uint32(i), witness)
}

// RawTx return the raw transaction
func (b *WrapTx) RawTx() interface{} {
	return b.tx
}

// SetVersion set the tx's version
func (b *WrapTx) SetVersion(version uint64) {
	b.tx.Version = version
}

// SetTimeRange set the tx's time range
func (b *WrapTx) SetTimeRange(timeRange uint64) {
	b.tx.TimeRange = timeRange
}

// AddInput add a new input to tx
func (b *WrapTx) AddInput(input interface{}) {
	b.tx.Inputs = append(b.tx.Inputs, input.(*types.TxInput))
}

// AddOutput add a new output to tx
func (b *WrapTx) AddOutput(output interface{}) {
	b.tx.Outputs = append(b.tx.Outputs, output.(*types.TxOutput))
}

// New return a new instance of WrapTx
func (b *WrapTx) New() protocol.WrapTx {
	return &WrapTx{&types.Tx{}}
}

// Build used to build the txData
func (b *WrapTx) Build() {
	b.tx = types.NewTx(b.tx.TxData)
}
