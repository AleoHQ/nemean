package account

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/blake2s"
	"math"
)

var privateKeyPrefix = []byte{127, 134, 189, 116, 210, 221, 210, 137, 144}

var errInvalidSeed = errors.New("invalid seed")

type PrivateKey struct {
	Seed []byte
	RPkCounter uint16
}

// String implements the stringer interface for PrivateKey.
// Returns the base58 encoded string.
func (pk PrivateKey) String() string {
	var buf bytes.Buffer
	buf.Write(privateKeyPrefix)
	binary.Write(&buf, binary.LittleEndian, pk.RPkCounter)
	buf.Write(pk.Seed)
	return base58.Encode(buf.Bytes())
}

func ParsePrivateKey(key string) (*PrivateKey, error) {
	// An account private key is formatted as a Base58 string, comprised of 58 characters.
	buf := base58.Decode(key)

	if keyLen := len(buf); keyLen != 43 {
		return nil, fmt.Errorf("invalid key length : got %d", keyLen)
	}

	if !bytes.Equal(buf[0:9], privateKeyPrefix) {
		return nil, errors.New("invalid prefix")
	}

	// First 2 bytes are the counter.
	// Last 32 bytes are the seed.
	counter, seed := buf[9:11], buf[11:]

	return &PrivateKey{
		Seed:       seed,
		RPkCounter: binary.LittleEndian.Uint16(counter),
	}, nil
}


func NewPrivateKey() {
	// Create a uniformly random 32-byte account seed.
	d := make([]byte, 32)
	rand.Read(d)
}

type Account struct {
	PrivateKey PrivateKey
	Address Address
}

func FromSeed(seed []byte) (*Account, error) {
	// Generate the SIG key pair.
	skSig, err := blake2s.New256(append(seed, 0x00))
	if err != nil {
		return nil, err
	}

	// Generate the PRF key pair.
	skPrf, err := blake2s.New256(append(seed, 0x01))
	if err != nil {
		return nil, err
	}

	// counter is a u16 value that is iterated on until a valid view_key
	// can be derived from private_key.
	// TODO
	var counter uint16 = 2

	for {
		if counter > math.MaxUint16 {
			return nil, errInvalidSeed
		}

		buf, err := blake2s.New256(append(seed, 0x00))
		if err != nil {
			return nil, err
		}

		if ValidPrivateKey(buf.Sum(nil)) {
			break
		}

		counter += 1
	}

	print("%v %v",skSig, skPrf)
	return &Account{
		PrivateKey: PrivateKey{Seed: seed, RPkCounter: counter},
		Address:    Address{},
	}, nil

}

func ValidPrivateKey(key []byte) (bool) {
	// Generate a ViewKey from the private key.
	// A ViewKey is a Schnorr Public Key.
	// For now, return true.
	return true
}

// An account private key is constructed from a randomly-sampled account seed. This account seed is used to generate:
//a secret key for the account signature scheme,
//a pseudorandom function seed for transaction serial numbers, and
//a commitment randomness for the account commitment scheme.