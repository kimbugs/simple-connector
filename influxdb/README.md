# INFLUXDB

Simple Influxdb package

## Usage

This simple example code. Close() function is used for finish processe.

```golang
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kimbugs/simple-connector/influxdb"
)

const (
	url    = "localhost:8086"
	bucket = "test_bucket"
	org    = "your organization"
	token  = "your token"
)

func main() {
	c, err := influxdb.NewClient(url, bucket, org, token)

	if err != nil {
		fmt.Println("influxdb not connected!")
		panic(err.Error())
	}

	// Ensures background processes finishes
	defer c.Close()

	for {
		t := time.NewTicker(1 * time.Second)
		for tc := range t.C {
			fmt.Println(tc.String())
			sec := tc.Second()
			i := 1000
			measurement := "test_measurement"
			tags := map[string]string{
				"test_id":   fmt.Sprintf("rack_%v", sec%3),
				"test_host": fmt.Sprintf("host_%v", i),
			}

			fields := map[string]interface{}{
				"temperature": rand.Float32() * 80.0,
				"humidity":    rand.Float32() * 40,
				"abc":         rand.Float32() * 80.0,
				"def":         rand.Float32() * 40,
			}

			c.Write(measurement, tags, fields)
		}
	}
}
```
