package goconfig

import "time"

// Config is the interface used by subpackages to retrieve config values.
// It is not coincidental that this looks like github.com/spf13/viper.
type Config interface {
	Get(key string) interface{}
	GetBool(key string) bool
	GetInt(key string) int
	GetFloat64(key string) float64
	GetString(key string) string
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
	GetStringSlice(key string) []string
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
	IsSet(key string) bool
}
