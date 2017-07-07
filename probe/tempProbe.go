package probe

import (
	"strconv"
	"time"

	"github.com/dlaize/homedatakeeper/database"
	client "github.com/influxdata/influxdb/client/v2"
)

// Probe Struct for probe object
type TempProbe struct {
	Name  string  `json:"name"`
	Temp  float64 `json:"temp"`
	Hygro float64 `json:"hygro"`
	Floor int     `json:"floor"`
}

// Add probe information to database
func (p *TempProbe) createTempProbe() error {

	tags := map[string]string{
		"floor": strconv.Itoa(p.Floor),
		"type":  "temp",
		"room":  p.Name,
	}
	fields := map[string]interface{}{
		"temp":  p.Temp,
		"hygro": p.Hygro,
	}

	bps, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  database.Dbname,
		Precision: "ms",
	})

	if err != nil {
		return err
	}

	point, err := client.NewPoint(
		"probe",
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
