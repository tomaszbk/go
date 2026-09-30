package mlkem_test

import (
	"crypto/mlkem"
	"log"
)

func Example() {
	// Alice generates a new key pair and sends the encapsulation key to Bob.
	dk, err := mlkem.GenerateKey768()
	if err != nil {
		log.Fatal(err)
	}
	encapsulationKey := dk.EncapsulationKey().Bytes()

	// Bob uses the encapsulation key to encapsulate a shared secret, and sends
	// back the ciphertext to Alice.
	ciphertext := Bob(encapsulationKey)

	// Alice decapsulates the shared secret from the ciphertext.
	sharedSecret, err := dk.Decapsulate(ciphertext)
	if err != nil {
		log.Fatal(err)
	}

	// Alice and Bob now share a secret.
	_ = sharedSecret
}

func Bob(encapsulationKey []byte) (ciphertext []byte) {
	// Bob encapsulates a shared secret using the encapsulation key.
	ek, err := mlkem.NewEncapsulationKey768(encapsulationKey)
	if err != nil {
		log.Fatal(err)
	}
	sharedSecret, ciphertext := ek.Encapsulate()

	// Alice and Bob now share a secret.
	_ = sharedSecret

	// Bob sends the ciphertext to Alice.
	return ciphertext
}
