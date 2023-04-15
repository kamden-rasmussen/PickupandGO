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
			prices := GetPricesByDate(user, dates, destination)

			if ForceBool{
				email.SendEmail(user, prices, destination, dates)
				continue
			} else if CheckAllPrices(prices) {
				log.Println("Price is good")
				email.SendEmail(user, prices, destination, dates)
			}
		}
	}
}

func GetPricesByDate(user data.User, dates []string, destination int) []float64 {
	// get price for each destination
	var prices []float64
	for _, date := range dates {

		price := GetPriceQuery(date, destination)

		// if price == 0 {
		// 	home := data.GetAirportById(user.Home)
		// 	dest := data.GetAirportById(destination)
		// 	log.Println("Getting Data for ", home, dest, date, data.GetReturnDateWithDate(date), "2", "DL", user.Currency, "USD")
		// 	data.GetData(home, dest, date, data.GetReturnDateWithDate(date), "2", "DL", user.Currency)
		// 	price = GetPriceQuery(date, destination)
		// }
		prices = append(prices, price)
	}
	return prices
}

func GetPriceQuery(date string, destination int) float64 {
	rows, err := mydatabase.MyDB.Query("SELECT price FROM flights WHERE departure_date = ? AND arrival_location = ? order by price asc limit 1", date, destination)
	if err != nil {
		log.Fatal("error getting price: ", err)
	}
	var price float64
	for rows.Next() {
		err = rows.Scan(&price)
		if err != nil {
			log.Fatal("error scanning price: ", err)
		}
	}
	return price
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