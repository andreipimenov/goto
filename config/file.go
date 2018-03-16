package config

import (
	"io/ioutil"
	"os"
)

//FileConfig contains specific configuration filename.
type FileConfig struct {
	file string
}

//NewFileConfig returns *FileConfig.
func NewFileConfig(file string) *FileConfig {
	return &FileConfig{
		file: file,
	}
}

//Get reads config file and return RawConfig.
func (c *FileConfig) Get() (*Raw, error) {
	_, err := os.Stat(c.file)
	if err != nil {
		return nil, err
	}
	r, err := ioutil.ReadFile(c.file)
	if err != nil {
		return nil, err
	}
	return &Raw{r}, nil
}

//GetJSON reads config file, unmarshal to value pointed to by v.
func (c *FileConfig) GetJSON(v interface{}) error {
	r, err := c.Get()
	if err != nil {
		return err
	}
	return r.UnmarshalJSON(v)
}
