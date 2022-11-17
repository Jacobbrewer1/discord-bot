package config

import "testing"

func TestReadConfig(t *testing.T) {
	var tests = []struct {
		name         string
		input        string
		isSuccessful bool
	}{
		{"success", "./../../config/config.json", true},
		{"success", "./config/config.json", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ReadConfig(tt.input)

			if tt.isSuccessful != (err == nil) {
				t.Errorf("ReadConfig(%s) = %t, expected %t", tt.input, err == nil, tt.isSuccessful)
			}
		})
	}
}
