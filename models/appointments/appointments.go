package appointments

import (
	"Scheduler/models/db"
	"errors"
	"fmt"
	"log"
	mathRand "math/rand"
	"time"
)

type Appointment struct {
	ID        int
	UserID    int
	StartTime time.Time
	EndTime   time.Time
	Date      time.Time
	Active    bool
}

var Appointments []Appointment

func GetAppointments(date time.Time) (err error) {
	fmt.Println("Fuck!")
	rows, err := db.Conn.Query(`SELECT id, userid, starttime, endtime, date, active FROM appointments WHERE date = ?`, string(date.Year())+"-"+string(date.Month())+"-"+string(date.Day()))
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer rows.Close()
	for rows.Next() {
		fmt.Println("rows")
		var appointment Appointment
		err = rows.Scan(&appointment.ID, &appointment.UserID, &appointment.StartTime, &appointment.EndTime, &appointment.Date, &appointment.Active)
		if err != nil {
			log.Println(err.Error())
			return
		}
		fmt.Println(appointment.ID, appointment.StartTime, appointment.Date, appointment.Active)
	}
	if err = rows.Err(); err != nil {
		log.Println(err.Error())
		return
	}
	return
}

func CheckGeneratedIDExits(ID int) (err error) {
	var ExistingID int
	stmt, err := db.Conn.Prepare(`SELECT id FROM appointments WHERE id=?`)
	if err != nil {
		err = errors.New("ea1001")
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(ID).Scan(&ExistingID)
	if err != nil {
		err = errors.New("ea1002")
		return
	}
	err = errors.New("ea1003")
	return
}

func GenerateAppointmentID() (ID int, err error) {
	min := 100000000
	max := 999999999
	mathRand.Seed(time.Now().UnixNano())
	for {
		ID = mathRand.Intn(max-min+1) + min
		err = CheckGeneratedIDExits(ID)
		if err.Error() == "ea1001" {
			return
		}
		if err.Error() == "ea1002" {
			log.Println("Generated user ID:", ID)
			break
		}
	}
	return
}
