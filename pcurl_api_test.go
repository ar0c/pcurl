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
			curl: `curl -X 'GET' -d '' -H 'Accept-Language: zh' -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjU2OTg4NDI3NjkwNDk2MTA3MSIsInR5cGUiOiJVU0VSIiwiZGV2aWNlSWQiOiIxMDAyIiwiZXhwIjoxNzY2Mzg2NzQ4LCJpYXQiOjE3NjYxMjc1NDh9.CzuNevCp6J1Dj_1Y-bJkCUu3URMZIFR14zUxKCcE-00' -H 'Content-Type: application/json' -H 'Cookie: ' -H 'Referer: https://qiandao.cn' -H 'Remoteip: 192.168.10.6' -H 'User-Agent: Kuril/6.26.0 (tech.echoing.kuril; build:685; iOS 18.6.2)' -H 'X-B3-Sampled: 1' -H 'X-B3-Spanid: cc0e6aef3a61ff06' -H 'X-Client-Package-Id: 1002' -H 'X-Device-Id: 9F9F9B92-A0FD-4280-8A88-0D89227B560B' -H 'X-Echo-Audit: false' -H 'X-Echo-City-Code: 381768123839202353' -H 'X-Echo-Install-Id: OTAzNzU1MjU4MzE5MTI3MzUw' -H 'X-Echo-Region: CN' -H 'X-Echo-Teen-Mode: false' -H 'X-Echoing-Env: test-x' -H 'X-Envoy-Attempt-Count: 1' -H 'X-Envoy-External-Address: 192.168.10.6' -H 'X-Forwarded-Client-Cert: By=spiffe://cluster.local/ns/dev/sa/default;Hash=e4c0a24cec8a617d1166836133813e0133f2e4c7417490686b08362da1d83911;Subject=\"\";URI=spiffe://cluster.local/ns/ingress-istio/sa/internal-default-service-account' -H 'X-Forwarded-For: 192.168.10.6,100.122.17.21' -H 'X-Forwarded-Proto: https' -H 'X-Request-Channel: AppStore' -H 'X-Request-Device: ios' -H 'X-Request-Id: 12e7183a-5f9c-9e87-be46-cd7cfe0a0930' -H 'X-Request-Mock-Id: 1541' -H 'X-Request-Package-Id: 1002' -H 'X-Request-Package-Sign-Version: 0.0.3' -H 'X-Request-Sign: gmpL9EtTX/HQ1ifeFpVFrRtnnDVY+Tq+r3nkSZjMnbJm/j8YTakoCfWikZbE2aZYWsArCYe3S9YE7q6JI+s8GT5FmSDf984uf2rKCGEgurz4EUv3oMVlpqxoZv4A9WjnEWjh6tyy62tuB5UScal8NMUGwO0YFKeTDI2Ui7Wy1PQYMuFNaYs2T0cmw2DAzlXQxsLzPXr3HX5mRnJcGRChV7ToyImlG+lA0CQydxXDrUTVJnuCagaNxj0I+NinIA3cV8+K21k4doVRomuRp1NcSqsod8rkd01M6csEf0NOqGpUWq5b/rpN578PonrSnEuIgKJOQANbt0fVWZs64aJoxg==' -H 'X-Request-Sign-Type: RSA2' -H 'X-Request-Sign-Version: v1' -H 'X-Request-Timestamp: 1766127853780' -H 'X-Request-Utm-Source: appstore' -H 'X-Request-Version: 6.26.0' 'https://dev-api.qiandao.cn/c2c-web/v1/luckybag/host-products?hostId=107763688714290865&offset=0&limit=20&salableonly=true&typePropertyTagIds=17&typePropertyTagIds=16&typePropertyTagIds=53193&typePropertyTagIds=1023855&typePropertyTagIds=1164028&typePropertyTagIds=55537&typePropertyTagIds=1236295&typePropertyTagIds=1598689&typePropertyTagIds=1599706&typePropertyTagIds=1598690&typePropertyTagIds=1598692'`,
		},
	}

	for _, tc := range tab {

		all, err := ParseAndObj(tc.curl)
		assert.Equal(t, all.URL, tc.url)
		assert.NoError(t, err)
	}
}
