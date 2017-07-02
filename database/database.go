package database

// import (
// 	_ "github.com/lib/pq"
// )

// var (
// 	DB *sql.DB
// )

//Connect to PostGres DB
// func Connect(port string) {
// 	host := os.Getenv("APP_HOST")
// 	user := os.Getenv("APP_DB_USERNAME")
// 	password := os.Getenv("APP_DB_PASSWORD")
// 	dbname := os.Getenv("APP_DB_NAME")
// 	connectionString :=
// 		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

// 	var err error = nil
// 	DB, err = sql.Open("postgres", connectionString)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

// func Close() {
// 	DB.Close()
// }
