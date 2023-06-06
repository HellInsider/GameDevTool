package main

import (
	"../SQLFunctions"
	InputDataModel "../models/InputDataModels"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "secure-password"
	dbname   = "connect-db"
)

// const SqlConnect = "user=GameDevWorker password=password dbname=GameDevDatabase sslmode=disable"
const SqlConnect = "user=postgres password=admin dbname=GameDevDatabase sslmode=disable"

func WriteBaseGameInfo(games InputDataModel.AllGamesRequest, db *sql.DB) error {
	for _, app := range games.Apps {
		if app.Name != "" {
			tx, err := db.Begin()
			if err != nil {
				log.Println(fmt.Errorf("transaction error %s", err))
				panic(err)
			}
			defer tx.Rollback()

			_, err = tx.Exec(SQLFunctions.AddBaseGameInfo, app.AppId, app.Name)
			if err != nil {
				panic(err)
				return fmt.Errorf("wrong data: %s", err)
			}

			err = tx.Commit()
			if err != nil {
				panic(err)
				return fmt.Errorf("wrong data: %s", err)
			}
		}
	}

	return nil
}

func GetBaseGamesData(lim, offset int, db *sql.DB) (error, []InputDataModel.App) {
	tx, err := db.Begin()
	if err != nil {
		log.Println(fmt.Errorf("transaction error %s", err))
		panic(err)
	}
	defer tx.Rollback()
	var orderRows *sql.Rows
	if lim != 0 {
		orderRows, err = db.Query(`select appid, name from "DB_schema"."Games" limit $1 offset $2`, lim, offset)
		if err != nil {
			return fmt.Errorf("error getting rows from games: %s", err), nil
		}
	} else {
		orderRows, err = db.Query(`select appid, name from "DB_schema"."Games"`)
		if err != nil {
			return fmt.Errorf("error getting rows from games: %s", err), nil
		}
	}

	defer orderRows.Close()

	var data InputDataModel.App
	var res []InputDataModel.App
	for orderRows.Next() {
		err = orderRows.Scan(&data.AppId, &data.Name)
		if err != nil {
			panic(err)
		}
		//fmt.Printf("%+v", data)
		res = append(res, data)
	}

	return err, res
}

func WriteGameDetails(details InputDataModel.AppDetailsRequest, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		log.Println(fmt.Errorf("transaction error %s", err))
		panic(err)
	}
	defer tx.Rollback()

	//fmt.Println("	Price insertion")
	/*_, err = tx.Exec(SQLFunctions.AddPrice, details.Data.Price.AppId, details.Data.Price.Currency,
	details.Data.Price.Initial, details.Data.Price.Final, details.Data.Price.DiscountPercent,
	details.Data.Price.Country)
	*/

	_, err = tx.Exec(`select "DB_schema"."addPrice"($1,$2,$3,$4,$5,$6)`, details.Data.Price.AppId, details.Data.Price.Currency,
		details.Data.Price.Initial, details.Data.Price.Final, details.Data.Price.DiscountPercent,
		details.Data.Price.Country)
	if err != nil {
		panic(err)
		return fmt.Errorf("wrong data: %s", err)
	}

	//fmt.Println("	Critic insertion")
	_, err = tx.Exec(SQLFunctions.AddCriticScore, details.Data.Critic.AppId, details.Data.Critic.CriticScore,
		details.Data.Critic.UsersScore, details.Data.Critic.Url)
	if err != nil {
		panic(err)
		return fmt.Errorf("wrong data: %s", err)
	}

	//fmt.Println("	Platforms insertion")
	_, err = tx.Exec(SQLFunctions.AddPlatforms, details.Data.Platform.AppId, details.Data.Platform.PlatformWin,
		details.Data.Platform.PlatformMac, details.Data.Platform.PlatformLinux)
	if err != nil {
		panic(err)
		return fmt.Errorf("wrong data: %s", err)
	}

	//fmt.Println("	Release date insertion")
	_, err = tx.Exec(SQLFunctions.AddReleaseDate, details.Data.ReleaseDate.AppId,
		details.Data.ReleaseDate.CommingSoon, details.Data.ReleaseDate.Date)
	if err != nil {
		panic(err)
		return fmt.Errorf("wrong data: %s", err)
	}

	//fmt.Println("	Requirements insertion")
	_, err = tx.Exec(SQLFunctions.AddRequirements, details.Data.PCRequirements.AppId,
		details.Data.PCRequirements.Minimum, details.Data.PCRequirements.Recommended)

	if err != nil {
		panic(err)
		return fmt.Errorf("wrong data: %s", err)
	}

	//fmt.Println("	Categories insertion")
	for _, e := range details.Data.Categories {
		//_, err = tx.Exec(SQLFunctions.AddCategory, e.CategoryId, e.Description)
		_, err = tx.Exec(`select addCategory($1,$2,$3)`, e.CategoryId, e.Description, details.AppId)
		if err != nil {
			panic(err)
			return fmt.Errorf("wrong data: %s", err)
		}
	}

	//fmt.Println("	Genres insertion")
	for _, e := range details.Data.Genres {
		num, _err := strconv.Atoi(e.GenreId)
		if _err != nil {
			panic(err)
			return fmt.Errorf("wrong data: %s", err)
		}
		//_, err = tx.Exec(SQLFunctions.AddGenre, num, e.Description)
		_, err = tx.Exec(`select addGenre($1,$2,$3)`, num, e.Description, details.AppId)
		if err != nil {
			panic(err)
			return fmt.Errorf("wrong data: %s", err)
		}
	}

	//fmt.Println("	Details insertion")
	age, err := strconv.Atoi(details.Data.RequiredAge)
	if err != nil {
		log.Info(details.AppId, ": ", err)
	}
	_, err = tx.Exec(SQLFunctions.UpdateGameDetails, details.AppId, age,
		details.Data.IsFree, pq.Array(details.Data.DLC), details.Data.AboutTheGame, details.Data.DetailedDescription,
		details.Data.ShortDescription, pq.Array(details.Data.Developers), pq.Array(details.Data.Publishers),
		pq.Array(details.Data.Packages), details.Data.Recommendations.Total, details.Data.AppType, time.Now())

	if err != nil {
		panic(err)
		return fmt.Errorf("wrong data: %s", err)
	}

	return tx.Commit()
}

func WriteGameReviews(review InputDataModel.Review, db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		log.Println(fmt.Errorf("transaction error %s", err))
		panic(err)
	}
	defer tx.Rollback()

	_, err = tx.Exec(SQLFunctions.AddReviews, review.AppId, review.QuerySummary.TotalReviews,
		review.QuerySummary.TotalPositive, review.QuerySummary.TotalNegative, review.QuerySummary.ReviewScore,
		review.QuerySummary.ReviewScoreDesc)
	if err != nil {
		panic(err)
		return fmt.Errorf("wrong data: %s", err)
	}

	return tx.Commit()
}

func OpenBD() (*sql.DB, error) {
	db, err := sql.Open("postgres", SqlConnect)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	return db, err
}
