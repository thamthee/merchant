package auth

import (
	"crypto/rsa"
	"sync"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/pkg/errors"
)

// These are the expected values for Claims.Roles.
const (
	RoleAdmin    = "ADMIN"
	RoleSeller   = "SELLER"
	RoleCustomer = "CUSTOMER"
)

type ctxKey int

const Key ctxKey = 1

type Claims struct {
	jwt.StandardClaims
	Roles []string `json:"roles"`
}

func (c Claims) Authorized(roles ...string) bool {
	for _, has := range c.Roles {
		for _, want := range roles {
			if has == want {
				return true
			}
		}
	}
	return false
}

type Keys map[string]*rsa.PrivateKey

type PublicKeyLookup func(kid string) (*rsa.PublicKey, error)

type Auth struct {
	mu        sync.RWMutex
	algorithm string
	method    jwt.SigningMethod
	keyFunc   func(t *jwt.Token) (interface{}, error)
	parser    *jwt.Parser
	keys      Keys
}

func New(algorithm string, lookup PublicKeyLookup, keys Keys) (*Auth, error) {
	method := jwt.GetSigningMethod(algorithm)
	if method == nil {
		return nil, errors.Errorf("unknow algorithm %v", algorithm)
	}

	keyFunc := func(t *jwt.Token) (interface{}, error) {
		kid, ok := t.Header["kid"]
		if !ok {
			return nil, errors.New("missing key id (kid) in token header")
		}
		kidID, ok := kid.(string)
		if !ok {
			return nil, errors.New("user token key id (kid) must be string")
		}
		return lookup(kidID)
	}

	parser := jwt.NewParser(jwt.WithValidMethods([]string{algorithm}), jwt.WithAudience("merchant"))

	a := Auth{
		algorithm: algorithm,
		method:    method,
		keyFunc:   keyFunc,
		parser:    parser,
		keys:      keys,
	}

	return &a, nil
}

func (a *Auth) AddKey(privateKey *rsa.PrivateKey, kid string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.keys[kid] = privateKey
}

func (a *Auth) RemoveKey(kid string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	delete(a.keys, kid)
}

func (a *Auth) GenerateToken(kid string, claims Claims) (string, error) {
	token := jwt.NewWithClaims(a.method, claims)
	token.Header["kid"] = kid

	var privateKey *rsa.PrivateKey
	a.mu.RLock()
	{
		var ok bool
		privateKey, ok = a.keys[kid]
		if !ok {
			return "", errors.New("kid lookup failed")
		}
	}
	a.mu.RUnlock()

	str, err := token.SignedString(privateKey)
	if err != nil {
		return "", errors.Wrap(err, "signing token")
	}

	return str, nil
}

func (a *Auth) ValidateToken(tokenStr string) (Claims, error) {
	var claims Claims
	token, err := a.parser.ParseWithClaims(tokenStr, &claims, a.keyFunc)
	if err != nil {
		return Claims{}, errors.Wrap(err, "parsing token")
	}

	if !token.Valid {
		return Claims{}, errors.New("invalid token")
	}

	return claims, nil
}
