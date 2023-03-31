package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gen2brain/beeep"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("usage, time in minutes: command [sit] [stand] ")
		return
	}

	sitTime, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	standTime, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		fmt.Printf("%v : sit down\n", time.Now())
		beeep.Alert("SIT DOWN", "You can sit down now :)", "")
		time.Sleep(time.Duration(sitTime) * time.Minute)
		fmt.Printf("%v : stand up\n", time.Now())
		beeep.Alert("STAND UP", "Please stand up", "")
		time.Sleep(time.Duration(standTime) * time.Minute)
	}
}
