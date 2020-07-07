package httputil

import (
	"alma-server/ap/src/common/error/chk"
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
)

// HTTPResult .
type HTTPResult struct {
	StatusCode int32
	Result     interface{}
}

// GetJSON getする
func GetJSON(ctx context.Context, url string, paramMap map[string][]string, reflectType reflect.Type) (*HTTPResult, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	params := req.URL.Query()
	for k, list := range paramMap {
		for _, v := range list {
			params.Add(k, v)
		}
	}
	req.URL.RawQuery = params.Encode()

	return doRequest(ctx, req, "application/json", reflectType)
}

// PostJSON .
func PostJSON(ctx context.Context, url string, param interface{}, reflectType reflect.Type) (*HTTPResult, error) {
	jsonStrBytes, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStrBytes))
	if err != nil {
		return nil, err
	}

	return doRequest(ctx, req, "application/json", reflectType)
}

func doRequest(ctx context.Context, req *http.Request, contentType string, reflectType reflect.Type) (*HTTPResult, error) {
	var err error

	req.Header.Set("Content-Type", contentType)

	req = req.WithContext(ctx)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		chk.SE(resp.Body.Close())
	}()

	// status codeが200じゃない、または、reflectTypeが文字列指定の場合
	if resp.StatusCode != http.StatusOK || reflectType.Kind() == reflect.String {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		result := &HTTPResult{
			StatusCode: int32(resp.StatusCode),
			Result:     string(body),
		}
		return result, nil
	}

	rv := reflect.New(reflectType).Elem().Addr() // newしてポインタ化する
	chk.SE(json.NewDecoder(resp.Body).Decode(rv.Interface()))

	result := &HTTPResult{
		StatusCode: http.StatusOK,
		Result:     rv.Elem().Interface(), // ポインタ解除しつつ、interface{}化する
	}

	return result, nil
}
