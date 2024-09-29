package utils

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	json "github.com/json-iterator/go"
)

// Struct APIRES dengan tag JSON yang benar
type APIRES struct {
	Status    int    `json:"status"`
	Name      string `json:"name"`
	ExpiredOn string `json:"expiredOn"`
	Message   string `json:"message"`
	Token     string `json:"token"`
}

// Fungsi untuk memeriksa token JWT
func CheckJWTTOKEN(tokenstr string) {
	mikey := "9297519a9e99804dc27282fae5ef9edfa87907c9"
	secretKey := []byte(mikey)
	token, err := jwt.Parse(tokenstr, func(token *jwt.Token) (interface{}, error) {
		// Pastikan algoritma yang digunakan adalah HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		iat := int64(claims["iat"].(float64))
		exp := int64(claims["exp"].(float64))
		now := time.Now().Unix()

		if now < iat {
			fmt.Println("Token issued in the future!")
		} else if now > exp {
			fmt.Println("Token has expired!")
		} else {
			fmt.Println("Token is valid!")
		}
	} else {
		fmt.Println("Invalid token!")
	}
}

// Fungsi untuk memeriksa whitelist address
func CheckWhitelistAddr(address string) {
	var apires APIRES
	resp, err := http.Get("https://liongfamily.net/check?address=" + address)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
		os.Exit(0)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
		os.Exit(0)
	}

	// Unmarshal JSON ke struct APIRES
	err = json.Unmarshal(body, &apires)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
		os.Exit(0)
	}

	if apires.Status == 200 {
		fmt.Println(apires.Message)
		CheckJWTTOKEN(apires.Token)
	} else {
		os.Exit(0)
	}
}
