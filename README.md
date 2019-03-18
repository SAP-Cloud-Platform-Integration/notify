# notify

[![CircleCI](https://circleci.com/gh/SAP-Cloud-Platform-Integration/notify.svg?style=shield)](https://circleci.com/gh/SAP-Cloud-Platform-Integration/notify)
[![codecov](https://codecov.io/gh/SAP-Cloud-Platform-Integration/notify/branch/master/graph/badge.svg)](https://codecov.io/gh/SAP-Cloud-Platform-Integration/notify)

Send email notify when any integration messages failed

## Project target

We also know that SAP Cloud Integration provides integration-related components, but when the deployed iflows fail, no one knows this information unless the administrator user logs in to the "UI shell".

Sometimes, error messages are so important that administrators want to get them as quickly as possible.

Therefore, the project hopes to provide a way to send notifications when an error occurs.

## Current status

- [x] setup proejct
- [x] setup CI
- [x] impl odata api invocation
- [x] parse config json
- [x] impl the periodic logic (schedule)
- [x] impl email template
- [ ] define the `Dockerfile` for deployement
- [ ] user document
- [ ] unit tests

## Hign level design

This application will periodic fetch CPI processing log (based on the [SAP CPI OData API](https://api.sap.com/package/CloudIntegrationAPI)), when some processing message `FAILED`, send email to notify user.

### Graphical description

![schematic diagram](https://assets.processon.com/chart_image/5c873b53e4b0ab74ecd43269.png)

## Failed codition

* timeout/connection reject(maybe the CPI tenant down)
* some integration messages failed(maybe the external/internal system down)

## The other way to notify user error happened

The CPI developer could define the `exception subprocess`, and invoke rest API to send email (or use the SMTP adapter).

## Email template

![](https://res.cloudinary.com/digf90pwi/image/upload/v1552638002/2019-03-15_16-19-05_bdxheh.png)
