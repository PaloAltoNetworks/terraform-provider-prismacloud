package timerange

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// TimeRange is the time range model for Prisma Cloud.
type TimeRange struct {
	Type             string      `json:"type"`
	Value            interface{} `json:"value"`
	RelativeTimeType string      `json:"relativeTimeType,omitempty"`
}

// SetType sets the time range type based on the value specified.
func (o *TimeRange) SetType() error {
	switch v := o.Value.(type) {
	case Absolute:
		o.Type = TypeAbsolute
	case Relative:
		o.Type = TypeRelative
	case string:
		o.Type = TypeToNow
	case nil:
		return fmt.Errorf("time range must be specified")
	default:
		return fmt.Errorf("invalid time range type: %v", v)
	}

	return nil
}

// SetValue sets the Value based on what Type is specified.
func (o *TimeRange) SetValue() error {
	var body []byte
	var err error

	if body, err = json.Marshal(o.Value); err != nil {
		return err
	}

	switch o.Type {
	case TypeAbsolute:
		v := Absolute{}
		if err = json.Unmarshal(body, &v); err != nil {
			return err
		}
		o.Value = v
	case TypeRelative:
		v := Relative{}
		if err = json.Unmarshal(body, &v); err != nil {
			return err
		}
		o.Value = v
	case TypeToNow:
		v := ToNow{}
		v.Unit, _ = strconv.Unquote(string(body))
		o.Value = v
	}

	return nil
}

type Absolute struct {
	Start int `json:"startTime"`
	End   int `json:"endTime"`
}

type Relative struct {
	Amount int    `json:"amount"`
	Unit   string `json:"unit"`
}

type ToNow struct {
	Unit string `json:"unit"`
}
