package database

import (
	"context"
	"crypto/tls"
	"crypto/x509"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	ErrGSSAPINoneSupport = errors.New("GSSAPI mechanism none support yet")
)

const (
	Merchant = "merchant"

	GASSAPI = "GASSAPI"
)

func Open(ctx context.Context, config MGConfig) (*mongo.Client, error) {
	if config.AuthMechanism == GASSAPI {
		return nil, ErrGSSAPINoneSupport
	}

	cred := options.Credential{
		AuthMechanism:           config.AuthMechanism,
		AuthMechanismProperties: nil,
		AuthSource:              config.AuthSource,
		Username:                config.Username,
		Password:                config.Password,
		PasswordSet:             false,
	}

	clientOption := options.Client().
		SetHosts(config.Hosts).
		SetAuth(cred).
		SetReplicaSet(config.ReplicaName).
		SetTLSConfig(ssl(config.SSL)).
		SetDirect(config.IsDirect).
		SetReadPreference(config.ReadPref)

	return mongo.Connect(ctx, clientOption)
}

func ssl(mgssl *MGSSL) *tls.Config {
	if mgssl == nil {
		return nil
	}

	switch mgssl.Type {
	case sslInsecure:
		return &tls.Config{
			InsecureSkipVerify: true,
		}
	case sslCAFile, sslPemFile:
		roots := x509.NewCertPool()
		roots.AppendCertsFromPEM([]byte(mgssl.Cert))

		config := &tls.Config{}
		config.RootCAs = roots

		return config
	default:
		return nil
	}
}

// StatusCheck healthy check for mongodb.
func StatusCheck(ctx context.Context, client *mongo.Client) error {
	return client.Ping(ctx, readpref.Primary())
}
