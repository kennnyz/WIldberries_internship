package main

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func GetTheCurrentTime() string {
	NtpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return NtpTime.String()
}
