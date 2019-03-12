# notify

[![CircleCI](https://circleci.com/gh/SAP-Cloud-Platform-Integration/notify.svg?style=shield)](https://circleci.com/gh/SAP-Cloud-Platform-Integration/notify)

Send email notify when any integration messages failed

## Project target

We also know that SAP Cloud Integration provides integration-related components, but when the deployed iflows fail, no one knows this information unless the administrator user logs in to the "UI shell".

Sometimes, error messages are so important that administrators want to get them as quickly as possible.

Therefore, the project hopes to provide a way to send notifications when an error occurs.

## Hign level design

Based on the [SAP CPI OData API](https://api.sap.com/package/CloudIntegrationAPI).

![schematic diagram](https://assets.processon.com/chart_image/5c873b53e4b0ab74ecd43269.png)

## Failed codition

* timeout/connection reject
* some integration messages failed
