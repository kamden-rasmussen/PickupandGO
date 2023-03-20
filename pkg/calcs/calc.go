package calcs

import (
	"log"

	"github.com/Kamden-Rasmussen/PickupandGO/pkg/data"
	"github.com/Kamden-Rasmussen/PickupandGO/pkg/mydatabase"
)



func BeginCalc(){
	log.Print("Starting Calculations")
	
	oneWeek := GetOneWeekAgo()
	twoWeeks := GetTwoWeeksAgo()
	oneMonth := GetOneMonthAgo()
	threeMonths := GetThreeMonthsAgo()
	sixMonths := GetSixMonthsAgo()
	oneYear := GetOneYearAgo()
	log.Println(oneWeek, twoWeeks, oneMonth, threeMonths, sixMonths, oneYear)
	

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
			oneWeekPrice := GetPriceByDate(oneWeek, destination)
			twoWeeksPrice := GetPriceByDate(twoWeeks, destination)
			oneMonthPrice := GetPriceByDate(oneMonth, destination)
			threeMonthsPrice := GetPriceByDate(threeMonths, destination)
			sixMonthsPrice := GetPriceByDate(sixMonths, destination)
			oneYearPrice := GetPriceByDate(oneYear, destination)
			log.Println(oneWeekPrice, twoWeeksPrice, oneMonthPrice, threeMonthsPrice, sixMonthsPrice, oneYearPrice)
			// TODO: start calcing
		}
	}
}

func GetPriceByDate(date string, destination int) float32 {
	// get price for each destination
	rows, err := mydatabase.MyDB.Query("SELECT * FROM flights WHERE departure_date = ? AND arrival_location = ? order by price asc limit 1", date, destination)
	if err != nil {
		panic(err.Error())
	}
	var price float32
	for rows.Next() {
		err = rows.Scan(&price)
		if err != nil {
			panic(err.Error())
		}
	}
	return price
}