package config

import "fmt"

type Config struct {
	current                map[string]string
	old                    map[string]string
	lastRanResetCallAge    int
	setTemporarilyAge      int
	setTemporarilyIsActive bool
}

func New() *Config {
	return &Config{
		current:                make(map[string]string),
		old:                    make(map[string]string),
		lastRanResetCallAge:    0,
		setTemporarilyAge:      1,
		setTemporarilyIsActive: false,
	}
}

func WithSetTemporarilyAge(age int) *Config {
	c := New()
	c.setTemporarilyAge = age
	return c
}

func (c *Config) Get(key string) (string, error) {
	if val, ok := c.current[key]; ok {
		return val, nil
	}
	return "", fmt.Errorf("Can't find value for %s", key)
}

func (c *Config) SetTemporarily(key, value string) {
	c.Reset()
	c.setTemporarilyIsActive = true
	if oldValue, ok := c.current[key]; ok {
		c.old[key] = oldValue
	}
	c.current[key] = value
}

func (c *Config) SetPermanently(key, value string) {
	c.current[key] = value
}

func (c *Config) Reset() {
	for name, val := range c.old {
		c.current[name] = val
	}
	c.old = make(map[string]string)
}

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
