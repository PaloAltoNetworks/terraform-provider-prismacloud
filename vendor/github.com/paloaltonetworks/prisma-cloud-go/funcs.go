package prismacloud

import (
	"regexp"
)

// scrubSensitiveData removes sensitive stuff from send/receive logging.
func scrubSensitiveData(b []byte) string {
	s := string(b)

	for _, val := range SensitiveKeys {
		hdr := `"` + val + `":`
		pat := regexp.MustCompile(hdr + `".*?"`)
		s = pat.ReplaceAllString(s, hdr+`"********"`)
	}

	return s
}
