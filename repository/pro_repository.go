package repository

import (
	"database/sql"
	"fmt"
	"go-laundry/model"
	"log"

	"github.com/google/uuid"
)

type ProRepository interface {
	Create(newPro *model.Product) error
	FindAll() ([]model.Product, error)
	FindOne(id string) (*model.Product, error)
	Update(updatedPro *model.Product) error
	Delete(id string) error
}

type proRepository struct {
	db *sql.DB
}

func (r *proRepository) Create(newPro *model.Product) error {
	// memeriksa apakah uom_id yang diberikan sudah ada di tabel "uom"
	var uomCount int
	err := r.db.QueryRow("SELECT COUNT(*) FROM uom WHERE id = $1", newPro.Uom.Id).Scan(&uomCount)
	if err != nil {
		return err
	}
	if uomCount == 0 {
		return fmt.Errorf("invalid uom_id: %s", newPro.Uom.Id)
	}

	newId := uuid.New().String()
	query := `INSERT INTO product (id,name,price,uom_id) Values ($1,$2,$3,$4)`
	_, err = r.db.Exec(query, newId, newPro.Name, newPro.Price, newPro.Uom.Id)

	if err != nil {
		log.Println(err)
	} else {
		log.Printf("Product %s added successfully", newPro.Name)
	}

	return err
}

func (r *proRepository) FindAll() ([]model.Product, error) {
	query := `SELECT id, name, price, uom_id FROM product`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	proList := []model.Product{}
	for rows.Next() {
		var pro model.Product
		err := rows.Scan(&pro.Id, &pro.Name, &pro.Price, &pro.Uom.Id)
		if err != nil {
			return nil, err
		}
		proList = append(proList, pro)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return proList, err
}

func (r *proRepository) FindOne(id string) (*model.Product, error) {
	query := `SELECT name,price,uom_id FROM product WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var pro model.Product
	err := row.Scan(&pro.Name, &pro.Price, &pro.Uom.Id)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("Product with id %s not found", id)
	} else if err != nil {
		return nil, err
	}

	pro.Id = id

	return &pro, nil
}

func (r *proRepository) Update(updatedPro *model.Product) error {
	query := `UPDATE product SET name = $1, price = $2, uom_id = $3 WHERE id = $4`
	_, err := r.db.Exec(query, updatedPro.Name, updatedPro.Price, updatedPro.Uom.Id, updatedPro.Id)

	if err == nil {
		log.Printf("Product with id %s updated succesfully", updatedPro.Id)
	} else {
		log.Println(err)
	}

	return err
}

func (r *proRepository) Delete(id string) error {
	query := `UPDATE product SET is_deleted = true WHERE id = $1`
	_, err := r.db.Exec(query, id)

	if err == nil {
		log.Printf("Product with id %s deleted succesfully", id)
	} else {
		log.Println(err)
	}

	return err
}

func NewProRepository(db *sql.DB) ProRepository {
	proRepo := proRepository{
		db,
	}

	return &proRepo
}
