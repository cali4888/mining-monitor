package strategies

import (
	"github.com/cali4888/mining-monitor"
	"github.com/stianeikeland/go-rpio"
	"log"
	"time"
)

var (
	pin = rpio.Pin(4)

	waitTimeAfterRecovery = 3 * time.Minute

	waitTimeBeforeTurnOn = 30 * time.Second
)

type RouterRecovery struct{}

func (rrs *RouterRecovery) Run(m monitor.Monitor) {
	for {
		rrs.waitMonitorFail(m)
		err := rrs.recovery()
		if err != nil {
			log.Printf("[ERROR] RouterRecovery recovery err: %s\n", err)
		}
		time.Sleep(waitTimeAfterRecovery)
	}
}

func (rrs *RouterRecovery) waitMonitorFail(m monitor.Monitor) error {
	monitorErrC := make(chan error, 1)
	monitorStorC := make(chan bool, 1)

	go m.Start(monitorErrC, monitorStorC)

	return <-monitorErrC
}

func (rrs *RouterRecovery) recovery() error {
	log.Println("[INFO] RouterRecovery recovery started")

	if err := rpio.Open(); err != nil {
		return err
	}
	defer rpio.Close()
	pin.Output()

	pin.Low()

	time.Sleep(waitTimeBeforeTurnOn)

	pin.High()

	return nil
}
