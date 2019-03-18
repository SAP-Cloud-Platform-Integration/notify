package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/imroc/req"
)

// GetFailedInformationFor specific tenant
func GetFailedInformationFor(t Tenant, from time.Time) *MessageProcessingLog {
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
		log.Println(err)
	}

	statusCode := res.Response().StatusCode

	if statusCode != 200 {
		log.Printf("access cpi data failed, request url: %s, response code: %d", res.Request().URL.String(), statusCode)
	}

	if statusCode == 200 {

		if err := res.ToJSON(rt); err != nil {
			log.Println(err)
		}

	}

	return rt

}
