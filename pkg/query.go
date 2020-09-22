package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/ishangupta-ds/plot/pkg/types"
)

func (v *validator) loadQuery(query string) {

	//retrieve response
	if _, found := v.values[query]; !found {
		u, err := url.Parse("http://aed80e4edc6a54c6a968bd62d4d1d599-76265430.us-east-2.elb.amazonaws.com:9090/api/v1/query_range")
		if err != nil {
			log.Println(err)
			return
		}
		q := u.Query()
		//q.Add("query", `heptio_eventrouter_normal_total{reason="ChaosInject",involved_object_name="catalogue-pod-cpu-hog", involved_object_namespace="litmus", involved_object_kind="ChaosEngine"} - on () (heptio_eventrouter_normal_total{reason="ChaosEngineCompleted",involved_object_name="catalogue-pod-cpu-hog", involved_object_namespace="litmus", involved_object_kind="ChaosEngine"} OR on() vector(0))`)
		q.Add("query", `chaosengine_experiments_count{engine_name="catalogue-node-cpu-hog"}`)
		//q.Add("query", query)
		q.Add("start", strconv.Itoa(int(v.startTime.Unix())))
		q.Add("end", strconv.Itoa(int(time.Now().Unix())))
		q.Add("step", "15")
		u.RawQuery = q.Encode()
		fmt.Println("url: ", u)

		log.Println(query)

		//ur := "http://aed80e4edc6a54c6a968bd62d4d1d599-76265430.us-east-2.elb.amazonaws.com:9090/api/v1/query_range?query=chaosengine_experiments_count%7Bengine_name%3D%22catalogue-node-cpu-hog%22%2C%7D&start=1600777115&end=1600777815&step=15"

		//ur2 := "http://aed80e4edc6a54c6a968bd62d4d1d599-76265430.us-east-2.elb.amazonaws.com:9090/api/v1/query_range?end=1600784487&query=chaosengine_experiments_count%7Bengine_name%3D%22catalogue-node-cpu-hog%22%2C%7D&start=1600783887&step=15"

		resp, err := http.DefaultClient.Get(u.String())
		if err != nil {
			log.Println(err)
			return
		}
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		log.Println(bodyString)

		/*
			queryRange := v1.Range{
					Start: v.startTime,
					End:   v.startTime.Add(time.Hour),
					Step:  15* time.Second,
				}
			value, warn, err := v.client.QueryRange(context.Background(), query, queryRange)
			if err != nil {
				log.Println(err)
				return
			}
			if warn != nil {
				log.Println("warn: ", warn)
				return
			}
					log.Println(value.String())

				v.values[query] = value.(model.Matrix)*/

		var response types.ApiResponse

		json.Unmarshal(bodyBytes, &response)
		var queryResult types.QueryResult

		err = json.Unmarshal(response.Data, &queryResult)

		if err != nil {
			log.Println(err)
			return
		}

		m := queryResult.Result
		v.values[query] = m
	}
}
