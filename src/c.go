package src

import "fmt"

func main2() {
	var resp, err = getConfig()
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println(resp.Code)
	fmt.Println(resp.Message)
}
