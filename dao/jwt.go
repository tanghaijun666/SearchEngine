package dao

import (
	"SimpleTikTok/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

/***?
author : xieteng
date :2022/06/5
*/
type Claims struct {
	UerId int64
	jwt.StandardClaims
}

var jwtKey = []byte("xieteng")

//颁发token
func GenerateToken(userId int64) (string, error) {
	expretime := time.Now().Add(time.Hour * 24 * 7)
	claims := &Claims{
		UerId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expretime.Unix(),
			Issuer:    "xieteng",
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return token, nil
}

//解析token
func ParesToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

//JWt 用于验证token，并返回token对应的userId
func JwtAuth(token string) (int64, error) {
	if token == "" {
		return 0, errors.New("token为空")
	}
	claim, err := ParesToken(token)
	if err != nil {
		return 0, errors.New("token过期")
	}
	//验证user是否存在
	if err := NewUserLoginDao().QueryUserbyId(claim.UerId, &model.Users{}); err != nil {
		return 0, err
	}
	return claim.UerId, nil
}
