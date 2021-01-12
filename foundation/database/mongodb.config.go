package database

import (
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MGConfig represent mongodb configuration properties to use the db specifically.
type MGConfig struct {
	AuthMechanism string
	AuthSource    string
	Username      string
	Password      string
	Hosts         []string
	ReplicaName   string
	SSL           *MGSSL
	IsDirect      bool
	ReadPref      *readpref.ReadPref
}

const (
	sslInsecure = "INSECURE"
	sslCAFile   = "CA_FILE"
	sslPemFile  = "PEN_FILE"
)

// MGSSL secure options for the database connection.
type MGSSL struct {
	Type string
	Cert string
}
