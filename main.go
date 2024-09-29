package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"new-moonbix/lib"
	"new-moonbix/utils"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

var (
	yellow = color.New(color.FgYellow).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	cyan   = color.New(color.FgCyan).SprintFunc()
	taskid = []int{2058, 2059, 2060, 2061}
	mu     sync.Mutex // Mutex for synchronization
)

type LogInfo struct {
	FirstName   string
	UserId      string
	AccessToken string
	Logs        []string
	Head        *lib.SetHeader
}

var ReponseText = []LogInfo{}
var currentIndex int = 0
var sleepCounter = 0
var logUpdateChan = make(chan bool) // Channel to signal log updates
func init() {
	utils.CheckWhitelistAddr("0xmoonbix")
	utils.InitializeVipers()
	// utils.LicenseCheck()
}

func main() {
	fmt.Print("\033[H\033[2J")
	fmt.Println("DMO ~ @ruc77h")
	query_list := LoadQueriesFromFile()
	prompt := promptui.Select{
		Label: "Time Bypass Playgame ?",
		Items: []string{"Yes", "No"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	var boolResult bool
	if result == "Yes" {
		boolResult = true
	} else {
		boolResult = false
	}
	var wg sync.WaitGroup
	var maxConcurrent int
	if viper.GetString("PROXY") != "" {
		maxConcurrent = viper.GetInt("THREAD_COUNT")
	} else {
		maxConcurrent = 3
	}
	sem := make(chan struct{}, maxConcurrent)
	defer close(sem)
	reff := viper.GetString("REFFCODE")
	processQueries := func(logconsole bool) {
		for index, query := range query_list {
			sem <- struct{}{}
			wg.Add(1)
			go func(index int, query string) {
				defer wg.Done()
				defer func() { <-sem }()
				var logs []string
				query = strings.TrimSpace(query)
				query = strings.Replace(query, "\r", "", -1)
				query = strings.Replace(query, "\n", "", -1)
				if query == "" {
					return
				}
				parsedData, err := url.ParseQuery(query)
				if err != nil {
					fmt.Println("Error parsing data:", err)
					return
				}
				userField := parsedData.Get("user")
				userDecoded, err := url.QueryUnescape(userField)
				if err != nil {
					fmt.Println("Error decoding user data:", err)
					return
				}
				head, err := lib.NewSetHeader()
				if err != nil {
					fmt.Println("Error setting header:", err)
					return
				}
				userID := utils.ExtractField(userDecoded, `"id":`, `,`)
				firstName := utils.ExtractField(userDecoded, `"first_name":"`, `"`)
				GetAccessToken, err := head.GetAccessTokenMoonbix(query)
				if err != nil {
					log.Printf("%s %v\n", red("Err GetAccessTokenMoonbix : "), err)
					logs = []string{"Failed Login " + err.Error()}
					return
				}
				accessToken := GetAccessToken.Data.AccessToken
				if logconsole {
					log.Printf("[ %d ] Using User ID: %s, First Name: %s, Access Token: %s \n", index+1, cyan(userID), yellow(firstName), cyan(accessToken))
					logs = []string{"Success Login..."}
				} else {
					logs = []string{"Success Reloggin..."}

				}

				mu.Lock() // Lock for ReponseText access
				ReponseText = append(ReponseText, LogInfo{
					FirstName:   firstName,
					UserId:      userID,
					AccessToken: accessToken,
					Logs:        logs,
					Head:        head,
				})
				mu.Unlock()

				time.Sleep(2 * time.Second)
			}(index, query)
		}
		wg.Wait()
	}
	processQueries(true)

	// Initialize UI
	if err := ui.Init(); err != nil {
		log.Fatalf("%s %v\n", red("Failed to initialize UI: "), err)
		return
	}
	defer ui.Close()

	// Create a paragraph to display data
	p := widgets.NewParagraph()
	p.SetRect(0, 0, 85, 65)
	p.BorderStyle.Fg = ui.ColorWhite
	p.BorderStyle.Bg = ui.ColorCyan
	updateAccountInfo(p)
	ui.Render(p)
	go handleUIEvents(p)
	go func() {
		for {
			<-logUpdateChan // Wait for log update signal
			updateAccountInfo(p)
			ui.Render(p) // Render UI after log update
		}
	}()

	// Main loop

	for {
		maxConcurrent = viper.GetInt("THREAD_COUNT")
		sem = make(chan struct{}, maxConcurrent)
		for i, resp := range ReponseText {
			sem <- struct{}{}
			wg.Add(1)
			go func(i int, resp LogInfo) {
				defer wg.Done()
				defer func() { <-sem }()

				userInfo, err := resp.Head.GetUserInfo()
				if err != nil {
					mu.Lock()
					ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("Error Get User Info %v", err))
					mu.Unlock()
					logUpdateChan <- true // Signal to update UI
					return
				}
				time.Sleep(2 * time.Second)
				mu.Lock()
				ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("User ID: %s", userInfo.Data.UserID))
				mu.Unlock()
				logUpdateChan <- true // Signal to update UI

				if userInfo.Data.UserID == "" {
					mu.Lock()
					ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("User %s Has Not Yet Registered", userInfo.Data.UserID))
					ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("Register Using Reff Code: %s", reff))
					mu.Unlock()
					logUpdateChan <- true // Signal to update UI
					time.Sleep(2 * time.Second)
					_, err = resp.Head.AcceptRefferall(reff)
					if err != nil {
						mu.Lock()
						ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("Error Accept Refferal %v", err))
						mu.Unlock()
						logUpdateChan <- true // Signal to update UI
						return
					}
					mu.Lock()
					ReponseText[i].Logs = append(ReponseText[i].Logs, "Success Accept Refferal")
					mu.Unlock()
					logUpdateChan <- true // Signal to update UI
					time.Sleep(1 * time.Second)

					participate, err := resp.Head.Participate()
					if err != nil {
						mu.Lock()
						ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("Error Participate %v", err))
						mu.Unlock()
						logUpdateChan <- true // Signal to update UI
						return
					}
					mu.Lock()
					ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("Participate %t", participate.Success))
					mu.Unlock()
					logUpdateChan <- true // Signal to update UI
					time.Sleep(1 * time.Second)
				}
				checkTaskList, err := resp.Head.GetTaskList()
				if err != nil {
					mu.Lock()
					ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("Error Get Task List %v", err))
					mu.Unlock()
					logUpdateChan <- true // Signal to update UI
					return
				}
				if checkTaskList.Data.Data[0].TaskList.Data[0].Status != "COMPLETED" {
					time.Sleep(3 * time.Second)
					checkin, err := resp.Head.CompleteTask("2057", reff)
					if err != nil {
						mu.Lock()
						ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("Error Complete Task %v", err))
						mu.Unlock()
						logUpdateChan <- true // Signal to update UI
						return
					}
					mu.Lock()
					ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("Status Daily Checkin %t", checkin.Success))
					mu.Unlock()
				}
				userInfo, err = resp.Head.GetUserInfo()
				if err != nil {
					mu.Lock()
					ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("Error Get User Info %v", err))
					mu.Unlock()
					logUpdateChan <- true // Signal to update UI
					return
				}

				mu.Lock()
				totalBalance := userInfo.Data.MetaInfo.TotalGrade + userInfo.Data.MetaInfo.ReferralTotalGrade
				totalAttempts := userInfo.Data.MetaInfo.TotalAttempts - userInfo.Data.MetaInfo.ConsumedAttempts
				qualified := userInfo.Data.Qualified
				ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("Balance: %v  | Attempts: %v | Qualified: %t", totalBalance, totalAttempts, qualified))
				ReponseText[i].Logs = append(ReponseText[i].Logs, "Waiting 5 Seconds Before Playing Game...")
				mu.Unlock()
				logUpdateChan <- true // Signal to update UI
				if userInfo.Data.UserID != "" && !userInfo.Data.Qualified {
					mu.Lock()
					ReponseText[i].Logs = append(ReponseText[i].Logs, "Skipping Play Game, User Got Banned")
					mu.Unlock()
					logUpdateChan <- true // Signal to update UI
					return
				}

				if totalAttempts < 1 {
					mu.Lock()
					ReponseText[i].Logs = append(ReponseText[i].Logs, "Total Attempts Less Than 1, Skipping Play Game")
					mu.Unlock()
					logUpdateChan <- true // Signal to update UI
				}
				time.Sleep(5 * time.Second)
				for j := 0; j < totalAttempts; j++ {
					go func() {
						if len(ReponseText[i].Logs) >= 25 {
							mu.Lock()
							ReponseText[i].Logs = []string{"Logs cleared after reaching 25 entries"}
							mu.Unlock()
							logUpdateChan <- true // Signal to update UI
						}
					}()
					startgame, err := resp.Head.StartGame()
					if err != nil {
						mu.Lock()
						ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("Error Start Game %v", err))
						mu.Unlock()
						logUpdateChan <- true // Signal to update UI
						continue
					}
					gameTag := startgame.Data.GameTag
					mu.Lock()
					ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("Start Game With Game Tag %s", gameTag))
					mu.Unlock()
					logUpdateChan <- true // Signal to update UI
					toJson, err := json.Marshal(startgame.Data.CryptoMinerConfig.ItemSettingList)
					if err != nil {
						mu.Lock()
						ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("Error Marshal %v", err))
						mu.Unlock()
						logUpdateChan <- true // Signal to update UI
						continue
					}
					keygame, err := lib.GetKey(gameTag, string(toJson))
					if err != nil {
						mu.Lock()
						ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("Error Get Key %v", err))
						mu.Unlock()
						logUpdateChan <- true // Signal to update UI
						continue
					}
					mu.Lock()
					ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("Encrypted Game Result Payload %s", keygame.Encrypted[:25]))
					if boolResult {
						ReponseText[i].Logs = append(ReponseText[i].Logs, "Sleeping 5 Seconds...")
					} else {
						ReponseText[i].Logs = append(ReponseText[i].Logs, "Sleeping 45 Seconds...")
					}
					mu.Unlock()
					logUpdateChan <- true // Signal to update UI
					if boolResult {
						time.Sleep(5 * time.Second)
					} else {
						time.Sleep(46 * time.Second)
					}

					gamesubmit, err := resp.Head.CompleteGame(keygame.Encrypted, strconv.Itoa(keygame.Point))
					if err != nil {
						mu.Lock()
						ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("Error Complete Game %v", err))
						mu.Unlock()
						logUpdateChan <- true // Signal to update U
						continue
					}
					mu.Lock()
					ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("Status Play Game %t", gamesubmit.Success))
					ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("Success Play Game: You Got %s Point", strconv.Itoa(keygame.Point)))
					mu.Unlock()
					logUpdateChan <- true // Signal to update UI
					mu.Lock()
					userInfo, _ = resp.Head.GetUserInfo()
					totalBalance := userInfo.Data.MetaInfo.TotalGrade + userInfo.Data.MetaInfo.ReferralTotalGrade
					totalAttempts := userInfo.Data.MetaInfo.TotalAttempts - userInfo.Data.MetaInfo.ConsumedAttempts
					qualified := userInfo.Data.Qualified
					ReponseText[i].Logs = append(ReponseText[i].Logs, fmt.Sprintf("Balance: %v  | Attempts: %v | Qualified: %t", totalBalance, totalAttempts, qualified))
					mu.Unlock()
					logUpdateChan <- true // Signal to update U
					time.Sleep(5 * time.Second)
				}
				mu.Lock()
				ReponseText[i].Logs = append(ReponseText[i].Logs, "Sleeping 32 Minutes...")
				mu.Unlock()
				logUpdateChan <- true // Signal to update UI
			}(i, resp)
		}
		wg.Wait()
		sleepCounter++
		time.Sleep(30 * time.Minute)
		if sleepCounter == 3 {
			ReponseText = []LogInfo{}
			processQueries(false)
			sleepCounter = 0
			time.Sleep(30 * time.Second)
		}
	}
}

// Goroutine for handling UI events
func handleUIEvents(p *widgets.Paragraph) {
	uiEvents := ui.PollEvents()
	for {
		e := <-uiEvents
		switch e.ID {
		case "q", "<C-c>":
			ui.Clear()
			ui.Close()
			fmt.Println("\nExit....")
			os.Exit(1)
		case "<Up>":
			if currentIndex < len(ReponseText)-1 {
				currentIndex++
				updateAccountInfo(p)
			}
		case "<Down>":
			if currentIndex > 0 {
				currentIndex--
				updateAccountInfo(p)
			}
		}
	}
}

func LoadQueriesFromFile() []string {
	prompt := promptui.Prompt{
		Label: "Query File: ",
	}
	queryfile, err := prompt.Run()
	if err != nil {
		log.Fatalf("%s %v\n", red("Fatal Err: "), err)
		os.Exit(0)
	}
	file, err := os.ReadFile(queryfile)
	if err != nil {
		log.Fatalf("%s %v\n", red("Fatal Err: "), err)
	}
	splfilequery := strings.Split(string(file), "\n")
	fmt.Println("Load Total Query", len(splfilequery))
	return splfilequery
}

func updateAccountInfo(p *widgets.Paragraph) {
	if len(ReponseText) > 0 {
		account := ReponseText[currentIndex]
		mu.Lock() // Ensure safe access to ReponseText
		defer mu.Unlock()

		header := fmt.Sprintf("\nFirst Name: %s\nUser ID: %s\nAccessToken: %s \n\n", account.FirstName, account.UserId, account.AccessToken)
		logs := ""
		for _, log := range account.Logs {
			logs += fmt.Sprintf("  %s\n", log)
		}
		p.Text = fmt.Sprintf("Account %d of %d\n\n%s%s", currentIndex+1, len(ReponseText), header, logs)
		ui.Render(p)
	}
}
