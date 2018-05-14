package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/splunk/ssc-client-go/model"
	"github.com/splunk/ssc-client-go/service"
)

func main() {
	var token = "eyJraWQiOiJTVGR3WXFCVnFQZlpkeXNNUTQyOElEWTQ5VzRZQzN5MzR2ajNxSl9LRjlvIiwiYWxnIjoiUlMyNTYifQ.eyJ2ZXIiOjEsImp0aSI6IkFULlY0VVYzR3ZzakhPdzg0T2w1N1dPN1hLQVFjWEJQOTVLU3V6RHJoNWh1SFEiLCJpc3MiOiJodHRwczovL3NwbHVuay1jaWFtLm9rdGEuY29tL29hdXRoMi9kZWZhdWx0IiwiYXVkIjoiYXBpOi8vZGVmYXVsdCIsImlhdCI6MTUyNTk4OTI4MCwiZXhwIjoxNTI1OTkyODgwLCJjaWQiOiIwb2FwYmcyem1MYW1wV2daNDJwNiIsInVpZCI6IjAwdXpsMHdlZFdxM2tvWEFDMnA2Iiwic2NwIjpbImVtYWlsIiwib3BlbmlkIiwicHJvZmlsZSJdLCJzdWIiOiJ4Y2hlbmdAc3BsdW5rLmNvbSJ9.jcoLF3FoqNNPVzQQ7yozTDbj6wVcAv0UEUCopvx_VU0DLHzXwJhIqADAL5qr1qlGfkdtgGWDuxkpsxEMIj2SmvX7aSUDGJ95c8mElx-WftOxJH7JRjrqufjFWb5ZeyYNDY7YVI2_9ojV2ED9MfCIU2WcPsgqjeWoKyEC9jB3tj1KWf5J9YlOy2E2VxIou9MCYxHqPtP8hCcFcmSra0LsJwoCMd0MuJdQtu3r8XXqXnc0PHht6dFH512kMVbariU8erCC-JeUoHfJPecT1S-c4vMpFTvxaM6YFqjfDjSCOsv_3tU3JyQrw5GvDAywmr63lYgFBa0sB2BOGSl31FtRtw"
	var tenantID = "e5d62915-23fc-4434-8d26-cb42062c4c1c"
	var url = "https://next.splunknovadev-playground.com:443"
	finish := make(chan bool)
	client := service.NewClient(tenantID, token, url, time.Duration(10)*time.Second)
	batchSender, _ := client.NewBatchEventsSender(3, 5000)
	batchSender.Run()
	server8001 := http.NewServeMux()
	server8001.HandleFunc("/foo", func(writer http.ResponseWriter, request *http.Request) {
		fooHandler(writer, request, batchSender)
	})
	go func() {
		http.ListenAndServe(":8001", server8001)
	}()
	<-finish
	batchSender.Stop()
}

func fooHandler(w http.ResponseWriter, r *http.Request, bs *service.BatchEventsSender) {
	w.Write([]byte("Listening on 8001: foo "))
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	event := model.HecEvent{Host: "xcheng", Index: "main", Sourcetype: "myserver:example", Source: "go", Event: bodyString}
	fmt.Printf("%v", event.Event)
	bs.AddEvent(event)
}
