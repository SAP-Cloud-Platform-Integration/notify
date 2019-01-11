package main

import (
	"encoding/json"
	"io/ioutil"
)

// Config Type
type Config struct {
	Interval int64    `json:"interval"`
	SMTP     SMTP     `json:"smtp"`
	Tenants  []Tenant `json:"tenants"`
}

// SMTP Config
type SMTP struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     string `json:"port"`
}

// Tenant Information
type Tenant struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Contact  string `json:"contact"`
}

// ParseConfigFromPath and return content
func ParseConfigFromPath(path string) *Config {
	rt := &Config{}

	if content, err := ioutil.ReadFile(path); err == nil {
		if err = json.Unmarshal(content, rt); err != nil {
			panic(err)
		}
	} else {
		panic(err)
	}

	return rt
}
