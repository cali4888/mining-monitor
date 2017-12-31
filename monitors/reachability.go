package monitors

import (
	"errors"
	"log"
	"net/http"
	"time"
)

var (
	defaultPingUrl = "https://google.com"
	defaultTimeout = 5 * time.Second

	maxRetries         = 10
	waitBeforeRetry    = 5 * time.Second
	waitBeforeRequests = 2 * time.Second
)

type Reachability struct {
	client http.Client
}

func (hm *Reachability) Start(c chan error, stop chan bool) {
	retries := 0
	for {
		log.Println("Ping")

		select {
		case <-stop:
			c <- errors.New("ReachabilityMonitor was stopped")
			return
		default:
		}

		err := hm.ping(defaultPingUrl, defaultTimeout)
		if err != nil {
			log.Printf("[ERROR] Reachability err: %s\n", err)
			if retries >= maxRetries {
				c <- err
				return
			}

			retries = retries + 1
			time.Sleep(time.Duration(retries) * waitBeforeRetry)
		} else {
			retries = 0
			time.Sleep(waitBeforeRequests)
		}
	}

	return
}

func (hm *Reachability) ping(dest string, timeout time.Duration) error {
	if timeout == 0 {
		hm.client.Timeout = defaultTimeout
	} else {
		hm.client.Timeout = timeout
	}

	_, err := hm.client.Get(dest)
	if err != nil {
		return err
	}

	return nil
}
