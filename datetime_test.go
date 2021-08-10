package phpgo

import (
	"log"
	"testing"
	"time"
)

func TestStrToTime(t *testing.T) {
	format := "1 year +5 month -1 day +10 hours -23 minutes +5 seconds"
	res, err := StrToTime(format)
	log.Println(format, time.Now(), res, err)
	format = "+1 year +2 month"
	res, err = StrToTime(format)
	log.Println(format, time.Now(), res, err)
	format = "+1 year + 2 month"
	res, err = StrToTime(format)
	log.Println(format, time.Now(), res, err)
	format = "+1 years +2 months -10 days"
	res, err = StrToTime(format)
	log.Println(format, time.Now(), res, err)
	format = "+1 years +2 months -10 days +5 days"
	res, err = StrToTime(format)
	log.Println(format, time.Now(), res, err)
	format = "+1 years +2 months -10 days +5 days +5 hours -30 minute +10 second"
	res, err = StrToTime(format)
	log.Println(format, time.Now(), res, err)
	format = "-1 years +2 months"
	res, err = StrToTime(format, time.Now().AddDate(-1, 0, 0).Unix())
	log.Println(format, time.Now(), res, err)
}

func TestDate(t *testing.T) {
	format := "Y-m-d H:i:s"
	log.Printf("format:%s; result:%s", format, Date(format))
	format = "y-M-d h:i:s"
	log.Printf("format:%s; result:%s", format, Date(format))
	format = "y-M-d h:i:s.u P O e"
	log.Printf("format:%s; result:%s", format, Date(format))
	format = "Y-y-m-M-F-d-D-j-H-h-g-i-su-a-A-e-O-P"
	log.Printf("format:%s; result:%s", format, Date(format))
}
