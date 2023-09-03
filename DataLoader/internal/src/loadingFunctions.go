package main

import (
	"../models/InputDataModels"
	"database/sql"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	allGames          = "..\\BashScripts\\GamesInfo\\AllGames.txt"
	apikey            = "33CF43A0E8B0A89B488CCE3063DED7FC"
	outputDetailsFile = "..\\BashScripts\\GamesInfo\\GameInfoId_"
	outputReview      = "..\\BashScripts\\GamesInfo\\Reviews\\"
)

/*
	func downloadGamesData(lim, offset, delay int) error {
		err, games := GetBaseGamesData(lim, offset)
		if err != nil {
			panic(err)
		}

		for i := 0; i < len(games); i++ {
			resPath := "C:/Dev/Projects/GameDevTool/BashScripts/GamesInfo/" + "GameInfoId_" + strconv.Itoa(int(games[i].AppId)) + ".txt"
			if _, err := os.Stat(resPath); err != nil {
				loadDetails(int(games[i].AppId), delay)
			} else {
				//fmt.Println(":Details already exists")
			}

			resPath = "C:/Dev/Projects/GameDevTool/BashScripts/GamesInfo/Reviews" + strconv.Itoa(int(games[i].AppId)) + ".txt"
			if _, err := os.Stat(resPath); err != nil {
				loadReviews(int(games[i].AppId), delay)
			} else {
				//fmt.Println(": Review already exists")
			}

		}
		return err
	}
*/
func loadDetails(appid, delay int) {
	//fmt.Println("Game with id ", appid)
	url := "http://store.steampowered.com/api/appdetails?appids=" + strconv.Itoa(appid)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(delay) * time.Millisecond)
	if resp.StatusCode != http.StatusOK {
		fmt.Println(time.Now(), "Blocked connection. \n 	Waiting 5 min...")
		time.Sleep(5 * time.Minute)
		return
	}

	var data []byte
	data, err = io.ReadAll(resp.Body)
	f, err := os.OpenFile(outputDetailsFile+strconv.Itoa(appid)+".txt",
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

func loadReviews(appid, delay int) {
	//fmt.Println("Game with id ", appid)
	url := "https://store.steampowered.com/appreviews/" + strconv.Itoa(appid) + "?json=1"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Duration(delay) * time.Millisecond)
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

func updateGamesList() error {
	log.Info("Updating games list")
	url := "http://api.steampowered.com/ISteamApps/GetAppList/v0002/?format=json"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(2000 * time.Millisecond)

	if resp.StatusCode != http.StatusOK {
		fmt.Println(time.Now(), "Blocked connection. \n 	Waiting 5 min...")
		time.Sleep(5 * time.Minute)
		return err
	}

	var data []byte
	data, err = io.ReadAll(resp.Body)
	f, err := os.OpenFile(allGames,
		os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		fmt.Println(err)
	}

	//
	var allGames InputDataModel.AllGamesRequest
	log.Info("Unmarshalling games list")
	data = skipOneJsonLvl(data)
	err = json.Unmarshal(data, &allGames)

	if err != nil {
		print("err!")
	}
	log.Info("Writing to db games list")
	db, err := OpenBD()
	defer db.Close()
	err = WriteBaseGameInfo(allGames, db)

	if err != nil {
		panic(err)
	}
	//fmt.Printf("%+v", allGames)
	return err
}

func updateGamesDetails(lim, offset, rpm int) error {
	count := 0
	db, err := OpenBD()
	defer db.Close()
	err, games := GetBaseGamesData(lim, offset, db)
	delay := 1000 * rpm / 60
	if err != nil {
		panic(err)
	}

	for _, g := range games {
		count++
		//fmt.Println(g.AppId)
		if count%100 == 0 {
			fmt.Println("Count: ", count)
		}

		getDetails(g, delay, db)
		getReviews(g, delay, db)
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

func getDetails(g InputDataModel.App, delay int, db *sql.DB) error {
	resPath := outputDetailsFile + strconv.Itoa(int(g.AppId)) + ".txt"

	if _, err := os.Stat(resPath); err != nil {
		loadDetails(int(g.AppId), delay)
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
		err = WriteGameDetails(appDetails, db)
		if err != nil {
			log.Info(g.AppId, ": ", err)
		}
	} else {
		//log.Info(g.AppId, ": not success, deleting")
		os.Remove(resPath)
	}

	return nil
}

func getReviews(g InputDataModel.App, delay int, db *sql.DB) error {
	resPath := outputReview + strconv.Itoa(int(g.AppId)) + ".txt"

	if _, err := os.Stat(resPath); err != nil {
		loadReviews(int(g.AppId), delay)
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
		err = WriteGameReviews(appReview, db)
		if err != nil {
			log.Info(g.AppId, ": ", err)
		}
	} else {
		//log.Info(g.AppId, ": not success, deleting")
		os.Remove(resPath)
	}

	return nil
}
