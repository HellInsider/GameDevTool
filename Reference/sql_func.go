package main

import (
	"Client/SQLFunctions"
	InputDataModel "Client/models/InputDataModels"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"strconv"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "secure-password"
	dbname   = "connect-db"
)

const SqlConnect = "user=GameDevWorker password=password dbname=GameDevDatabase sslmode=disable"

func WriteBaseGameInfo(games InputDataModel.AllGamesRequest) error {
	db, err := sql.Open("postgres", SqlConnect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

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

	return err
}

func GetBaseGamesData(lim, offset int) (error, []InputDataModel.App) {
	db, err := sql.Open("postgres", SqlConnect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

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

func WriteGameDetails(details InputDataModel.AppDetailsRequest) error {
	db, err := sql.Open("postgres", SqlConnect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

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
	_, err = tx.Exec(`select addPrice($1,$2,$3,$4,$5,$6)`, details.Data.Price.AppId, details.Data.Price.Currency,
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
		pq.Array(details.Data.Packages), details.Data.Recommendations.Total, details.Data.AppType)

	if err != nil {
		panic(err)
		return fmt.Errorf("wrong data: %s", err)
	}

	return tx.Commit()
}

/*
func SetDataToDB(data model.Order) error {
	db, err := sql.Open("postgres", SqlConnect)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		log.Println(fmt.Errorf("transaction error %s", err))
	}
	defer tx.Rollback()

	_, err = tx.Exec("select adddeliverydata($1,$2,$3,$4,$5,$6,$7)", data.Delivery.Name,
		data.Delivery.Phone, data.Delivery.Zip, data.Delivery.City, data.Delivery.Address, data.Delivery.Region,
		data.Delivery.Email)
	if err != nil {
		return fmt.Errorf("wrong data: %s", err)
	}
	//fmt.Println(result.RowsAffected())

	_, err = tx.Exec("select addpaymentdata($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)", data.Payment.Transaction,
		data.Payment.RequestId, data.Payment.Currency, data.Payment.Provider, data.Payment.Amount,
		data.Payment.PaymentDt, data.Payment.Bank, data.Payment.DeliveryCost, data.Payment.GoodsTotal,
		data.Payment.CustomFee)
	if err != nil {
		return fmt.Errorf("wrong data: %s", err)
	}
	//fmt.Println(result.RowsAffected())

	_, err = tx.Exec("select addorderdata($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)", data.OrderUid,
		data.TrackNumber, data.Entry, data.Payment.Transaction, data.Locale, data.InternalSignature,
		data.CustomerId, data.DeliveryService, data.Shardkey, data.SmId, data.DateCreated, data.OofShard)
	if err != nil {
		return fmt.Errorf("wrong data: %s", err)
	}
	//fmt.Println(result.RowsAffected())

	for i, _ := range data.Items {
		_, err = tx.Exec("select additemdata($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)", data.OrderUid,
			data.Items[i].ChrtId, data.Items[i].TrackNumber, data.Items[i].Price, data.Items[i].Rid, data.Items[i].Name,
			data.Items[i].Sale, data.Items[i].Size, data.Items[i].TotalPrice, data.Items[i].NmId,
			data.Items[i].Brand, data.Items[i].Status)
		if err != nil {
			return fmt.Errorf("wrong data: %s", err)
		}
		//fmt.Println(result.RowsAffected())
	}

	return tx.Commit()
}

func GetDataFromDB(cache *model.Cashe) error {
	db, err := sql.Open("postgres", SqlConnect)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	orderRows, err := db.Query(`select * from "order"`)
	if err != nil {
		return fmt.Errorf("error getting rows from orders: %s", err)
	}
	defer orderRows.Close()

	var data model.Order
	var paymentId string
	var deliveryId string
	for orderRows.Next() {
		err = orderRows.Scan(&data.OrderUid, &data.TrackNumber, &data.Entry, &deliveryId, &paymentId, &data.Locale,
			&data.InternalSignature, &data.CustomerId, &data.DeliveryService, &data.Shardkey, &data.SmId, &data.DateCreated,
			&data.OofShard)
		if err != nil {
			return fmt.Errorf("error reading order from db: %s", err)
		}

		deliveryRows, err := db.Query("select * from delivery where delivery.id = $1", deliveryId)
		if err != nil {
			return fmt.Errorf("error getting rows from delivery: %s", err)
		}

		deliveryRows.Next()
		err = deliveryRows.Scan(&data.Delivery.Phone, &data.Delivery.Zip, &data.Delivery.City, &data.Delivery.Address,
			&data.Delivery.Region, &data.Delivery.Email, &data.Delivery.Name, &deliveryId)
		if err != nil {
			return fmt.Errorf("error reading delivery from db: %s", err)
		}
		deliveryRows.Close()

		paymentRows, err := db.Query("select * from payment where payment.transaction = $1", paymentId)
		if err != nil {
			return fmt.Errorf("error getting rows from payment: %s", err)
		}

		paymentRows.Next()
		err = paymentRows.Scan(&data.Payment.Transaction, &data.Payment.RequestId, &data.Payment.Currency,
			&data.Payment.Provider, &data.Payment.Amount, &data.Payment.PaymentDt, &data.Payment.Bank,
			&data.Payment.DeliveryCost, &data.Payment.GoodsTotal, &data.Payment.CustomFee)
		if err != nil {
			return fmt.Errorf("error reading payment from db: %s", err)
		}
		paymentRows.Close()

		itemsRows, err := db.Query("select chrt_id, track_number, price, rid, name, sale, size,total_price, "+
			"nm_id, brand, status from items where items.order_uid = $1", data.OrderUid)
		if err != nil {
			return fmt.Errorf("error getting rows from items: %s", err)
		}
		data.Items = []model.Items{}
		for itemsRows.Next() {
			item := model.Items{}
			err = itemsRows.Scan(&item.ChrtId, &item.TrackNumber, &item.Price, &item.Rid, &item.Name,
				&item.Sale, &item.Size, &item.TotalPrice, &item.NmId, &item.Brand, &item.Status)
			if err != nil {
				return fmt.Errorf("error reading item from db: %s", err)
			}
			data.Items = append(data.Items, item)
		}
		itemsRows.Close()
		fmt.Println(data.OrderUid)
		cache.Lock()
		cache.Memory[data.OrderUid] = data
		cache.Unlock()
	}
	return nil
}
*/
