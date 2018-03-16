package config

import (
	"encoding/json"
	"errors"
)

//Configer interface.
type Configer interface {
	Get() (*Raw, error)
}

//Raw contains raw configuration data from any source.
type Raw struct {
	Data interface{}
}

//UnmarshalJSON unmarshalls RawConfig.Data as JSON and store it in value pointed to by v.
func (r *Raw) UnmarshalJSON(v interface{}) error {
	b, ok := r.Data.([]byte)
	if !ok {
		return errors.New("error assertion RawConfig.Data to []byte")
	}
	err := json.Unmarshal(b, v)
	if err != nil {
		return err
	}
	return nil
}
