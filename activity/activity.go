package activity

import (
	"time"

	"github.com/dlaize/homedatakeeper/database"
	client "github.com/influxdata/influxdb/client/v2"
)

type Activity struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
	Unit  string  `json:"unit"`
}

// Add activity information to database
func (a *Activity) createActivity() error {

	if a.Unit == "min" {
		a.Value = a.Value * 60
	}

	tags := map[string]string{
		"unit": a.Unit,
		"type": "activity",
	}
	fields := map[string]interface{}{
		"value": a.Value,
	}

	bps, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  database.Dbname,
		Precision: "s",
	})

	if err != nil {
		return err
	}

	point, err := client.NewPoint(
		a.Name,
		tags,
		fields,
		time.Now(),
	)

	if err != nil {
		return err
	}

	bps.AddPoint(point)

	err = database.InfluxDBcon.Write(bps)

	if err != nil {
		return err
	}

	return nil
}
