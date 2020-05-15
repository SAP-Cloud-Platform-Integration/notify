package main

import "github.com/urfave/cli"

var flags = []cli.Flag{
	cli.StringFlag{
		Name:   "config, c",
		Value:  "notify.json",
		EnvVar: "CONFIG_PATH",
		Usage:  "config file path",
	},
	cli.StringFlag{
		Name:   "ravendsn",
		EnvVar: "RAVEN_DSN",
		Usage:  "Sentry (Raven) Logger DSN",
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
		Value:  "465",
	},
	cli.StringFlag{
		Name:   "smtpuser",
		EnvVar: "SMTP_USER",
		Usage:  "SMTP Server User",
	},
	cli.StringFlag{
		Name:   "smtppass",
		EnvVar: "SMTP_PASSWORD",
		Usage:  "SMTP Server User Password",
	},
	cli.StringFlag{
		Name:   "smtpfrom",
		EnvVar: "SMTP_FROM",
		Value:  "CPI System Notification",
		Usage:  "The 'from' header for the smtp transfer, format: Theo Sun <theo.sun@abd.com>",
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
		Usage:  "The CPI Instance Name",
	},
	cli.StringFlag{
		Name:   "cpihost",
		EnvVar: "CPI_HOST",
		Usage:  "The CPI Host Name",
	},
	cli.StringFlag{
		Name:   "cpiuser",
		EnvVar: "CPI_USER",
		Usage:  "The CPI Tech Credential",
	},
	cli.StringFlag{
		Name:   "cpipassword",
		EnvVar: "CPI_PASSWORD",
		Usage:  "The password of CPI Tech Credential",
	},
	cli.StringFlag{
		Name:   "contact",
		EnvVar: "CONTACT_NAME",
		Value:  "User",
		Usage:  "The name of receiver, split with ',' if multi contacts existed",
	},
	cli.StringFlag{
		Name:   "email",
		EnvVar: "CONTACT_EMAIL",
		Usage:  "The email address of receiver, split with ',' if multi contacts existed",
	},
}
