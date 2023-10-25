package main

import qrcode "github.com/skip2/go-qrcode"

func main() {
	err := qrcode.WriteFile("https://cyberpunk-attack-linker.web.app/", qrcode.Medium, 1024, "qr.png")
	if err != nil {
		panic(err)
	}
}
