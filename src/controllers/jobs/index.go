package controllers

import (
	"danspro/src/helpers"

	"github.com/dgrijalva/jwt-go"
)

type ControllerStructure struct {
	helpers.ELK
	helpers.Response
}

type ControllerGetJobsReq struct {
	Description string `json:"description"`
	Location    string `json:"location"`
	FullTime    string `json:"full_time"`
	Pagination  string `json:"page"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type ControllerGetJobsListRes struct {
	Id          string `json:"id"`
	Type        string `json:"type"`
	Url         string `json:"url"`
	CreatedAt   string `json:"created_at"`
	Company     string `json:"company"`
	CompanyUrl  string `json:"company_url"`
	Location    string `json:"location"`
	Title       string `json:"title"`
	Description string `json:"description"`
	HowToApply  string `json:"how_to_apply"`
	CompanyLogo string `json:"company_logo"`
}
