package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Dari contoh yang diberikan, cobalah untuk mengimplementasikan sebuah http client sederhana.
// Buatlah sebuah http client dan lakukan request GET ke API https://www.metaweather.com/api/.
// Buatlah request get ke endpoint /api/location/(woeid)/(date)/ dengan nilai woeid 1047378.
// Untuk date, gunakan format YYYY/MM/dd

func main() {
	// TODO: answer here
	var (
		woeid, year, month, day int
	)

	woeid = 1047378
	year = 2022
	month = 05
	day = 13

	c := http.Client{}

	resp, err := c.Get("https://www.metaweather.com/api/location/" + strconv.Itoa(woeid) + "/" + strconv.Itoa(year) + "/" + strconv.Itoa(month) + "/" + strconv.Itoa(day))

	if err != nil {
		fmt.Printf("Errpr: %s", err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("err : %s", err)
		return
	}
	fmt.Println(string(body))

}
