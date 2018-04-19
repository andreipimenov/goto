// Package config implements configuration.
// This can be useful when your application needs to be configured with some data.
// Currently supported only json local file for configurtion.
package config

import (
	"encoding/json"
	"errors"
)

// Configer interface describes method to get configuration data of any type.
type Configer interface {
	Get() (*Raw, error)
}

// Raw contains raw configuration data received from any source.
type Raw struct {
	Data interface{}
}

// JSON unmarshalls Raw.Data as JSON and store it in value pointed to by v.
func (r *Raw) JSON(v interface{}) error {
	b, ok := r.Data.([]byte)
	if !ok {
		return errors.New("error assertion Raw.Data to []byte")
	}
	err := json.Unmarshal(b, v)
	if err != nil {
		return err
	}
	return nil
}
