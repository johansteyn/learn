package main

import log "github.com/sirupsen/logrus"

func main() {
	//log.SetFormatter(&log.TextFormatter{})
	//log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, DisableLevelTruncation: true})
	//log.SetFormatter(&log.JSONFormatter{})

	//log.SetReportCaller(true)

	log.Info("Information about flight departures and arrivals")
	log.Warn("Bill Stickers will be prosecuted")
	log.Error("Don't DO that!")
	log.Fatal("Watch out for the big pffft... What pfft?") // Will exit with 1
}
