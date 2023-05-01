package cron

import (
	"log"
	"testing"
	"time"
)

// Many tests schedule a job for every second, and then wait at most a second
// for it to run.  This amount is just slightly larger than 1 second to
// compensate for a few milliseconds of runtime.
const OneSecond = 1*time.Second + 50*time.Millisecond

func TestOnce(t *testing.T) {
	cron := New()
	log.Println("start...")
	_, err := cron.AddFunc("@delay 15s", func(e *Entry) {
		log.Printf("delay:xxxxx")
	})
	if err != nil {
		panic(err)
	}
	_, err = cron.AddFunc("@every 15s", func(e *Entry) {
		log.Printf("every:xxxxx")
		cron.AddFunc("@delay 10s", func(e *Entry) {
			log.Printf("delay:xxxxx")
		})
	})
	if err != nil {
		panic(err)
	}
	cron.Start()
	time.Sleep(1 * time.Minute)
	for {
		select {
		case <-cron.Stop().Done():
			break
		}
	}

}

func TestDelayIfStillRunningInEntry(t *testing.T) {
	cron := New(WithChain(DelayIfStillRunningInEntry(DefaultLogger)))
	log.Println("start...")
	_, err := cron.AddFunc("@every 10s", func(e *Entry) {
		log.Printf("1开始执行")
		time.Sleep(1 * time.Minute)
		log.Printf("1执行结束")
	})
	_, err = cron.AddFunc("@every 5s", func(e *Entry) {
		log.Printf("2开始执行")
		time.Sleep(1 * time.Minute)
		log.Printf("2执行结束")
	})
	if err != nil {
		panic(err)
	}
	cron.Start()
	for {

	}
}
