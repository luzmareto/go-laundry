package repository

import (
	"database/sql"
	"fmt"
	"go-enigma-laundry/model"
	"log"

	"github.com/google/uuid"
)

type UomRepository interface {
	Create(newUom *model.Uom) error
	FindAll() ([]model.Uom, error)
	FindOne(id string) (*model.Uom, error)
	Update(updatedUom *model.Uom) error
	Delete(id string) error
}

type uomRepository struct {
	db *sql.DB
}

func (r *uomRepository) Create(newUom *model.Uom) error {
	if newUom.Name == "" {
		return fmt.Errorf("UOM Name cannot be empty")
	}
	newId := uuid.New().String()
	query := `INSERT INTO uom(id,name) Values ($1,$2)`
	_, err := r.db.Exec(query, newId, newUom.Name)

	if err != nil {
		log.Println(err)
	} else {
		log.Printf("UOM %s added successfully", newUom.Name)
	}

	return err
}

func (r *uomRepository) FindAll() ([]model.Uom, error) {
	query := `SELECT id, name FROM uom`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	uomList := []model.Uom{}
	for rows.Next() {
		var uom model.Uom
		err := rows.Scan(&uom.Id, &uom.Name)
		if err != nil {
			return nil, err
		}
		uomList = append(uomList, uom)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return uomList, err
}

func (r *uomRepository) FindOne(id string) (*model.Uom, error) {
	query := `SELECT name FROM uom WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var uom model.Uom
	err := row.Scan(&uom.Name)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("UOM with id %s not found", id)
	} else if err != nil {
		return nil, err
	}

	uom.Id = id

	return &uom, nil
}

func (r *uomRepository) Update(updatedUom *model.Uom) error {
	query := `UPDATE uom SET name = $1 WHERE id = $2`
	_, err := r.db.Exec(query, updatedUom.Name, updatedUom.Id)

	if err == nil {
		log.Printf("UOM with id %s updated succesfully", updatedUom.Id)
	} else {
		log.Println(err)
	}

	return err
}

func (r *uomRepository) Delete(id string) error {
	query := `UPDATE uom SET is_deleted = true WHERE id = $1`
	_, err := r.db.Exec(query, id)

	if err == nil {
		log.Printf("UOM with id %s deleted succesfully", id)
	} else {
		log.Println(err)
	}

	return err
}

func NewUomRepository(db *sql.DB) UomRepository {
	uomRepo := uomRepository {
		db,
	}

	return &uomRepo
}