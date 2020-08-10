package datetime

import "time"

func Get(df, tf, location string) (string, error) {
	cd := time.Now()
	tz, err := time.LoadLocation(location)
	if err != nil {
		return "", err
	}
	switch df + tf {
	case "YY-MM-DD" + "SS:MM:HH":
		return cd.In(tz).Format("2006-01-02 03:04:05PM"), nil
	case "DD-MM-YY" + "HH:MM:SS":
		return cd.In(tz).Format("02-01-2006 03:04:05PM"), nil
	case "DD/MM/YY" + "MM:HH:SS":
		return cd.In(tz).Format("02/01/2006 04:03:05PM"), nil
	case "YY/MM/DD" + "HH:MM:SS":
		return cd.In(tz).Format("2006/01/02 15:04:05"), nil
	case "YY/DD/MM" + "HH:SS:MM":
		return cd.In(tz).Format("2006/02/01 15:04:05"), nil
	case "DD/MM/YYYY" + "":
		return cd.In(tz).Format("02/01/2006"), nil
	case "MM/DD/YYYY" + "":
		return cd.In(tz).Format("01/02/2006"), nil
	case "YY/MM/DD" + "":
		return cd.In(tz).Format("2006/01/02"), nil
	case "MM/DD/YY" + "":
		return cd.In(tz).Format("01/02/2006"), nil
	case "YYYY-MM-DD" + "":
		return cd.In(tz).Format("2006-01-02"), nil
	case "YYYYMMDD" + "HHMMSS":
		return cd.In(tz).Format("20060102150405"), nil
	case "YYYYMMDD" + "":
		return cd.In(tz).Format("20060102"), nil
	case "CYYMMDD" + "":
		return "1" + cd.In(tz).Format("060102"), nil
	case "DDMMYYYY" + "":
		return cd.In(tz).Format("02012006"), nil
		//For Telecash
	case "MMDDYY" + "HHMMSS":
		return cd.In(tz).Format("010206150405"), nil
	case "DD.MM.YYYY" + "HH:MM:SS":
		return cd.In(tz).Format("02.01.2006 15:04:05"), nil
	default:
		return cd.In(tz).Format("2006-01-02T15:04:05"), nil
	}
}
