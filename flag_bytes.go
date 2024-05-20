package cli

import (
	"fmt"
)

type BytesFlag = FlagBase[[]byte, BytesConfig, bytesValue]

// BytesConfig defines the configuration for bytes flags
type BytesConfig struct {
}

// -- bytes Value
type bytesValue struct {
	destination *[]byte
}

// Below functions are to satisfy the ValueCreator interface

func (i bytesValue) Create(val []byte, p *[]byte, c BytesConfig) Value {
	*p = val
	return &bytesValue{
		destination: p,
	}
}

func (i bytesValue) ToString(b []byte) string {
	// if b == "" {
	// 	return b
	// }
	return fmt.Sprintf("%q", b)
}

// Below functions are to satisfy the flag.Value interface

var _ Value = &bytesValue{}

func (s *bytesValue) Set(val string) error {
	*s.destination = []byte(val)
	return nil
}

func (s *bytesValue) Get() any { return *s.destination }

func (s *bytesValue) String() string {
	if s.destination != nil {
		return string(*s.destination)
	}
	return ""
}

// func (cmd *Command) String(name string) string {
// 	if v, ok := cmd.Value(name).(string); ok {
// 		tracef("string available for flag name %[1]q with value=%[2]v (cmd=%[3]q)", name, v, cmd.Name)
// 		return v
// 	}

// 	tracef("string NOT available for flag name %[1]q (cmd=%[2]q)", name, cmd.Name)
// 	return ""
// }
