package main

import (
	"fmt"
	"time"

	"github.com/cyberpunkattack/pkg/dispatcher"
)

func main() {
	em := dispatcher.New()

	em.AddListener("app.group.session", "aboba", func() (time.Time, error) {
		fmt.Println("aboaobobaoabobaobaoabaaboaobbaobaoaboaboaob")
		return time.Now(), nil
	})

	em.AddListener("app.group.session", "aboba2", func() (time.Time, error) {
		fmt.Println("aboaobobaoabobaobaoabaaboaobbaobaoaboaboaob")
		return time.Now(), nil
	})

	em.AddListener("app.group.session.sesh", "aboba", func() (time.Time, error) {
		fmt.Println("hello nigger")
		return time.Now(), nil
	})

	em.RemoveListener("app.group.session.sesh", "aboba")

	for {
		time.Sleep(time.Second * 2)
		em.Execute("aboba", func(err error) {})
	}

}
