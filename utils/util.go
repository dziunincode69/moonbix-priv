package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"new-moonbix/models"
	"os"
	"strings"

	json "github.com/json-iterator/go"
	"github.com/spf13/viper"
)

const MAIN_URL = "https://www.binance.info/bapi/growth/v1/friendly/growth-paas/mini-app-activity/third-party/"

func ExtractField(data, start, end string) string {
	s := strings.Index(data, start)
	if s == -1 {
		return ""
	}
	s += len(start)
	e := strings.Index(data[s:], end)
	if e == -1 {
		return ""
	}
	return data[s : s+e]
}
func InitializeVipers() {
	viper.SetConfigFile("config.yml")
	if err := viper.ReadInConfig(); err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			fmt.Println("File Config Does not exist")
			var license string
			fmt.Print("Enter License: ")
			fmt.Scanln(&license)
			WriteConfig(license)
		}
	}
}
func WriteConfig(license string) {
	file, err := os.Create("config.yml")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	confstring := `
PROXY:  #HOST:PORT:USER:PASS
REFFCODE: 5684555036
THREAD_COUNT: 5
LICENSE: ` + license + `
`

	// Write some data to the file
	_, err = file.WriteString(confstring)
	if err != nil {
		panic(err)
	}
}

func LicenseCheck() {
	url := "https://dmo.wtf/api/check/" + viper.GetString("LICENSE")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching the URL: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}
	var response models.LicenseChecker
	json.Unmarshal(body, &response)
	if response.Message != "Whitelist found" {
		os.Exit(0)
	}
}
