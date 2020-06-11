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
		count int
		err   error
	)

	switch len(os.Args) {
	case 1:
		count = 1
	case 2:
		count, err = strconv.Atoi(os.Args[1])
		if err != nil {
			logrus.Fatalf("invalid count '%s'", os.Args[1])
		}
	}

	for i := 0; i < count; i++ {
		fmt.Printf("%s\n", uuid.NewV4())
	}
}
