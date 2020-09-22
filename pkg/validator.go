package pkg

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

type validator struct {
	client    v1.API
	startTime time.Time
	values    map[string][]*model.SampleStream
	out       *os.File
}

type logRt struct {
	transport http.RoundTripper
}

const (
	gaugeStr   = "gauge"
	counterStr = "counter"
	histStr    = "hist"
)

func (rt *logRt) RoundTrip(r *http.Request) (*http.Response, error) {
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	log.Println(bodyString)
	return rt.transport.RoundTrip(r)
}

func NewValidator(address string) (*validator, error) {
	f, err := os.Create("./answer.txt")
	check(err)
	logRt := &logRt{transport: api.DefaultRoundTripper}
	c, err := api.NewClient(
		api.Config{
			Address:      address,
			RoundTripper: logRt,
		})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	v := &validator{
		client:    v1.NewAPI(c),
		values:    map[string][]*model.SampleStream{},
		startTime: time.Now().Add(-10 * time.Minute),
		out:       f,
	}
	return v, nil
}

func (v *validator) ValidateAndFetch(filePath string) {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")
		params := strings.Split(line, ",")
		// validate response
		if strings.HasPrefix(line, gaugeStr) || strings.HasPrefix(line, counterStr) || strings.HasPrefix(line, histStr) {
			labels := strings.Split(params[2], ",")
			matcher := ""
			for _, lb := range labels {
				parts := strings.Split(lb, ":")
				matcher += parts[0] + "="
				matcher += "\"" + parts[1] + "\""
				matcher += ","
			}
			matcher += "}"
			query := strings.Trim(params[1], " ") + strings.Trim(matcher, " ")
			v.loadQuery(query)

			v.writeOne(params[0], query)
		}
		break
	}
}
