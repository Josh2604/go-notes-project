package auth

import (
	"context"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/Josh2604/go-notes-project/core/apierrors"
	"github.com/Josh2604/go-notes-project/core/entities"
	"github.com/Josh2604/go-notes-project/core/providers/mongo"
	"github.com/dgrijalva/jwt-go/v4"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *entities.User `json:"user"`
}

type AuthRepositoryImplementation struct {
	UserRepo       mongo.UserRepository
	HashSalt       string
	SigningKey     []byte
	ExpireDuration time.Duration
}

func (a *AuthRepositoryImplementation) SingUp(ctx context.Context, username, password string) error {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.HashSalt))

	user := &entities.User{
		Username: username,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
	}

	return a.UserRepo.CreateUser(ctx, user)
}

func (a *AuthRepositoryImplementation) SingIn(ctx context.Context, username, password string) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(a.HashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := a.UserRepo.GetUser(ctx, username, password)
	if err != nil {
		return "", apierrors.ErrUserNotFound
	}

	claims := AuthClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(a.ExpireDuration)),
		},
		User: user,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	return token.SignedString(a.SigningKey)

}

func (a *AuthRepositoryImplementation) ParseToken(ctx context.Context, accessToken string) (*entities.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return a.SigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, apierrors.ErrInvalidAccessToken
}
