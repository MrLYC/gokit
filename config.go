package gokit

import (
	"time"

	"github.com/MrLYC/gokit/builtins"
)

//
var (
	Version = ""
	Mode    = "debug"
)

// Configuration : viper interface
type Configuration interface {
	// AllSettings merges all settings and returns them as a map[string]interface{}.
	AllSettings() map[string]interface{}

	// SetDefault sets the default value for this key.
	// SetDefault is case-insensitive for a key.
	// Default only used when no value is provided by the user via flag, config or ENV.
	SetDefault(key string, value interface{})

	// IsSet checks to see if the key has been set in any of the data locations.
	// IsSet is case-insensitive for a key.
	IsSet(key string) bool

	// Set sets the value for the key in the override regiser.
	// Set is case-insensitive for a key.
	Set(key string, value interface{})

	// Get can retrieve any value given the key to use.
	// Get is case-insensitive for a key.
	Get(key string) interface{}

	// GetString returns the value associated with the key as a string.
	GetString(key string) string

	// GetBool returns the value associated with the key as a boolean.
	GetBool(key string) bool

	// GetInt returns the value associated with the key as an integer.
	GetInt(key string) int

	// GetInt32 returns the value associated with the key as an integer.
	GetInt32(key string) int32

	// GetInt64 returns the value associated with the key as an integer.
	GetInt64(key string) int64

	// GetFloat64 returns the value associated with the key as a float64.
	GetFloat64(key string) float64

	// GetTime returns the value associated with the key as time.
	GetTime(key string) time.Time

	// GetDuration returns the value associated with the key as a duration.
	GetDuration(key string) time.Duration

	// GetStringSlice returns the value associated with the key as a slice of strings.
	GetStringSlice(key string) []string

	// GetStringMap returns the value associated with the key as a map of interfaces.
	GetStringMap(key string) map[string]interface{}

	// GetStringMapString returns the value associated with the key as a map of strings.
	GetStringMapString(key string) map[string]string

	// GetStringMapStringSlice returns the value associated with the key as a map to a slice of strings.
	GetStringMapStringSlice(key string) map[string][]string

	// GetSizeInBytes returns the size of the value associated with the given key
	// in bytes.
	GetSizeInBytes(key string) uint
}

// SysConfiguration : System configuration
var (
	SysConfiguration Configuration
)

func init() {
	SysConfiguration = builtins.NewMemConfiguration()
}
