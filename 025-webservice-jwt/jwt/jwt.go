package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"strings"
)

type header struct {
	alg string `json:"alg"`
	typ string `json:"typ"`
}

type Input struct {
	Name string `json:"name"`
	Sub  string `json:"sub"`
	Iat  int64  `json:"iat"`
}

type Validate struct {
	encodedStr string `json:encodedStr`
}

const secret = "my-long-long-jwt-secret-key-to-generate-jwt-token"

func Generate(inp *Input) (string, error) {

	header, err := json.Marshal(header{
		alg: "HS256",
		typ: "JWT",
	})
	if err != nil {
		return "", err
	}

	input, err := json.Marshal(inp)
	if err != nil {
		return "", err
	}

	head := base64.StdEncoding.EncodeToString(header)
	mid := base64.StdEncoding.EncodeToString(input)

	hmc := hmac.New(sha256.New, []byte(secret))
	hmc.Write([]byte(head + "." + mid))
	sign := hmc.Sum(nil)

	return head + "." + mid + "." + string(sign), nil
}

func Validate(encStr string) (bool, error) {

	arr := strings.Split(encStr, ".")
	if len(arr) != 3 {
		return false, errors.New("invalid token")
	}

	head := base64.URLEncoding arr[0]



}
