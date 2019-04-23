# notify

[![CircleCI](https://circleci.com/gh/SAP-Cloud-Platform-Integration/notify.svg?style=shield)](https://circleci.com/gh/SAP-Cloud-Platform-Integration/notify)
[![codecov](https://codecov.io/gh/SAP-Cloud-Platform-Integration/notify/branch/master/graph/badge.svg)](https://codecov.io/gh/SAP-Cloud-Platform-Integration/notify)
[![Size](https://shields.beevelop.com/docker/image/image-size/theosun/cpi-notify/latest.svg?style=flat-square)](https://cloud.docker.com/repository/docker/theosun/cpi-notify)
[![Layers](https://shields.beevelop.com/docker/image/layers/theosun/cpi-notify/latest.svg?style=flat-square)](https://cloud.docker.com/repository/docker/theosun/cpi-notify)
![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/SAP-Cloud-Platform-Integration/notify.svg)



Send notifications when any integration messages failed

## Project target

We also know that SAP Cloud Integration provides integration-related components, but when the deployed iflows fail, no one knows this information unless the administrator user logs in to the "UI shell".

Sometimes, error messages are so important that administrators want to get them as quickly as possible.

Therefore, the project hopes to provide a way to send notifications when an error occurs.

## Deploy with docker

Tranditional deploy approach is deprecated, only support docker deployment now.

With docker just run with: 

```bash
docker run -d theosun/cpi-notify
```

**Required** env variables: 

* CPI_HOST
* CPI_USER	
* CPI_PASSWORD	

**Optional** env variables:

* CHECK_INTERVAL - default `60` seconds

* SMTP_SERVER	- used for email integration
* SMTP_PORT	
* SMTP_USER	
* SMTP_PASSWORD	
* CONTACT_NAME	
* CONTACT_EMAIL

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

## Hign level design

This application will periodic fetch CPI processing log (based on the [SAP CPI OData API](https://api.sap.com/package/CloudIntegrationAPI)), when some processing message `FAILED`, send email to notify user.

### Graphical description

![schematic diagram](https://res.cloudinary.com/digf90pwi/image/upload/v1555907777/CPI-notify_qshvgp.png)

## Failed codition

* timeout/connection reject(maybe the CPI tenant down)
* some integration messages failed(maybe the external/internal system down)

## [CHANGELOG](./CHANGELOG)

## [LICENSE](./LICENSE)