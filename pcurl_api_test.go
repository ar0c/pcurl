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
			curl:`curl --location 'https://dev-df.echo.tech/common/create_or_update_api' \
--header 'accept: application/json, text/plain, */*' \
--header 'accept-language: zh-CN,zh;q=0.9,en;q=0.8,en-GB;q=0.7,en-US;q=0.6' \
--header 'content-type: application/json' \
--header 'origin: https://qa.echo.tech' \
--header 'priority: u=1, i' \
--header 'referer: https://qa.echo.tech/' \
--header 'sec-ch-ua: "Microsoft Edge";v="137", "Chromium";v="137", "Not/A)Brand";v="24"' \
--header 'sec-ch-ua-mobile: ?0' \
--header 'sec-ch-ua-platform: "Windows"' \
--header 'sec-fetch-dest: empty' \
--header 'sec-fetch-mode: cors' \
--header 'sec-fetch-site: same-site' \
--header 'user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36 Edg/137.0.0.0' \
--data '{"name":"待助力列表","desc":"","maintainer":"lrc","group":"C2C","url":"https://dev-df.echo.tech/c2c/get_boost_fes","req_param":[{"key":"user_id","name":"用户id","type":"input","required":true,"default":{"name":"","value":""},"value":[]},{"key":"act_id","name":"活动id","type":"input","required":true,"default":{"name":"","value":"877893809604676009"},"value":[]}],"effic_imp_time":5,"extra":null,"is_using":1}'`,
url: `https://dev-df.echo.tech/common/create_or_update_api`,
		},
	}

	for _, tc := range tab {

		all, err := ParseAndObj(tc.curl)
		assert.Equal(t, all.URL, tc.url)
		assert.NoError(t, err)
	}
}
