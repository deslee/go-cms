package data

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const SigningKey = "abcdefgjasiojdaoidjabcdefgjasiojdaoidjabcdefsgjasiojdaoidj" // TODO: configure!!
const UnauthenticatedMsg = "Not Authenticated"

type UserContextKey string

type User struct {
	ID       string     `gorm:"type:text;primary_key;column:Id"`
	Email    string     `gorm:"type:text;column:Email"`
	Password string     `gorm:"type:text;column:Password"`
	Salt     string     `gorm:"type:text;column:Salt"`
	Data     JSONObject `gorm:"type:text;column:Data"`
	AuditFields
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResult struct {
	GenericResult
	Data  *User   `json:"data"`
	Token *string `json:"token"`
}

type RegisterInput struct {
	Email    string `json:"email"`
	Data     string `json:"data"`
	Password string `json:"password"`
}

type UserInput struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Data  string `json:"data"`
}

type UserResult struct {
	GenericResult
	Data *User `json:"data"`
}

func UserFromContext(ctx context.Context, db *gorm.DB) (*User, error) {
	userId, ok := ctx.Value(UserContextKey("sub")).(string)
	if !ok || len(userId) == 0 {
		return nil, nil
	}
	return getUserById(ctx, db, userId)
}

func (user User) Sites(ctx context.Context, db *gorm.DB) ([]Site, error) {
	panic("not implemented")
}

func UpdateUser(ctx context.Context, db *gorm.DB, user UserInput) (UserResult, error) {
	existingUser, err := getUserByEmail(ctx, db, user.Email)
	die(err)

	if existingUser == nil {
		return UserResult{GenericResult: GenericErrorMessage(fmt.Sprintf("User %s not found", user.Email))}, nil
	}

	return UserResult{
		GenericResult: GenericSuccess(),
		Data:          existingUser,
	}, nil
}

func Register(ctx context.Context, db *gorm.DB, registration RegisterInput) (UserResult, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(registration.Password), bcrypt.DefaultCost)
	die(err)

	if existingUser, _ := getUserByEmail(ctx, db, registration.Email); existingUser != nil {
		return UserResult{GenericResult: GenericErrorMessage(fmt.Sprintf("User %s already exists", registration.Email))}, nil
	}

	id := generateId()
	auditFields := CreateAuditFields(ctx, nil)

	user := User{
		ID:          id,
		Email:       registration.Email,
		Password:    registration.Password,
		Salt:        string(hash),
		Data:        registration.Data,
		AuditFields: auditFields,
	}

	err = db.Create(user).Error
	die(err)

	return UserResult{
		GenericResult: GenericSuccess(),
		Data: &User{
			ID:          id,
			Email:       registration.Email,
			Data:        registration.Data,
			AuditFields: auditFields,
		},
	}, nil
}

func Login(ctx context.Context, db *gorm.DB, login LoginInput) (LoginResult, error) {
	user, err := getUserByEmail(ctx, db, login.Email)
	die(err)
	if user == nil {
		return LoginResult{
			GenericResult: GenericErrorMessage("Failed to login"),
			Data:          nil,
			Token:         nil,
		}, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Salt), []byte(user.Password))
	if err != nil {
		return LoginResult{
			GenericResult: GenericErrorMessage("Failed to login"),
			Data:          nil,
			Token:         nil,
		}, nil
	}

	token := generateToken(user)

	return LoginResult{
		GenericResult: GenericSuccess(),
		Data:          user,
		Token:         &token,
	}, nil
}

func getUserByEmail(ctx context.Context, db *gorm.DB, email string) (*User, error) {
	user := User{}

	if err := db.Where("Email = ?", email).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func getUserById(ctx context.Context, db *gorm.DB, id string) (*User, error) {
	user := User{}

	if err := db.Where("Id = ?", id).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func ParseTokenToContext(tokenString string, ctx context.Context) (context.Context, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validated the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(SigningKey), nil
	})

	if err != nil {
		return ctx, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sub, present := claims["sub"]
		if present == false {
			return ctx, fmt.Errorf("subject not found in token")
		}
		return context.WithValue(ctx, UserContextKey("sub"), sub.(string)), nil
	} else {
		return ctx, fmt.Errorf("invalid token")
	}
}

func generateToken(user *User) string {
	const hoursExpire = 7 * 24
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"exp": time.Hour * hoursExpire, // expires at
		"sub": user.ID,                 // subject
		"iat": time.Now().Unix(),       // issued at
	})

	tokenString, err := token.SignedString([]byte(SigningKey))
	die(err)
	return tokenString
}
