package prismacloud

import (
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"log"
	"math/rand"
	"strings"
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

func PollApiUntilSuccessCustom(p Poller) diag.Diagnostics {
	for {
		if err := p(); err == nil {
			break
		} else if "429" == strings.Split(err.Error(), " ")[0] {
			waitingTime := rand.Intn(5) + 1
			time.Sleep(time.Duration(waitingTime) * time.Second)
		} else {
			return diag.FromErr(err)
		}
	}
	return nil
}

func (b *BackOffRetry) PollApiByBackoffUntilSuccess(p Poller) diag.Diagnostics {
	var delay int
	var retries = 0
	for {
		delay = 1 << retries
		if err := p(); err == nil {
			log.Printf("Error is nil")
			break
		} else if retries <= b.maxRetries && errors.Is(err, pc.ObjectNotFoundError) {
			log.Printf("Re-trying API call with exponential backoff, delay of %d seconds", delay)
			time.Sleep(time.Duration(delay) * time.Second)
			retries++
		} else if retries > b.maxRetries {
			log.Printf("Exhasted all the retries, still encountered error: %v", err)
			return diag.FromErr(err)
		} else {
			log.Printf("Encountered error: %v", err)
			return diag.FromErr(err)
		}
	}
	return nil
}

type BackOffRetry struct {
	maxRetries int
}
