package main

import (
	"bytes"
	"log"
	"text/template"
)

var funcs = template.FuncMap{
	"formatDate": func(s string) string {
		return formatTime(ParseODataTimeStamp(s))
	},
}

var tpl = template.Must(template.New("email").Funcs(funcs).Parse(`
<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">
<html lang="en">

<head>
    <meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1"> <!-- So that mobile will display zoomed in -->
    <meta http-equiv="X-UA-Compatible" content="IE=edge"> <!-- enable media queries for windows phone 8 -->
    <meta name="format-detection" content="telephone=no"> <!-- disable auto telephone linking in iOS -->
    <title>SAP Cloud Integration Error Notification</title>

    <style type="text/css">
        body {
            margin: 0;
            padding: 0;
            -ms-text-size-adjust: 100%;
            -webkit-text-size-adjust: 100%;
        }

        table {
            border-spacing: 0;
        }

        table td {
            border-collapse: collapse;
        }

        .ExternalClass {
            width: 100%;
        }

        .ExternalClass,
        .ExternalClass p,
        .ExternalClass span,
        .ExternalClass font,
        .ExternalClass td,
        .ExternalClass div {
            line-height: 100%;
        }

        .ReadMsgBody {
            width: 100%;
            background-color: #ebebeb;
        }

        table {
            mso-table-lspace: 0pt;
            mso-table-rspace: 0pt;
        }

        img {
            -ms-interpolation-mode: bicubic;
        }

        .yshortcuts a {
            border-bottom: none !important;
        }

        @media screen and (max-width: 599px) {

            .force-row,
            .container {
                width: 100% !important;
                max-width: 100% !important;
            }
        }

        @media screen and (max-width: 400px) {
            .container-padding {
                padding-left: 12px !important;
                padding-right: 12px !important;
            }
        }

        .ios-footer a {
            color: #aaaaaa !important;
            text-decoration: underline;
        }
    </style>
</head>

<body style="margin:0; padding:0;" bgcolor="#F0F0F0" leftmargin="0" topmargin="0" marginwidth="0" marginheight="0">

    <!-- 100% background wrapper (grey background) -->
    <table border="0" width="100%" height="100%" cellpadding="0" cellspacing="0" bgcolor="#F0F0F0">
        <tr>
            <td align="center" valign="top" bgcolor="#F0F0F0" style="background-color: #ffffff;">

                <br>

                <!-- 600px container (white background) -->
                <table border="0" width="600" cellpadding="0" cellspacing="0" class="container"
                    style="width:600px;max-width:600px">
                    <tr>
                        <td class="container-padding header" align="left"
                            style="font-family:Helvetica, Arial, sans-serif;font-size:22px;font-weight:bold;padding-bottom:12px;color:#2f3c48;padding-left:5px;padding-right:24px">
                            <br>
                            SAP Cloud Integration Error Notification
                            <br>
                            <br>
                        </td>

                    </tr>
                    <tr>
                        <td class="container-padding content" align="left"
                            style="padding-left:24px;padding-right:24px;padding-top:12px;padding-bottom:12px;background-color:#ffffff">
                            <div class="body-text"
                                style="font-family:Helvetica, Arial, sans-serif;font-size:14px;line-height:20px;text-align:left;color:#333333">

                                Dear {{.ContactName}}:
                                <br>
                                For tenant: {{.Tenant.Host}}
                                <br>
                                From UTC time: {{.LastRun}}
                                <br>
                                To UTC time: {{.Now}}
                                <br>
                                Following integration messages are failed

                            </div>

                        </td>
                    </tr>

                    {{range .Artifacts}}

                    
                    <div>
                        <tr>
                            <td class="container-padding header" align="left"
                                style="font-family:Helvetica, Arial, sans-serif;font-size:18px;font-weight:bold;padding-bottom:12px;color:#2f3c48;padding-left:5px;padding-right:24px">

                            </td>
                        </tr>

                        <tr>
                            <td class="container-padding content" align="left"
                                style="padding-left:24px;padding-right:24px;padding-top:12px;padding-bottom:12px;background-color:#ffffff">
                                <br>
                                <div>
                                    <div class="title"
                                        style="font-family:Helvetica, Arial, sans-serif;font-size:18px;font-weight:600;color:#374550">
                                        Artifact: {{.ArtifactName}}</div>
                                    <br>

                                    <div class="body-text"
                                        style="font-family:Helvetica, Arial, sans-serif;font-size:14px;line-height:20px;text-align:left;color:#333333">
                                        {{range .Errors}}
                                        [{{.LogEnd | formatDate}} {{.Status}}]: <a href='{{.AlternateWebLink}}'>{{.TransactionID}}</a>
                                        <br>
                                        {{end}}


                                        <br>
                                    </div>
                                </div>
                            </td>
                        </tr>

                    </div>

                    {{end}}



                    <tr>
                        <td class="container-padding footer-text" align="left"
                            style="font-family:Helvetica, Arial, sans-serif;font-size:12px;line-height:16px;color:#aaaaaa;padding-left:5px;padding-right:24px">
                            <br><br>
                            Server Time (UTC): {{.Now}}
                            <br><br>

                            <a href="https://github.com/SAP-Cloud-Platform-Integration/notify">
                                <strong>CPI notify project</strong>
                            </a>

                            <br>

                            <br><br>

                        </td>
                    </tr>


                </table>
                <!--/600px container -->


            </td>
        </tr>
    </table>
    <!--/100% background wrapper-->

</body>

</html>
`))

func FormatTemplate(data NotificationModel) string {
	buf := new(bytes.Buffer)
	if err := tpl.Execute(buf, data); err != nil {
		log.Println(err)
	}
	return buf.String()
}

// NotificationModel type
type NotificationModel struct {
	Tenant      Tenant
	ContactName string
	LastRun     string
	Now         string
	Artifacts   []Artifact
}

// Artifact type
type Artifact struct {
	ArtifactName string
	Errors       []Result
}
