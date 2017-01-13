package db

import (
	"os"
	"strings"

	"gopkg.in/mgo.v2"
)

const DB = "catchphrase"

func CreateSession() (*mgo.Session, error) {
	addrs := []string{"127.0.0.1"}
	if mu := os.Getenv("MONGO_URL"); mu != "" {
		addrs = strings.Split(mu, ",")
	}
	info := &mgo.DialInfo{
		Addrs:    addrs,
		Username: os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PASS"),
		Database: DB,
		Source:   "admin",
	}
	return mgo.DialWithInfo(info)
}
