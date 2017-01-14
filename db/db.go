package db

import (
	"os"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
)

const DB = "catchphrase"

func CreateSession() (*mgo.Session, error) {
	addrs := []string{"127.0.0.1"}
	if mu := os.Getenv("MONGO_URL"); mu != "" {
		addrs = strings.Split(mu, ",")
	}
	source := os.Getenv("MONGO_ADMIN")
	if source == "" {
		source = "admin"
	}
	info := &mgo.DialInfo{
		Addrs:    addrs,
		Username: os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PASS"),
		Database: DB,
		Source:   source,
		Timeout:  time.Second,
	}
	return mgo.DialWithInfo(info)
}
