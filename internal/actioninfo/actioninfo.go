package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, str := range dataset {
		if err := dp.Parse(str); err != nil {
			log.Println(err.Error())
			continue
		}

		line, err := dp.ActionInfo()
		if err != nil {
			log.Println(err.Error())
			continue
		}

		fmt.Println(line)
	}
}
