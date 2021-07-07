package searchads

import (
	"encoding/json"
	"fmt"
)

// PricingModel type to represent enum of PricingModel (CPC/CPM)
type PricingModel byte

// CPC and CPM enum values
const (
	CPC PricingModel = iota + 1
	CPM
)

var (
	_PricingModelNameToValue = map[string]PricingModel{
		"CPC": CPC,
		"CPM": CPM,
	}

	_PricingModelValueToName = map[PricingModel]string{
		CPC: "CPC",
		CPM: "CPM",
	}
)

func init() {
	var v PricingModel
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_PricingModelNameToValue = map[string]PricingModel{
			interface{}(CPC).(fmt.Stringer).String(): CPC,
			interface{}(CPM).(fmt.Stringer).String(): CPM,
		}
	}
}

// MarshalJSON is generated so PricingModel satisfies json.Marshaler.
func (r PricingModel) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _PricingModelValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid PricingModel: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so PricingModel satisfies json.Unmarshaler.
func (r *PricingModel) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("PricingModel should be a string, got %s", data)
	}
	v, ok := _PricingModelNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid PricingModel %q", s)
	}
	*r = v
	return nil
}

// ParsePricingModel to turn a String into the PricingModel
func ParsePricingModel(name string) (PricingModel, error) {
	v, ok := _PricingModelNameToValue[name]
	if ok {
		return v, nil
	}
	return PricingModel(0), fmt.Errorf("invalid PricingModel: %s", name)
}

// String to return the String of a PricingModel
func (r PricingModel) String() (string, error) {
	s, ok := _PricingModelValueToName[r]
	if !ok {
		return "", fmt.Errorf("invalid PricingModel: %d", r)
	}
	return s, nil
}
