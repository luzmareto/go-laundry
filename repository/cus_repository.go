package repository

import (
	"database/sql"
	"fmt"
	"go-laundry/model"
	"log"

	"github.com/google/uuid"
)

type CusRepository interface {
	Create(newCus *model.Customer) error
	FindAll() ([]model.Customer, error)
	FindOne(id string) (*model.Customer, error)
	Update(updatedCus *model.Customer) error
	Delete(id string) error
}

type cusRepository struct {
	db *sql.DB
}

func (r *cusRepository) Create(newCus *model.Customer) error {
	newId := uuid.New().String()
	query := `INSERT INTO customer(id,name,phone_number) Values ($1,$2,$3)`
	_, err := r.db.Exec(query, newId, newCus.Name, newCus.PhoneNumber)

	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Customer %s added successfully", newCus.Name)
	}

	return err
}

func (r *cusRepository) FindAll() ([]model.Customer, error) {
	query := `SELECT id, name, phone_number FROM customer`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cusList := []model.Customer{}
	for rows.Next() {
		var cus model.Customer
		err := rows.Scan(&cus.Id, &cus.Name, &cus.PhoneNumber)
		if err != nil {
			return nil, err
		}
		cusList = append(cusList, cus)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return cusList, err
}

func (r *cusRepository) FindOne(id string) (*model.Customer, error) {
	query := `SELECT name,phone_number FROM customer WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var cus model.Customer
	err := row.Scan(&cus.Name, &cus.PhoneNumber)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("Customer with id %s not found", id)
	} else if err != nil {
		return nil, err
	}

	cus.Id = id

	return &cus, nil
}

func (r *cusRepository) Update(updatedCus *model.Customer) error {
	query := `UPDATE customer SET name = $1, phone_number = $2 WHERE id = $3`
	_, err := r.db.Exec(query, updatedCus.Name, updatedCus.PhoneNumber, updatedCus.Id)

	if err == nil {
		log.Printf("Customer with id %s updated succesfully", updatedCus.Id)
	} else {
		log.Println(err)
	}

	return err
}

func (r *cusRepository) Delete(id string) error {
	query := `UPDATE customer SET is_deleted = true WHERE id = $1`
	_, err := r.db.Exec(query, id)

	if err == nil {
		log.Printf("Customer with id %s deleted succesfully", id)
	} else {
		log.Println(err)
	}

	return err
}

func NewCusRepository(db *sql.DB) CusRepository {
	cusRepo := cusRepository{
		db,
	}

	return &cusRepo
}
