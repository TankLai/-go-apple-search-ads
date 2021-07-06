package searchads

import (
	"encoding/json"
	"fmt"
)

type AdChannelType byte

const (
	SEARCH AdChannelType = iota
	DISPLAY
)

var (
	_AdChannelTypeNameToValue = map[string]AdChannelType{
		"SEARCH":  SEARCH,
		"DISPLAY": DISPLAY,
	}

	_AdChannelTypeValueToName = map[AdChannelType]string{
		SEARCH:  "SEARCH",
		DISPLAY: "DISPLAY",
	}
)

func init() {
	var v AdChannelType
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_AdChannelTypeNameToValue = map[string]AdChannelType{
			interface{}(SEARCH).(fmt.Stringer).String():  SEARCH,
			interface{}(DISPLAY).(fmt.Stringer).String(): DISPLAY,
		}
	}
}

// MarshalJSON is generated so AdChannelType satisfies json.Marshaler.
func (r AdChannelType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _AdChannelTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid AdChannelType: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so AdChannelType satisfies json.Unmarshaler.
func (r *AdChannelType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("AdChannelType should be a string, got %s", data)
	}
	v, ok := _AdChannelTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid AdChannelType %q", s)
	}
	*r = v
	return nil
}

// ParseAdChannelType to turn a String into the AdChannelType
func ParseAdChannelType(name string) (AdChannelType, error) {
	v, ok := _AdChannelTypeNameToValue[name]
	if ok {
		return v, nil
	}
	return AdChannelType(0), fmt.Errorf("invalid AdChannelType: %s", name)
}

// String to return the String of a AdChannelType
func (r AdChannelType) String() (string, error) {
	s, ok := _AdChannelTypeValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid AdChannelType: %d", r)
	}
	return s, nil
}
