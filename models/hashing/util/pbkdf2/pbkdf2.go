package pbkdf2

import (
	"Scheduler/models/hashing/crypto/aes"
	"bytes"
	"crypto/sha256"
	"golang.org/x/crypto/pbkdf2"
)

type Pbkdf2 struct {
	Key    []byte
	Salt   []byte
	Plain  []byte
	Cipher []byte
	Itr    int
	KeyLen int
}

func (obj *Pbkdf2) Encrypt() (err error) {
	obj.Key = pbkdf2.Key(obj.Plain, obj.Salt, obj.Itr, obj.KeyLen, sha256.New)
	obj.Cipher, err = aes.Encrypt(obj.Plain, obj.Key)
	if err != nil {
		return
	}
	return
}

func (obj *Pbkdf2) Compare() (result bool, err error) {
	obj.Key = pbkdf2.Key(obj.Plain, obj.Salt, obj.Itr, obj.KeyLen, sha256.New)

	cp, err := aes.Encrypt(obj.Plain, obj.Key)
	if err != nil {
		return
	}

	if !bytes.Equal(cp, obj.Cipher) {
		result = false
		return
	}
	result = true
	return
}
