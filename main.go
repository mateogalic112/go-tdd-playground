package main

import (
	"os"
	"time"

	"example.com/hello/mocking"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func main() {
	sleeper := &DefaultSleeper{}
	mocking.Countdown(os.Stdout, sleeper)

}
