package btm

import (
	"github.com/bytom/bytom/protocol/bc/types"
	"github.com/pkg/errors"
)

// UnserializeTx unserialize btyom tx
func UnserializeTx(rawTx interface{}) (*types.Tx, error) {
	txStr, ok := rawTx.(string)
	if !ok {
		return nil, errors.New("decode raw_tx")
	}

	tx := &types.Tx{}
	if err := tx.UnmarshalText([]byte(txStr)); err != nil {
		return nil, err
	}

	return tx, nil
}
