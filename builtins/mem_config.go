package builtins

import (
	"sync"
	"time"
)

// MemConfiguration :
type MemConfiguration struct {
	configurations sync.Map
}

// AllSettings :
func (c *MemConfiguration) AllSettings() map[string]interface{} {
	configurations := make(map[string]interface{})
	c.configurations.Range(func(k, v interface{}) bool {
		configurations[k.(string)] = v
		return true
	})
	return configurations
}

// SetDefault :
func (c *MemConfiguration) SetDefault(key string, value interface{}) {
	if !c.IsSet(key) {
		c.Set(key, value)
	}
}

// IsSet :
func (c *MemConfiguration) IsSet(key string) bool {
	_, ok := c.configurations.Load(key)
	return ok
}

// Set :
func (c *MemConfiguration) Set(key string, value interface{}) {
	c.configurations.Store(key, value)
}

// Get :
func (c *MemConfiguration) Get(key string) interface{} {
	value, _ := c.configurations.Load(key)
	return value
}

// GetString :
func (c *MemConfiguration) GetString(key string) string {
	value := c.Get(key)
	if value == nil {
		var defaults string
		return defaults // zero value
	}
	return value.(string)
}

// GetBool :
func (c *MemConfiguration) GetBool(key string) bool {
	value := c.Get(key)
	if value == nil {
		var defaults bool
		return defaults // zero value
	}
	return value.(bool)
}

// GetInt :
func (c *MemConfiguration) GetInt(key string) int {
	value := c.Get(key)
	if value == nil {
		var defaults int
		return defaults // zero value
	}
	return value.(int)
}

// GetInt32 :
func (c *MemConfiguration) GetInt32(key string) int32 {
	value := c.Get(key)
	if value == nil {
		var defaults int32
		return defaults // zero value
	}
	return value.(int32)
}

// GetInt64 :
func (c *MemConfiguration) GetInt64(key string) int64 {
	value := c.Get(key)
	if value == nil {
		var defaults int64
		return defaults // zero value
	}
	return value.(int64)
}

// GetFloat64 :
func (c *MemConfiguration) GetFloat64(key string) float64 {
	value := c.Get(key)
	if value == nil {
		var defaults float64
		return defaults // zero value
	}
	return value.(float64)
}

// GetTime :
func (c *MemConfiguration) GetTime(key string) time.Time {
	value := c.Get(key)
	if value == nil {
		var defaults time.Time
		return defaults // zero value
	}
	return value.(time.Time)
}

// GetDuration :
func (c *MemConfiguration) GetDuration(key string) time.Duration {
	value := c.Get(key)
	if value == nil {
		var defaults time.Duration
		return defaults // zero value
	}
	return value.(time.Duration)
}

// GetStringSlice :
func (c *MemConfiguration) GetStringSlice(key string) []string {
	value := c.Get(key)
	if value == nil {
		return nil // zero value
	}
	return value.([]string)
}

// GetStringMap :
func (c *MemConfiguration) GetStringMap(key string) map[string]interface{} {
	value := c.Get(key)
	if value == nil {
		return nil // zero value
	}
	return value.(map[string]interface{})
}

// GetStringMapString :
func (c *MemConfiguration) GetStringMapString(key string) map[string]string {
	value := c.Get(key)
	if value == nil {
		return nil // zero value
	}
	return value.(map[string]string)
}

// GetStringMapStringSlice :
func (c *MemConfiguration) GetStringMapStringSlice(key string) map[string][]string {
	value := c.Get(key)
	if value == nil {
		return nil // zero value
	}
	return value.(map[string][]string)
}

// GetSizeInBytes :
func (c *MemConfiguration) GetSizeInBytes(key string) uint {
	value := c.Get(key)
	if value == nil {
		var defaults uint
		return defaults // zero value
	}
	return value.(uint)
}

// NewMemConfiguration :
func NewMemConfiguration() *MemConfiguration {
	return &MemConfiguration{}
}
