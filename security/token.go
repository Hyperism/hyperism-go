package security

import (
	"github.com/hyperism/hyperism-go/models"
	"github.com/hyperism/hyperism-go/utix"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func NewToken(user models.User) (string, int64, error) {

	secret := os.Getenv("JWT_SECRET_KEY")

	token := jwt.New(jwt.SigningMethodHS256)
	exp := time.Now().Add(time.Hour * 24).Unix()

	claims := token.Claims.(jwt.MapClaims)

	claims["Id"] = user.ID
	claims["exp"] = exp
	claims["Issuer"] = user.ID
	claims["IssueAt"] = time.Now().Unix()

	claims["dumbfuck"] = true

	t, err := token.SignedString([]byte(secret))
	utix.CheckErorr(err)

	return t, exp, err

}
