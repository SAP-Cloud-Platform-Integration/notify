// To parse and unparse this JSON data, add this code to your project and do:
//
//    config, err := UnmarshalConfig(bytes)
//    bytes, err = config.Marshal()

package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/urfave/cli"
)

func UnmarshalConfig(data []byte) (Config, error) {
	var r Config
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Config) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Config struct {
	SMTP    SMTP     `json:"smtp"`
	Tenants []Tenant `json:"tenants"`
}

type SMTP struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Server   string `json:"server"`
	Port     string `json:"port"`
}

type Tenant struct {
	Interval int64     `json:"interval"`
	Name     string    `json:"name"`
	Host     string    `json:"host"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Contact  []Contact `json:"contact"`
}

type Contact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// ParseConfigFromPath and return content
func ParseConfigFromPath(path string) *Config {
	rt := &Config{}

	if content, err := ioutil.ReadFile(path); err == nil {
		if err = json.Unmarshal(content, rt); err != nil {
			log.Println(err)
			panic(err)
		}
	} else {
		log.Println(err)
		panic(err)
	}

	return rt
}

func ParseConfigFromEnv(c *cli.Context) *Config {
	return &Config{
		SMTP{
			Server:   c.GlobalString("smtpserver"),
			Port:     c.GlobalString("smtpport"),
			Username: c.GlobalString("smtpuser"),
			Password: c.GlobalString("smtppass"),
		},
		[]Tenant{
			Tenant{
				Interval: c.GlobalInt64("interval"),
				Name:     c.GlobalString("cpiname"),
				Host:     c.GlobalString("cpihost"),
				Username: c.GlobalString("cpiuser"),
				Password: c.GlobalString("cpipassword"),
				Contact: []Contact{
					Contact{
						Name:  c.GlobalString("contact"),
						Email: c.GlobalString("email"),
					},
				},
			},
		},
	}
}
