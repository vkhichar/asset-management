package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/vkhichar/asset-management/config"
)

const (
	userid  = "user_id"
	isAdmin = "is_admin"
)

var signingMethod = jwt.GetSigningMethod("HS256")

type JwtService struct {
}

func NewJwtService() TokenService {
	return &JwtService{}
}

func (jwtService *JwtService) GenerateToken(c *Claims) (string, error) {

	token := jwt.NewWithClaims(signingMethod, generateJwtClaims(c))
	return token.SignedString([]byte(config.GetJwtConfig().Secret))
}

func (jwtService *JwtService) ValidateToken(token string) (*Claims, error) {

	jwtToken, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return []byte(config.GetJwtConfig().Secret), nil
	})
	if err != nil {
		return nil, InvalidTokenError{message: "Invalid token"}
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)

	if !ok {
		return nil, NewInvalidTokenError("Invalid token")
	}
	userId, _ := claims[userid].(float64)
	isAdmin, _ := claims[isAdmin].(bool)
	return &Claims{UserID: int(userId), IsAdmin: isAdmin}, nil
}

func generateJwtClaims(c *Claims) jwt.Claims {
	now := jwt.TimeFunc()
	claims := jwt.MapClaims{}
	claims[userid] = c.UserID
	claims[isAdmin] = c.IsAdmin
	claims["iat"] = now.Unix() // issuedAt
	claims["exp"] = now.Add(time.Minute * time.Duration(config.GetJwtConfig().TokenExpiry)).Unix()
	// expiresAt = issuedAt + tokenExpiry in minute
	return claims
}
