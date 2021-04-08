package prismacloud

import (
	"time"
)

type Poller func() error

// PollApiUntilSuccess is a function to wait until an API call that should
// succeed actually does.
func PollApiUntilSuccess(p Poller) {
	for {
		if err := p(); err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
