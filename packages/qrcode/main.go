package main

import qrcode "github.com/skip2/go-qrcode"

func main() {
	err := qrcode.WriteFile("https://redirect.cyberpunkattack.xyz/?redirect=nJ6L9fg6hEDIey5pJHSwg97r0gGWvABv", qrcode.Medium, 1024, "qr.png")
	if err != nil {
		panic(err)
	}
}
