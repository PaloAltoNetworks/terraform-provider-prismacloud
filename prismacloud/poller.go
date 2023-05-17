package prismacloud

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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

func PollApiUntilSuccessCustom(p Poller, errorCode string) diag.Diagnostics {
	for {
		if err := p(); err == nil {
			break
		} else if errorCode == strings.Split(err.Error(), " ")[0] {
			waitingTime := rand.Intn(5) + 1
			time.Sleep(time.Duration(waitingTime) * time.Second)
		} else {
			return diag.FromErr(err)
		}
	}
	return nil
}
