package metadome_users

import (
	"encoding/json"
	"fmt"
	"github.com/eduhub/helper"
	"io/ioutil"
	"net/http"
)

func FetchUserDetails(token string) (*helper.UserResponse, error) {
	url := "http://localhost:8888/metadome-api/user" // Put it in a config file
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}
	fmt.Println("-->", token)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil, err
	}

	var user helper.UserResponse
	err = json.Unmarshal(body, &user)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return nil, err
	}
	return &user, nil
}
