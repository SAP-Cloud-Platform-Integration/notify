package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

// Version string, in release version
// This variable will be overwrite by complier
var Version = "SNAPSHOT"

// AppName of this application
var AppName = "CPI Notifier"

// AppUsage of this application
var AppUsage = "Send email notify when any integration messages failed"

// main entry
func main() {

	app := cli.NewApp()
	app.Version = Version
	app.Name = AppName
	app.Usage = AppUsage
	app.EnableBashCompletion = true

	app.Commands = []cli.Command{
		cli.Command{
			Name: "start",
			Action: func(c *cli.Context) error {
				if c.GlobalBool("env") {
					config := Config{
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
					StartAllJobs(config)
				} else {
					configPath := c.GlobalString("config")
					log.Printf("start notify with config %s", configPath)
					config := ParseConfigFromPath(configPath)
					StartAllJobs(*config)
				}
				return nil
			},
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "notify.json",
			Usage: "config file path",
		},
		cli.BoolFlag{
			Name:  "env",
			Usage: "use env configuration",
		},
		cli.StringFlag{
			Name:   "smtpserver",
			EnvVar: "SMTP_SERVER",
			Usage:  "SMTP Server Host",
		},
		cli.StringFlag{
			Name:   "smtpport",
			EnvVar: "SMTP_PORT",
			Usage:  "SMTP Server Port",
		},
		cli.StringFlag{
			Name:   "smtpuser",
			EnvVar: "SMTP_USER",
			Usage:  "SMTP User",
		},
		cli.StringFlag{
			Name:   "smtppass",
			EnvVar: "SMTP_PASSWORD",
			Usage:  "SMTP Password",
		},
		cli.Int64Flag{
			Name:   "interval",
			Value:  60,
			EnvVar: "CHECK_INTERVAL",
			Usage:  "Check interval",
		},
		cli.StringFlag{
			Name:   "cpiname",
			EnvVar: "CPI_NAME",
		},
		cli.StringFlag{
			Name:   "cpihost",
			EnvVar: "CPI_HOST",
		},
		cli.StringFlag{
			Name:   "cpiuser",
			EnvVar: "CPI_USER",
		},
		cli.StringFlag{
			Name:   "cpipassword",
			EnvVar: "CPI_PASSWORD",
		},
		cli.StringFlag{
			Name:   "contact",
			EnvVar: "CONTACT_NAME",
		},
		cli.StringFlag{
			Name:   "email",
			EnvVar: "CONTACT_EMAIL",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
