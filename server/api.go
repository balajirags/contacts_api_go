package server

import (
	"strconv"
	"github.com/contacts_api_go/config"
	"net/http"
	"github.com/contacts_api_go/logger"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"fmt"
	"strings"
	"github.com/contacts_api_go/appcontext"
	statsdv2 "gopkg.in/alexcesaro/statsd.v2"
	"time"
	"github.com/sirupsen/logrus"
	"gopkg.in/tylerb/graceful.v1"
)

func StartServer() {
	router := InitRouter()
	handlerFunc := router.ServeHTTP
	n := negroni.New(negroni.NewRecovery())
	n.Use(httpStatLogger())
	n.Use(StatsDMiddlewareLogger())
	n.UseHandlerFunc(handlerFunc)

	port := strconv.Itoa(config.GetAppPort())
	server := &graceful.Server{
		Timeout: 0,
		Server: &http.Server{
			Addr:    ":"+port,
			Handler: n,
		},
	}
	err := server.ListenAndServe()
	if err != http.ErrServerClosed {
		logger.Log.Fatal(err.Error())
	}
}
func httpStatLogger() negroni.Handler {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		startTime := time.Now()
		next(rw, r)
		responseTime := time.Now()
		deltaTime := responseTime.Sub(startTime).Seconds() * 1000
		logger.Log.WithFields(logrus.Fields{
			"RequestTime":   startTime.Format(time.RFC3339),
			"ResponseTime":  responseTime.Format(time.RFC3339),
			"DeltaTime":     deltaTime,
			"RequestUrl":    r.URL.Path,
			"RequestMethod": r.Method,
			"RequestProxy":  r.RemoteAddr,
		}).Info("request")
	})
}
func StatsDMiddlewareLogger() negroni.HandlerFunc {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		vars := mux.Vars(r)
		path := r.URL.Path
		fmt.Println(vars)
		for _, v := range vars {
			fmt.Println(v)
			path = strings.Replace(path, v, "", len(path))
		}
		fmt.Println("ch", path)
		key := GetKeyStructure(r.URL.Path)
		t := TimingInStatsD()
		next(rw, r)
		SendInStatsD(key+".time", t)
		IncrementInStatsD(key + ".calls")
	})
}


func IncrementInStatsD(s string) {
	t := appcontext.GetStatsDClient()
	t.Increment(s)

}
func SendInStatsD(s string, timing *statsdv2.Timing) {
	timing.Send(s)
}

func TimingInStatsD() *statsdv2.Timing {
	t := appcontext.GetStatsDClient().NewTiming()
	return &t
}

func GetKeyStructure(url string) string {
	baseKey := "go.response"
	basePath := strings.Split(url, "/GF")[0]
	keyBasePath := strings.Replace(basePath, "/", ".", len(basePath))
	key := baseKey + keyBasePath
	return key
}
