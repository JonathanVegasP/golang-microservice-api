package auth

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

var privateKey *rsa.PrivateKey

var publicKey *rsa.PublicKey

func Init() {
	key, err := rsa.GenerateKey(rand.Reader, 2048)

	if err != nil {
		panic(err)
	}

	privateKey = key

	publicKey = &key.PublicKey
}

func createJwt(id interface{}, result chan string, claim map[string]interface{}) {
	h, err := json.Marshal(map[string]string{
		"alg": "RS256",
		"typ": "JWT",
	})

	if err != nil {
		panic(err)
	}

	date := time.Now()

	claims := map[string]interface{}{
		"sub": id,
		"iat": date.Unix(),
		"nbf": date.Add(time.Second * 3).Unix(),
		"exp": date.Add(time.Hour * 3).Unix(),
	}

	for key, value := range claim {
		claims[key] = value
	}

	p, err := json.Marshal(claims)

	if err != nil {
		panic(err)
	}

	data := base64.RawURLEncoding.EncodeToString(h) + "." + base64.RawURLEncoding.EncodeToString(p)

	encrypted, err := rsa.EncryptOAEP(crypto.SHA256.New(), rand.Reader, publicKey, []byte(data), []byte(""))

	if err != nil {
		panic(err)
	}

	result <- data + "." + base64.RawURLEncoding.EncodeToString(encrypted)
}

func CreateJWT(id interface{}, claim map[string]interface{}) string {
	channel := make(chan string)

	go createJwt(id, channel, claim)

	result := <-channel

	close(channel)

	return result
}

func validateToken(r *http.Request, result chan bool) {
	h := r.Header["Authorization"]

	if size := len(h); size != 1 {
		result <- false
		return
	}

	auth := strings.Split(h[0], " ")

	if auth[0] != "Bearer" {
		result <- false
		return
	}

	token := strings.Split(auth[1], ".")[2]

	encrypted, err := base64.RawURLEncoding.DecodeString(token)

	if err != nil {
		result <- false
		return
	}

	decrypted, err := privateKey.Decrypt(nil, encrypted, &rsa.OAEPOptions{Hash: crypto.SHA256})

	if err != nil {
		result <- false
		return
	}

	t := string(decrypted)

	body := strings.Split(t, ".")[1]

	jsonString, err := base64.RawURLEncoding.DecodeString(body)

	if err != nil {
		result <- false
		return
	}

	var jwt map[string]interface{}

	err = json.Unmarshal(jsonString, &jwt)

	if err != nil {
		result <- false
		return
	}

	now := time.Now().Unix()

	result <- int64(jwt["nbf"].(float64)) <= now && int64(jwt["exp"].(float64)) > now
}

func ValidateToken(r *http.Request) bool {
	channel := make(chan bool)

	go validateToken(r, channel)

	result := <-channel

	close(channel)

	return result
}

func getClaims(r *http.Request, result chan map[string]interface{}) {
	token := strings.Split(r.Header["Authorization"][0], " ")[1]

	encrypted, _ := base64.RawURLEncoding.DecodeString(strings.Split(token, ".")[2])

	decrypted, _ := privateKey.Decrypt(nil, encrypted, &rsa.OAEPOptions{Hash: crypto.SHA256})

	t := string(decrypted)

	body := strings.Split(t, ".")[1]

	jsonString, _ := base64.RawURLEncoding.DecodeString(body)

	var jwt map[string]interface{}

	json.Unmarshal(jsonString, &jwt)

	result <- jwt
}

func GetClaims(r *http.Request) map[string]interface{} {
	channel := make(chan map[string]interface{})

	go getClaims(r, channel)

	result := <-channel

	close(channel)

	return result
}
