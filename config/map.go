package config

import "time"

// MapConfiguration :
type MapConfiguration struct {
	configurations map[string]interface{}
}

// AllSettings :
func (p *MapConfiguration) AllSettings() map[string]interface{} {
	configurations := make(map[string]interface{})
	for k, v := range p.configurations {
		configurations[k] = v
	}
	return configurations
}

// SetDefault :
func (p *MapConfiguration) SetDefault(key string, value interface{}) {
	if !p.IsSet(key) {
		p.Set(key, value)
	}
}

// IsSet :
func (p *MapConfiguration) IsSet(key string) bool {
	_, ok := p.configurations[key]
	return ok
}

// Set :
func (p *MapConfiguration) Set(key string, value interface{}) {
	p.configurations[key] = value
}

// Get :
func (p *MapConfiguration) Get(key string) interface{} {
	return p.configurations[key]
}

// GetString :
func (p *MapConfiguration) GetString(key string) string {
	return p.Get(key).(string)
}

// GetBool :
func (p *MapConfiguration) GetBool(key string) bool {
	return p.Get(key).(bool)
}

// GetInt :
func (p *MapConfiguration) GetInt(key string) int {
	return p.Get(key).(int)
}

// GetInt32 :
func (p *MapConfiguration) GetInt32(key string) int32 {
	return p.Get(key).(int32)
}

// GetInt64 :
func (p *MapConfiguration) GetInt64(key string) int64 {
	return p.Get(key).(int64)
}

// GetFloat64 :
func (p *MapConfiguration) GetFloat64(key string) float64 {
	return p.Get(key).(float64)
}

// GetTime :
func (p *MapConfiguration) GetTime(key string) time.Time {
	return p.Get(key).(time.Time)
}

// GetDuration :
func (p *MapConfiguration) GetDuration(key string) time.Duration {
	return p.Get(key).(time.Duration)
}

// GetStringSlice :
func (p *MapConfiguration) GetStringSlice(key string) []string {
	return p.Get(key).([]string)
}

// GetStringMap :
func (p *MapConfiguration) GetStringMap(key string) map[string]interface{} {
	return p.Get(key).(map[string]interface{})
}

// GetStringMapString :
func (p *MapConfiguration) GetStringMapString(key string) map[string]string {
	return p.Get(key).(map[string]string)
}

// GetStringMapStringSlice :
func (p *MapConfiguration) GetStringMapStringSlice(key string) map[string][]string {
	return p.Get(key).(map[string][]string)
}

// GetSizeInBytes :
func (p *MapConfiguration) GetSizeInBytes(key string) uint {
	return p.Get(key).(uint)
}

// NewMapConfiguration :
func NewMapConfiguration() Configuration {
	return &MapConfiguration{
		configurations: make(map[string]interface{}),
	}
}
