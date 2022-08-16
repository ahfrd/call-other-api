package controllers

import (
	"danspro/constant"
	"danspro/src/helpers"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

func (o ControllerStructure) ControllerGetJobsList(w http.ResponseWriter, req *http.Request) {
	var bodyReq ControllerGetJobsReq
	res := helpers.Response{}
	err := json.NewDecoder(req.Body).Decode(&bodyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	tokenJwt, err := req.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			res.Body.Code = "99"
			res.Body.Msg = "No Cookie detect"
			res.Reply(w)
			return
		}
		res.Body.Code = constant.GeneralErrorCode
		res.Body.Msg = constant.GeneralErrorDesc
		res.Reply(w)
		return
	}
	tokenStr := tokenJwt.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return helpers.JwtKey(), nil
		})
	if err != nil {
		res.Body.Code = constant.GeneralErrorCode
		res.Body.Msg = constant.GeneralErrorDesc
		res.Reply(w)
		return
	}
	if !tkn.Valid {
		res.Body.Code = "99"
		res.Body.Msg = "Token Not Valid"
		res.Reply(w)
		return
	}
	description := bodyReq.Description
	location := bodyReq.Location
	fulltime := bodyReq.FullTime
	paging := bodyReq.Pagination
	URL := os.Getenv("DATADANS")
	searchLocation := ""
	searchFullTime := ""
	searchDescription := ""
	pagination := ""
	sampleRegexp := regexp.MustCompile(`\d`)
	regexForPage := sampleRegexp.MatchString(paging)
	if !regexForPage {
		res.Body.Code = "99"
		res.Body.Msg = "Page must be numeric"
		res.Reply(w)
		return
	}
	if paging != "" {
		pagination = fmt.Sprintf("page=%s", paging)
	}
	if location != "" {
		searchLocation = fmt.Sprintf("&location=%s", location)

	}
	if fulltime != "" {
		searchFullTime = fmt.Sprintf("&full_time=%s", fulltime)

	}
	if description != "" {
		searchDescription = fmt.Sprintf("&description=%s", description)

	}
	if description != "" || location != "" || fulltime != "" || pagination != "" {
		URL = os.Getenv("DATADANS") + pagination + strings.ToLower(searchDescription) + strings.ToLower(searchLocation) + strings.ToLower(searchFullTime)
	}
	api, err := helpers.FormatCallApiResult(URL)
	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = api
	res.Reply(w)
	return
}
func (o ControllerStructure) ControllerGetJobsDetail(w http.ResponseWriter, req *http.Request) {
	res := helpers.Response{}
	slug := mux.Vars(req)

	tokenJwt, err := req.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			res.Body.Code = "99"
			res.Body.Msg = "No Cookie detect"
			res.Reply(w)
			return
		}
		res.Body.Code = constant.GeneralErrorCode
		res.Body.Msg = constant.GeneralErrorDesc
		res.Reply(w)
		return
	}
	tokenStr := tokenJwt.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return helpers.JwtKey(), nil
		})
	if err != nil {
		res.Body.Code = constant.GeneralErrorCode
		res.Body.Msg = constant.GeneralErrorDesc
		res.Reply(w)
		return
	}
	if !tkn.Valid {
		res.Body.Code = "99"
		res.Body.Msg = "Token Not Valid"
		res.Reply(w)
		return
	}
	if slug["id"] == "" {
		res.Body.Code = "99"
		res.Body.Msg = "Slug not found"
		res.Reply(w)
		return
	}
	URL := os.Getenv("DATADETAILDANS") + slug["id"]
	fmt.Println(URL)
	api, err := helpers.FormatCallApiRow(URL)
	if err != nil {
		res.Body.Code = "99"
		res.Body.Msg = err.Error()
		res.Reply(w)
		return
	}
	res.Body.Code = constant.RCSuccess
	res.Body.Msg = constant.RCSuccessMsg
	res.Body.Data = api
	res.Reply(w)
	return
}
