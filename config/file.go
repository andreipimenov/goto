package config

import (
	"io/ioutil"
	"os"
)

//File contains specific configuration filename.
type File struct {
	file string
}

//NewFile returns *File.
func NewFile(file string) *File {
	return &File{
		file: file,
	}
}

//Get reads config file and return RawConfig.
func (c *File) Get() (*Raw, error) {
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
func (c *File) GetJSON(v interface{}) error {
	r, err := c.Get()
	if err != nil {
		return err
	}
	return r.UnmarshalJSON(v)
}
