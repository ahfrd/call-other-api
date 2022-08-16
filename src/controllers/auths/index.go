package controllers

import (
	"danspro/src/database"
	"danspro/src/helpers"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
)

// ControllerNamaFungsiObjectRes is a
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
type ControllerUserReq struct {
	Username        string `son:"username"`
	Password        string `json:"password"`
	Login_Retry     int    `json:"login_retry"`
	Next_Login_date string `json:"next_login_date"`
	Email           string ` json:"email"`
}

type LoginStructRes struct {
	Username string `json:"username"`
	Email    string ` json:"email"`
}
type ControllerStructure struct {
	helpers.ELK
	helpers.Response
	database.TblUser
}
