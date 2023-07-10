package helper

import (
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func CreateToken(guid, secret string, duration time.Duration, m map[string]string) (string, error) {
	claims := jwt.MapClaims{}
	for k, v := range m {
		claims[k] = v
	}

	timeNow := time.Now()

	claims["iss"] = guid
	claims["exp"] = timeNow.Add(duration).Unix()
	claims["iat"] = timeNow.Unix()
	claims["alg"] = "HS256"
	// claims["typ"] = "JWT"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func VerifyToken(secret, tokenStr string) error {
	if !(len(tokenStr) > 8 && tokenStr[:7] == "Bearer ") {
		return fmt.Errorf("invalid token")
	}

	token, err := jwt.Parse(tokenStr[7:], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		return err
	}

	timeNow := time.Now()
	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return fmt.Errorf("error parse claims")
	}

	expDate, err := claims.GetExpirationTime()
	if err != nil {
		return err
	}
	if expDate.Unix() < timeNow.Unix() {
		return fmt.Errorf("token expired")
	}

	iatDate, err := claims.GetIssuedAt()
	if err != nil {
		return err
	}
	if expDate.Unix() < iatDate.Unix() {
		return fmt.Errorf("token is invalid")
	}
	return nil
}
