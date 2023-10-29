package config

import (
	"context"
	"io"
	"log"
	"os"
	"time"
)

var (
	info *log.Logger
	lerr *log.Logger
)

func InitLogger(location string) {

	// for local only
	file, err := os.OpenFile(location, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	// file, err := os.OpenFile("/var/log/cakap-payment/logger.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	mw := io.MultiWriter(os.Stdout, file)

	info = log.New(mw, "INFO: ", log.LstdFlags|log.Lshortfile)
	lerr = log.New(mw, "ERROR: ", log.LstdFlags|log.Lshortfile)

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(mw)
}

func Info(ctx context.Context) {
	now := time.Now()
	message, _ := MsgFromContext(ctx)
	hashcode, _ := HashKeyFromContext(ctx)
	info.Printf("%v (%s) [Info] %s\n", now, hashcode, message)

}

func Error(ctx context.Context) {
	now := time.Now()
	message, _ := MsgFromContext(ctx)
	hashcode, _ := HashKeyFromContext(ctx)
	lerr.Printf("%v (%s) [Error] %s\n", now, hashcode, message)
}