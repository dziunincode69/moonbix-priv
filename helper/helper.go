package helper

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/corpix/uarand"
	"golang.org/x/exp/rand"
)

type IdentityCookie struct {
	IdentityCookieID string `json:"$identity_cookie_id"`
}
type WebGLVendorRenderer struct {
	Vendor   string
	Renderer string
}

func RandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(uint64(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomScreenResolution() string {
	resolutions := []string{
		"1920,1080", "1366,768", "1440,900", "1536,864", "1280,720",
		"2560,1440", "3840,2160", "1600,900", "1680,1050",
	}
	return resolutions[rand.Intn(len(resolutions))]
}

func RandomTimezone() string {
	timezones := []string{
		"GMT-05:00", "GMT+01:00", "GMT+03:00", "GMT+07:00", "GMT+08:00",
		"GMT-08:00", "GMT+10:00", "GMT+00:00", "GMT+09:00",
	}
	return timezones[rand.Intn(len(timezones))]
}

func RandomSystemVersion() string {
	versions := []string{
		"Windows 10", "Windows 11", "macOS 12.0", "Ubuntu 20.04", "Fedora 34",
	}
	return versions[rand.Intn(len(versions))]
}

func RandomWebTimezone() string {
	timezone := []string{
		"Asia/Jakarta", "Asia/Tokyo", "Asia/Shanghai", "Asia/Singapore", "Asia/Kolkata",
		"Asia/Dubai", "Asia/Seoul", "Asia/Hong_Kong", "Asia/Taipei",
		"Asia/Kuala_Lumpur", "Asia/Bangkok", "Asia/Manila", "Asia/Kabul", "Asia/Baghdad",
		"US/Pacific", "US/Eastern", "US/Central", "US/Mountain", "US/Alaska", "US/Hawaii",
		"Europe/London", "Europe/Paris", "Europe/Berlin", "Europe/Moscow", "Europe/Istanbul",
		"Africa/Cairo", "Africa/Nairobi", "Africa/Johannesburg",
		"Australia/Sydney", "Australia/Melbourne", "Australia/Brisbane", "Australia/Perth",
		"America/New_York", "America/Los_Angeles", "America/Chicago", "America/Denver",
		"America/Toronto", "America/Vancouver", "America/Mexico_City", "America/Bogota",
		"America/Caracas", "America/Santiago", "America/Buenos_Aires", "America/Sao_Paulo",
		"America/Lima", "America/La_Paz", "America/Manaus", "America/Godthab",
		"America/St_Johns", "America/Adak", "Pacific/Honolulu", "Pacific/Midway",
		"Pacific/Auckland", "Pacific/Fiji", "Pacific/Guam", "Pacific/Tongatapu",
		"Pacific/Port_Moresby", "Pacific/Norfolk", "Pacific/Palau", "Pacific/Nauru",
		"Pacific/Majuro", "Pacific/Kwajalein", "Pacific/Funafuti", "Pacific/Wake",
	}
	return timezone[rand.Intn(len(timezone))]
}

func RandomizeDeviceName() string {
	deviceNames := []string{
		"Edge V91.0.864.64 (Windows)", "Chrome V91.0.4472.124 (Windows)",
		"Firefox V89.0 (Windows)", "Safari V14.1.1 (MacOS)", "Opera V76.0.4017.123 (Windows)",
		"Brave V1.26.74 (Windows)", "Vivaldi V3.8.2259.42 (Windows)",
		"Chromium V91.0.4472.124 (Windows)", "Yandex V21.6.3.757 (Windows)",
		"Samsung Internet V14.0.1.76 (Android)", "UC Browser V12.12.8.1206 (Android)",
		"Opera Mini V58.0.2254.56692 (Android)", "Firefox V89.1.1 (Android)",
		"Chrome V91.0.4472.120 (Android)", "Brave V1.26.74 (Android)",
		"Vivaldi V3.8.2259.42 (Android)", "Chromium V91.0.4472.124 (Android)",
		"Yandex V21.6.3.757 (Android)", "Safari V14.1.1 (iOS)", "Chrome V91.0.4472.124 (iOS)",
		"Firefox V89.0 (iOS)", "Opera V76.0.4017.123 (iOS)", "Brave V1.26.74 (iOS)",
		"Vivaldi V3.8.2259.42 (iOS)", "Chromium V91.0.4472.124 (iOS)", "Yandex V21.6.3.757 (iOS)",
		"Samsung Internet V14.0.1.76 (iOS)",
	}
	return deviceNames[rand.Intn(len(deviceNames))]
}
func generateRandomIdentityCookieID() string {
	rand.Seed(uint64(time.Now().UnixNano()))
	b := make([]byte, 16)
	for i := range b {
		b[i] = byte(rand.Intn(256)) // Random byte between 0-255
	}
	return base64.RawURLEncoding.EncodeToString(b)
}
func GenerateCSRFToken() string {
	// Use current Unix timestamp as a base for token generation
	currentTime := time.Now().UnixNano()

	// Convert the timestamp to a string
	data := fmt.Sprintf("%d", currentTime)

	// Generate MD5 hash
	hash := md5.Sum([]byte(data))

	// Convert the hash to a hexadecimal string
	token := hex.EncodeToString(hash[:])

	return token
}

// Function to generate identity_cookie and return its base64 string and identity_cookie_id
func GenerateIdentityCookie() (string, string) {
	// Generate random identity_cookie_id
	identityCookieID := generateRandomIdentityCookieID()

	// Here we are creating a sample identity_cookie structure with identityCookieID
	identityCookie := fmt.Sprintf(`{"$identity_cookie_id":"%s"}`, identityCookieID)

	// Convert the identityCookie JSON to Base64
	encodedIdentityCookie := base64.StdEncoding.EncodeToString([]byte(identityCookie))

	// Return both the encoded identity cookie and the generated ID
	return encodedIdentityCookie, identityCookieID
}
func GenerateUUID() string {
	rand.Seed(uint64(time.Now().UnixNano()))

	// Template for UUID: "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx"
	uuid := "xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx"

	// Replace "x" and "y" with appropriate random values
	result := strings.Map(func(r rune) rune {
		switch r {
		case 'x':
			return rune(fmt.Sprintf("%x", rand.Intn(16))[0])
		case 'y':
			return rune(fmt.Sprintf("%x", 3&rand.Intn(16)|8)[0]) // 'y' should be one of [8, 9, A, or B]
		default:
			return r
		}
	}, uuid)

	// Replace the last character with "b"
	return result[:len(result)-1] + "b"
}
func RandomizeWebGLVendor() WebGLVendorRenderer {
	// Daftar vendor dan renderer yang umum
	vendors := []WebGLVendorRenderer{
		{"Google Inc.", "ANGLE (Intel(R) UHD Graphics 620 Direct3D11 vs_5_0 ps_5_0)"},
		{"Google Inc.", "ANGLE (NVIDIA GeForce GTX 1050 Ti Direct3D11 vs_5_0 ps_5_0)"},
		{"Google Inc.", "ANGLE (AMD Radeon RX 580 Direct3D11 vs_5_0 ps_5_0)"},
		{"Google Inc.", "ANGLE (NVIDIA GeForce RTX 2060 Direct3D11 vs_5_0 ps_5_0)"},
		{"Google Inc.", "ANGLE (Intel(R) Iris(R) Plus Graphics Direct3D11 vs_5_0 ps_5_0)"},
		{"Google Inc.", "ANGLE (NVIDIA GeForce GTX 1080 Direct3D11 vs_5_0 ps_5_0)"},
		{"Mozilla", "Mozilla (Intel(R) HD Graphics 630)"},
		{"Mozilla", "Mozilla (NVIDIA GeForce GTX 1060)"},
		{"Apple Inc.", "Apple GPU"},
		{"Microsoft Corporation", "Microsoft Basic Render Driver"},
		{"Intel Inc.", "Intel(R) HD Graphics 620"},
	}

	// Seed untuk random generator
	rand.Seed(uint64(time.Now().UnixNano()))
	// Random dari vendor dan renderer yang tersedia
	randomVendor := vendors[rand.Intn(len(vendors))]

	return randomVendor
}
func RandomizeDeviceInfo() (string, string) {
	data := make(map[string]interface{})
	useragent := uarand.GetRandom()
	webGLInfo := RandomizeWebGLVendor()
	data["screen_resolution"] = RandomScreenResolution()
	data["available_screen_resolution"] = RandomScreenResolution()
	data["system_version"] = RandomSystemVersion()
	data["brand_model"] = "unknown"
	data["system_lang"] = "en-us"
	data["timezone"] = RandomTimezone()
	data["timezoneOffset"] = rand.Intn(1440) - 720
	data["user_agent"] = useragent
	data["list_plugin"] = RandomString(30)
	data["canvas_code"] = RandomString(8)
	data["webgl_vendor"] = webGLInfo.Vendor
	data["webgl_renderer"] = webGLInfo.Renderer
	data["audio"] = fmt.Sprintf("%.10f", rand.Float64()*200)
	data["platform"] = "Win64"
	data["web_timezone"] = RandomWebTimezone()
	data["device_name"] = RandomizeDeviceName()
	data["fingerprint"] = RandomString(32)
	data["device_id"] = RandomString(16)
	data["related_device_ids"] = RandomString(16)
	randomizedJSON, err := json.Marshal(data)
	if err != nil {
		return "", ""
	}
	encoded := base64.StdEncoding.EncodeToString(randomizedJSON)

	return encoded, useragent
}

func GetRandomUA() string {
	uaarray := []string{
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.136 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.60 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/101.0.1210.39",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.67 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/45.0.2454.93 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/48.0.2564.116 UBrowser/5.6.13705.206 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.87 Safari/537.36 OPR/36.0.2130.46",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/53.0.2785.90 Safari/537.36 Vivaldi/1.4.589.11",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 AOL/11.0 AOLBUILD/11.0.1305 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.78 Safari/537.36 OPR/47.0.2631.55",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.25 Safari/537.36 Core/1.70.3883.400 QQBrowser/10.8.4559.400",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36 SE 2.X MetaSr 1.0",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.66 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.88 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36 SLBrowser/8.0.0.3161 SLBChan/10",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.145 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.119 YaBrowser/22.3.0.2539 Yowser/2.5 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.141 YaBrowser/22.3.4.734 Yowser/2.5 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Maxthon/5.0.4.3000 Chrome/47.0.2526.73 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36 AtContent/96.5.9594.95",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36 Edg/100.0.1185.44",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36 Edg/100.0.1185.50",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36 OPR/86.0.4363.50",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.127 Safari/537.36 OPR/86.0.4363.59",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.136 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.143 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.143 Safari/537.36 Edg/100.0.1185.57",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.147 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.151 Whale/3.14.134.62 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.60 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.60 Safari/537.36 Edg/100.0.1185.29",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.75 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.75 Safari/537.36 Edg/100.0.1185.39",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.88  Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.4896.88 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4950.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.26 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.41 Safari/537.36 Edg/101.0.1210.32",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Config/99.2.4111.12",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36 Edg/101.0.1210.39",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.64 Safari/537.36 Edg/101.0.1210.47",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.67 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.5005.40 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5030.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.5042.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36 Edge/12.0",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/55.0.2869.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3191.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36 Edge/17.17134",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.121 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.71 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/78.0.3904.108 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.136 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.79 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149  Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.87 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/81.0.4044.92 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.61 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36 Edg/84.0.522.52",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.67 Safari/537.36 Edg/87.0.664.52",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.150 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.90 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36 Edg/90.0.818.56",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.114 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.164 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.77 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36 HBPC/12.0.0.300",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.113 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4624.0 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.846.563 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.45 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.119 Safari/537.36 Edg/98.0.1108.76",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.141 YaBrowser/22.3.3.852 Yowser/2.5 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.141 YaBrowser/22.3.4.731 Yowser/2.5 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.80 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.80 Safari/537.36 Edg/98.0.1108.43",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.82 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.87 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36 Edg/99.0.1150.30",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36 OPR/85.0.4341.22",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.74 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.74 Safari/537.36 Edg/99.0.1150.46",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.82 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36 OPR/85.0.4341.72",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36 OPR/85.0.4341.75",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36 OPR/85.0.4341.79",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.84 Safari/537.36 OPR/85.0.4341.79 (Edition Campaign 76)",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.88 Safari/537.36",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.100.4896.127 Safari/537.36",
	}
	return uaarray[rand.Intn(len(uaarray))]
}

func ParseProxy(proxy string) (string, string, string, string, error) {
	// Memisahkan proxy berdasarkan karakter ":"
	parts := strings.Split(proxy, ":")
	if len(parts) != 4 {
		return "", "", "", "", fmt.Errorf("invalid proxy format, expected format: host:port:username:password")
	}
	host := parts[0]
	port := parts[1]
	username := parts[2]
	password := parts[3]

	return host, port, username, password, nil
}
