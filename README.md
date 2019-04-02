# notify

[![CircleCI](https://circleci.com/gh/SAP-Cloud-Platform-Integration/notify.svg?style=shield)](https://circleci.com/gh/SAP-Cloud-Platform-Integration/notify)
[![codecov](https://codecov.io/gh/SAP-Cloud-Platform-Integration/notify/branch/master/graph/badge.svg)](https://codecov.io/gh/SAP-Cloud-Platform-Integration/notify)
[![](https://images.microbadger.com/badges/image/theosun/cpi-notify.svg)](https://microbadger.com/images/theosun/cpi-notify "Get your own image badge on microbadger.com")
[![](https://images.microbadger.com/badges/version/theosun/cpi-notify.svg)](https://microbadger.com/images/theosun/cpi-notify "Get your own version badge on microbadger.com")

Send email notify when any integration messages failed

## Project target

We also know that SAP Cloud Integration provides integration-related components, but when the deployed iflows fail, no one knows this information unless the administrator user logs in to the "UI shell".

Sometimes, error messages are so important that administrators want to get them as quickly as possible.

Therefore, the project hopes to provide a way to send notifications when an error occurs.

## Run

You need a smtp mail server & some CPI infromation.

As you see, you can add many tenants to monitor.

The `interval` unit is `second`.

Following is a sample `notify.json` configuration file.

```json
{
  "$schema": "https://raw.githubusercontent.com/SAP-Cloud-Platform-Integration/notify/master/config_schema.json",
  "smtp": {
    "username": "username",
    "password": "password",
    "server": "1.2.3.4",
    "port": "465"
  },
  "tenants": [
    {
      "interval": 60,
      "name": "sample",
      "host": "mock-tmn.hci.cn1.hana.ondemand.com",
      "username": "username",
      "password": "password",
      "contact": [
        {
          "name": "Theo Sun",
          "email": "theo.sun@outlook.com"
        }
      ]
    }
  ]
}
```

then run 

```bash
./notify start
2019/03/27 13:07:19 start notify with config notify.json
2019/03/27 13:07:21 setup job for xxxxxxx-tmn.hci.eu1.hana.ondemand.com tenant
2019/03/27 13:07:21 starting jobs
```

## Deploy with docker

just run 

```bash
docker run -d theosun/cpi-notify
```

required env variables: 

* SMTP_SERVER	
* SMTP_PORT	
* SMTP_USER	
* SMTP_PASSWORD	
* CHECK_INTERVAL
* CPI_HOST
* CPI_USER	
* CPI_PASSWORD	
* CONTACT_NAME	
* CONTACT_EMAIL

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

![schematic diagram](https://assets.processon.com/chart_image/5c873b53e4b0ab74ecd43269.png)

## Failed codition

* timeout/connection reject(maybe the CPI tenant down)
* some integration messages failed(maybe the external/internal system down)

## The other way to notify user error happened

The CPI developer could define the `exception subprocess`, and invoke rest API to send email (or use the SMTP adapter).

## Email notification example

![](https://res.cloudinary.com/digf90pwi/image/upload/v1553665838/%E5%BE%AE%E4%BF%A1%E6%88%AA%E5%9B%BE_20190327133545_a5mj0f.png)
