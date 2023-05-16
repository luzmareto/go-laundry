package repository

import (
	"database/sql"
	"fmt"
	"go-enigma-laundry/model"
	"log"

	"github.com/google/uuid"
)

type EmpRepository interface {
	Create(newEmp *model.Employee) error
	FindAll() ([]model.Employee, error)
	FindOne(id string) (*model.Employee, error)
	Update(updatedEmp *model.Employee) error
	Delete(id string) error
}

type empRepository struct {
	db *sql.DB
}

func (r *empRepository) Create(newEmp *model.Employee) error {
	if newEmp.Name == "" {
		return fmt.Errorf("UOM Name cannot be empty")
	}
	newId := uuid.New().String()
	query := `INSERT INTO employee(id,name) Values ($1,$2)`
	_, err := r.db.Exec(query, newId, newEmp.Name)

	if err != nil {
		log.Println(err)
	} else {
		log.Printf("UOM %s added successfully", newEmp.Name)
	}

	return err
}

func (r *empRepository) FindAll() ([]model.Employee, error) {
	query := `SELECT id, name FROM employee`
	rows, err := r.db.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	empList := []model.Employee{}
	for rows.Next() {
		var emp model.Employee
		err := rows.Scan(&emp.Id, &emp.Name)
		if err != nil {
			return nil, err
		}
		empList = append(empList, emp)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return empList, err
}

func (r *empRepository) FindOne(id string) (*model.Employee, error) {
	query := `SELECT name FROM employee WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var emp model.Employee
	err := row.Scan(&emp.Name)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("Employee with id %s not found", id)
	} else if err != nil {
		return nil, err
	}

	emp.Id = id

	return &emp, nil
}

func (r *empRepository) Update(updatedEmp *model.Employee) error {
	query := `UPDATE employee SET name = $1 WHERE id = $2`
	_, err := r.db.Exec(query, updatedEmp.Name, updatedEmp.Id)

	if err == nil {
		log.Printf("Employee with id %s updated succesfully", updatedEmp.Id)
	} else {
		log.Println(err)
	}

	return err
}

func (r *empRepository) Delete(id string) error {
	query := `UPDATE employee SET is_deleted = true WHERE id = $1`
	_, err := r.db.Exec(query, id)

	if err == nil {
		log.Printf("Employee with id %s deleted succesfully", id)
	} else {
		log.Println(err)
	}

	return err
}

func NewEmpRepository(db *sql.DB) EmpRepository {
	empRepo := empRepository{
		db,
	}

	return &empRepo
}
