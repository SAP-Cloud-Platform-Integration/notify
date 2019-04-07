package main

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"time"

	"github.com/imroc/req"
)

// CheckAPIAvailable func
func CheckAPIAvailable(t Tenant) (bool, string) {
	avaible := true
	msg := ""
	link := fmt.Sprintf("https://%s/api/v1/", t.Host)
	if res, err := req.Head(
		link,
		req.Header{
			"Authorization": fmt.Sprintf(
				"Basic %s",
				base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", t.Username, t.Password))),
			),
		},
	); err != nil {
		avaible = false
		msg = err.Error()
	} else if res.Response().StatusCode != 200 {
		avaible = false
		msg = fmt.Sprintf("access CPI metadata failed, please check user & privileges for %s", t.Host)
	}
	return avaible, msg
}

// GetErrorLogFor msg
func GetErrorLogFor(t Tenant, m Result) string {
	rt := ""
	res, _ := req.Get(
		fmt.Sprintf("%s/$value", m.ErrorInformation.Deferred.URI),
		req.Header{
			"Authorization": fmt.Sprintf(
				"Basic %s",
				base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", t.Username, t.Password))),
			),
		},
	)
	statusCode := res.Response().StatusCode

	switch statusCode {
	case 200:
		rt, _ = res.ToString()
	default:
		rt = "Error log not found."
	}

	return rt
}

// GetFailedInformationFor specific tenant
func GetFailedInformationFor(t Tenant, from time.Time) (msg *MessageProcessingLog, err error) {

	// set timeout
	req.SetTimeout(time.Duration(t.Interval) * time.Second)

	rt := &MessageProcessingLog{}

	res, err := req.Get(
		fmt.Sprintf("https://%s/api/v1/MessageProcessingLogs", t.Host),
		req.Header{
			"Authorization": fmt.Sprintf(
				"Basic %s",
				base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", t.Username, t.Password))),
			),
		},
		req.QueryParam{
			"$orderby":     "LogEnd desc",
			"$inlinecount": "allpages",
			// in json format
			"$format": "json",
			// only fetch 100 records
			"$top": 100,
			// only fetch failed logs
			"$filter": fmt.Sprintf("Status eq 'FAILED' and LogEnd ge datetime'%s'", formatTime(from)),
		},
	)
	if err != nil {
		return nil, err
	}

	statusCode := res.Response().StatusCode

	if statusCode != 200 {
		u, _ := url.QueryUnescape(res.Request().URL.String())
		return nil, fmt.Errorf("access cpi data failed, request url: %s, response code: %d", u, statusCode)
	}

	if statusCode == 200 {

		if err := res.ToJSON(rt); err != nil {
			return nil, err
		}

	}

	return rt, nil

}
