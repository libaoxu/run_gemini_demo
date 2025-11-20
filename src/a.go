package src

import "fmt"

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func getConfig() (*Response, error) {
	return nil, nil
}

func main() {
	var resp, err = getConfig()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(resp.Code)
	fmt.Println(resp.Message)
}
