package data

import (
	"context"
	"fmt"
	. "github.com/deslee/cms/model"
	"github.com/deslee/cms/repository"
	"github.com/dgrijalva/jwt-go"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const SigningKey = "abcdefgjasiojdaoidjabcdefgjasiojdaoidjabcdefsgjasiojdaoidj" // TODO: configure!!
const UnauthenticatedMsg = "Not Authenticated"

type UserContextKey string

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

func UserIdFromContext(ctx context.Context) string {
	// try to get the user id from the context
	userId, ok := ctx.Value(UserContextKey("sub")).(string)

	// if unsuccessful, return an empty string
	if !ok || len(userId) == 0 {
		return ""
	}

	return userId
}

func QueryGetCurrentUser(ctx context.Context, db *sqlx.DB) (*User, error) {
	// try to get the user id
	userId := UserIdFromContext(ctx)

	// if unsuccessful, return nil
	if len(userId) == 0 {
		return nil, nil
	}

	// delegate to getUserById
	return repository.FindUserById(ctx, db, userId)
}

func SitesFromUser(ctx context.Context, db *sqlx.DB, user User) ([]Site, error) {
	return repository.ScanSiteList(ctx, db, `SELECT S.* FROM SiteUsers SU INNER JOIN Sites S ON S.Id=SU.SiteId WHERE SU.UserId=?`, user.Id)
}

func MutationUpdateUser(ctx context.Context, db *sqlx.DB, user UserInput) (UserResult, error) {
	// get the user by email, make sure it exists
	existingUser, err := repository.FindUserByEmail(ctx, db, user.Email)
	if err != nil {
		return UnexpectedErrorUserResult(err), nil
	}
	if existingUser == nil {
		return UserResult{GenericResult: ErrorGenericResult(fmt.Sprintf("User %s not found", user.Email))}, nil
	}

	// make sure the current user is the user
	currentUserId := UserIdFromContext(ctx)
	if currentUserId != existingUser.Id {
		return UserResult{GenericResult: ErrorGenericResult(fmt.Sprintf("You do not have authorization to update %s", user.Email))}, nil
	}

	// return the result
	return UserResult{
		GenericResult: GenericSuccess(),
	}, nil
}

func MutationRegister(ctx context.Context, db *sqlx.DB, registration RegisterInput) (UserResult, error) {
	// get the existing user, make sure it doesnt already exist
	existingUser, err := repository.FindUserByEmail(ctx, db, registration.Email)
	if err != nil {
		return UnexpectedErrorUserResult(err), nil
	}
	if existingUser != nil {
		return UserResult{GenericResult: ErrorGenericResult(fmt.Sprintf("User %s already exists", registration.Email))}, nil
	}

	// generate the salt
	hash, err := bcrypt.GenerateFromPassword([]byte(registration.Password), bcrypt.DefaultCost)
	if err != nil {
		return UnexpectedErrorUserResult(err), nil
	}

	// generate an id
	id := GenerateId()

	// construct the user object
	user := User{
		Id:          id,
		Email:       registration.Email,
		Password:    registration.Password,
		Salt:        string(hash),
		Data:        registration.Data,
		AuditFields: CreateAuditFields(ctx),
	}

	// insert the user into the table
	err = repository.UpsertUser(ctx, db, user)
	if err != nil {
		return UnexpectedErrorUserResult(err), nil
	}

	// grab the user back out
	createdUser, err := repository.FindUserById(ctx, db, id)
	if err != nil {
		return UnexpectedErrorUserResult(err), nil
	}

	return UserResult{
		GenericResult: GenericSuccess(),
		Data:          createdUser,
	}, nil
}

func Login(ctx context.Context, db *sqlx.DB, login LoginInput) (LoginResult, error) {
	// get the user from the database, make sure it's not nil
	user, err := repository.FindUserByEmail(ctx, db, login.Email)
	if err != nil {
		return UnexpectedErrorLoginResult(err), nil
	}
	if user == nil {
		return UnexpectedErrorLoginResult(err), nil
	}

	// check the password
	err = bcrypt.CompareHashAndPassword([]byte(user.Salt), []byte(user.Password))
	if err != nil {
		return UnexpectedErrorLoginResult(err), nil
	}

	// generate a token and respond
	token := generateToken(*user)

	return LoginResult{
		GenericResult: GenericSuccess(),
		Data:          user,
		Token:         &token,
	}, nil
}

/**
Given a token string, return a new context with the user identity embedded
*/
func ParseTokenToContext(ctx context.Context, tokenString string) (context.Context, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validated the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

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

/**
Creates a JWT given a user
*/
func generateToken(user User) string {
	const hoursExpire = 7 * 24
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"exp": time.Hour * hoursExpire, // expires at
		"sub": user.Id,                 // subject
		"iat": time.Now().Unix(),       // issued at
	})

	tokenString, err := token.SignedString([]byte(SigningKey))
	die(err)
	return tokenString
}

/**
Returns true if user has access to the site, otherwise false
*/
func AssertContextUserHasAccessToSite(ctx context.Context, db *sqlx.DB, siteId string) (bool, error) {
	userId := UserIdFromContext(ctx)
	if len(userId) == 0 {
		return false, nil
	}

	var count int
	err := db.Get(&count, "SELECT count(*) FROM SiteUsers WHERE UserId=? AND SiteId=?", userId, siteId)
	if err != nil {
		return false, err
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}
