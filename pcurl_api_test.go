package pcurl

import (
	"fmt"
	"testing"

	"github.com/ar0c/pcurl/body"
	"github.com/stretchr/testify/assert"
)

func TestParseAndObj(t *testing.T) {
	// TODO 如果没有加右'接尾，报错
	all, err := ParseAndObj(`curl https://api.openai.com/v1/completions -H 'Content-Type: application/json' -H 'Authorization: Bearer YOUR_API_KEY' -d '{ "model": "text-davinci-003", "prompt": "Say this is a test", "max_tokens": 7, "temperature": 0 }'`)
	assert.NoError(t, err)
	assert.Equal(t, all.Encode.Body, body.EncodeJSON)
	assert.Equal(t, all.Method, "POST")
	fmt.Printf("%#v\n", all)
}

func TestParseAndJSON(t *testing.T) {
	// TODO 如果没有加右'接尾，报错
	all, err := ParseAndJSON(`curl https://api.openai.com/v1/completions -H 'Content-Type: application/json' -H 'Authorization: Bearer YOUR_API_KEY' -d '{ "model": "text-davinci-003", "prompt": "Say this is a test", "max_tokens": 7, "temperature": 0 }'`)
	assert.NoError(t, err)
	fmt.Printf("%s\n", all)
}

type testCaseObj struct {
	curl string
	url  string
	body string
}

func TestPaserAndObj(t *testing.T) {

	tab := []testCaseObj{
		// {
		// 	curl: `curl -X POST -d '{"a":"b"}' www.qq.com/test`,
		// 	url:  `www.qq.com/test`,
		// },
		// {
		// 	curl: `curl -X POST -i 'http://{{.Host}}/{{.OrgName}}/{{.AppName}}/messages/users' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Authorization: Bearer <YourAppToken>' -d '{"from": "user1","to": ["user2"],"type": "txt","body": {"msg": "testmessages"}}'`,
		// 	url:  `http://{{.Host}}/{{.OrgName}}/{{.AppName}}/messages/users`,
		// },
		// {
		// 	curl: `curl -X DELETE -H 'Accept: application/json' -H 'Authorization: Bearer <YourAppToken> ' https://{{.Host}}/{{.OrgName}}/{{.AppName}}/chatgroups/{{.GroupID}}`,
		// 	url:  `https://{{.Host}}/{{.OrgName}}/{{.AppName}}/chatgroups/{{.GroupID}}`,
		// },
		{
			curl:`curl --location 'https://dev-api.qiandao.cn/auctioneer/bid/list?auction_id=628859035487220257' \
--header 'authority: dev-api.qiandao.cn' \
--header 'x-echoing-env: test-c' \
--header 'x-request-package-id: 1000' \
--header 'x-request-circle: chaowanzu' \
--header 'x-request-sign: Njc3MmZmNTY2MDI0ZGY2N2ZkMTA0MWViMzU5YmFjZTNhZDg4NmFiZDdlZGI2MmEyYmUyNjk3YmIwY2JhMjg0Yg==' \
--header 'x-request-timestamp: 1692694161474' \
--header 'authorization: Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Ijc3MTM5MTQ2MzY4OTg2NDY2MSIsInR5cGUiOiJVU0VSIiwiZXhwIjoxNzUyOTEyMzY0LCJpYXQiOjE3NTI2NTMxNjR9.HkVZJE0UCh_vv88krfEWLeRBXuH5PMlwguZJbmQSAE0VGG-6OxnQGPwe5nEzL17c5MKizsPeNoqfhz8UKdT715XVEBKf6tCG1wE3cLeTOPELddOMC7yI6YACKnBrLhOg96eVuAt7VyHoNn0sHsEDIOZhn_yn_qHaTZnSX8BOMoYFpBGLnTRQtLzUrXmL3BTz0Y2PUuiaVGRFIf6_xxF3FUOiaahvbzlOhFlLXswGLusD694qf15Gxtc1vdZps25WaJi-l7_Li5eJh3BptxZx9H0fTUpH7AYnp2dce9Kb18c77GKecqwR47n6hPw9uFaC2bzkG7XIw-JAErrtAUIAQA' \
--header 'content-type: application/json' \
--header 'x-package-id: 1000' \
--header 'user-agent: Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.0 Mobile/15E148 Safari/604.1 wechatdevtools/1.06.2307250 MicroMessenger/8.0.5 webview/' \
--header 'x-request-sign-type: HMAC_SHA256' \
--header 'x-request-version: 2.26.33' \
--header 'x-request-package-sign-version: 0.0.2' \
--header 'x-device-id: 9c6ae14c-591c-4964-b09b-a40ce527ac99' \
--header 'x-request-sign-version: v1' \
--header 'accept: */*' \
--header 'sec-fetch-site: cross-site' \
--header 'sec-fetch-mode: cors' \
--header 'sec-fetch-dest: empty' \
--header 'referer: https://servicewechat.com/wxaf7362726f135b5c/devtools/page-frame.html' \
--header 'Cookie: et_user_dev=eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Ijc3MTM5MTQ2MzY4OTg2NDY2MSIsInR5cGUiOiJVU0VSIiwiZXhwIjoxNzUyOTEyMzY0LCJpYXQiOjE3NTI2NTMxNjR9.HkVZJE0UCh_vv88krfEWLeRBXuH5PMlwguZJbmQSAE0VGG-6OxnQGPwe5nEzL17c5MKizsPeNoqfhz8UKdT715XVEBKf6tCG1wE3cLeTOPELddOMC7yI6YACKnBrLhOg96eVuAt7VyHoNn0sHsEDIOZhn_yn_qHaTZnSX8BOMoYFpBGLnTRQtLzUrXmL3BTz0Y2PUuiaVGRFIf6_xxF3FUOiaahvbzlOhFlLXswGLusD694qf15Gxtc1vdZps25WaJi-l7_Li5eJh3BptxZx9H0fTUpH7AYnp2dce9Kb18c77GKecqwR47n6hPw9uFaC2bzkG7XIw-JAErrtAUIAQA'`,
url: `https://dev-df.echo.tech/common/create_or_update_api`,
		},
	}

	for _, tc := range tab {

		all, err := ParseAndObj(tc.curl)
		assert.Equal(t, all.URL, tc.url)
		assert.NoError(t, err)
	}
}
