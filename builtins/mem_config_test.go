package builtins_test

import (
	"testing"
	"time"

	"github.com/MrLYC/gokit/builtins"
	"github.com/stretchr/testify/assert"
)

// TestAllSettings :
func TestAllSettings(t *testing.T) {
	dob, _ := time.Parse(time.RFC3339, "1979-05-27T07:32:00Z")
	all := map[string]interface{}{"owner": map[string]interface{}{"organization": "MongoDB", "bio": "MongoDB Chief Developer Advocate & Hacker at Large", "dob": dob}, "title": "TOML Example", "ppu": 0.55, "eyes": "brown", "clothing": map[string]interface{}{"trousers": "denim", "jacket": "leather", "pants": map[string]interface{}{"size": "large"}}, "id": "0001", "batters": map[string]interface{}{"batter": []interface{}{map[string]interface{}{"type": "Regular"}, map[string]interface{}{"type": "Chocolate"}, map[string]interface{}{"type": "Blueberry"}, map[string]interface{}{"type": "Devil's Food"}}}, "hacker": true, "beard": true, "hobbies": []interface{}{"skateboarding", "snowboarding", "go"}, "age": 35, "type": "donut", "newkey": "remote", "name": "Cake", "p_id": "0001", "p_ppu": "0.55", "p_name": "Cake", "p_batters": map[string]interface{}{"batter": map[string]interface{}{"type": "Regular"}}, "p_type": "donut", "foos": []map[string]interface{}{map[string]interface{}{"foo": []map[string]interface{}{map[string]interface{}{"key": 1}, map[string]interface{}{"key": 2}, map[string]interface{}{"key": 3}, map[string]interface{}{"key": 4}}}}}

	config := builtins.NewMemConfiguration()
	for k, v := range all {
		config.Set(k, v)
	}

	assert.Equal(t, all, config.AllSettings())
}

// TestNotFound :
func TestNotFound(t *testing.T) {
	config := builtins.NewMemConfiguration()
	testcases := []struct {
		result, excepted interface{}
	}{
		{config.GetInt("x"), 0},
		{config.GetFloat64("x"), 0.0},
		{config.GetString("x"), ""},
		{config.GetBool("x"), false},
		{config.GetStringMap("x"), map[string]interface{}(nil)},
	}

	for _, cases := range testcases {
		assert.Equal(t, cases.excepted, cases.result)
	}
}
