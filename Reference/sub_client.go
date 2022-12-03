package main

import (
	"Client/models/InputDataModels"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	clusterName    = "test-cluster"
	clientName     = "test-client"
	subscriberName = "subscriber"
)

var (
	apikey = "33CF43A0E8B0A89B488CCE3063DED7FC"
	output = "C:\\Dev\\Projects\\GameDevTool\\BashScripts\\GamesInfo\\GameInfoId_"
)

func ConnectNatsStream() (stan.Conn, error) {
	sc, err := stan.Connect(clusterName, clientName,
		stan.NatsURL(stan.DefaultNatsURL),
		stan.Pings(2, 5),
		stan.SetConnectionLostHandler(func(con stan.Conn, err error) {
			log.Printf("Connection nats lost: %s", err)
		}))
	if err != nil {
		return sc, err
	}
	return sc, nil
}

/*
func MsgProcessing(sc stan.Conn) error { //, cache *model.Cashe

	handler := func(msg *stan.Msg) {
		var newItem model.Order
		if err := msg.Ack(); err != nil {
			log.Printf("error ack msg:%v", err)
		}
		err := json.Unmarshal(msg.Data, &newItem)
		if err != nil {
			log.Println(fmt.Errorf("Bad nats message %s", err))
			return
		}
		err = SetDataToDB(newItem)
		if err != nil {
			fmt.Println("error with wriring to DB", err)
		}
		//cache.Lock()
		//cache.Memory[newItem.OrderUid] = newItem
		//cache.Unlock()
		fmt.Println("msg got")
	}

	_, err := sc.Subscribe("myChannel", handler, stan.DurableName(subscriberName), stan.SetManualAckMode())
	if err != nil {
		return fmt.Errorf("error subcribe: %s", err)
	}
	return nil

}
*/

func downloadGamesDataShell(lim, offset int) error {
	err, games := GetBaseGamesData(lim, offset)
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	step := 20
	for i := 0; i < len(games); i += step {
		for j := 0; j < step; j++ {
			wg.Add(1)
			go loaderGoroutine(int(games[i+j].AppId), &wg)
		}
		wg.Wait()
		time.Sleep(3)
	}
	return err
}

func downloadGamesData(lim, offset int) error {
	err, games := GetBaseGamesData(lim, offset)
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	step := 1

	for i := 0; i < len(games); i += step {
		for j := 0; j < step; j++ {
			resPath := "C:/Dev/Projects/GameDevTool/BashScripts/GamesInfo/" + "GameInfoId_" + strconv.Itoa(int(games[i+j].AppId)) + ".txt"
			if _, err := os.Stat(resPath); err != nil {
				wg.Add(1)
				go loaderGoroutine(int(games[i+j].AppId), &wg)
			} else {
				fmt.Println(":Already exists")
			}

		}
		wg.Wait()
		time.Sleep(1000 * time.Millisecond)
	}
	return err
}

func loaderGoroutine(appid int, wg *sync.WaitGroup) {
	defer wg.Done()
	//fmt.Println("Game with id ", appid)
	url := "http://store.steampowered.com/api/appdetails?appids=" + strconv.Itoa(appid)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(time.Now(), "Blocked connection. \n 	Waiting 5 min...")
		time.Sleep(5 * time.Minute)
		return
	}

	var data []byte
	data, err = io.ReadAll(resp.Body)
	f, err := os.OpenFile(output+strconv.Itoa(appid)+".txt",
		os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		fmt.Println(err)
	}
}

func processAllGames() error {
	var allGames InputDataModel.AllGamesRequest
	resPath := "C:/Dev/Projects/GameDevTool/BashScripts/GamesInfo/" + "AllGames.txt"

	inputFile, err := os.Open(resPath)
	if err != nil {
		panic(err)
		return err
	}

	var data []byte
	data, err = ioutil.ReadAll(inputFile)
	if err != nil {
		panic(err)
		return err
	}
	data = skipOneJsonLvl(data)
	err = json.Unmarshal(data, &allGames)

	if err != nil {
		print("err!")
	}
	err = WriteBaseGameInfo(allGames)

	if err != nil {
		panic(err)
	}
	//fmt.Printf("%+v", allGames)
	return err
}

func updateGamesDetails(lim, offset int) error {
	count := 0
	err, games := GetBaseGamesData(lim, offset)
	if err != nil {
		panic(err)
	}

	for _, g := range games {
		resPath := "C:/Dev/Projects/GameDevTool/BashScripts/GamesInfo/" + "GameInfoId_" + strconv.Itoa(int(g.AppId)) + ".txt"

		if _, err := os.Stat(resPath); err == nil {
			count++
			if count%100 == 0 {
				fmt.Println("Count: ", count)
			}

			inputFile, err := os.Open(resPath)
			if err != nil {
				panic(err)
				return err
			}

			var data []byte
			data, err = ioutil.ReadAll(inputFile)
			if err != nil {
				inputFile.Close()
				panic(err)
				return err
			}
			inputFile.Close()

			data = skipOneJsonLvlDetails(data, int(g.AppId))
			if data == nil {
				os.Remove(resPath)
				continue
			}
			var appDetails InputDataModel.AppDetailsRequest
			appDetails.Success = false

			err = json.Unmarshal(data, &appDetails)
			if err != nil {
				log.Info(g.AppId, ": ", err)
			}

			appDetails.Data.PCRequirements.AppId = g.AppId
			appDetails.Data.Price.AppId = g.AppId
			appDetails.Data.Price.Country = "Russia"
			appDetails.Data.ReleaseDate.AppId = g.AppId
			appDetails.Data.Platform.AppId = g.AppId
			appDetails.Data.Critic.AppId = g.AppId
			appDetails.AppId = g.AppId

			if appDetails.Success {
				err = WriteGameDetails(appDetails)
				if err != nil {
					log.Info(g.AppId, ": ", err)
				}
			} else {
				//log.Info(g.AppId, ": not success, deleting")
				os.Remove(resPath)
			}

		}

		/*
			if _, err := os.Stat(resPath); errors.Is(err, os.ErrNotExist) {
				cmd := exec.Command("powershell",
					"C:/Dev/Projects/GameDevTool/BashScripts/GetGameInfo.sh",
					strconv.Itoa(int(g.AppId)))

				err := cmd.Start()
				if err != nil {
					return err
				}

				err = cmd.Wait()
				if err != nil {
					return err
				}
				time.Sleep(3 * time.Second)


			}*/

		//fmt.Printf("%+v", appDetails)
	}

	return err
}

func skipOneJsonLvl(data []byte) []byte {
	i := 0
	for _, c := range data {
		if c == '{' && i != 0 {
			break
		}
		i++
	}
	res := data[i : len(data)-1]
	return res
}

func skipOneJsonLvlDetails(data []byte, id int) []byte {
	if !strings.Contains(string(data), strconv.Itoa(id)) {
		return nil
	}
	i := 0
	for _, c := range data {
		if c == '{' && i != 0 {
			break
		}
		i++
	}
	res := data[i : len(data)-1]
	return res
}
