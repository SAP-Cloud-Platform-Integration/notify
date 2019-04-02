package main

import "github.com/urfave/cli"

var flags = []cli.Flag{
	cli.StringFlag{
		Name:   "config, c",
		Value:  "notify.json",
		EnvVar: "CONFIG_PATH",
		Usage:  "config file path",
	},
	cli.BoolFlag{
		Name:   "env",
		EnvVar: "ENV_CONFIG",
		Usage:  "use env configuration",
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
