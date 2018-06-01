// Personal Access Token:
// URL: https://api.github.com/repos/Hendrix5/test/stats/commit_activity
//resp.Header.Set("Authorization", "token: ")

package main

import (
	"fmt"
	"log"
	"net/http"
)

//creates a struct to encapsulate the json payload
type commit_activity struct {
	Days  string `json: "days"`
	Total string `json: "total"`
	Week  string `json: "week"`
}

func main() {
	var err error
	var u commit_activity

	// request github api
	req, err := http.NewRequest("GET", "https://api.github.com/repos/Hendrix5/test/stats/commit_activity", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("", "x-oauth-basic")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("StatusCode:", res.StatusCode)

	// read body
	/*body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&u)
	if err != nil {
		log.Fatal(err)
	} */

	fmt.Printf("%+v\n", u)
}
