package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"time"
)

var (
	DATE_FORMAT_DEFAULT       = "2006-01-02 15:04:05"
	DATE_FORMAT_SIMPLE        = "20060102150405"
	DATE_FORMAT_WITH_TIMEZONE = "2006-01-02 15:04:05-0700"
)

func Md5Hex(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func FormatNow(layout string) string {
	return time.Now().Format(layout)
}

func FormatUTCTime(t time.Time) string {
	loc, _ := time.LoadLocation("UTC")
	return t.In(loc).Format(DATE_FORMAT_DEFAULT)
}

func ArrayContains(arr []string, text string) bool {
	if nil == arr || 0 == len(arr) {
		return false
	}
	for _, val := range arr {
		if val == text {
			return true
		}
	}
	return false
}

func GenerateUrlParams(paramMap map[string]string) string {
	var buffer bytes.Buffer
	var index int

	for key, value := range paramMap {
		if index > 0 {
			buffer.WriteString("&")
		}
		buffer.WriteString(key)
		buffer.WriteString("=")
		buffer.WriteString(value)
		index = index + 1
	}
	return buffer.String()
}
