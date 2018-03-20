package pow

import (
	"math"
	"math/big"

	"go-blockchain/utils"
)

const complexity = 16
const maxNonce = math.MaxInt64

type ProofOfWork struct {
	target *big.Int
}

func NewProofOfWork() *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-complexity))

	pow := &ProofOfWork{target}

	return pow
}

func (pow *ProofOfWork) Run(header [32]byte) ([8]byte, [32]byte) {

	var nonceInt int64
	var nonce [8]byte
	var hash [32]byte

	var hashInt big.Int

	for nonceInt = 0; nonceInt < maxNonce; nonceInt++ {
		n := utils.IntToHex(nonceInt)
		copy(nonce[:], n)
		hash = utils.Hash([]interface{}{
			header,
			nonce,
		})
		hashInt.SetBytes(hash[:])
		if hashInt.Cmp(pow.target) == -1 {
			break
		}
	}
	return nonce, hash
}

func (pow *ProofOfWork) IsValid(header [32]byte, nonce [8]byte) bool {

	var hash [32]byte
	var hashInt big.Int

	hash = utils.Hash([]interface{}{
		header,
		nonce,
	})
	hashInt.SetBytes(hash[:])

	if hashInt.Cmp(pow.target) == -1 {
		return true
	}
	return false
}
