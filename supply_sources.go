package searchads

import (
	"encoding/json"
	"fmt"
)

type SupplySources byte

const (
	APPSTORE_SEARCH_RESULTS SupplySources = iota + 1
	APPSTORE_SEARCH_TAB
)

var (
	_SupplySourcesNameToValue = map[string]SupplySources{
		"APPSTORE_SEARCH_RESULTS": APPSTORE_SEARCH_RESULTS,
		"APPSTORE_SEARCH_TAB":     APPSTORE_SEARCH_TAB,
	}

	_SupplySourcesValueToName = map[SupplySources]string{
		APPSTORE_SEARCH_RESULTS: "APPSTORE_SEARCH_RESULTS",
		APPSTORE_SEARCH_TAB:     "APPSTORE_SEARCH_TAB",
	}
)

func init() {
	var v SupplySources
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_SupplySourcesNameToValue = map[string]SupplySources{
			interface{}(APPSTORE_SEARCH_RESULTS).(fmt.Stringer).String(): APPSTORE_SEARCH_RESULTS,
			interface{}(APPSTORE_SEARCH_TAB).(fmt.Stringer).String():     APPSTORE_SEARCH_TAB,
		}
	}
}

// MarshalJSON is generated so SupplySources satisfies json.Marshaler.
func (r SupplySources) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _SupplySourcesValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid SupplySources: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so SupplySources satisfies json.Unmarshaler.
func (r *SupplySources) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("SupplySources should be a string, got %s", data)
	}
	v, ok := _SupplySourcesNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid SupplySources %q", s)
	}
	*r = v
	return nil
}

// ParseSupplySources to turn a String into the SupplySources
func ParseSupplySources(name string) (SupplySources, error) {
	v, ok := _SupplySourcesNameToValue[name]
	if ok {
		return v, nil
	}
	return SupplySources(0), fmt.Errorf("invalid SupplySources: %s", name)
}

// String to return the String of a SupplySources
func (r SupplySources) String() (string, error) {
	s, ok := _SupplySourcesValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid SupplySources: %d", r)
	}
	return s, nil
}
