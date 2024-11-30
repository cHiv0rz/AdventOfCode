package main

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
        log.SetFormatter(&log.TextFormatter{
                FullTimestamp:   true,
                TimestampFormat: "2006-01-02 15:04:05",
        })
        log.SetOutput(os.Stdout)

        data, err := os.Open("tu_archivo.txt")
        if err != nil {
                log.WithError(err).Error("Error al leer el archivo")
                return
        }

        log.Info(data.Name())
}