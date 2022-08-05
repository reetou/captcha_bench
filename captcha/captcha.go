package captcha

import (
	"github.com/google/uuid"
	"github.com/mojocn/base64Captcha"
)

//configJsonBody json request body.
type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

func Generate() map[string]interface{} {
	var param configJsonBody
	param.CaptchaType = "string"
	param.DriverString = &base64Captcha.DriverString{
		Height:          100,
		Width:           200,
		ShowLineOptions: 8,
		Length:          5,
		Source:          "abcdefghijklmnopqrstvwxyz",
	}
	param.Id = uuid.NewString()
	param.VerifyValue = uuid.NewString()
	var driver base64Captcha.Driver

	//choose driver
	switch param.CaptchaType {
	case "audio":
		driver = param.DriverAudio
	case "string":
		driver = param.DriverString.ConvertFonts()
	case "math":
		driver = param.DriverMath.ConvertFonts()
	case "chinese":
		driver = param.DriverChinese.ConvertFonts()
	default:
		driver = param.DriverDigit
	}
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}

	return body
}
