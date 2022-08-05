package dchest

import (
	"bytes"
	"encoding/base64"

	"github.com/dchest/captcha"
	"github.com/google/uuid"
)

func Generate() map[string]interface{} {
	id := captcha.New()

	digits := captcha.RandomDigits(5)

	img := captcha.NewImage(id, digits, 200, 100)

	b := new(bytes.Buffer)

	img.WriteTo(b)

	body := map[string]interface{}{"code": 1, "data": base64.StdEncoding.EncodeToString(b.Bytes()), "captchaId": uuid.NewString(), "msg": "success"}

	return body
}
