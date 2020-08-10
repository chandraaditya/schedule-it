package user

import (
	"Scheduler/models/db"
	"Scheduler/models/hashing/encoding/base64"
	"Scheduler/models/hashing/util/pbkdf2"
	"crypto/rand"
	"errors"
	"log"
)

type User struct {
	ID        int
	Email     string
	Username  string
	PlainPass string
	HashPass  string
	HouseNo   string
}

func (user User) Authenticate() (err error) {
	err = db.Err
	if err != nil {
		log.Println(err.Error())
		return
	}
	stmt, err := db.Conn.Prepare(`SELECT id, username, password FROM users WHERE email=?`)
	if err != nil {
		log.Println("el1002: error connecting to db:", err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(user.Email).Scan(&user.ID, &user.Username, &user.HashPass)
	if err.Error() == "sql: no rows in result set" {
		err = errors.New("el1004")
		return
	}
	if err != nil {
		log.Println("el1003: error querying from db:", err)
		return
	}
	err = Verify(user.PlainPass, user.HashPass)
	if err != nil {
		err = errors.New("ea1001: error authenticating user")
		return
	}
	return
}

func (user User) CreateNewUser() (err error) {
	return
}

func Verify(Plain, Hash string) (err error) {
	//TODO need to finish hash function
	cp, err1 := base64.Decode([]byte(Hash))
	if err1 != nil {
		log.Println(err.Error())
		return
	}
	var pbkdf pbkdf2.Pbkdf2
	pbkdf.Itr = 32
	pbkdf.KeyLen = 32
	pbkdf.Plain = []byte(Plain)
	pbkdf.Salt = cp[:32]
	pbkdf.Cipher = cp[32:]
	result, err := pbkdf.Compare()
	if err != nil {
		log.Println(err.Error())
		return
	}
	if !result {
		err = errors.New("unable to verify password")
		log.Println(err.Error())
		return
	}
	return
}

func Hash(Plain string) (Hash string, err error) {
	b := make([]byte, 32)
	_, err = rand.Read(b)
	var pbkdf pbkdf2.Pbkdf2
	pbkdf.Itr = 32
	pbkdf.KeyLen = 32
	pbkdf.Plain = []byte(Plain)
	pbkdf.Salt = b
	err = pbkdf.Encrypt()
	if err != nil {
		err = errors.New("password encryption failed")
		return
	}
	var tmp []byte
	tmp = append(tmp, pbkdf.Salt...)
	tmp = append(tmp, pbkdf.Cipher...)
	out, err1 := base64.Encode(tmp)
	if err1 != nil {
		err = errors.New("password encryption failed")
		return
	}
	Hash = string(out)
	return
}
