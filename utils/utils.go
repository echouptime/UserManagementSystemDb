package utils

import (
	"fmt"
	"os"
	"unicode/utf8"
)

var Errors = make(map[string]string)

func Input(prompt string) string {
	var text string
	fmt.Printf(prompt)
	fmt.Scan(&text)
	return text
}

func FileIsExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		panic(err)
		return false
	}
}

func DataCheck(field, data string, limit int) {
	dataLength := utf8.RuneCountInString(data)
	if dataLength == 0 {
		Errors[field] = field + "不能为空"
	} else if dataLength > limit {
		message := fmt.Sprintf("[%s] 字段不能超过 [%d]", field, limit)
		Errors[field] = message
	}
}

//func SaveDb(tasks []*models.User) {
//	os.Remove("userInfo.json")
//	file, err := os.OpenFile("userInfo.json", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0777)
//
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer file.Close()
//	jsonCode := json.NewEncoder(file)
//	jsonCode.Encode(&tasks)
//}
