package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {

	for _, data := range dataset {

		err := dp.Parse(data)
		if err != nil {
			log.Printf("Parse error for '%s': %v", data, err)
			continue
		}
	}

	result, err := dp.ActionInfo()
	if err != nil {
		log.Printf("ActionInfo error: %v", err)
		return
	}
	fmt.Println(result)
}
