package calcs

import (
	"log"

	"github.com/Kamden-Rasmussen/PickupandGO/pkg/data"
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

}


