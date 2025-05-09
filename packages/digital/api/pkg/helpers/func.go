package helpers

import (
	"bytes"
	"crypto/rand"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	random "math/rand"
	"regexp"
	"strings"
	"time"
)

type TagsInterests interface {
	GetId() int
}

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

type listOf []TagsInterests

func Includes[T comparable](list []T, item T) (bool, T, int) {
	if len(list) == 0 {
		return false, item, -1
	}
	for k, v := range list {
		if v == item {
			return true, item, k
		}
	}
	return false, item, -1
}

func IncludesBool[T comparable](list []T, item T) bool {
	if len(list) == 0 {
		return false
	}
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

func ForEachOrError[T comparable](list []T, callback func(item T, index int, list []T) error) error {
	if len(list) == 0 {
		return nil
	}

	for k, v := range list {
		if err := callback(v, k, list); err != nil {
			return err
		}
	}

	return nil
}

func GetIntListFromModel(list listOf) []int {
	result := []int{}
	for _, v := range list {
		result = append(result, v.GetId())
	}
	return result
}

func ShortInclude[T comparable](list []T, item T) bool {
	if len(list) == 0 {
		return false
	}
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

func Some[T comparable](list []T, callback func(item T, index int) bool) bool {
	result := true
	if len(list) == 0 {
		return false
	}
	for k, v := range list {
		if callback(v, k) {
			result = true
		}
	}
	return result
}

func Every[T comparable](list []T, callback func(item T, index int) bool) bool {
	result := true
	if len(list) == 0 {
		return false
	}
	for k, v := range list {
		if !callback(v, k) {
			result = false
		}
	}
	return result
}

func IsArrayIncludes[T comparable](list, list2 []T) bool {
	return Every(list, func(item T, index int) bool {
		inc, _, _ := Includes(list2, item)
		return inc
	})
}

func HandleNilValues(val any, defaultVal any) any {
	if val == nil {
		return defaultVal
	}
	return val
}

func HandleStringValues(val any, defaultVal string) string {
	castedVal, ok := val.(string)
	if val == nil || !ok {
		return defaultVal
	}
	return castedVal
}

func CheckIsValidContentType(needable string, list []string) bool {
	for _, v := range list {
		subStr := strings.Contains(v, needable)
		if subStr {
			return true
		}
	}
	return false
}

func GetFileExtensionFromFile(fileName string) string {
	splitted := strings.Split(fileName, ".")
	if len(splitted) < 2 {
		return ""
	}
	return splitted[1]
}

func EncodeToString(max int) string {
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

func HasPhoneNumber(s string) bool {
	phoneRegex := `(?i)\+?\d[\d\s\-\(\)]{7,}\d`
	r, err := regexp.Compile(phoneRegex)
	if err != nil {
		return false
	}
	return r.MatchString(s)
}

func ParseByteToMap(data []byte) map[string]any {
	result := map[string]any{}
	err := json.Unmarshal(data, &result)
	if err != nil {
		return map[string]any{}
	}
	return result
}

func I(value any) int {
	if value == nil {
		return -1
	}
	return value.(int)
}

func B(value any) bool {
	if value == nil {
		return false
	}
	return value.(bool)
}

func S(value any) string {
	if value == nil {
		return ""
	}
	return value.(string)
}

func F[T any](value any) T {
	var val T
	if value == nil {
		return val
	}
	return value.(T)
}

func AnyToTypeSlice[T any](anySlice []any) []T {
	stringSlice := make([]T, len(anySlice))
	for i, v := range anySlice {
		str, ok := v.(T)
		if !ok {
			return []T{}
		}
		stringSlice[i] = str
	}
	return stringSlice
}

func AnyToStringSlice(anySlice []any) []string {
	stringSlice := make([]string, len(anySlice))
	for i, v := range anySlice {
		str, ok := v.(string)
		if !ok {
			return []string{}
		}
		stringSlice[i] = str
	}
	return stringSlice
}

func WrapToBytes(value any) []byte {
	data, err := json.Marshal(value)
	if err != nil {
		fmt.Println(err)
	}
	return data
}

func Capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}

func Ternary[T comparable](v bool, v1 T, v2 T) T {
	if v {
		return v1
	} else {
		return v2
	}
}

func Duplicate[T any](item T, count int) []T {
	if count <= 1 {
		return []T{item}
	}

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
