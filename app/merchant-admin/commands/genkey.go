package commands

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

// ErrHelp provides context that help was given.
var ErrHelp = errors.New("provided help")

func GenKey() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return ErrHelp
	}

	privateFile, err := os.Create("private.pem")
	if err != nil {
		return errors.Wrap(err, "creating private")
	}
	defer privateFile.Close()

	privateBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   x509.MarshalPKCS1PrivateKey(privateKey),
	}

	if err := pem.Encode(privateFile, &privateBlock); err != nil {
		return errors.Wrap(err, "encoding to private file")
	}

	asn1Bytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return errors.Wrap(err, "marshaling public key")
	}

	publicFile, err := os.Create("public.pem")
	if err != nil {
		return errors.Wrap(err, "creating public file")
	}
	defer publicFile.Close()

	publicBlock := pem.Block{
		Type:    "RSA PUBLIC KEY",
		Headers: nil,
		Bytes:   asn1Bytes,
	}

	if err := pem.Encode(publicFile, &publicBlock); err != nil {
		return errors.Wrap(err, "encoding to public")
	}

	fmt.Println("private and public key files generated")
	return nil
}
