package security

import (
	"context"
	"encoding/hex"
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/FirdavsMF/gosql/pkg/types"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrNoSuchUser      = errors.New("no such user")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInternal        = errors.New("internal error")
	ErrExpired         = errors.New("expired token")
)

//Service - describes managers service.
type Service struct {
	pool *pgxpool.Pool
}

//newService - create a service.
func NewService(pool *pgxpool.Pool) *Service {
	return &Service{pool: pool}
}

//Auth - checks compatibility for the presence of a username and password in database.
func (s *Service) Auth(login string, password string) (ok bool) {

	item := &types.Managers{}
	ctx := context.Background()
	sql := `SELECT id, name, login, password, salary, plan, active, created 
			FROM managers WHERE login = $1 AND password = $2;`
	err := s.pool.QueryRow(ctx, sql, login, password).Scan(
		&item.ID,
		&item.Name,
		&item.Login,
		&item.Password,
		&item.Salary,
		&item.Plan,
		&item.Active,
		&item.Created,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		log.Println("No Rows")
		return
	}

	if err != nil {
		log.Println("Incorrect Password and Login:", err)
		return false
	}
	return true
}

//TokernForCustomer - generates a token for customer.
// Returns ErrNoSuchUser if user is not found.
// Returns ErrInvalidPassword if password incorrect.
// Returns ErrInternal if another error occurs.
func (s *Service) TokenForCustomer(ctx context.Context, phone string, password string,
) (token string, err error) {

	item := &types.Customer{}
	id := item.ID
	hash := item.Password

	sql1 := `SELECT id, password FROM customers WHERE phone = $1;`
	err = s.pool.QueryRow(ctx, sql1, phone).Scan(&id, &hash)

	if err == pgx.ErrNoRows {
		return "", ErrNoSuchUser
	}

	if err != nil {
		return "", ErrInternal
	}

	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return "", ErrInvalidPassword
	}

	buffer := make([]byte, 256)
	n, err := rand.Read(buffer)
	if n != len(buffer) || err != nil {
		return "", ErrInternal
	}

	token = hex.EncodeToString(buffer)
	sql2 := `INSERT INTO customers_tokens(token, customer_id) VALUES($1, $2);`
	_, err = s.pool.Exec(ctx, sql2, token, id)
	if err != nil {
		return "", ErrInternal
	}
	return token, nil
}

// AuthenticateCustomer - performs the customers authentication procedure,
// if successfull returns its id.
// Returns ErrNoSuchUser if user is not found.
// Returns ErrInternal if another error occurs.
func (s *Service) AuthenticateCustomer(ctx context.Context, token string) (int64, error) {

	item := &types.CustomerToken{}
	id := item.Customer_id
	expire := item.Expire

	sql := `SELECT customer_id, expire FROM customers_tokens WHERE token = $1;`
	err := s.pool.QueryRow(ctx, sql, token).Scan(&id, &expire)

	if err == pgx.ErrNoRows {
		log.Println(err)
		return 0, ErrNoSuchUser
	}

	if err != nil {
		log.Println(err)
		return 0, ErrInternal
	}

	currentTime := time.Now().Unix()
	expiredTime := expire.Unix()
	if currentTime > expiredTime {
		return 0, ErrExpired
	}

	return id, nil
}
