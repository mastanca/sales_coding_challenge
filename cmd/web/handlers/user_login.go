package handlers

import (
	"net/http"
	"time"

	"github.com/mastanca/SALES_MARTIN_STANCANELLI/usecases"

	"github.com/mastanca/SALES_MARTIN_STANCANELLI/cmd/web/middleware"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginHandler interface {
	Handle(c *gin.Context)
}

type loginHandlerImpl struct {
	getUser usecases.GetUser
}

func NewLoginHandlerImpl(getUser usecases.GetUser) *loginHandlerImpl {
	return &loginHandlerImpl{getUser: getUser}
}

func (l *loginHandlerImpl) Handle(c *gin.Context) {
	var credentials middleware.Credentials
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user, err := l.getUser.Execute(c, credentials.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if passOK := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); passOK != nil {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	// Set token's expiration time to 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &middleware.Claims{
		Username: credentials.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(middleware.JwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	// We should send expire time and build regen token endpoint to avoid asking
	// for pass and username again each 5 mins, but lets keep it this way for this test task
	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})
}
