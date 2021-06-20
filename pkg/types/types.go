package types

import "time"

//Customer - represents information about customer.
type Customer struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	Phone    string    `json:"phone"`
	Password string    `json:"password"`
	Active   bool      `json:"active"`
	Created  time.Time `json:"created"`
}

//CustomerToken - represents information about customers token.
type CustomerToken struct {
	Token       string    `json:"token"`
	Customer_id int64     `json:"customer_id"`
	Expire      time.Time `json:"expire"`
	Created     time.Time `json:"created"`
}

//Managers - represents information about customers.
type Managers struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Login      string    `json:"login"`
	Password   string    `json:"password"`
	Salary     int64     `json:"salary"`
	Plan       int64     `json:"plan"`
	Boss_id    int64     `json:"boss_id"`
	Department string    `json:"department"`
	Active     bool      `json:"active"`
	Created    time.Time `json:"created"`
}
