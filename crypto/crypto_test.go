package crypto
import (
	"fmt"
	_"github.com/hybridnetwork/hxd/chaincfg/chainec"
	"testing"
)

func TestCrypto(t *testing.T) {
	fmt.Println("test start")
	var pk PublicKey
	pk = new(PublicKeyAdapter)
	fmt.Println(pk.GetType())
}