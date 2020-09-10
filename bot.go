package dingo

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var urlBase = `https://oapi.dingtalk.com/robot/send?access_token=%s&timestamp=%d&sign=%s`

// NewBot ...
func NewBot(token string, secret string) *Bot {
	return &Bot{
		token,
		secret,
	}
}

// Bot ...
type Bot struct {
	token  string
	secret string
}

// Send ...
func (b *Bot) Send(d interface{}) (err error) {

	ts := time.Now().UnixNano() / 1000000

	signature := fmt.Sprintf("%d\n%s", ts, b.secret)

	mac := hmac.New(sha256.New, []byte(b.secret))
	mac.Write([]byte(signature))

	b64 := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	url := fmt.Sprintf(urlBase, b.token, ts, b64)

	ab, _ := json.Marshal(d)

	var rsp *http.Response
	rsp, err = http.Post(url, `application/json`, bytes.NewBuffer(ab))
	if err != nil {
		return
	}
	rsp.Body.Close()

	return
}
