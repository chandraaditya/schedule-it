package base64

import (
	"Scheduler/models/hashing/util/log"
	"encoding/base64"
	"errors"
)

func Encode(in []byte) (out []byte, err error) {
	tmp := base64.StdEncoding.EncodeToString(in)
	out = []byte(tmp)
	return
}

func Decode(in []byte) (out []byte, err error) {

	out, err = base64.StdEncoding.DecodeString(string(in))
	if err != nil {
		log.Println("Error", "Error", err)
		err = errors.New("base64 decoding fail")
		return
	}
	return
}
