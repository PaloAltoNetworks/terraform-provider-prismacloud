package prismacloud

import (
	"errors"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	pc "github.com/paloaltonetworks/prisma-cloud-go"
	"log"
	"math/rand"
	"strconv"
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

// isRetryableError checks if an error is retryable (429 Too Many Requests or 5xx Server Errors).
// Handles multiple SDK error formats:
//   - "503 error, and could not unmarshal header..." (unmarshal failure)
//   - "503 error without the \"X-Redlock-Status\" header..." (missing header)
//   - "503/path Error(msg:... severity:...)" (PrismaCloudErrorList format)
//   - "429", "500", etc. (plain status codes)
func isRetryableError(err error) bool {
	if err == nil {
		return false
	}
	errMsg := err.Error()

	// Extract the status code from the error message.
	// The first token may be "503" or "503/some/path", so split on "/" first
	// to handle the PrismaCloudErrorList format (e.g., "503/compliance/...").
	firstToken := strings.Split(errMsg, " ")[0]
	code := strings.Split(firstToken, "/")[0]

	// Validate that code is a number to avoid false positives
	if _, parseErr := strconv.Atoi(code); parseErr != nil {
		return false
	}

	switch code {
	case "429", "500", "502", "503", "504":
		return true
	}
	return false
}

// RetryWithBackoff retries a Poller function using the provider's configured
// retry settings (max_retries, retry_max_delay, retry_type) for retryable
// HTTP errors such as 429 (Too Many Requests) and 5xx (Server Errors).
func RetryWithBackoff(client *pc.Client, p Poller) diag.Diagnostics {
	maxRetries := client.MaxRetries
	retryMaxDelay := client.RetryMaxDelay
	retryType := client.RetryType

	if maxRetries <= 0 {
		maxRetries = 5
	}
	if retryMaxDelay <= 0 {
		retryMaxDelay = 30
	}
	if retryType == "" {
		retryType = "exponential_backoff"
	}

	var retries int
	for {
		err := p()
		if err == nil {
			return nil
		}

		if !isRetryableError(err) {
			return diag.FromErr(err)
		}

		retries++
		if retries > maxRetries {
			log.Printf("[ERROR] Exhausted all %d retries, last error: %v", maxRetries, err)
			return diag.FromErr(err)
		}

		var delay int
		if retryType == "exponential_backoff" {
			delay = 1 << retries
		} else {
			delay = 1 + retries
		}
		if delay > retryMaxDelay {
			delay = retryMaxDelay
		}

		log.Printf("[WARN] Retryable error encountered (attempt %d/%d), retrying after %d seconds: %v", retries, maxRetries, delay, err)
		time.Sleep(time.Duration(delay) * time.Second)
	}
}

func (b *BackOffRetry) PollApiByBackoffUntilSuccess(p Poller) diag.Diagnostics {
	var delay int
	var retries = 0
	for {
		delay = 1 << retries
		if err := p(); err == nil {
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
