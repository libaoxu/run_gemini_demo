package src

import "fmt"

type Response2 struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func getConfig2() (*Response2, error) {
	return nil, nil
}

func main3() {
	var resp, err = getConfig2()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(resp.Code)
	fmt.Println(resp.Message)
}
