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

	tags := map[string]string{
		"unit": a.Unit,
		"type": "activity",
	}
	fields := map[string]interface{}{
		"value": a.Value,
	}

	bps, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  database.Dbname,
		Precision: "ms",
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
