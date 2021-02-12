package teamwork

import (
	"fmt"
	"os"
)

func ExampleConnect() {
	baseURL := ""
	apiToken := ""
	conn, err := Connect(baseURL, apiToken)
	if err != nil {
		fmt.Printf("Error connecting to TeamWork: %s", err.Error())
		os.Exit(1)
	}
	fmt.Printf("%+v", conn)
}
