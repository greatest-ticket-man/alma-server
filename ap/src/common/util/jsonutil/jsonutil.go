package jsonutil

import (
	"alma-server/ap/src/common/error/chk"
	"encoding/json"
	"io"
)

// Marshal オブジェクトをjson文字列にする
func Marshal(obj interface{}) string {
	b, err := json.Marshal(obj)
	chk.SE(err)
	return string(b)
}

// Unmarshal json文字列をオブジェクトにマッピング
func Unmarshal(b []byte, obj interface{}) {
	chk.SE(json.Unmarshal(b, obj))
}

// Write 書き込み
func Write(w io.Writer, data interface{}) error {
	return json.NewEncoder(w).Encode(data)
}
