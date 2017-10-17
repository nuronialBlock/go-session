package main

import "go-session/session"

var globalSessions *session.Manager

func init() {
	globalSessions = NewManager("memory", "gosessionid", 3600)
}
