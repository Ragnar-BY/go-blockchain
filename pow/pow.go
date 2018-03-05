package pow

import (
	"bytes"
	"math"
	"math/big"
	"time"

	"go-blockchain/utils"
)

const complexity = 16
const maxNonce = math.MaxInt64

type ProofOfWork struct {
	headerByte []byte
	target     *big.Int
}

func NewProofOfWork(b []byte) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-complexity))

	pow := &ProofOfWork{b, target}

	return pow
}

func (pow *ProofOfWork) Run() (int64, [32]byte, int64) {

	t := time.Now().UnixNano()
	timeByte := utils.IntToHex(t)

	bh := pow.headerByte
	data := bytes.Join(
		[][]byte{
			bh, timeByte,
		},
		[]byte{},
	)

	var nonce int64
	var nonceByte []byte
	var hash [32]byte

	var hashInt big.Int

	for nonce = 0; nonce < maxNonce; nonce++ {
		nonceByte = utils.IntToHex(nonce)
		data = bytes.Join(
			[][]byte{
				data, nonceByte,
			},
			[]byte{},
		)
		hash = utils.Hash(data)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		}
	}
	return nonce, hash, t

}
