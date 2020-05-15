// To parse and unparse this JSON data, add this code to your project and do:
//
//    config, err := UnmarshalConfig(bytes)
//    bytes, err = config.Marshal()

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

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
	From     string `json:"from"`
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

// ParseConfigFromEnv func
func ParseConfigFromEnv(c *cli.Context) (*Config, error) {
	// support multi contacts
	contactNames := strings.Split(c.GlobalString("contact"), ",")
	emails := strings.Split(c.GlobalString("email"), ",")

	if len(contactNames) != len(emails) {
		return nil, fmt.Errorf("contact length is not match email length")
	}

	contacts := []Contact{}

	for i := 0; i < len(emails); i++ {
		contacts = append(contacts, Contact{
			Name:  contactNames[i],
			Email: emails[i],
		})
	}

	// create config
	return &Config{
		SMTP{
			From:     c.GlobalString("smtpfrom"),
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
				Contact:  contacts,
			},
		},
	}, nil
}
