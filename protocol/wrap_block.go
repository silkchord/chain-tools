package protocol

// WrapBlock represent the wrap of raw block, it provide consistent method to different coin blocks
type WrapBlock interface {
	Hash() string
	Height() uint64
	Timestamp() uint64
	PrevBlockHash() string
	RawBlock() interface{}
	Transactions() []BaseTx
}
