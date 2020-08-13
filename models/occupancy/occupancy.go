package occupancy

import (
	"Scheduler/models/appointments"
	"time"
)

type Occupancy struct {
	Appointments[] appointments.Appointment
	CurrentOccupancy []int
	SessionStartTime time.Time
	SessionEndTime time.Time
	MinimumAppointmentDuration time.Duration
}

func New(Appointments[] appointments.Appointment, SessionStartTime time.Time, SessionEndTime time.Time, MinimumAppointmentDuration time.Duration) Occupancy {
	occupancy := Occupancy {
		Appointments:               Appointments,
		SessionStartTime:           SessionStartTime,
		SessionEndTime:             SessionEndTime,
		MinimumAppointmentDuration: MinimumAppointmentDuration,
	}
	occupancy.CalculateOccupancy()
	return occupancy
}

func (occupancy* Occupancy) CalculateOccupancy(){
	k := occupancy.MinimumAppointmentDuration

	newT := occupancy.SessionStartTime

	mp := make(map[time.Time]int)
	for i := 0; i < len(occupancy.Appointments); i++ {
		st := occupancy.Appointments[i].StartTime
		et := occupancy.Appointments[i].EndTime
		if stVal, ok := mp[st]; ok {
			mp[st] = stVal + 1
		} else {
			mp[st] = 1
		}
		if etVal, ok:= mp[et]; ok {
			mp[et] = etVal - 1
		} else {
			mp[et] = -1
		}
	}
	cVal := 0
	for newT.Before(occupancy.SessionEndTime) || newT == occupancy.SessionEndTime {
		if val, ok := mp[newT]; ok {
			cVal = cVal + val
		}
		occupancy.CurrentOccupancy = append(occupancy.CurrentOccupancy, cVal)
		newT = newT.Add(time.Minute * k)
	}
}

func (occupancy* Occupancy) CalculateIndexFromTime(indexForTime time.Time) (index int) {
	duration := indexForTime.Sub(occupancy.SessionStartTime)
	index = int(duration.Minutes()) / int(occupancy.MinimumAppointmentDuration)
	return
}

func (occupancy* Occupancy) GetOccupancyAt(Time time.Time) (numOccupants int) {
	numOccupants = occupancy.CurrentOccupancy[occupancy.CalculateIndexFromTime(Time)]
	return
}
