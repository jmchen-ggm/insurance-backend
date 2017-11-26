package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"reflect"
	"unsafe"
)

func BytesToString(b []byte) (s string) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pstring.Data = pbytes.Data
	pstring.Len = pbytes.Len
	return
}

func StringToBytes(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}

func ObjToString(v interface{}) string {
	var bytes, _ = json.Marshal(v)
	return BytesToString(bytes)
}

func MD5(s string) string {
	h := md5.New()
	h.Write(StringToBytes(s))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
