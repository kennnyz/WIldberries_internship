package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

func GetTheCurrentTime() string {
	NtpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return NtpTime.String()
}

func main() {
	timeNTP := GetTheCurrentTime()
	time.Sleep(time.Second * 2)
	now := time.Now().String()

	timeNTP = timeNTP[0:16]
	now = now[0:16]

	fmt.Println(now)
	fmt.Println(timeNTP)
}
