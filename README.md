# notify

[![CircleCI](https://circleci.com/gh/SAP-Cloud-Platform-Integration/notify.svg?style=shield)](https://circleci.com/gh/SAP-Cloud-Platform-Integration/notify)
[![codecov](https://codecov.io/gh/SAP-Cloud-Platform-Integration/notify/branch/master/graph/badge.svg)](https://codecov.io/gh/SAP-Cloud-Platform-Integration/notify)
![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/SAP-Cloud-Platform-Integration/notify.svg)
[![Docker Cloud Build Status](https://img.shields.io/docker/cloud/build/thedockerimages/cpi-notify)](https://hub.docker.com/repository/docker/thedockerimages/cpi-notify)
[![Size](https://shields.beevelop.com/docker/image/image-size/thedockerimages/cpi-notify/latest.svg?style=flat-square)](https://hub.docker.com/repository/docker/thedockerimages/cpi-notify)
[![Layers](https://shields.beevelop.com/docker/image/layers/thedockerimages/cpi-notify/latest.svg?style=flat-square)](https://hub.docker.com/repository/docker/thedockerimages/cpi-notify)

Send notifications when any integration messages failed

## Project target

We also know that SAP Cloud Integration provides integration-related components, but when the deployed iflows fail, no one knows this information unless the administrator user logs in to the "UI shell".

Sometimes, error messages are so important that administrators want to get them as quickly as possible.

Therefore, the project hopes to provide a way to send notifications when an error occurs.

## Deploy with docker

With docker just run with: 

```bash
docker run -d --restart=always theosun/cpi-notify:latest
```

**Required** env variables: 

* CPI_HOST -- cpi tmn server, `{Account Short Name}-tmn.{SSL Host}.{region}.hana.ondemand.com`
* CPI_USER -- SAP ID User
* CPI_PASSWORD -- SAP ID Password

**Optional** env variables:

* CHECK_INTERVAL - default `60` seconds

* SMTP_SERVER	-- used for email integration, `tls` must be enabled
* SMTP_PORT	-- SMTP Server PORT, default as `465`
* SMTP_USER
* SMTP_PASSWORD
* SMTP_FROM -- 'FROM' header for SFTP	
* CONTACT_NAME	-- user1,user2
* CONTACT_EMAIL -- user1@corp.com,user2.corp.com

* RAVEN_DSN - used for sentry integration

## Current status

- [x] setup proejct
- [x] setup CI
- [x] impl odata api invocation
- [x] parse config json
- [x] impl the periodic logic (schedule)
- [x] impl email template
- [x] define the `Dockerfile` for deployement
- [x] daemon support (by docker)
- [x] user document
- [x] unit tests (private)

## High level design

This application will periodic fetch CPI processing log (based on the [SAP CPI OData API](https://api.sap.com/package/CloudIntegrationAPI)), when some processing message `FAILED`, send email to notify user.

### Graphical description

![schematic diagram](https://res.cloudinary.com/digf90pwi/image/upload/v1555907777/CPI-notify_qshvgp.png)

## Failed codition

* timeout/connection reject(maybe the CPI tenant down)
* some integration messages failed(maybe the external/internal system down)

## [CHANGELOG](./CHANGELOG.md)

## [LICENSE](./LICENSE)