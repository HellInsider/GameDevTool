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
	"time"
)

const (
	clusterName    = "test-cluster"
	clientName     = "test-client"
	subscriberName = "subscriber"
)

var (
	apikey       = "33CF43A0E8B0A89B488CCE3063DED7FC"
	output       = "C:\\Dev\\Projects\\GameDevTool\\BashScripts\\GamesInfo\\GameInfoId_"
	outputReview = "C:\\Dev\\Projects\\GameDevTool\\BashScripts\\GamesInfo\\Reviews\\"
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

func downloadGamesData(lim, offset int) error {
	err, games := GetBaseGamesData(lim, offset)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(games); i++ {
		resPath := "C:/Dev/Projects/GameDevTool/BashScripts/GamesInfo/" + "GameInfoId_" + strconv.Itoa(int(games[i].AppId)) + ".txt"
		if _, err := os.Stat(resPath); err != nil {
			loadDetails(int(games[i].AppId))
		} else {
			//fmt.Println(":Details already exists")
		}

		resPath = "C:/Dev/Projects/GameDevTool/BashScripts/GamesInfo/Reviews" + strconv.Itoa(int(games[i].AppId)) + ".txt"
		if _, err := os.Stat(resPath); err != nil {
			loadReviews(int(games[i].AppId))
		} else {
			//fmt.Println(": Review already exists")
		}

	}
	return err
}

func loadDetails(appid int) {
	//fmt.Println("Game with id ", appid)
	url := "http://store.steampowered.com/api/appdetails?appids=" + strconv.Itoa(appid)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(2000 * time.Millisecond)
	if resp.StatusCode != http.StatusOK {
		fmt.Println(time.Now(), "Blocked connection. \n 	Waiting 5 min...")
		time.Sleep(5 * time.Minute)
		return
	}

	var data []byte
	data, err = io.ReadAll(resp.Body)
	f, err := os.OpenFile(output+strconv.Itoa(appid)+".txt",
		os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		fmt.Println(err)
	}
}

func loadReviews(appid int) {
	//fmt.Println("Game with id ", appid)
	url := "https://store.steampowered.com/appreviews/" + strconv.Itoa(appid) + "?json=1"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(2000 * time.Millisecond)
	if resp.StatusCode != http.StatusOK {
		fmt.Println(time.Now(), "Blocked connection. \n 	Waiting 5 min...")
		time.Sleep(5 * time.Minute)
		return
	}

	var data []byte
	data, err = io.ReadAll(resp.Body)
	f, err := os.OpenFile(outputReview+strconv.Itoa(appid)+".txt",
		os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
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
		count++
		//fmt.Println(g.AppId)
		if count%100 == 0 {
			fmt.Println("Count: ", count)
		}

		getDetails(g)
		getReviews(g)
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

func getDetails(g InputDataModel.App) error {
	resPath := "C:/Dev/Projects/GameDevTool/BashScripts/GamesInfo/" + "GameInfoId_" + strconv.Itoa(int(g.AppId)) + ".txt"

	if _, err := os.Stat(resPath); err != nil {
		loadDetails(int(g.AppId))
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
		log.Info("Bad details file", g.AppId)
		return nil
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

	return nil
}

func getReviews(g InputDataModel.App) error {
	resPath := outputReview + strconv.Itoa(int(g.AppId)) + ".txt"

	if _, err := os.Stat(resPath); err != nil {
		loadReviews(int(g.AppId))
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

	if data == nil {
		os.Remove(resPath)
		log.Info("Bad review file ", g.AppId)
		return nil
	}

	var appReview InputDataModel.Review
	appReview.Success = 0
	appReview.AppId = g.AppId

	err = json.Unmarshal(data, &appReview)
	if err != nil {
		log.Info(g.AppId, ": ", err)
	}

	if appReview.Success == 1 {
		err = WriteGameReviews(appReview)
		if err != nil {
			log.Info(g.AppId, ": ", err)
		}
	} else {
		//log.Info(g.AppId, ": not success, deleting")
		os.Remove(resPath)
	}

	return nil
}
