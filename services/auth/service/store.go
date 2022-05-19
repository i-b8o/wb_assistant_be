package authservice

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	// "github.com/dgrijalva/jwt-go"

	"github.com/bogach-ivan/wb_assistant_be/pb"
	"github.com/bogach-ivan/wb_assistant_be/services/auth/repo"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	// "github.com/i-rm/wb/be/helpers"
	"golang.org/x/crypto/bcrypt"
)

type Store interface {
	Login(email string, pass string) *pb.LoginResponse
	Register(username string, email string, pass, token string) *pb.RegisterResponse
	Confirm(token string) *pb.AuthConfirmResponse
	UpdateVerificationToken(email, pass, token string) *pb.UpdateVerificationTokenResponse
	SetTokenToPassReset(email, token string) *pb.SetTokenToPassResetResponse
	PassReset(email, newPass, token string) *pb.PassResetResponse
}

// DBStore ...
type DBStore struct {
	db *sql.DB
}

// NewDBStore ...
func NewDBStore(host, username, password, name string) *DBStore {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	db, err := repo.NewMySQLDB(repo.Config{
		Host:     viper.GetString("db.host"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()
	return &DBStore{
		db: db,
	}
}

func (store *DBStore) Login(email, pass string) *pb.LoginResponse {

	// Validation
	valid := helpers.Validation(
		[]helpers.ValidationType{
			{Value: email, Valid: "email"},
			{Value: pass, Valid: "password"},
		})

	resp := &pb.LoginResponse{}
	if !valid {
		resp.StatusCode = http.StatusUnprocessableEntity
		return resp
	}
	// Connect DB
	// db := repo.ConnectDB(store.username, store.password, store.name, store.host)

	// User not found
	user := &User{}
	if store.db.Where("email = ? ", email).First(&user).RecordNotFound() {
		resp.StatusCode = http.StatusNotFound
		return resp
	}

	// Verify password
	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	// Wrong password
	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		resp.StatusCode = http.StatusUnauthorized
		return resp
	}

	// Verify confirmed Email
	ver := &Verified{}
	// Email not confirmed
	if !store.db.Where("user_id = ?", user.ID).First(&ver).RecordNotFound() {
		resp.StatusCode = http.StatusForbidden
		return resp
	}

	// Find accounts for the user
	// accounts := []Account{}
	// store.db.Table("accounts").Select("id, type, expires").Where("user_id = ? ", user.ID).Scan(&accounts)
	// // Convert them to pb view
	// accs := []*pb.ResponseAccount{}
	// for _, a := range accounts {
	// 	acc := &pb.ResponseAccount{
	// 		ID:      uint64(a.ID),
	// 		Type:    a.Type,
	// 		Expires: a.Expires,
	// 	}
	// 	accs = append(accs, acc)
	// }

	responseUser := &pb.User{
		Username: user.Username,
		Email:    user.Email,
		Password: pass,
		Type:     user.Type,
		Expires:  user.Expires,
	}
	resp.User = responseUser
	// var jwtToken = prepareJWTToken(user)
	// resp.Jwt = jwtToken
	resp.StatusCode = http.StatusOK
	return resp

}

// TODO Drop jwt token from response
func (store *DBStore) Register(username, email, pass, token string) *pb.RegisterResponse {
	resp := &pb.RegisterResponse{}
	// validation
	valid := helpers.Validation(
		[]helpers.ValidationType{
			{Value: username, Valid: "username"},
			{Value: email, Valid: "email"},
			{Value: pass, Valid: "password"},
		})

	if !valid {
		resp.StatusCode = http.StatusUnprocessableEntity
		return resp
	}

	// Connect to db

	defer store.db.Close()

	// if Email already exists in db
	emailExists := !store.db.Where("email = ?", email).First(&User{}).RecordNotFound()
	if emailExists {
		resp.StatusCode = http.StatusConflict
		return resp
	}

	// Hash password and then create user in db
	generatedPassword := helpers.HashAndSalt([]byte(pass))
	// Create a 7 days free account
	var datetime = time.Now()
	t2 := datetime.AddDate(0, 0, 7)
	dt := t2.Format(time.RFC3339)
	user := &User{Username: username, Email: email, Password: generatedPassword, Type: "free", Expires: dt}
	store.db.Create(&user)

	// // Create a 7 days free account
	// var datetime = time.Now()
	// t2 := datetime.AddDate(0, 0, 7)
	// dt := t2.Format(time.RFC3339)
	// account := &Account{Type: "free", Expires: dt, UserID: user.ID}
	// store.db.Create(&account)

	// Create verified
	verified := Verified{UserID: user.ID, Token: token}
	store.db.Create(&verified)

	// Create response with all fields

	// User
	respUser := &pb.User{
		Username: user.Username,
		Email:    user.Email,
		Password: pass,
		Type:     user.Type,
		Expires:  user.Expires,
	}
	resp.StatusCode = http.StatusOK
	resp.User = respUser
	return resp
}

func (store *DBStore) UpdateVerificationToken(email, pass, token string) *pb.UpdateVerificationTokenResponse {
	resp := &pb.UpdateVerificationTokenResponse{}
	// validation
	valid := helpers.Validation(
		[]helpers.ValidationType{
			{Value: email, Valid: "email"},
			{Value: pass, Valid: "password"},
		})

	if !valid {
		resp := &pb.UpdateVerificationTokenResponse{
			StatusCode: http.StatusUnprocessableEntity,
		}
		return resp
	}

	// Connect to db

	// User not found
	user := &User{}
	if store.db.Where("email = ? ", email).First(&user).RecordNotFound() {
		fmt.Println(user.Model)
		resp.StatusCode = http.StatusNotFound
		return resp
	}

	// Verify password
	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	// Wrong password
	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		resp.StatusCode = http.StatusUnauthorized
		return resp
	}

	ver := &Verified{}

	store.db.Model(&Verified{}).Where("user_id = ?", user.ID).First(&ver).Update("token", token)
	resp.StatusCode = http.StatusOK
	return resp

}

func (store *DBStore) Confirm(token string) *pb.AuthConfirmResponse {

	// Connect DB

	defer store.db.Close()

	resp := &pb.AuthConfirmResponse{}
	// Check token exists
	verified := &Verified{}
	if store.db.Where("token = ? ", token).First(&verified).RecordNotFound() {
		resp.StatusCode = http.StatusNotFound
		return resp
	}

	store.db.Unscoped().Delete(&verified)
	resp.StatusCode = http.StatusOK
	return resp
}

func (store *DBStore) SetTokenToPassReset(email, token string) *pb.SetTokenToPassResetResponse {

	// Connect DB

	defer store.db.Close()
	resp := &pb.SetTokenToPassResetResponse{}
	// User not found
	user := &User{}
	if store.db.Where("email = ? ", email).First(&user).RecordNotFound() {
		resp.StatusCode = http.StatusNotFound
		return resp
	}
	// Create pass
	reset := Reset{UserID: user.ID, Token: token}
	store.db.Create(&reset)
	resp.StatusCode = http.StatusOK
	return resp
}

func (store *DBStore) PassReset(email, newPass, token string) *pb.PassResetResponse {

	// Connect DB

	defer store.db.Close()
	resp := &pb.PassResetResponse{}
	// User not found
	user := &User{}
	if store.db.Where("email = ? ", email).First(&user).RecordNotFound() {
		resp.StatusCode = http.StatusNotFound
		return resp
	}
	// Row not found
	reset := &Reset{
		UserID: user.ID,
	}

	if store.db.Where("user_id = ? and token = ?", user.ID, token).First(&reset).RecordNotFound() {
		resp.StatusCode = http.StatusConflict
		return resp
	}

	store.db.Unscoped().Delete(&reset)

	generatedPassword := helpers.HashAndSalt([]byte(newPass))
	store.db.Model(user).Update("password", generatedPassword)
	resp.StatusCode = http.StatusOK
	return resp
}

// func prepareJWTToken(user *User) string {
// 	tokenContent := jwt.MapClaims{
// 		"user_id": user.ID,
// 		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
// 	}
// 	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
// 	token, err := jwtToken.SignedString([]byte("TokenPassword"))
// 	helpers.HandleErr(err)

// 	return token
// }
