package customers

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"

	"github.com/FirdavsMF/gosql/pkg/types"
)

var (
	ErrNotFound = errors.New("item not found") //retrun when customer not found
	ErrInternal = errors.New("internal error") //return when an internal error occurred
)

//Service - describes customer service.
type Service struct {
	pool *pgxpool.Pool
}

//newService - create a service.
func NewService(pool *pgxpool.Pool) *Service {
	return &Service{pool: pool}
}

//All - returns all existing customers.
func (s *Service) All(ctx context.Context) ([]*types.Customer, error) {

	items := make([]*types.Customer, 0)
	sql := `SELECT id, name, phone, active, created FROM customers;`
	rows, err := s.pool.Query(ctx, sql)
	if err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}
	defer rows.Close()

	for rows.Next() {
		item := &types.Customer{}
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.Phone,
			&item.Active,
			&item.Created,
		)

		if err != nil {
			log.Println(err)
			return nil, err
		}
		items = append(items, item)
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(items)
	return items, nil
}

//AllActive - returns all existing customers.
func (s *Service) AllActive(ctx context.Context) ([]*types.Customer, error) {

	items := make([]*types.Customer, 0)
	sql := `SELECT id, name, phone, active, created FROM customers WHERE active = TRUE;`
	rows, err := s.pool.Query(ctx, sql)
	if err != nil {
		log.Println(err)
		return nil, ErrNotFound
	}
	defer rows.Close()

	for rows.Next() {
		item := &types.Customer{}
		err = rows.Scan(
			&item.ID,
			&item.Name,
			&item.Phone,
			&item.Active,
			&item.Created,
		)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		items = append(items, item)
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println(items)
	return items, nil
}

//Save - creates, saves and updates customer lists.
func (s *Service) Save(ctx context.Context, items *types.Customer) (*types.Customer, error) {

	item := &types.Customer{}

	if items.ID == 0 {
		sql1 := `INSERT INTO customers (name, phone, password) VALUES ($1, $2, $3) ON CONFLICT (phone) DO UPDATE SET name= excluded.name
		RETURNING id, name, phone, password, active, created;`
		err := s.pool.QueryRow(ctx, sql1, items.Name, items.Phone, items.Password).Scan(
			&item.ID,
			&item.Name,
			&item.Phone,
			&item.Password,
			&item.Active,
			&item.Created,
		)
		if err != nil {
			log.Println(err)
			return nil, ErrInternal
		}
		return item, nil
	}

	sql2 := `UPDATE customers SET name = $2, phone = $3, password = $4 WHERE id = $1 
	RETURNING  id, name, phone, password, active, created;`
	err := s.pool.QueryRow(ctx, sql2, items.ID, items.Name, items.Phone, items.Password).Scan(
		&item.ID,
		&item.Name,
		&item.Phone,
		&item.Password,
		&item.Active,
		&item.Created,
	)

	if err != nil {
		log.Println(err)
		return nil, ErrInternal
	}
	return item, nil
}

//Delete - deletes customer lists.
func (s *Service) Delete(ctx context.Context, id int64) (*types.Customer, error) {
	item := &types.Customer{}

	sql := `DELETE FROM customers WHERE id = $1 RETURNING id, name, phone, active, created;`
	err := s.pool.QueryRow(ctx, sql, id).Scan(
		&item.ID,
		&item.Name,
		&item.Phone,
		&item.Active,
		&item.Created,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		log.Println("No Rows")
		return nil, ErrNotFound
	}

	if err != nil {
		log.Println(err)
		return nil, ErrInternal
	}
	return item, nil
}

//Block - block customers list(sets false status).
func (s *Service) Block(ctx context.Context, id int64, active bool) (*types.Customer, error) {
	item := &types.Customer{}

	sql := `UPDATE customers SET active = $2 WHERE id = $1 RETURNING id, name, phone, active, created;`
	err := s.pool.QueryRow(ctx, sql, id, active).Scan(
		&item.ID,
		&item.Name,
		&item.Phone,
		&item.Active,
		&item.Created,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		log.Println("No Rows")
		return nil, ErrNotFound
	}

	if err != nil {
		log.Println(err)
		return nil, ErrInternal
	}
	return item, nil
}

//Unblock - unblock customers list(sets true status).
func (s *Service) Unblock(ctx context.Context, id int64, active bool) (*types.Customer, error) {
	item := &types.Customer{}

	sql := `UPDATE customers SET active = $2 WHERE id = $1 RETURNING id, name, phone, active, created;`
	err := s.pool.QueryRow(ctx, sql, id, active).Scan(
		&item.ID,
		&item.Name,
		&item.Phone,
		&item.Active,
		&item.Created,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		log.Println("No Rows")
		return nil, ErrNotFound
	}

	if err != nil {
		log.Println(err)
		return nil, ErrInternal
	}
	return item, nil
}

//ByID - returns customer by ID.
func (s *Service) ByID(ctx context.Context, id int64) (*types.Customer, error) {
	item := &types.Customer{}

	sql := `SELECT id, name, phone, active, created FROM customers WHERE id = $1;`
	err := s.pool.QueryRow(ctx, sql, id).Scan(
		&item.ID,
		&item.Name,
		&item.Phone,
		&item.Active,
		&item.Created,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		log.Println("No Rows")
		return nil, ErrNotFound
	}

	if err != nil {
		log.Println(err)
		return nil, ErrInternal
	}
	return item, nil
}
