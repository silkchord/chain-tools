package singer

import (
	"bytes"
	"encoding/hex"

	"github.com/bytom/bytom/crypto/ed25519/chainkd"
	mnem "github.com/bytom/bytom/wallet/mnemonic"
	"github.com/pkg/errors"

	"github.com/songqianba/chain-tools/protocol"
)

// Signer tx signer
type Signer struct {
	xprv chainkd.XPrv
}

// NewSigner return Signer
func NewSigner(mnemonic string) (*Signer, error) {
	xprv, _, err := chainkd.NewXKeys(bytes.NewBuffer(mnem.NewSeed(mnemonic, "")))
	if err != nil {
		return nil, err
	}

	return &Signer{
		xprv: xprv,
	}, nil
}

// Sign sign tx
func (s *Signer) Sign(tx protocol.WrapTx, signInstructions ...SigningInstructionTemp) (witnesses [][]string, err error) {
	for i, signingInstruction := range signInstructions {
		if signingInstruction.GetPubKey() == "" {
			witnesses = append(witnesses, []string{})
			continue
		}

		path := make([][]byte, len(signingInstruction.GetDerivationPath()))
		for i, p := range signingInstruction.GetDerivationPath() {
			if path[i], err = hex.DecodeString(p); err != nil {
				return nil, errors.Wrap(err, "decode path err")
			}
		}

		xprv := s.xprv.Derive(path)
		if signingInstruction.GetPubKey() != hex.EncodeToString(xprv.XPub().PublicKey()) {
			return nil, errors.New("didn't find Signer")
		}

		signData := tx.SignHashBytes(i)
		witness := []string{hex.EncodeToString(xprv.Sign(signData))}
		witnesses = append(witnesses, witness)
	}
	return witnesses, nil
}
