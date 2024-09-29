package lib

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"new-moonbix/helper"
	"new-moonbix/models"
	"strconv"
	"strings"
	"time"

	"github.com/corpix/uarand"
	"github.com/spf13/viper"
)

type SetHeader struct {
	client  *http.Client
	client2 *http.Client
	header  map[string]string
}

func NewSetHeader() (*SetHeader, error) {
	var client *http.Client
	var client2 *http.Client
	header := make(map[string]string)
	// uuu := uuid.New().String()

	proxyConfig := viper.GetString("PROXY")
	host, port, username, password, _ := helper.ParseProxy(proxyConfig)
	proxyURL := "http://" + username + ":" + password + "@" + host + ":" + port
	u, err := url.Parse(proxyURL)
	if err != nil {
		log.Fatal(err)
	}
	// Proxy: http.ProxyURL(u),

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client = &http.Client{
		Transport: transport,
		Timeout:   35 * time.Second,
	}
	client2 = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(u),
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: 35 * time.Second,
	}
	// uuidg := helper.GenerateUUID()
	// uuid2 := uuid.NewString()
	// RandomDevice, useragent := helper.RandomizeDeviceInfo()
	// sensor_b64, ident := helper.GenerateIdentityCookie()
	// base_sensorsdata2015jssdkcross := `{"distinct_id":"` + ident + `","first_id":"","props":{"$latest_traffic_source_type":"直接流量","$latest_search_keyword":"未取到值_直接打开","$latest_referrer":""},"identities":"` + sensor_b64 + `","history_login_id":{"name":"","value":""}}`
	// fvideo := helper.GetVideoToken(RandomDevice)
	// encryptedfvideo := helper.EncryptFvideo(fvideo.Dfp)
	// fvideo_token := helper.CreateFvideoTkn(fvideo.Dt)

	header["authority"] = "www.binance.info"
	header["Content-Type"] = "application/json"
	header["User-Agent"] = uarand.GetRandom()
	header["clienttype"] = "web"
	header["Bnc-location"] = "BINANCE"
	header["accept"] = "*/*"
	header["dnt"] = "1"
	header["accept-language"] = "en-US,en;q=0.9"
	header["priority"] = "u=1, i"
	header["accept-language"] = "en-US,en;q=0.9"
	header["origin"] = "https://www.binance.info"
	header["referer"] = "https://www.binance.info/en/game/tg/moon-bix"
	// header["device-info"] = RandomDevice
	header["lang"] = "en"
	header["sec-ch-ua-mobile"] = "?0"
	header["sec-ch-ua-platform"] = `"Windows"`
	header["sec-fetch-dest"] = "empty"
	header["sec-fetch-mode"] = "cors"
	header["sec-fetch-site"] = "same-origin"
	header["x-kl-saas-ajax-request"] = "Ajax_Request"
	header["x-passthrough-token"] = ""
	header["sec-ch-ua-platform"] = "Windows"
	// header["Fvideo-Id"] = fvideo.Dfp
	// header["Fvideo-Token"] = fvideo_token.FvideoToken
	header["Csrftoken"] = helper.GenerateCSRFToken()
	// header["Bnc-Uuid"] = uuidg
	// header["X-Ui-Request-Trace"] = uuid2
	// header["X-Trace-Id"] = uuid2
	// header["Cookie"] = `bnc-uuid=` + uuidg + `; sensorsdata2015jssdkcross=` + base_sensorsdata2015jssdkcross + `; sajssdk_2015_cross_new_user=1; BNC_FV_KEY=` + fvideo.Dfp + "; BNC_FV_KEY_T=" + encryptedfvideo.FvideoToken + `; userPreferredCurrency=USD_USD;`
	return &SetHeader{
		client:  client,
		header:  header,
		client2: client2,
	}, nil
}

func (S *SetHeader) addHeaders(req *http.Request) {
	for key, value := range S.header {
		req.Header.Set(key, value)
	}
}

func (S *SetHeader) CompleteGame(payload, log string) (*models.StandardResp, error) {
	logtoi, _ := strconv.Atoi(log)
	body := models.CompleteGameBody{
		ResourceID: "2056",
		Payload:    payload,
		Log:        logtoi,
	}
	reqbody, _ := json.Marshal(body)
	var data = strings.NewReader(string(reqbody))
	var response models.StandardResp
	var err error
	req, _ := http.NewRequest("POST", "https://www.binance.info/bapi/growth/v1/friendly/growth-paas/mini-app-activity/third-party/game/complete", data)
	S.addHeaders(req)
	var resp *http.Response
	if viper.GetString("PROXY") == "" {
		resp, err = S.client.Do(req)
	} else {
		resp, err = S.client2.Do(req)
	}
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyText, &response)
	if err != nil {
		return nil, err
	}

	if response.Success {
		return &response, nil
	}
	fmt.Println(response)
	time.Sleep(5 * time.Second)
	return nil, errors.New("Failed to complete game" + string(bodyText))
}

func GetKey(gametag, gamedata string) (*models.GameKey, error) {
	flag := false
	minpoint := 75
	var response models.GameKey
	for !flag {
		client := &http.Client{}
		var data = strings.NewReader(`{
		"key": "` + gametag + `",
		"gamedata": ` + gamedata + `
	}`)
		req, err := http.NewRequest("POST", "http://38.46.221.46:3000/encrypt", data)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("User-Agent", "insomnia/10.0.0")
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(bodyText, &response)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		if response.Point > minpoint {
			flag = true
		}
	}
	return &response, nil

}

func (S *SetHeader) StartGame() (*models.StartGame, error) {
	var data = strings.NewReader(`{"resourceId":2056}`)
	var response models.StartGame
	const maxRetry = 3

	for attempt := 0; attempt < maxRetry; attempt++ {
		req, err := http.NewRequest("POST", "https://www.binance.info/bapi/growth/v1/friendly/growth-paas/mini-app-activity/third-party/game/start", data)
		if err != nil {
			continue
		}
		S.addHeaders(req)
		var resp *http.Response
		if viper.GetString("PROXY") == "" {
			resp, err = S.client.Do(req)
		} else {
			resp, err = S.client2.Do(req)
		}
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			continue
		}

		err = json.Unmarshal(bodyText, &response)
		if err != nil {
			continue
		}

		if response.Success {
			return &response, nil
		}

		time.Sleep(10 * time.Second)
	}

	return nil, errors.New("failed to start game after multiple attempts")
}

func (S *SetHeader) CompleteTask(resourceid, reff string) (*models.CompleteTask, error) {
	var data = strings.NewReader(`{"resourceIdList":[` + resourceid + `],"referralCode":null}`)
	var response models.CompleteTask
	const maxRetry = 5

	for attempt := 0; attempt < maxRetry; attempt++ {
		req, err := http.NewRequest("POST", "https://www.binance.info/bapi/growth/v1/friendly/growth-paas/mini-app-activity/third-party/task/complete", data)
		if err != nil {
			return nil, err
		}
		S.addHeaders(req)

		resp, err := S.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyText, &response)
		if err != nil {
			return nil, err
		}
		if response.Success {
			return &response, nil
		}
		time.Sleep(10 * time.Second)
	}

	return nil, errors.New("Failed to complete task after multiple attempts")
}

func (S *SetHeader) Participate() (*models.UserInfo, error) {
	var data = strings.NewReader(`{"resourceId":2056}`)
	var response models.UserInfo
	const maxRetry = 5

	for attempt := 0; attempt < maxRetry; attempt++ {
		req, err := http.NewRequest("POST", "https://www.binance.info/bapi/growth/v1/friendly/growth-paas/mini-app-activity/third-party/game/participated", data)
		if err != nil {
			return nil, err
		}
		S.addHeaders(req)

		resp, err := S.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyText, &response)
		if err != nil {
			return nil, err
		}

		// Jika request berhasil, keluar dari loop dan kembalikan response
		if response.Success {
			return &response, nil
		}

		// Jika gagal, coba lagi setelah menunggu 1 detik
		time.Sleep(10 * time.Second)
	}

	return nil, errors.New("Failed to participate after multiple attempts")
}

func (S *SetHeader) AcceptRefferall(reff string) (*models.StandardResp, error) {
	var data = strings.NewReader(`{"resourceId":2056,"agentId":"` + reff + `"}`)
	var response models.StandardResp
	// const maxRetry = 50
	var err error

	req, err := http.NewRequest("POST", "https://www.binance.info/bapi/growth/v1/friendly/growth-paas/mini-app-activity/third-party/referral", data)
	if err != nil {
		return nil, err
	}
	S.addHeaders(req)

	resp, err := S.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bodyText, &response)
	if err != nil {
		return nil, err
	}

	if response.Success {
		return &response, nil
	}

	return nil, err
}
func (S *SetHeader) GetTaskList() (*models.TaskList, error) {
	var data = strings.NewReader(`{"resourceId":2056}`)
	var response models.TaskList
	const maxRetry = 5

	for attempt := 0; attempt < maxRetry; attempt++ {
		req, err := http.NewRequest("POST", "https://www.binance.info/bapi/growth/v1/friendly/growth-paas/mini-app-activity/third-party/task/list", data)
		if err != nil {
			return nil, err
		}
		S.addHeaders(req)

		resp, err := S.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyText, &response)
		if err != nil {
			return nil, err
		}
		if response.Success {
			return &response, nil
		}
		time.Sleep(10 * time.Second)
	}

	return nil, errors.New("Failed to get user info after multiple attempts")
}

func (S *SetHeader) GetUserInfo() (*models.UserInfo, error) {
	var data = strings.NewReader(`{"resourceId":2056}`)
	var response models.UserInfo
	const maxRetry = 5

	for attempt := 0; attempt < maxRetry; attempt++ {
		req, err := http.NewRequest("POST", "https://www.binance.info/bapi/growth/v1/friendly/growth-paas/mini-app-activity/third-party/user/user-info", data)
		if err != nil {
			return nil, err
		}
		S.addHeaders(req)

		resp, err := S.client.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyText, &response)
		if err != nil {
			return nil, err
		}
		if response.Success {
			return &response, nil
		}
		time.Sleep(10 * time.Second)
	}

	return nil, errors.New("Failed to get user info after multiple attempts")
}

func (S *SetHeader) GetAccessTokenMoonbix(q string) (*models.GetAccessToken, error) {
	body := `{"queryString":"` + q + `","socialType":"telegram"}`
	var data = strings.NewReader(body)

	var response models.GetAccessToken
	var attempt int
	const maxRetry = 5

	for attempt = 0; attempt < maxRetry; attempt++ {
		req, err := http.NewRequest("POST", "https://www.binance.info/bapi/growth/v1/friendly/growth-paas/third-party/access/accessToken", data)
		if err != nil {
			return nil, err
		}
		S.addHeaders(req)

		var resp *http.Response
		if viper.GetString("PROXY") == "" {
			resp, err = S.client.Do(req)
		} else {
			resp, err = S.client2.Do(req)
		}
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		bodyText, err := io.ReadAll(resp.Body)
		if err != nil {
			continue
		}

		err = json.Unmarshal(bodyText, &response)
		if err != nil {
			continue
		}

		if response.Data.AccessToken != "" && response.Success {
			// Jika token ditemukan dan sukses, simpan dan keluar dari loop
			S.header["x-growth-token"] = response.Data.AccessToken
			return &response, nil
		} else {
			return nil, errors.New("Cannot login")
		}

	}
	time.Sleep(5 * time.Second)
	return nil, errors.New("Failed to get access token after multiple attempts")
}
