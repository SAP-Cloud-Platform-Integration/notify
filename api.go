package main

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/imroc/req"
)

func GetFailedInformationFor(t Tenant, from time.Time) {
	if res, err := req.Get(
		fmt.Sprintf("https://%s/api/v1/MessageProcessingLogs", t.Host),
		req.Header{
			"Authorization": fmt.Sprintf(
				"Basic %s",
				base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", t.Username, t.Password))),
			),
		},
		req.QueryParam{
			"$format":      "json",
			"$inlinecount": "allpages",
			"$filter":      "Status eq 'FAILED'",
		},
	); err != nil {
		panic(err)
	} else {
		logs := &MessageProcessingLog{}
		if err := res.ToJSON(logs); err != nil {
			panic(err)
		}
	}

}
