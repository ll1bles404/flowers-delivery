package pgsql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/ll1bles404/flowers-delivery/internal/config"
)

type Storage struct {
	db *sql.DB
}

type Order struct {
	ID        int
	VendorID  int
	CourierID int
	ClientID  int
	Address   string
	TimeDate  string //пока string
}

type Client struct {
	ID    int
	Name  string
	Phone string //пока string
}

type Vendor struct {
	ID      int
	Name    string
	Phone   string // пока string
	Address string
}

type CourierID struct {
	ID     int
	Name   string
	Phone  string
	Status string
}

func New(cfg config.Config) (*Storage, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to connect: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping: %w", err)
	}

	log.Print("successful connection to database")

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS clients (
		id int not null,
		name text,
		phone text
	)`)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create table `clients`: %w", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS vendors (
		id int not null,
		name text,
		phone text,
		address text
	)`)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create table `vendors`: %w", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS couriers (
		id int not null, 
		name text,
		phone text,
		status text
	)`)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to create table `couriers`: %w", err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) GetClient(id int) (*Client, error) {
	const op = "storage.pgsql.GetClient"
	var client Client
	stmt, err := s.db.Prepare("SELECT id, name, phone FROM clients WHERE id = $1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err = stmt.QueryRow(id).Scan(&client.ID, &client.Name, &client.Phone)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("client is not found")
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &client, nil
}
