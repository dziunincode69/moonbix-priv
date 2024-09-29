package helper

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"new-moonbix/models"
	"strings"
)

type EncryptFvideoModel struct {
	FvideoToken string `json:"fvideoToken"`
}

func CreateFvideoTkn(enc string) EncryptFvideoModel {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://38.46.221.46:3000/create-fvideo?dt="+enc, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "insomnia/10.0.0")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response EncryptFvideoModel
	json.Unmarshal(bodyText, &response)
	return response
}
func EncryptFvideo(basekey string) EncryptFvideoModel {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://38.46.221.46:3000/fvideo?secretKey="+basekey, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "insomnia/10.0.0")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response EncryptFvideoModel
	json.Unmarshal(bodyText, &response)
	return response
}

func GetVideoToken(device string) models.VideoTkn {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	var data = strings.NewReader(device)
	req, err := http.NewRequest("POST", "https://www.binance.info/fvideo/dt/sign/web?en=CXU&t=binance", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Host", "www.binance.info")
	req.Header.Set("Content-Length", "1160")
	req.Header.Set("Sec-Ch-Ua", `"Chromium";v="128", "Not;A=Brand";v="24", "Microsoft Edge";v="128", "Microsoft Edge WebView2";v="128"`)
	req.Header.Set("Sec-Ch-Ua-Platform", `"Windows"`)
	req.Header.Set("Csrftoken", "d41d8cd98f00b204e9800998ecf8427e")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/128.0.0.0 Safari/537.36 Edg/128.0.0.0")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", "https://www.binance.info")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://www.binance.info/en/game/tg/moon-bix")
	req.Header.Set("Accept-Language", "en-GB,en;q=0.9,en-US;q=0.8")
	req.Header.Set("Priority", "u=1, i")
	req.Header.Set("Cookie", "bnc-uuid=1db6435f-beb9-4c93-8cb2-2a3573b5dabb; sajssdk_2015_cross_new_user=1; sensorsdata2015jssdkcross=%7B%22distinct_id%22%3A%221921f924ca9128-0948c2f7a46a79-4c657b58-4953600-1921f924caa19f1%22%2C%22first_id%22%3A%22%22%2C%22props%22%3A%7B%22%24latest_traffic_source_type%22%3A%22%E7%9B%B4%E6%8E%A5%E6%B5%81%E9%87%8F%22%2C%22%24latest_search_keyword%22%3A%22%E6%9C%AA%E5%8F%96%E5%88%B0%E5%80%BC_%E7%9B%B4%E6%8E%A5%E6%89%93%E5%BC%80%22%2C%22%24latest_referrer%22%3A%22%22%7D%2C%22identities%22%3A%22eyIkaWRlbnRpdHlfY29va2llX2lkIjoiMTkyMWY5MjRjYTkxMjgtMDk0OGMyZjdhNDZhNzktNGM2NTdiNTgtNDk1MzYwMC0xOTIxZjkyNGNhYTE5ZjEifQ%3D%3D%22%2C%22history_login_id%22%3A%7B%22name%22%3A%22%22%2C%22value%22%3A%22%22%7D%7D; theme=dark")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response models.VideoTkn
	err = json.Unmarshal(bodyText, &response)
	if err != nil {
		log.Fatal(err)
	}
	return response
}
