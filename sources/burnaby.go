package sources

import (
	"net/http"
	"net/url"
	"strings"
)

// const burnabyURL = `
// curl -H 'Authorization: Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6IjIyMkY3QzY3MTZGRkY5N0U5QzNGNzIyMjZGMDBBM0NEQkVGNTlEOEFSUzI1NiIsInR5cCI6ImF0K2p3dCIsIng1dCI6IklpOThaeGJfLVg2Y1AzSWlid0NqemI3MW5ZbyJ9.eyJuYmYiOjE2OTQ3Mjg1MDAsImV4cCI6MTY5NDczMjEwMCwiaXNzIjoiaHR0cHM6Ly9nb2xmYnVybmFieS5jcHMuZ29sZi9pZGVudGl0eWFwaSIsImF1ZCI6WyJjdXN0b21lciIsImVtYWlsIiwiaW52ZW50b3J5Iiwib25saW5lcmVzZXJ2YXRpb24iLCJyZWNvbW1lbmQiLCJyZWZlcmVuY2VzIiwic2FsZSIsInNoIl0sImNsaWVudF9pZCI6ImpzMSIsInN1YiI6IjEyMTcwNCIsImF1dGhfdGltZSI6MTY5NDcyODUwMCwiaWRwIjoibG9jYWwiLCJyb2xlIjoidXNlciIsInByZWZlcnJlZF91c2VybmFtZSI6IjEyMTcwNCIsImlzX3Bza3VzZXIiOmZhbHNlLCJlbXBsb3llZV90eXBlIjoiTm9uZSIsImZpcnN0X25hbWUiOiJTSEFXTiBIT09ORyBUWkUiLCJsYXN0X25hbWUiOiJMSU0iLCJuYW1lIjoiTElNIFNIQVdOIEhPT05HIFRaRSIsImVtYWlsIjoiU0hBV05fTElNQE9VVExPT0suQ09NIiwibW9iaWxlX3Bob25lIjoiNzc4MjQ2MjY2OCIsInNleCI6IlUiLCJwbGF5ZXJfaWQiOiIwIiwiZGVwZW5kZW50X2lkIjoiMCIsImNsYXNzQ29kZSI6IkFEIiwic3RvcmVfaWQiOiIxIiwiZ29sZmVySWQiOiI2MjYiLCJ1c2VyX3R5cGUiOiJNZW1iZXIiLCJhY2N0IjoiMTIxNzA0IiwibWVtYmVyX3R5cGVfY29kZSI6IkNPQiIsInVzZXJuYW1lIjoiMTIxNzA0IiwibG9naW5fbWV0aG9kIjoiMCIsImNvbXBvbmVudF9pZCI6IjEiLCJpc3N1ZXJfaG9zdCI6ImdvbGZidXJuYWJ5LmNwcy5nb2xmIiwianRpIjoiMzJEQzFBQzM1NjUzOTQwOTRFNDI2NzBCRjk3MzBGNkUiLCJpYXQiOjE2OTQ3Mjg1MDAsInNjb3BlIjpbImN1c3RvbWVyIiwiZW1haWwiLCJpbnZlbnRvcnkiLCJvbmxpbmVyZXNlcnZhdGlvbiIsIm9wZW5pZCIsInByb2ZpbGUiLCJyZWNvbW1lbmQiLCJyZWZlcmVuY2VzIiwic2FsZSIsInNoIl0sImFtciI6WyJwd2QiXX0.kHvdoCsrpc0UvZOea_QfC_CsCXJmh1Pmja_YvbGfNA9v8Lw_WCbQoY7ah-yYO66W3AIOqKV19ARrB-CefBBWcDblRQlfzuUXM2dd8pR3B13yF_69v-vjEElb7kR1bai2VYJ3I-ZV-oWDcE_fwS52sUKQ9bK9kfM6vAjisDZ0Ah24tNFPuvDdA0icZVoLxnaBhKNdaLzih9O-0AF4UlKgwMhTVKDKQ83uZbIGCxmSP38O6MzrxTYBpmb8nml1vTFxKXrawpR585uEr0h_V3JuJyTkuTxFDTbfWcV4uoSRNM70aqzDI3rkwvC-IJGjehDhTrdk6fT55xMM8Ihi7MVODa6mF0f3PKosTlUQ2SPf5RR_IaGlpQOb_QQ2rymVD-ABA8xCMzjRXoybAiCZG4w2qLOuhDCROase_TNX1xzUyK1qiY-wNAsSDwUpDMmXypMMUy62EXiFbxzbYIjtMBnJz_x4i62s06jyIf-GAN8n5DTS0ml8x6vNrqJ-66cyMj3IANTJRNeM91sY4M_7grIQVFkyfbVjczPwLeEbYJAX8EIDBJslrN_E34DdXEUwJxKrAYILbg76Rw6oMovk-W2pmmFXJSICC28TGNaHrAGWBQMU4BA7-IqaPMPw67zvMJlGI43CrBccbTVnT_4mHUEgTewIlumi_ASYNAc57XNcs1I' -H 'X-Componentid: 1' -H "Content-type: application/json" 'https://golfburnaby.cps.golf/onlineres/onlineapi/api/v1/onlinereservation/TeeTimes?searchDate=Thu%20Sep%2014%202023&holes=0&numberOfPlayer=0&courseIds=1%2C2&searchTimeType=0&teeOffTimeMin=0&teeOffTimeMax=23&isChangeTeeOffTime=true&teeSheetSearchView=5&classCode=AD&defaultOnlineRate=N&isUseCapacityPricing=false&memberStoreId=1&searchType=1'
// `

var burnabyURL = `https://golfburnaby.cps.golf/onlineres/onlineapi/api/v1/onlinereservation/TeeTimes`

type Burnaby struct {
}

func (b *Burnaby) Login(username, password string) (token string, err error) {
	const loginURL = "https://golfburnaby.cps.golf/identityapi/connect/token"

	client := &http.Client{}

	form := url.Values{}
	form.Add("grant_type", "password")
	form.Add("scope", "openid profile onlinereservation sale inventory sh customer email recommend references")
	form.Add("username", username)
	form.Add("password", password)
	form.Add("client_id", "js1")
	form.Add("client_secret", "v4secret")

	request, err := http.NewRequest("POST", loginURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", nil
	}
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}

	// pase token
	return
}

func (b *Burnaby) BookTeeTime() {

}
