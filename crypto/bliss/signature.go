package bliss

import (
	"github.com/LoCCS/bliss"
	dcrcrypto "github.com/decred/dcrd/crypto"
)

type Signature struct{
	dcrcrypto.SignatureAdapter
	bliss.Signature
}

func (s Signature) GetType() int {
	return pqcTypeBliss
}

func (s Signature) Serialize() []byte{
	return s.Signature.Serialize()
}