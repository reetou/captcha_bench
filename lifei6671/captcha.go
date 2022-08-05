package lifei6671

import (
	"bytes"
	"encoding/base64"

	"github.com/google/uuid"
	"github.com/lifei6671/gocaptcha"
)

func Generate() map[string]interface{} {
	//初始化一个验证码对象
	captchaImage := gocaptcha.NewCaptchaImage(200, 100, gocaptcha.RandLightColor())

	captchaImage.DrawLine(3)
	//画边框
	captchaImage.DrawBorder(gocaptcha.ColorToRGB(0x17A7A7A))

	//画随机噪点
	captchaImage.DrawNoise(gocaptcha.CaptchaComplexHigh)

	//画随机文字噪点
	captchaImage.DrawTextNoise(gocaptcha.CaptchaComplexLower)
	//画验证码文字，可以预先保持到Session种或其他储存容器种
	captchaImage.DrawText(gocaptcha.RandText(4))
	//将验证码保持到输出流种，可以是文件或HTTP流等

	b := new(bytes.Buffer)

	captchaImage.SaveImage(b, gocaptcha.ImageFormatJpeg)

	// fmt.Printf("Base64 data %s", base64.StdEncoding.EncodeToString(b.Bytes()))

	// os.WriteFile("image.txt", []byte(base64.StdEncoding.EncodeToString(b.Bytes())), os.ModePerm)

	body := map[string]interface{}{"code": 1, "data": base64.StdEncoding.EncodeToString(b.Bytes()), "captchaId": uuid.NewString(), "msg": "success"}

	return body
}
