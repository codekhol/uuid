package main

import (
	"fmt"
	"os"
	"strconv"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

func main() {
	var (
		qty int
		err error
	)

	switch len(os.Args) {
	case 1:
		qty = 1
	case 2:
		qty, err = strconv.Atoi(os.Args[1])
		if err != nil {
			logrus.Fatalf("invalid qty '%s'", os.Args[1])
		}
	default:
		logrus.Infof("usage: uuid [qty]")
		return
	}

	for i := 0; i < qty; i++ {
		fmt.Printf("%s\n", uuid.NewV4())
	}
}
