package main

import (
	"os"
	"time"

	"example.com/hello/mocking"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{Duration: 1 * time.Second, SleepFunc: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)

}
