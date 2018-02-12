package bliss

import (
	dcrcrypto "github.com/decred/dcrd/crypto"
	"github.com/hybridnetwork/bliss"
)

type Signature struct {
	dcrcrypto.SignatureAdapter
	bliss.Signature
}

func (s Signature) GetType() int {
	return pqcTypeBliss
}

func (s Signature) Serialize() []byte {
	return s.Signature.Serialize()
}
