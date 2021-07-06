package searchads

import (
	"encoding/json"
	"fmt"
)

type BillingEvent byte

const (
	TAPS BillingEvent = iota
	IMPRESSIONS
)

var (
	_BillingEventNameToValue = map[string]BillingEvent{
		"TAPS":        TAPS,
		"IMPRESSIONS": IMPRESSIONS,
	}

	_BillingEventValueToName = map[BillingEvent]string{
		TAPS:        "TAPS",
		IMPRESSIONS: "IMPRESSIONS",
	}
)

func init() {
	var v BillingEvent
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_BillingEventNameToValue = map[string]BillingEvent{
			interface{}(TAPS).(fmt.Stringer).String():        TAPS,
			interface{}(IMPRESSIONS).(fmt.Stringer).String(): IMPRESSIONS,
		}
	}
}

// MarshalJSON is generated so BillingEvent satisfies json.Marshaler.
func (r BillingEvent) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _BillingEventValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid BillingEvent: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so BillingEvent satisfies json.Unmarshaler.
func (r *BillingEvent) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("BillingEvent should be a string, got %s", data)
	}
	v, ok := _BillingEventNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid BillingEvent %q", s)
	}
	*r = v
	return nil
}

// ParseBillingEvent to turn a String into the BillingEvent
func ParseBillingEvent(name string) (BillingEvent, error) {
	v, ok := _BillingEventNameToValue[name]
	if ok {
		return v, nil
	}
	return BillingEvent(0), fmt.Errorf("invalid BillingEvent: %s", name)
}

// String to return the String of a BillingEvent
func (r BillingEvent) String() (string, error) {
	s, ok := _BillingEventValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid BillingEvent: %d", r)
	}
	return s, nil
}
