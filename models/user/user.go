package user

import (
	"Scheduler/models/db"
	"Scheduler/models/hashing/encoding/base64"
	"Scheduler/models/hashing/util/pbkdf2"
	"crypto/rand"
	"errors"
	"log"
	mathRand "math/rand"
	"time"
)

type User struct {
	ID        int
	Email     string
	Name      string
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
	stmt, err := db.Conn.Prepare(`SELECT id, name, password FROM users WHERE email=?`)
	if err != nil {
		log.Println("el1002: error connecting to db:", err)
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(user.Email).Scan(&user.ID, &user.Name, &user.HashPass)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			err = errors.New("el1004")
			return
		}
	}
	if err != nil {
		log.Println("el1003: error querying from db:", err)
		return
	}
	err = Verify(user.PlainPass, user.HashPass)
	user.PlainPass = ""
	if err != nil {
		err = errors.New("ea1001: error authenticating user")
		return
	}
	return
}

func (user User) CreateNewUser() (err error) {
	err = db.Err
	if err != nil {
		log.Println(err.Error())
		return
	}
	stmt, err := db.Conn.Prepare(`INSERT INTO users(id, email, name, password, houseno) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		log.Println(err.Error())
		return
	}
	user.ID, err = GenerateUserID()
	if err != nil {
		if err.Error() != "eu1002" {
			log.Println(err.Error())
			return
		}
	}
	user.HashPass, err = Hash(user.PlainPass)
	user.PlainPass = ""
	if err != nil {
		log.Println(err.Error())
		return
	}
	res, err := stmt.Exec(user.ID, user.Email, user.Name, user.HashPass, user.HouseNo)
	if err != nil {
		log.Println(err.Error())
		if err.Error()[:10] == "Error 1062" {
			err = errors.New("ea1001")
			log.Println(err.Error())
		}
		return
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	return
}

func CheckGeneratedIDExits(ID int) (err error) {
	var ExistingID int
	stmt, err := db.Conn.Prepare(`SELECT id FROM users WHERE id=?`)
	if err != nil {
		err = errors.New("eu1001")
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(ID).Scan(&ExistingID)
	if err != nil {
		err = errors.New("eu1002")
		return
	}
	err = errors.New("eu1003")
	return
}

func GenerateUserID() (ID int, err error) {
	min := 100000000
	max := 999999999
	mathRand.Seed(time.Now().UnixNano())
	for {
		ID = mathRand.Intn(max-min+1) + min
		err = CheckGeneratedIDExits(ID)
		if err.Error() == "eu1001" {
			return
		}
		if err.Error() == "eu1002" {
			log.Println("Generated user ID:", ID)
			break
		}
	}
	return
}

func Verify(Plain, Hash string) (err error) {
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

func FindTimeSlots(appointmentDuration time.Duration, minimumDuration time.Duration, currentOccupancy[] int, maximumOccupancy int) (possibleTimes[] time.Time){
	var count int
	var divisions int = int(appointmentDuration) / int(minimumDuration)
	Time, _ := time.Parse(time.Kitchen, "7:00AM")
	for i := 0; i < len(currentOccupancy) - divisions; i++ {
		count = 0
		for j := i; j < i + divisions; j++ {
			if currentOccupancy[j] >= maximumOccupancy {
				break
			}
			if currentOccupancy[j] < maximumOccupancy {
				count++
			}
			if count == divisions{
				x := time.Minute * (minimumDuration * time.Duration(i))
				possibleTimes = append(possibleTimes, Time.Add(x))
			}
		}
	}
	return
}