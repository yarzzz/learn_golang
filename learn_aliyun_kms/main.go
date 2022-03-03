package main

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var Endpoint = "https://kms.cn-hangzhou.aliyuncs.com"
var AccessKeyId = "LTAI5t8wWWJ4EoRLY9nr8Une"
var AccessKeySecret = "eUL4l6crr72b1J4jT3Sf8rtApaMNzE&"
var KeyId = "alias/fileCoin"

func HMACSHA1(keyStr, value string) string {
	key := []byte(keyStr)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(value))
	res := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return res
}

func percentEncode(x string) string {
	x = url.QueryEscape(x)
	x = strings.ReplaceAll(x, "+", "%20")
	x = strings.ReplaceAll(x, "*", "%2A")
	x = strings.ReplaceAll(x, "%7E", "~")
	return x
}

type GenerateDataKeyResponse struct {
	CiphertextBlob string
	Plaintext      []byte
}

func GenerateDataKey(pwd string) (*GenerateDataKeyResponse, error) {
	param := make(map[string]string)
	if len(pwd) > 0 {
		EncryptionContext, err := json.Marshal(map[string]string{
			"key": pwd,
		})
		if err != nil {
			return nil, err
		}
		param["EncryptionContext"] = string(EncryptionContext)
	}
	respData, err := Call("GenerateDataKey", param)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(respData))
	var d GenerateDataKeyResponse
	json.Unmarshal(respData, &d)
	return &d, nil
}

type DecryptResponse struct {
	Plaintext []byte
}

func Decrypt(blob string, pwd string) (*DecryptResponse, error) {
	param := make(map[string]string)
	param["CiphertextBlob"] = blob
	if len(pwd) > 0 {
		EncryptionContext, err := json.Marshal(map[string]string{
			"key": pwd,
		})
		if err != nil {
			return nil, err
		}
		param["EncryptionContext"] = string(EncryptionContext)
	}
	respData, err := Call("Decrypt", param)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(respData))
	var d DecryptResponse
	json.Unmarshal(respData, &d)
	return &d, nil
}

func Call(action string, param map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, Endpoint, nil)
	if err != nil {
		return nil, err
	}

	ts := time.Now().UTC().Format("2006-01-02T15:04:05Z")

	q := req.URL.Query()
	q.Set("AccessKeyId", AccessKeyId)
	q.Set("Action", action)
	q.Set("KeyId", KeyId)
	q.Set("SignatureMethod", "HMAC-SHA1")
	q.Set("SignatureVersion", "1.0")
	q.Set("Timestamp", ts)
	q.Set("Version", "2016-01-20")
	for k, v := range param {
		q.Set(k, v)
	}

	data := "GET&" + percentEncode("/") + "&" + percentEncode(q.Encode())
	sign := HMACSHA1(AccessKeySecret, data)
	q.Set("Signature", sign)

	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(resp.Body)
}

func main() {
	d1, err := GenerateDataKey("")
	if err != nil {
		fmt.Println(err)
	}
	d2, err := Decrypt(d1.CiphertextBlob, "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bytes.Equal(d1.Plaintext, d2.Plaintext))
}

// PS E:\gopath\src\learn\learn_aliyun_kms> go run .\main.go
// {
//         "KeyId": "8bbbf9af-c380-4729-af41-e4d948b8a1e8",
//         "KeyVersionId": "b67b8e62-f837-471e-bc76-7782e5c86f0b",
//         "CiphertextBlob": "YjY3YjhlNjItZjgzNy00NzFlLWJjNzYtNzc4MmU1Yzg2ZjBieG6PWUy9++2wnMxQP9ck8WToXNOqNbp53Fd7ZEV+Qe2BTBDtAbUYfME3eLWy2Qm5wiTzieY9KHyllHdFRVDHoOielrH2ejw5",
//         "Plaintext": "3USP77v3C11o/LIrB25x9DUN4gYWP3ZybDOFxx5zaPc=",
//         "RequestId": "c2bdfeba-f84c-4825-aca4-f92e4138ed88"
// }
// {
//         "KeyId": "8bbbf9af-c380-4729-af41-e4d948b8a1e8",
//         "Plaintext": "3USP77v3C11o/LIrB25x9DUN4gYWP3ZybDOFxx5zaPc=",
//         "RequestId": "18d687ff-f4b0-4b12-b6f5-33647e6ab9c8"
// }
// true
// PS E:\gopath\src\learn\learn_aliyun_kms> go run .\main.go
// {
//         "KeyId": "8bbbf9af-c380-4729-af41-e4d948b8a1e8",
//         "KeyVersionId": "b67b8e62-f837-471e-bc76-7782e5c86f0b",
//         "CiphertextBlob": "YjY3YjhlNjItZjgzNy00NzFlLWJjNzYtNzc4MmU1Yzg2ZjBi8irOjhuZG61exPemMss1RSyafZ25LHi7uhI59ulal2vcigI2VkpFXzGPKzgCzy6wjMjcPJDcnFZdW+T2UtifRwsbb/Ig164x",
//         "Plaintext": "VQJbrINHTNs3THRu6L5l1YVFcCCet7kdonpOwWcCPQI=",
//         "RequestId": "760c87a6-151d-41ad-a521-4ae0d0c74308"
// }
// {
//         "KeyId": "8bbbf9af-c380-4729-af41-e4d948b8a1e8",
//         "KeyVersionId": "b67b8e62-f837-471e-bc76-7782e5c86f0b",
//         "Plaintext": "VQJbrINHTNs3THRu6L5l1YVFcCCet7kdonpOwWcCPQI=",
//         "RequestId": "04fffe3b-be7f-4362-9b30-8a268672611a"
// }
// true
// PS E:\gopath\src\learn\learn_aliyun_kms> go run .\main.go
// {
//         "KeyId": "8bbbf9af-c380-4729-af41-e4d948b8a1e8",
//         "KeyVersionId": "b67b8e62-f837-471e-bc76-7782e5c86f0b",
//         "CiphertextBlob": "YjY3YjhlNjItZjgzNy00NzFlLWJjNzYtNzc4MmU1Yzg2ZjBi88JD8qN9PZKEdIpG7QWymo6O78fFTRKrn1ZV/EsKYJbUrcbYweTk1+EMC42BIGlpODv5yrxHOiJsUCqMxrpuOh+yHh3BDKrW",
//         "Plaintext": "Bq+PrqIAFGUpxC4mbPE+5z9aEQgrXCyNfYEw0UTSpJQ=",
//         "RequestId": "de2d0795-6bea-4e15-ac87-35bffa4e408d"
// }
// {
//         "KeyId": "8bbbf9af-c380-4729-af41-e4d948b8a1e8",
//         "KeyVersionId": "b67b8e62-f837-471e-bc76-7782e5c86f0b",
//         "Plaintext": "Bq+PrqIAFGUpxC4mbPE+5z9aEQgrXCyNfYEw0UTSpJQ=",
//         "RequestId": "5fd4ed1c-c105-4553-81da-2bf295dc24ae"
// }
// true
