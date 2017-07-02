package database

import (
	"fmt"
	"net/url"
	"os"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
)

var (
	// InfluxDBcon : Connection to infludb database
	InfluxDBcon client.Client
	// Dbname : Database name
	Dbname string
)

// Initialize DB connection
func Initialize() error {
	println("InfluxDB : Trying to Connect to database ")

	host := os.Getenv("APP_DB_HOST")
	Dbname = os.Getenv("APP_DB_NAME")
	port := os.Getenv("APP_DB_PORT")

	u, err := url.Parse(fmt.Sprintf("http://%s:%s", host, port))
	if err != nil {
		println("InfluxDB : Invalid Url,Please check domain name given in environnement variables \nError Details: ", err.Error())
		return err
	}

	conf := client.HTTPConfig{
		Addr: u.String(),
	}

	InfluxDBcon, err = client.NewHTTPClient(conf)

	if err != nil {
		println("InfluxDB : Failed to connect to Database . Please check the details entered in environnement variables\nError Details: ", err.Error())
		return err
	}

	_, ver, err := InfluxDBcon.Ping(10 * time.Second)

	if err != nil {
		println("InfluxDB : Failed to connect to Database . Please check the details entered in environnement variables\nError Details: ", err.Error())
		return err
	}

	println("InfluxDB: Successfuly connected . Version:", ver)

	return nil
}

// Close close db connection
func Close() {
	InfluxDBcon.Close()
}
