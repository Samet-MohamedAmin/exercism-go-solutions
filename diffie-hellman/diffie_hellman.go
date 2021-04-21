package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

var bigTwo = big.NewInt(2)

// PrivateKey returns priveate key
func PrivateKey(p *big.Int) *big.Int {
	max := new(big.Int).Sub(p, bigTwo)
	random, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return random.Add(bigTwo, random)
}

// PublicKey returns public key
func PublicKey(private, p *big.Int, g int64) (out *big.Int) {
	return new(big.Int).Exp(big.NewInt(g), private, p)
}

// NewPair returns new key pair
func NewPair(p *big.Int, g int64) (private, public *big.Int) {
	private = PrivateKey(p)
	public = PublicKey(private, p, g)
	return
}

// SecretKey returns secret key
func SecretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
