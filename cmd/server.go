package main

import (
	constants "assignment_2"
	"assignment_2/handlers"
	"net/http"
)

func main() {
	http.HandleFunc(constants.DEFAULT_PATH, handlers.DefaultHandler)
	http.HandleFunc(constants.CASES_PATH, handlers.CasesHandler)
	http.HandleFunc(constants.POLICY_PATH, handlers.PolicyHandler)
	http.HandleFunc(constants.STATUS_PATH, handlers.StatusHandler)
	http.HandleFunc(constants.NOTIFICATION_PATH, handlers.NotificationHandler)
	http.ListenAndServe(":8080", nil)
}
