package rsa_test

import (
	"crypto/rsa"
	"testing"

	ursa "github.com/spacelavr/pandora/pkg/utils/crypto/rsa"
	"github.com/stretchr/testify/assert"
)

func setup(t *testing.T) (*rsa.PrivateKey, *rsa.PublicKey) {
	private, public, err := ursa.GenerateKeys()
	assert.NoError(t, err)
	assert.NotNil(t, private)
	assert.NotNil(t, public)

	return private, public
}

func TestGenerateKeys(t *testing.T) {
	_, _ = setup(t)
}

func TestSignPSS(t *testing.T) {
	private, _ := setup(t)

	signature, err := ursa.SignPSS(private)
	assert.NoError(t, err)
	assert.NotNil(t, signature)
}

func TestVerifyPSS(t *testing.T) {
	private, public := setup(t)

	signature, err := ursa.SignPSS(private)
	assert.NoError(t, err)
	assert.NotNil(t, signature)

	cases := []struct {
		error     error
		signature []byte
		name      string
	}{{
		name:      "verifying sugnature success",
		error:     nil,
		signature: signature,
	}, {
		name:      "verifying signature failed",
		error:     rsa.ErrVerification,
		signature: []byte("signature"),
	}}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err = ursa.VerifyPSS(public, c.signature)
			assert.Equal(t, c.error, err)
		})
	}
}