package influxdb

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
)

// Influxdb : simple influxdb struct
type Influxdb struct {
	client   influxdb2.Client
	writeAPI api.WriteAPI
	queryAPI api.QueryAPI
}

// NewClient : simple influxdb client
// url - localhost:8086
// bucket - influxdb 2.0 bucket is like database name
// org - influxdb 2.0 organization
func NewClient(url string, bucket string, org string, token string) (
	*Influxdb, error) {
	influx := &Influxdb{}

	opts := influxdb2.DefaultOptions().
		SetMaxRetries(1).
		SetMaxRetryInterval(10).
		SetRetryInterval(5)
	influx.client = influxdb2.NewClientWithOptions("http://"+url, token, opts)
	influx.writeAPI = influx.client.WriteAPI(org, bucket)
	influx.queryAPI = influx.client.QueryAPI(org)

	if _, err := influx.client.Ready(context.Background()); err != nil {
		return nil, err
	}

	return influx, nil
}

// Close : influxdb instance close, ex) defer Close()
func (influx *Influxdb) Close() {
	influx.client.Close()
}

// QueryToBucket : influxdb Query To Bucket
func (influx *Influxdb) QueryToBucket(query string, bucket string, org string) error {
	plusQuery := fmt.Sprintf(`|> to(bucket: "%s", org: "%s")`, bucket, org)
	queryResult := query + plusQuery
	_, err := influx.queryAPI.QueryRaw(context.Background(), queryResult, influxdb2.DefaultDialect())

	return err
}

// Write : influxdb Write data
func (influx *Influxdb) Write(measurement string, tags map[string]string,
	fields map[string]interface{}) {
	p := influxdb2.NewPoint(measurement, tags, fields, time.Now())
	influx.writeAPI.WritePoint(p)
}

// ErrorCheck : check influxdb write error
// example code
// go func() {
// 	for err := range errorsCh {
// 		fmt.Printf("write error: %s\n", err.Error())
// 	}
// }()
func (influx *Influxdb) ErrorCheck() <-chan error {
	return influx.writeAPI.Errors()
}
