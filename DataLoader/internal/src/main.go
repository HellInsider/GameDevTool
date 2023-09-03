package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	f, err := os.OpenFile("logs.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err == nil {
		defer f.Close()
		log.SetOutput(f)
	}

	flags := ParseFlags()
	log.Info("Flags parsed")
	log.Info("Starting update cycle")
	for {
		updateGamesList()
		updateGamesDetails(flags.Limit, flags.Offset, flags.RPM)
	}

	/*	if flags.AllGames {
			pushAllGamesNames()
		} else if flags.GameDetail {
			updateGamesDetails(0, 0)
		} else if flags.LoadData {
			downloadGamesData(lim, offset)
		}
	*/
}
