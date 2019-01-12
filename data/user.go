package data

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID       string
	Email    string
	Password string
	Salt     string
	Data     JSONObject
	AuditFields
}

func (user User) Sites(ctx context.Context, db *sql.DB) ([]Site, error) {
	panic("not implemented")
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

func IsAuthenticated(ctx context.Context) bool {
	userId, ok := ctx.Value(UserContextKey("sub")).(string)
	return ok && len(userId) > 0
}

func UserFromContext(ctx context.Context, db *sql.DB) (*User, error) {
	userId, ok := ctx.Value(UserContextKey("sub")).(string)
	if !ok || len(userId) == 0 {
		return nil, nil
	}
	return getUserById(ctx, db, userId)
}

func UpdateUser(ctx context.Context, db *sql.DB, user UserInput) (UserResult, error) {
	panic("not implemented")
}

func Register(ctx context.Context, db *sql.DB, registration RegisterInput) (UserResult, error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	newUuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(registration.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	id := newUuid.String()
	auditFields := CreateAuditFields(ctx, nil)

	executeSql(
		tx,
		`INSERT INTO USERS (Id, Email, Password, Salt, Data, CreatedBy, LastUpdatedBy, CreatedAt, LastUpdatedAt) VALUES (?,?,?,?,?,?,?,?,?)`,
		id,
		registration.Email,
		registration.Password,
		string(hash),
		registration.Data,
		auditFields.CreatedBy,
		auditFields.LastUpdatedBy,
		auditFields.CreatedAt,
		auditFields.LastUpdatedAt,
	)

	err = tx.Commit()
	if err != nil {
		panic(err)
	}
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

func Login(ctx context.Context, db *sql.DB, login LoginInput) (LoginResult, error) {
	user, err := getUserByEmail(ctx, db, login.Email)
	if err != nil {
		panic(err)
	}
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

func AddUserToSite(ctx context.Context, db *sql.DB, userId string, siteId string) (GenericResult, error) {
	panic("not implemented")
}

func getUserByEmail(ctx context.Context, db *sql.DB, email string) (ret *User, err error) {
	user := User{}
	err = db.QueryRowContext(
		ctx,
		`SELECT Id, Email, Password, Salt, Data, CreatedBy, LastUpdatedBy, CreatedAt, LastUpdatedAt from Users where Email = ?`, email,
	).Scan(
		&user.ID, &user.Email, &user.Password, &user.Salt, &user.Data, &user.CreatedBy, &user.LastUpdatedBy, &user.CreatedAt, &user.LastUpdatedAt,
	)

	if err == sql.ErrNoRows {
		err = nil
		ret = nil
	} else if err == nil {
		ret = &user
	}

	return
}

func getUserById(ctx context.Context, db *sql.DB, id string) (ret *User, err error) {
	user := User{}
	err = db.QueryRowContext(
		ctx,
		`SELECT Id, Email, Password, Salt, Data, CreatedBy, LastUpdatedBy, CreatedAt, LastUpdatedAt from Users where Id = ?`, id,
	).Scan(
		&user.ID, &user.Email, &user.Password, &user.Salt, &user.Data, &user.CreatedBy, &user.LastUpdatedBy, &user.CreatedAt, &user.LastUpdatedAt,
	)

	if err == sql.ErrNoRows {
		err = nil
		ret = nil
	} else if err == nil {
		ret = &user
	}

	return
}

const privateKey = "abcdefgjasiojdaoidjabcdefgjasiojdaoidjabcdefgjasiojdaoidj" //

type UserContextKey string

func ParseTokenToContext(tokenString string, ctx context.Context) (context.Context, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(privateKey), nil
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
	const hoursExpire = 7 & 24
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"exp": time.Hour * hoursExpire, // expires at
		"sub": user.ID, // subject
		"iat": time.Now().Unix(), // issued at
	})

	tokenString, err := token.SignedString([]byte(privateKey))
	if err != nil {
		panic(err)
	}
	return tokenString
}
