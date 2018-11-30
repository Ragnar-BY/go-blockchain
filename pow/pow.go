package pow

import (
	"math"
	"math/big"

	"github.com/Ragnar-BY/go-blockchain/utils"
)

const complexity = 16
const maxNonce = math.MaxInt64

// ProofOfWork is PoW struct
type ProofOfWork struct {
	target *big.Int
}

// NewProofOfWork creates new PoW.
func NewProofOfWork() *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-complexity))

	pow := &ProofOfWork{target}

	return pow
}

// Run starts PoW process.
func (pow *ProofOfWork) Run(header [32]byte) ([8]byte, [32]byte, error) {

	var nonceInt uint64
	var nonce [8]byte
	var hash [32]byte

	var hashInt big.Int

	for nonceInt = 0; nonceInt < maxNonce; nonceInt++ {
		n := utils.UintToHex(nonceInt)
		copy(nonce[:], n)
		var err error
		hash, err = utils.EncodeAndHash([]interface{}{
			header,
			nonce,
		})
		if err != nil {
			return [8]byte{}, [32]byte{}, err
		}
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			break
		}
	}
	return nonce, hash, nil
}

// IsValid checks if header+nonce can be saved in blockchain.
func (pow *ProofOfWork) IsValid(header [32]byte, nonce [8]byte) (bool, error) {

	var hash [32]byte
	var hashInt big.Int

	hash, err := utils.EncodeAndHash([]interface{}{
		header,
		nonce,
	})
	if err != nil {
		return false, err
	}
	hashInt.SetBytes(hash[:])

	if hashInt.Cmp(pow.target) == -1 {
		return true, nil
	}
	return false, nil
}
