package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Coupon : Data structure for coupon
type Coupon struct {
	Brand string `json:"brand"`
	Value int    `json:"value"`
}

// CouponAPIResponse : An array of type Coupon
type CouponAPIResponse []Coupon

// CouponRequest : Coupon structure for sending a request
type CouponRequest struct {
	Brand string `json:"brand"`
	Value int    `json:"value"`
	Limit int    `json:"limit"`
}

func main() {
	runServer()
}

func runServer() {
	http.HandleFunc("/coupons", couponsHandler)
	http.ListenAndServe(":80", nil)
}

func couponsHandler(response http.ResponseWriter, request *http.Request) {
	/*
		Validate request - Check each parameter meets set of constraints
		If validation fails return with 422 and helpful error message per failure
		Validation rules:
		brand: ['required', 'string']
		value: ['required', 'integer', 'min:1']
		limit: ['required', 'integer', 'min:1']
	*/

	parameters := request.URL.Query()

	/*
		Authentication - Request JWT from service-b
	*/

	/*
		Send request to service B to retrieve Coupons
	*/
	requestData := CouponRequest{
		Brand: parameters.Get("brand"),
		Value: parameters.Get("value"),
		Limit: parameters.Get("limit"),
	}

	requestDataJson, err := json.Marshal(requestData)

	res, err := getRequest("http://service-b/api/get-coupons", requestDataJson, headers)

	if err != nil {
		panic(err.Error())
	}

	/*
		Response - process response from service-b to send back to client
	*/
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	c, err := getCoupons([]byte(body))

	response.Header().Set("Content-Type", "application/json")
	response.Write(c)

	/*
		response.Header().Set("Content-Type", "application/json")
		response.Write(couponsJSON)
	*/
}

func getRequest(url string, data, headers []string) (resp, err) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, data)

	// for i := 0; i < len(headers); i++ {
	// 	req.Header.Add()
	// }

	return resp, err
}

func getCoupons(body []byte) (*CouponAPIResponse, error) {
	var c = new(CouponAPIResponse)
	err := json.Unmarshal(body, &c)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return c, err
}
