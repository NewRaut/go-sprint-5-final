package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Printf("invalid data from '%s': %w", data, err)
			continue
		}
		infoAction, err := dp.ActionInfo()
		if err != nil {
			log.Printf("invalid data from '%s': %w", data, err)
			continue
		}

		fmt.Println(infoAction)
	}
}
