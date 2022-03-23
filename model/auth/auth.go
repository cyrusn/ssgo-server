package auth

import (
	"database/sql"
	"errors"
	"time"

	auth "github.com/cyrusn/goJWTAuthHelper"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var (
	lifeTime           int64
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
)

// UpdateLifeTime is used for update the lifeTime of jwt token
func UpdateLifeTime(time int64) {
	lifeTime = time
}

type DB struct {
	*sql.DB
	*auth.Secret
}

type Credential struct {
	UserAlias string `json:"userAlias"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	Name      string `json:"name"`
	CName     string `json:"cname"`
}

// Claims is jwt.Claims for authentication credential
type Claims struct {
	UserAlias string
	Role      string
	Name      string
	Cname     string
	jwt.StandardClaims
}

func expiresAfter(lifeTime int64) int64 {
	return time.Now().Add(time.Minute * time.Duration(lifeTime)).Unix()
}

// Update use to refresh token,
func (claims *Claims) Update(token *jwt.Token) {
	claims.ExpiresAt = expiresAfter(lifeTime)
	token.Claims = claims
}

func (db *DB) Insert(c *Credential) error {
	password := []byte(c.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		return err
	}

	_, err = db.Exec(
		`INSERT INTO Credential (
      userAlias,
			name,
			cname,
      password,
      role
    ) values (?, ?, ?, ?, ?)`,
		c.UserAlias,
		c.Name,
		c.CName,
		hashedPassword,
		c.Role,
	)
	return err
}

func (db *DB) Authenticate(userAlias, password string) (string, error) {
	hashedPassword := []byte{}
	role := ""
	name := ""
	cname := ""

	err := db.QueryRow(
		`SELECT name, cname, password, role FROM Credential WHERE useralias=?`,
		userAlias,
	).Scan(&name, &cname, &hashedPassword, &role)

	switch {
	case err == sql.ErrNoRows:
		return "", ErrUserNotFound
	case err != nil:
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password)); err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return "", ErrInvalidPassword
		}
		return "", err
	}

	claims := Claims{
		UserAlias: userAlias,
		Name:      name,
		Cname:     cname,
		Role:      role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAfter(lifeTime),
		},
	}
	return db.Secret.GenerateToken(claims)
}

func (db *DB) Refresh(jwt string) (string, error) {
	claims := new(Claims)
	return db.Secret.UpdateToken(jwt, claims)
}
