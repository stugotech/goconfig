package goconfig

import (
	"time"

	"github.com/spf13/viper"
)

// Config is the interface used by subpackages to retrieve config values.
// It is not coincidental that this looks like github.com/spf13/viper.
type Config interface {
	GetBool(key string) bool
	GetInt(key string) int
	GetFloat64(key string) float64
	GetString(key string) string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	IsSet(key string) bool
}

type vpr struct{}

// Viper creates a config provider from github.com/spf13/viper
func Viper() Config {
	return &vpr{}
}

func (vpr) GetBool(key string) bool {
	return viper.GetBool(key)
}
func (vpr) GetInt(key string) int {
	return viper.GetInt(key)
}
func (vpr) GetFloat64(key string) float64 {
	return viper.GetFloat64(key)
}
func (vpr) GetString(key string) string {
	return viper.GetString(key)
}
func (vpr) GetStringSlice(key string) []string {
	return viper.GetStringSlice(key)
}
func (vpr) GetTime(key string) time.Time {
	return viper.GetTime(key)
}
func (vpr) GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}
func (vpr) IsSet(key string) bool {
	return viper.IsSet(key)
}
