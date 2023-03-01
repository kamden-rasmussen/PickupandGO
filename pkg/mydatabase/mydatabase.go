package mydatabase

import (
	"database/sql"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

// global variable for database
var MyDB *sql.DB

func Init() (err error) {
	// connect to a database
	config := mysql.Config{
        User:      os.Getenv("DBACCESSUSERNAME"),
		Passwd:    os.Getenv("DBACCESSPASSWORD"),
		Net:       "tcp",
		Addr:      os.Getenv("DBHOST"),
		DBName:    os.Getenv("DBDATABASENAME"),
    }
	MyDB, err = sql.Open("mysql", config.FormatDSN())
	pingErr := MyDB.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
	// defer MyDB.Close()

	log.Println("Connected to database")

	return err

}

func SetUpTables(db *sql.DB) {

	// remove security for now
	_, err := db.Exec("SET SESSION sql_require_primary_key = 0;")
	if err != nil {
		log.Fatal(err.Error())
	}

	query := os.Getenv("CREATE_USERS_TABLE")
	// create users table
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}
	
	// create airports table
	query = os.Getenv("CREATE_AIRPORTS_TABLE")

	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	// create destinations table
	query = os.Getenv("CREATE_DESTINATIONS_TABLE")
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	// create companies table
	query = os.Getenv("CREATE_COMPANIES_TABLE")
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	// create airlines table
	query = os.Getenv("CREATE_AIRLINES_TABLE")
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}

	// create flights table
	query = os.Getenv("CREATE_FLIGHTS_TABLE")
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetFunTable(db *sql.DB) {
	results, err := db.Query("SELECT * FROM fun")
	if err != nil {
		log.Fatal(err.Error())
	}
	for results.Next() {
		var id int
		var name string
		var color string
		err = results.Scan(&id, &name, &color)
		if err != nil {
			log.Fatal(err.Error())
		}
		// log.Println(id, name, color)
	}	
}

func DbHealthCheck() error {
	return MyDB.Ping()
}

func GetDB() *sql.DB {
	return MyDB
}