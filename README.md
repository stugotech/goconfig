# goconfig 

An interface to retrieve config in GO.

The interface looks like this:

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

The easiest way of satisfying this interface is to use [viper](https://github.com/spf13/viper).