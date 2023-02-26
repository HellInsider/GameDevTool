package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {

	/*	mem := model.New()
		err := GetDataFromDB(mem)
		if err != nil {
			log.Println(fmt.Errorf("init cache error %s", err))
			return
		}
		sc, err := ConnectNatsStream()
		if err != nil {
			log.Fatal(err)
		}
		defer sc.Close()
		err = MsgProcessing(sc, mem)
		if err != nil {
			log.Fatal(err)
		}
		Start(mem)
		err = http.ListenAndServe(":2345", nil)
		if err != nil {
			fmt.Errorf(err.Error())
		}
	*/

	f, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		defer f.Close()
		log.SetOutput(f)
	}

	flags := ParseFlags()
	lim := 0
	offset := 68100
	if flags.AllGames {
		processAllGames()
	} else if flags.GameDetail {
		updateGamesDetails(0, 0)
	} else if flags.LoadData {
		downloadGamesData(lim, offset)
	}

}
