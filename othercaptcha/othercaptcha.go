package othercaptcha

import (
	"bytes"
	"encoding/base64"

	"github.com/google/uuid"
	"github.com/steambap/captcha"
)

func Generate() map[string]interface{} {

	// create a captcha of 150x50px
	data, err := captcha.New(200, 100)

	b := new(bytes.Buffer)

	data.WriteImage(b)

	// fmt.Printf("Base64 data %s", base64.StdEncoding.EncodeToString(b.Bytes()))

	// os.WriteFile("image.txt", []byte(base64.StdEncoding.EncodeToString(b.Bytes())), os.ModePerm)

	body := map[string]interface{}{"code": 1, "data": base64.StdEncoding.EncodeToString(b.Bytes()), "captchaId": uuid.NewString(), "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}

	return body
}
