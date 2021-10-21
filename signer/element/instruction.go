package element

// SigningInstructionTemp template
type SigningInstructionTemp interface {
	GetDerivationPath() []string
	GetSignData() []string
	GetDataWitness() []byte
	GetPubKey() string
}

// SigningInstruction represent the signing instruction
type SigningInstruction struct {
	DerivationPath []string `json:"derivation_path"`
	SignData       []string `json:"sign_data"`
	DataWitness    []byte   `json:"-"`

	// only shown for a single-signature tx
	Pubkey string `json:"pubkey,omitempty"`
}

// GetDerivationPath Get DerivationPath
func (s *SigningInstruction) GetDerivationPath() []string {
	return s.DerivationPath
}

// GetSignData Get SignData
func (s *SigningInstruction) GetSignData() []string {
	return s.SignData
}

// GetDataWitness Get DataWitness
func (s *SigningInstruction) GetDataWitness() []byte {
	return s.DataWitness
}

// GetPubKey get pub key
func (s *SigningInstruction) GetPubKey() string {
	return s.Pubkey
}
