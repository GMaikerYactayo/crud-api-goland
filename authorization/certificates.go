package authorization

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"io/ioutil"
	"os"
	"sync"
)

// Public and private certificate information
var (
	singKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
	once      sync.Once
)

// LoadFiles upload the generated files
func LoadFiles(privateFile, publicFile string) error {
	var err error
	once.Do(func() {
		err = loadFiles(privateFile, publicFile)
	})
	return err
}

// loadFiles read the content of the certificates
func loadFiles(privateFile, publicFile string) error {
	privateBytes, err := ioutil.ReadFile(privateFile)
	if err != nil {
		return err
	}
	publicBytes, err := os.ReadFile(publicFile)
	if err != nil {
		return err
	}
	return parseRSA(privateBytes, publicBytes)
}

// parseRSA parses private and public key bytes in PEM format and
// initializes the global variables singKey and verifyKey with the corresponding keys.
// If there are any errors during the process, the error is returned.
func parseRSA(privateBytes, publicBytes []byte) error {
	var err error
	singKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		return err
	}
	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		return err
	}
	return nil
}
