package risk

import (
	"fmt"
)

func Text() {
	fmt.Println("hello")
}

type RiskData struct {
	rawData       map[string]interface{}
	formattedData map[string]interface{}
}
