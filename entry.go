package main

import (
	"log"
	"os"

	"github.com/getsentry/raven-go"
	"github.com/imroc/req"

	"github.com/urfave/cli"
)

// Version string, in release version
// This variable will be overwrite by complier
var Version = "SNAPSHOT"

// AppName of this application
var AppName = "CPI Notifier"

// AppUsage of this application
var AppUsage = "Send email notifications when any integration messages failed"

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
				// avoid timeout
				req.SetTimeout(DefaultTimeout)
				ravenDSN := c.GlobalString("ravendsn")
				if ravenDSN != "" {
					log.Printf("setup sentry(raven) log with DSN: %s", ravenDSN)
					if err := raven.SetDSN(ravenDSN); err != nil {
						log.Println(err)
					}
				}
				if c.GlobalBool("env") {
					log.Printf("start notify with config from env")
					config := ParseConfigFromEnv(c)
					StartAllJobs(*config)
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

	app.Flags = flags

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
