package repository

import (
	"database/sql"
	"fmt"
	"go-laundry/model"
	"go-laundry/util"
	"time"

	"github.com/google/uuid"
)

type TraRepository interface {
	Create(newTra *model.Bill) error
	FindOne(id string) (*model.Bill, error)
}

type traRepository struct {
	db *sql.DB
}

func (r *traRepository) Create(newBill *model.Bill) error {
	// Mulai transaksi.
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}

	// Lakukan insert ke tabel bill.
	newId := uuid.New().String()
	res, err := tx.Exec("INSERT INTO bill (id, bill_date, finish_date, employee_id, customer_id) VALUES ($1, $2, $3, $4, $5)",
		newId, newBill.BillDate, newBill.FinishDate, newBill.Employee.Id, newBill.Customer.Id)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to insert into bill table: %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to get affected rows: %v", err)
	}
	if rowsAffected == 0 {
		tx.Rollback()
		return fmt.Errorf("failed to insert into bill table: no rows affected")
	}

	// Lakukan insert ke tabel bill_detail.
	for _, item := range newBill.Items {
		newId2 := uuid.New().String()
		res, err = tx.Exec("INSERT INTO bill_detail (id, bill_id, product_id, product_price, qty) VALUES ($1, $2, $3, $4, $5)",
			newId2, newId, item.Product.Id, item.ProductPrice, item.Qty)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to insert into bill_detail table: %v", err)
		}
		rowsAffected, err = res.RowsAffected()
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to get affected rows: %v", err)
		}
		if rowsAffected == 0 {
			tx.Rollback()
			return fmt.Errorf("failed to insert into bill_detail table: no rows affected")
		}
	}

	// Commit transaksi.
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (r *traRepository) FindOne(customerID string) (*model.Bill, error) {
	rows, err := r.db.Query(`
		SELECT
			b.bill_date,
			b.finish_date,
			e.name AS employee_name,
			c.name AS customer_name,
			c.phone_number AS customer_phone,
			bd.id AS bill_detail_id,
			p.name AS product_name,
			bd.qty,
			u.name AS uom_name,
			p.price AS product_price,
			bd.product_price AS bill_detail_product_price
		FROM
			bill b
			JOIN employee e ON e.id = b.employee_id
			JOIN customer c ON c.id = b.customer_id
			JOIN bill_detail bd ON bd.bill_id = b.id
			JOIN product p ON p.id = bd.product_id
			JOIN uom u ON u.id = p.uom_id
		WHERE
			b.customer_id = $1
			AND c.is_deleted = false
		ORDER BY
			b.bill_date DESC;
	`, customerID)
	util.CheckErr(err)
	defer rows.Close()

	var (
		bill          *model.Bill
		billItems     []model.Bill_detail
		billDate      time.Time
		finishDate    sql.NullTime
		employeeName  string
		customerName  string
		customerPhone string
	)

	for rows.Next() {
		var (
			billDetailID             string
			productName              string
			qty                      int
			uomName                  string
			productPrice             float64
			billDetailProductPrice   float64
			billDetailProductPrice64 int64
		)
		if err := rows.Scan(
			&billDate,
			&finishDate,
			&employeeName,
			&customerName,
			&customerPhone,
			&billDetailID,
			&productName,
			&qty,
			&uomName,
			&productPrice,
			&billDetailProductPrice,
		); err != nil {
			return nil, err
		}

		// Convert float64 billDetailProductPrice to int64
		billDetailProductPrice64 = int64(billDetailProductPrice)

		// Create new bill item and append to bill items slice
		billItems = append(billItems, model.Bill_detail{
			Id: billDetailID,
			Product: model.Product{
				Name:  productName,
				Uom:   model.Uom{Name: uomName},
				Price: fmt.Sprintf("%.2f", productPrice),
			},
			Qty:          qty,
			ProductPrice: billDetailProductPrice64,
		})
	}

	// Check if there's any error while reading rows
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// If there's no data returned from query, return error
	if len(billItems) == 0 {
		return nil, fmt.Errorf("no bill found")
	}

	// Create new bill object with customer, employee, and items detail
	bill = &model.Bill{
		BillDate:   billDate,
		FinishDate: finishDate.Time.Format("2006-01-02"),
		Employee: model.Employee{
			Name: employeeName,
		},
		Customer: model.Customer{
			Name:        customerName,
			PhoneNumber: customerPhone,
		},
		Items: billItems,
	}

	return bill, nil
}

func NewTraRepository(db *sql.DB) TraRepository {
	traRepo := traRepository{
		db,
	}

	return &traRepo
}
