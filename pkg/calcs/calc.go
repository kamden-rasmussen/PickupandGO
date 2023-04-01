package calcs

import (
	"log"

	"github.com/Kamden-Rasmussen/PickupandGO/pkg/data"
	"github.com/Kamden-Rasmussen/PickupandGO/pkg/email"
	"github.com/Kamden-Rasmussen/PickupandGO/pkg/mydatabase"
)

var ForceBool = false

func BeginCalc() {
	log.Print("Starting Calculations")

	dates := GetDates()
	log.Println(dates)

	// get all users
	users := data.GetAllUsers()
	log.Println(users)

	// Get data for each user
	for _, user := range users {

		// get data per destination 
		destinations := data.GetDestinationsForUser(user.Id)
		for _, destination := range destinations {
			log.Println(destination)

			// get price for each destination
			prices := GetPricesByDate(dates, destination)

			if ForceBool{
				email.SendEmail(user, prices, destination)
				continue
			} else if CheckAllPrices(prices) {
				log.Println("Price is good")
				email.SendEmail(user, prices, destination)
			}
			// TODO: start calcing
		}
	}
}

func GetPricesByDate(dates []string, destination int) []float64 {
	// get price for each destination
	var prices []float64
	for _, date := range dates {

		rows, err := mydatabase.MyDB.Query("SELECT price FROM flights WHERE departure_date = ? AND arrival_location = ? order by price asc limit 1", date, destination)
		if err != nil {
			panic(err.Error())
		}
		var price float64
		for rows.Next() {
			err = rows.Scan(&price)
			if err != nil {
				panic(err.Error())
			}
		}
		prices = append(prices, price)
	}
	return prices
}

func CheckPrice(currentPrice float64, checkingPrice float64) bool {
	return currentPrice < ( checkingPrice * .75 )
}

func CheckAllPrices(prices []float64) bool {
	for _, price := range prices {
		if !CheckPrice(prices[0], price) {
			return false
		}
	}
	return true
}

func CronBeginCalc() {
	log.Println("Starting Calcs")
	BeginCalc()
	log.Println("Finished Calcs")
}