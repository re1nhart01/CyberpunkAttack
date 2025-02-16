package main

import qrcode "github.com/skip2/go-qrcode"

func main() {
	err := qrcode.WriteFile("https://redirect.cyberpunkattack.xyz/?redirect=j9fegbGbH5VEApbeY0sEJv7PMas2Iwwt", qrcode.Medium, 1024, "qr.png")
	if err != nil {
		panic(err)
	}
}
