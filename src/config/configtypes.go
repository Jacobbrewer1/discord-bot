package config

var Cfg *Config

type (
	Config struct {
		ApplicationId string `json:"applicationId,omitempty"`
		Token         string `json:"token,omitempty"`
	}
)
