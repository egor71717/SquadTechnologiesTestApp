package handlers

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/egor71717/SquadTechnologiesTestApp/config"
)

//AuthJWTMiddleware validation middleware
var AuthJWTMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	// When set, the middleware verifies that tokens are signed with the specific signing algorithm
	// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return config.SigningKey, nil
	},
	// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries
	SigningMethod: jwt.SigningMethodHS256,
})
