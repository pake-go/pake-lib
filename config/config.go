// Package config implements the Config type allowing changes in runtime
// behavior.
package config

import "fmt"

// A Config represents a language's configuration, allowing you to change
// the behavior of the language during runtime.
type Config struct {
	// Current represents the current state of the configuration.
	current map[string]string
	// Old represents the previous state of the configuration.
	old map[string]string
	// How many times reset has been called.
	lastRanResetCallAge int
	// How many SmartReset() calls may be called before a flag that is set
	// temporarily is cleared.
	setTemporarilyAge int
	// Represents if a flag has been set temporarily.
	setTemporarilyIsActive bool
}

// New returns a Config which allows you to change the behavior of the language
// during runtime by retrieving and setting flags.
func New() *Config {
	return &Config{
		current:                make(map[string]string),
		old:                    make(map[string]string),
		lastRanResetCallAge:    0,
		setTemporarilyAge:      1,
		setTemporarilyIsActive: false,
	}
}

// WithSetTemporarily returns a Config with setTemporarilyAge set to the given age.
func WithSetTemporarilyAge(age int) *Config {
	c := New()
	c.setTemporarilyAge = age
	return c
}

// Get attempts to retrieve a value based on the given key.  It returns the value if
// the key is found and an error if the key is not found.
func (c *Config) Get(key string) (string, error) {
	if val, ok := c.current[key]; ok {
		return val, nil
	}
	return "", fmt.Errorf("Can't find value for %s", key)
}

// SetTemporarily sets the value of the given key until SmartReset() has been called
// setTemporarilyAge times.
func (c *Config) SetTemporarily(key, value string) {
	c.Reset()
	c.setTemporarilyIsActive = true
	if oldValue, ok := c.current[key]; ok {
		c.old[key] = oldValue
	}
	c.current[key] = value
}

// SetPermanently sets the value of the given key until a new value has been specified
// for the given key.
func (c *Config) SetPermanently(key, value string) {
	c.current[key] = value
}

// Reset would clear the effect of the SetTemporarily regardless of how many SmartReset()
// calls has been done.
func (c *Config) Reset() {
	for name, val := range c.old {
		c.current[name] = val
	}
	c.old = make(map[string]string)
}

// SmartReset checks to see if a flag that has been set temporarily should be cleared
// before clearing it.  This is meant to be ran every time a command has finished executing.
//
// Bug(pchan): If multiple SetTemporarily() has been called and setTemporarilyAge > 1, then all
// flags set by SetTemporarily would be cleared once the first one needs to be cleared.
func (c *Config) SmartReset() {
	if !c.setTemporarilyIsActive {
		return
	}
	if c.lastRanResetCallAge > c.setTemporarilyAge {
		c.Reset()
		c.lastRanResetCallAge = 0
		return
	}
	c.lastRanResetCallAge += 1
}
