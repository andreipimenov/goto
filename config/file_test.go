package config

import (
	"fmt"
	"testing"
)

func TestFileConfigGetJSON(t *testing.T) {
	tests := []struct {
		File          string
		ExpectedError bool
	}{
		{"invalid.json", true},
		{"valid.json", false},
	}

	cfg := &struct {
		status string
	}{}

	for _, test := range tests {
		driver := NewFileConfig(fmt.Sprintf("testdata/%s", test.File))
		err := driver.GetJSON(cfg)
		if (err == nil && test.ExpectedError) || (err != nil && !test.ExpectedError) {
			t.Errorf("Expected error: %t, received: %v", test.ExpectedError, err)
		}
	}
}
