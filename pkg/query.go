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
	"github.com/ishangupta-ds/plot/types"
)

func (v *validator) loadQuery(query string) {

	//retrieve response
	if _, found := v.values[query]; !found {
		u, err := url.Parse("http://aed80e4edc6a54c6a968bd62d4d1d599-76265430.us-east-2.elb.amazonaws.com:9090/api/prom/api/v1/query_range")
		if err != nil {
			log.Println(err)
			return
		}
		q := u.Query()
		q.Add("query", query)
		q.Add("start", strconv.Itoa(int(v.startTime.Unix())))
		q.Add("end", strconv.Itoa(int(time.Now().Unix())))
		q.Add("step", "15")
		u.RawQuery = q.Encode()
		fmt.Println("url: ", u)

		log.Println(query)

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

		var response apiResponse

		json.Unmarshal(bodyBytes, &response)
		var queryResult queryResult

		err = json.Unmarshal(response.Data, &queryResult)

		if err != nil {
			log.Println(err)
			return
		}

		m := queryResult.Result
		v.values[query] = m
	}
}
