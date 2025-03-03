package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	random "math/rand"
	"time"
)

type Aboba struct {
	Name string
}

func Shuffle[T comparable](arr []T) []T {
	shuffled := make([]T, len(arr))
	copy(shuffled, arr)

	randomizer := random.New(random.NewSource(time.Now().UnixMilli()))

	for i := len(shuffled) - 1; i > 0; i-- {
		j := randomizer.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	return shuffled
}

func Duplicate[T any](item T, count int) []T {
	var buf bytes.Buffer

	if err := gob.NewEncoder(&buf).Encode(item); err != nil {
		fmt.Println(err)
		return nil
	}

	result := make([]T, 0)
	for range count {
		var copy T
		decoder := gob.NewDecoder(bytes.NewReader(buf.Bytes()))
		if err := decoder.Decode(&copy); err != nil {
			return nil
		}
		result = append(result, copy)
	}

	return result
}

func main() {
	a := Aboba{Name: "Name of struct"}
	b := Aboba{Name: "Name of struct1"}
	c := Aboba{Name: "Name of struct2"}
	d := Aboba{Name: "Name of struct3"}
	e := Aboba{Name: "Name of struct4"}
	f := Aboba{Name: "Name of struct5"}

	fmt.Println(Shuffle([]Aboba{a, b, c, d, e, f}))

	fmt.Println(Duplicate(a, 5))
}
