package logproxy

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"gopkg.in/sakura-internet/go-rison.v3"
)

const kibanaFormat = "https://logs.adeo.no/app/kibana#/discover?_a=%s&_g=%s"
const searchQueryV0 = "+x_delivery_id:\"%s\" -level:\"Trace\""
const searchQueryV1 = "+x_correlation_id:\"%s\" -level:\"Trace\" -level:\"Debug\""

type query struct {
	Language string `json:"language"`
	Query    string `json:"query"`
}

type appState struct {
	Index string `json:"index"`
	Query query  `json:"query"`
}

type timeRange struct {
	From string `json:"from"`
	Mode string `json:"mode"`
	To   string `json:"to"`
}

type globalState struct {
	Time timeRange `json:"time"`
}

var uuidRegex = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$")

func oneDay(ts time.Time) timeRange {
	od := time.Hour * 24
	start := ts.Truncate(od)
	end := start.Add(od)

	startStr, _ := start.MarshalText()
	endStr, _ := end.MarshalText()

	return timeRange{
		From: string(startStr),
		Mode: "absolute",
		To:   string(endStr),
	}
}

func formatKibana(deliveryID string, ts time.Time, version int) string {
	var q string

	switch version {
	default:
		q = searchQueryV0
	case 1:
		q = searchQueryV1
	}

	as := appState{
		Index: "logstash-apps-*",
		Query: query{
			Language: "lucene",
			Query:    fmt.Sprintf(q, deliveryID),
		},
	}

	gs := globalState{
		Time: oneDay(ts),
	}

	b, _ := rison.Encode(as, rison.Rison)
	c, _ := rison.Encode(gs, rison.Rison)

	return fmt.Sprintf(kibanaFormat, string(b), string(c))
}

func MakeURL(baseURL, deliveryID string, timestamp time.Time) string {
	return fmt.Sprintf("%s/logs?delivery_id=%s&ts=%d&v=1", baseURL, deliveryID, timestamp.Unix())
}

func HandleFunc(w http.ResponseWriter, r *http.Request) {
	badrequest := func(err error) {
		w.WriteHeader(http.StatusBadRequest)
		log.Error(err)
		_, err = w.Write([]byte(err.Error() + "\n"))
		if err != nil {
			log.Errorf("unable to answer http request: %s", err)
		}
	}
	deliveryID := r.URL.Query().Get("delivery_id")
	timestamp := r.URL.Query().Get("ts")
	sversion := r.URL.Query().Get("v")
	version, _ := strconv.Atoi(sversion)

	if !uuidRegex.MatchString(deliveryID) {
		badrequest(fmt.Errorf("delivery_id '%s' is not a well-formed UUID", deliveryID))
		return
	}

	switch version {
	case 0:
	case 1:
	default:
		badrequest(fmt.Errorf("version '%s' is not supported", deliveryID))
		return
	}

	unixtime, err := strconv.Atoi(timestamp)
	if err != nil {
		badrequest(fmt.Errorf("ts '%s' is not a well-formed unix timestamp: %s", timestamp, err))
		return
	}

	utctime := time.Unix(int64(unixtime), 0).UTC()
	url := formatKibana(deliveryID, utctime, version)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
