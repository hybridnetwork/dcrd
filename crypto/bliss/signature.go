package bliss

import (
	dcrcrypto "github.com/hybridnetwork/hxd/crypto"
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
