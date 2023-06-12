package delivery

import (
	"bufio"
	"fmt"
	"go-laundry/config"
	"go-laundry/db"
	"go-laundry/model"
	"go-laundry/repository"
	"go-laundry/usecase"
	"go-laundry/util"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Console struct {
	uomUsecase usecase.UomUsecase
	cusUsecase usecase.CusUsecase
	empUsecase usecase.EmpUsecase
	proUsecase usecase.ProUsecase
	traUsecase usecase.TraUsecase
}

func (c *Console) MainMenuForm() {
	fmt.Println(strings.Repeat("=", 30))
	fmt.Println("Enigma Laundry")
	fmt.Println(strings.Repeat("=", 30))

	fmt.Println("1. Master UOM")
	fmt.Println("2. Master Product")
	fmt.Println("3. Master Employee")
	fmt.Println("4. Master Customer")
	fmt.Println("5. Menu Transaksi")
	fmt.Println("0. Exit")
	fmt.Print("Pilih Menu (0-5): ")
}

func (c *Console) UomMaster() {
	fmt.Println(strings.Repeat("=", 30))
	fmt.Println("Enigma Laundry - Master UOM")
	fmt.Println(strings.Repeat("=", 30))
	fmt.Println("1. CREATE")
	fmt.Println("2. READ All")
	fmt.Println("3. READ ONE")
	fmt.Println("4. Update")
	fmt.Println("5. Delete")
	fmt.Println("0. Back to main menu")
	fmt.Print("Pilih Menu (0-5): ")

	for {
		var selectedUom string
		fmt.Scanln(&selectedUom)
		switch selectedUom {
		case "1":
			//add
			newUom, ok := c.uomCreateForm()
			if ok {
				err := c.uomUsecase.Register(&newUom)
				util.CheckErr(err)
			}
			c.UomMaster()
		case "2":
			//viewall
			uoms, err := c.uomUsecase.FindAll()
			util.CheckErr(err)
			if len(uoms) == 0 {
				fmt.Println("Belum ada data UOM")
			} else {
				fmt.Println("Daftar UOM:")
				for _, uom := range uoms {
					fmt.Printf("%s - %s\n", uom.Id, uom.Name)
				}
			}
			c.UomMaster()
		case "3":
			//viewbyid
			var id string
			fmt.Print("Masukkan ID UOM: ")
			fmt.Scanln(&id)
			uom, err := c.uomUsecase.FindOne(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("UOM ditemukan: %s - %s\n", uom.Id, uom.Name)
			}
			c.UomMaster()
		case "4":
			//update
			var id string
			fmt.Print("Masukkan ID UOM yang ingin diupdate: ")
			fmt.Scanln(&id)
			uom, err := c.uomUsecase.FindOne(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Data UOM yang akan diupdate: %s - %s\n", uom.Id, uom.Name)
				reader := bufio.NewReader(os.Stdin)
				fmt.Print("UOM Name: ")
				uomName, _ := reader.ReadString('\n')
				uomName = strings.TrimSuffix(uomName, "\n")
				fmt.Printf("UOM %s akan diubah? (y/t): ", uomName)
				saveConfirmation, _ := reader.ReadString('\n')
				saveConfirmation = strings.TrimSpace(saveConfirmation)
				if saveConfirmation == "y" {
					uom.Name = uomName
					err = c.uomUsecase.Edit(uom)
					util.CheckErr(err)
					fmt.Println("Data UOM berhasil diupdate")
				} else if saveConfirmation == "t" {
					fmt.Println("Data tidak jadi disimpan")
				} else {
					fmt.Println("Input tidak valid, data tidak jadi disimpan")
				}
			}
			c.UomMaster()
		case "5":
			var id string
			fmt.Print("Masukkan ID UOM yang ingin dihapus: ")
			fmt.Scanln(&id)
			err := c.uomUsecase.Unreg(id)
			util.CheckErr(err)
			fmt.Println("Data UOM berhasil dihapus")
			c.UomMaster()
		case "0":
			c.MainMenuForm()
			return
		default:
			fmt.Println("Menu tidak tersedia")
			c.UomMaster()
		}
		break
	}
}

func (c *Console) uomCreateForm() (model.Uom, bool) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("UOM Name: ")
	uomName, _ := reader.ReadString('\n')
	uomName = strings.TrimSuffix(uomName, "\n")

	fmt.Printf("UOM %s akan disimpan? (y/t): ", uomName)
	saveConfirmation, _ := reader.ReadString('\n')
	saveConfirmation = strings.TrimSpace(saveConfirmation)

	if saveConfirmation == "y" {
		var uom model.Uom
		uom.Name = uomName
		return uom, true
	} else if saveConfirmation == "t" {
		fmt.Println("Data tidak jadi disimpan")
	} else {
		fmt.Println("Input tidak valid, data tidak jadi disimpan")
	}
	return model.Uom{}, false
}

func (c *Console) CusMaster() {
	fmt.Println(strings.Repeat("=", 30))
	fmt.Println("Enigma Laundry - Master Customer")
	fmt.Println(strings.Repeat("=", 30))
	fmt.Println("1. CREATE")
	fmt.Println("2. READ All")
	fmt.Println("3. READ ONE")
	fmt.Println("4. Update")
	fmt.Println("5. Delete")
	fmt.Println("0. Back to main menu")
	fmt.Print("Pilih Menu (0-5): ")

	for {
		var selectedCus string
		fmt.Scanln(&selectedCus)
		switch selectedCus {
		case "1":
			//add
			newCus, ok := c.cusCreateForm()
			if ok {
				err := c.cusUsecase.Register(&newCus)
				util.CheckErr(err)
			}
			c.CusMaster()
		case "2":
			//viewall
			cuss, err := c.cusUsecase.FindAll()
			util.CheckErr(err)
			if len(cuss) == 0 {
				fmt.Println("Belum ada data Customer")
			} else {
				fmt.Println("Daftar Customer:")
				for _, cus := range cuss {
					fmt.Printf("%s - %s - %s \n", cus.Id, cus.Name, cus.PhoneNumber)
				}
			}
			c.CusMaster()
		case "3":
			//viewbyid
			var id string
			fmt.Print("Masukkan ID Customer: ")
			fmt.Scanln(&id)
			cus, err := c.cusUsecase.FindOne(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Customer ditemukan: %s - %s - %s\n", cus.Id, cus.Name, cus.PhoneNumber)
			}
			c.CusMaster()
		case "4":
			//update
			var id string
			fmt.Print("Masukkan ID Customer yang ingin diupdate: ")
			fmt.Scanln(&id)
			cus, err := c.cusUsecase.FindOne(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Data Customer yang akan diupdate: %s - %s - %s \n", cus.Id, cus.Name, cus.PhoneNumber)
				reader := bufio.NewReader(os.Stdin)

				fmt.Print("Customer Name: ")
				cusName, _ := reader.ReadString('\n')
				cusName = strings.TrimSuffix(cusName, "\n")
				fmt.Print("Customer PhoneNumber: ")
				cusPhoneNumber, _ := reader.ReadString('\n')
				cusPhoneNumber = strings.TrimSuffix(cusPhoneNumber, "\n")

				fmt.Printf("Customer %s & %s akan diubah? (y/t): ", cusName, cusPhoneNumber)
				saveConfirmation, _ := reader.ReadString('\n')
				saveConfirmation = strings.TrimSpace(saveConfirmation)
				if saveConfirmation == "y" {
					cus.Name = cusName
					cus.PhoneNumber = cusPhoneNumber
					err = c.cusUsecase.Edit(cus)
					util.CheckErr(err)
					fmt.Println("Data Customer berhasil diupdate")
				} else if saveConfirmation == "t" {
					fmt.Println("Data tidak jadi disimpan")
				} else {
					fmt.Println("Input tidak valid, data tidak jadi disimpan")
				}
			}
			c.CusMaster()
		case "5":
			var id string
			fmt.Print("Masukkan ID Customer yang ingin dihapus: ")
			fmt.Scanln(&id)
			err := c.cusUsecase.Unreg(id)
			util.CheckErr(err)
			fmt.Println("Data Customer berhasil dihapus")
			c.CusMaster()
		case "0":
			c.MainMenuForm()
			return
		default:
			fmt.Println("Menu tidak tersedia")
			c.CusMaster()
		}
		break
	}
}

func (c *Console) cusCreateForm() (model.Customer, bool) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Customer Name: ")
	cusName, _ := reader.ReadString('\n')
	cusName = strings.TrimSuffix(cusName, "\n")
	fmt.Print("Customer PhoneNumber: ")
	cusPhoneNumber, _ := reader.ReadString('\n')
	cusPhoneNumber = strings.TrimSuffix(cusPhoneNumber, "\n")

	fmt.Printf("Customer %s & %s akan disimpan? (y/t): ", cusName, cusPhoneNumber)
	saveConfirmation, _ := reader.ReadString('\n')
	saveConfirmation = strings.TrimSpace(saveConfirmation)

	if saveConfirmation == "y" {
		var cus model.Customer
		cus.Name = cusName
		cus.PhoneNumber = cusPhoneNumber
		return cus, true
	} else if saveConfirmation == "t" {
		fmt.Println("Data tidak jadi disimpan")
	} else {
		fmt.Println("Input tidak valid, data tidak jadi disimpan")
	}
	return model.Customer{}, false
}

func (c *Console) EmpMaster() {
	fmt.Println(strings.Repeat("=", 30))
	fmt.Println("Enigma Laundry - Master Employee")
	fmt.Println(strings.Repeat("=", 30))
	fmt.Println("1. CREATE")
	fmt.Println("2. READ All")
	fmt.Println("3. READ ONE")
	fmt.Println("4. Update")
	fmt.Println("5. Delete")
	fmt.Println("0. Back to main menu")
	fmt.Print("Pilih Menu (0-5): ")

	for {
		var selectedEmp string
		fmt.Scanln(&selectedEmp)
		switch selectedEmp {
		case "1":
			//add
			newEmp, ok := c.empCreateForm()
			if ok {
				err := c.empUsecase.Register(&newEmp)
				util.CheckErr(err)
			}
			c.EmpMaster()
		case "2":
			//viewall
			emps, err := c.empUsecase.FindAll()
			util.CheckErr(err)
			if len(emps) == 0 {
				fmt.Println("Belum ada data Employee")
			} else {
				fmt.Println("Daftar Employee:")
				for _, emp := range emps {
					fmt.Printf("%s - %s\n", emp.Id, emp.Name)
				}
			}
			c.EmpMaster()
		case "3":
			//viewbyid
			var id string
			fmt.Print("Masukkan ID Employee: ")
			fmt.Scanln(&id)
			emp, err := c.empUsecase.FindOne(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Employee ditemukan: %s - %s\n", emp.Id, emp.Name)
			}
			c.EmpMaster()
		case "4":
			//update
			var id string
			fmt.Print("Masukkan ID Employee yang ingin diupdate: ")
			fmt.Scanln(&id)
			emp, err := c.empUsecase.FindOne(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Data Employee yang akan diupdate: %s - %s\n", emp.Id, emp.Name)
				reader := bufio.NewReader(os.Stdin)

				fmt.Print("Employee Name: ")
				empName, _ := reader.ReadString('\n')
				empName = strings.TrimSuffix(empName, "\n")

				fmt.Printf("Employee %s akan diubah? (y/t): ", empName)
				saveConfirmation, _ := reader.ReadString('\n')
				saveConfirmation = strings.TrimSpace(saveConfirmation)

				if saveConfirmation == "y" {
					emp.Name = empName
					err = c.empUsecase.Edit(emp)
					util.CheckErr(err)
					fmt.Println("Data Employee berhasil diupdate")
				} else if saveConfirmation == "t" {
					fmt.Println("Data tidak jadi disimpan")
				} else {
					fmt.Println("Input tidak valid, data tidak jadi disimpan")
				}
			}
			c.EmpMaster()
		case "5":
			var id string
			fmt.Print("Masukkan ID Employee yang ingin dihapus: ")
			fmt.Scanln(&id)
			err := c.empUsecase.Unreg(id)
			util.CheckErr(err)
			fmt.Println("Data Employee berhasil dihapus")
			c.EmpMaster()
		case "0":
			c.MainMenuForm()
			return
		default:
			fmt.Println("Menu tidak tersedia")
			c.EmpMaster()
		}
		break
	}
}

func (c *Console) empCreateForm() (model.Employee, bool) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Employee Name: ")
	empName, _ := reader.ReadString('\n')
	empName = strings.TrimSuffix(empName, "\n")

	fmt.Printf("Employee %s akan disimpan? (y/t): ", empName)
	saveConfirmation, _ := reader.ReadString('\n')
	saveConfirmation = strings.TrimSpace(saveConfirmation)

	if saveConfirmation == "y" {
		var emp model.Employee
		emp.Name = empName
		return emp, true
	} else if saveConfirmation == "t" {
		fmt.Println("Data tidak jadi disimpan")
	} else {
		fmt.Println("Input tidak valid, data tidak jadi disimpan")
	}
	return model.Employee{}, false
}

func (c *Console) ProMaster() {
	fmt.Println(strings.Repeat("=", 30))
	fmt.Println("Enigma Laundry - Master Product")
	fmt.Println(strings.Repeat("=", 30))
	fmt.Println("1. CREATE")
	fmt.Println("2. READ ONE")
	fmt.Println("3. READ All")
	fmt.Println("4. Update")
	fmt.Println("5. Delete")
	fmt.Println("0. Back to main menu")
	fmt.Print("Pilih Menu (0-5): ")

	for {
		var selectedProduk string
		fmt.Scanln(&selectedProduk)
		switch selectedProduk {
		case "1":
			newProduct, ok := c.proCreateForm()
			if ok {
				err := c.proUsecase.Register(&newProduct)
				util.CheckErr(err)
			}
			c.ProMaster()
		case "2":
			var id string
			fmt.Print("Masukkan ID Produk: ")
			fmt.Scanln(&id)
			product, err := c.proUsecase.FindOne(id)
			util.CheckErr(err)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("Produk terdaftar : %s - %s - %s - %s \n", product.Id, product.Name, product.Price, product.Uom.Id)
			}
			c.ProMaster()
		case "3":
			product, err := c.proUsecase.FindAll()
			util.CheckErr(err)
			if len(product) == 0 {
				fmt.Println("Belum ada data Produk!")
			} else {
				fmt.Println("Daftar Produk:")
				for _, product := range product {
					fmt.Printf("%s - %s - %s - %s \n", product.Id, product.Name, product.Price, product.Uom.Id)
				}
			}
			c.ProMaster()
		case "4":
			var id string
			fmt.Print("Masukkan ID Produk yang ingin diupdate: ")
			fmt.Scanln(&id)
			product, err := c.proUsecase.FindOne(id)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Data Produk yang akan diupdate: %s - %s - %s - %s \n", product.Id, product.Name, product.Price, product.Uom.Id)

			var ProductName, ProductPrice, ProductUom_Id string
			var saveConfirmation string
			fmt.Print("Product Name: ")
			scanner := bufio.NewScanner(os.Stdin)
			if scanner.Scan() {
				ProductName = scanner.Text()
			}
			fmt.Print("Product Price: ")
			fmt.Scanln(&ProductPrice)
			fmt.Print("Product Uom_Id: ")
			fmt.Scanln(&ProductUom_Id)
			fmt.Printf("Product %s - %s - %s akan disimpan? (y/t): ", ProductName, ProductPrice, ProductUom_Id)
			fmt.Scanln(&saveConfirmation)
			if saveConfirmation == "y" {
				product.Name = ProductName
				product.Price = ProductPrice
				product.Uom.Id = ProductUom_Id
				err = c.proUsecase.Edit(product)
				util.CheckErr(err)
			} else if saveConfirmation == "t" {
				fmt.Println("Data tidak jadi disimpan")
			} else {
				fmt.Println("Input tidak valid, data tidak jadi disimpan")
			}

			fmt.Println("Update Berhasil!")
			c.ProMaster()
		case "5":
			var id string
			fmt.Print("Masukkan ID Produk yang ingin dihapus: ")
			fmt.Scanln(&id)
			err := c.proUsecase.Unreg(id)
			util.CheckErr(err)
			fmt.Println("Data Produk berhasil dihapus !")
			c.ProMaster()
		case "0":
			c.MainMenuForm()
			return
		default:
			fmt.Println("Menu tidak tersedia")
		}
		break
	}

}

func (c *Console) proCreateForm() (model.Product, bool) {
	var ProductName, ProductPrice, ProductUom_Id string
	var saveConfirmation string
	fmt.Print("Product Name: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		ProductName = scanner.Text()
	}
	fmt.Print("Product Price: ")
	fmt.Scanln(&ProductPrice)
	fmt.Print("Product Uom_Id: ")
	fmt.Scanln(&ProductUom_Id)
	fmt.Printf("Product %s - %s - %s akan disimpan? (y/t): ", ProductName, ProductPrice, ProductUom_Id)
	fmt.Scanln(&saveConfirmation)

	if saveConfirmation == "y" {
		var product model.Product
		product.Name = ProductName
		product.Price = ProductPrice
		product.Uom.Id = ProductUom_Id
		return product, true
	} else if saveConfirmation == "t" {
		fmt.Println("Data tidak jadi disimpan")
	} else {
		fmt.Println("Input tidak valid, data tidak jadi disimpan")
	}
	return model.Product{}, false

}

func (c *Console) TraMaster() {
	fmt.Println(strings.Repeat("=", 30))
	fmt.Println("Enigma Laundry - Transaction")
	fmt.Println(strings.Repeat("=", 30))
	fmt.Println("1. CREATE")
	fmt.Println("2. READ ONE")
	fmt.Println("0. Back to main menu")
	fmt.Print("Pilih Menu (0-2): ")

	for {
		var selectedProduk string
		fmt.Scanln(&selectedProduk)
		switch selectedProduk {
		case "1":
			c.traCreateForm()
		case "2":
			c.traReadForm()
		case "0":
			c.MainMenuForm()
			return
		default:
			fmt.Println("Menu tidak tersedia")
		}
		break
	}
}

func (c *Console) traCreateForm() model.Bill {
	// Baca input dari user.
	var newBill model.Bill

	//menampilkan list dari customer
	fmt.Println("List Customer:")
	customers, err := c.cusUsecase.FindAll()
	util.CheckErr(err)
	for _, customer := range customers {
		fmt.Printf("%s - %s - %s\n", customer.Id, customer.Name, customer.PhoneNumber)
	}
	var CustomerID string
	fmt.Print("Masukan ID Customer: ")
	fmt.Scanln(&CustomerID)

	//menampilkan daftar product
	fmt.Println("List Product:")
	products, err := c.proUsecase.FindAll()
	util.CheckErr(err)
	for _, product := range products {
		fmt.Printf("%s - %s - %s\n", product.Id, product.Name, product.Price)
	}
	var ProductID string
	fmt.Print("Masukan ID Product: ")
	fmt.Scanln(&ProductID)

	//finish date
	fmt.Print("Tanggal Selesai (format: yyyy-mm-dd): ")
	fmt.Scanln(&newBill.FinishDate)

	//qty
	var quantity int
	for {
		fmt.Print("Masukan quantity (minimal 1): ")
		_, err := fmt.Scanln(&quantity)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if quantity < 1 {
			fmt.Println("Quantity harus minimal 1")
			continue
		}
		break
	}

	//menampilkan daftar employee
	fmt.Println("List Employee:")
	employees, err := c.empUsecase.FindAll()
	util.CheckErr(err)
	for _, emnploye := range employees {
		fmt.Printf("%s - %s \n", emnploye.Id, emnploye.Name)
	}
	var EmployeeID string
	fmt.Print("Masukan ID Employee: ")
	fmt.Scanln(&EmployeeID)

	Billdatee := time.Now()

	//menghitung total harga
	var hargaSatuan float64
	for _, product := range products {
		if product.Id == ProductID {
			price, err := strconv.ParseFloat(product.Price, 64)
			if err != nil {
				fmt.Println("Gagal melakukan konversi")
			}
			hargaSatuan = price
			break
		}
	}
	totalHarga := hargaSatuan * float64(quantity)

	item := model.Bill_detail{
		Id: uuid.New().String(),
		Bill: model.Bill{
			Id: newBill.Id,
		},
		Product: model.Product{
			Id: ProductID,
		},
		ProductPrice: int64(totalHarga),
		Qty:          quantity,
	}

	// membuat transaksi
	newTra := model.Bill{
		Id:         uuid.New().String(),
		BillDate:   Billdatee,
		FinishDate: newBill.FinishDate,
		Employee: model.Employee{
			Id: EmployeeID,
		},
		Customer: model.Customer{
			Id: CustomerID,
		},
		Items: []model.Bill_detail{item},
	}
	item.Bill = newTra
	item.Product = model.Product{
		Id:    ProductID,
		Name:  "",
		Price: "",
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Transaksi akan disimpan? (y/t): ")
	saveConfirmation, _ := reader.ReadString('\n')
	saveConfirmation = strings.TrimSpace(saveConfirmation)

	if saveConfirmation == "y" {
		err = c.traUsecase.Register(&newTra)
		util.CheckErr(err)
		fmt.Println("Transaksi berhasil disimpan")
	} else if saveConfirmation == "t" {
		fmt.Println("Transaksi tidak jadi disimpan")
	} else {
		fmt.Println("Input tidak valid, transaksi tidak jadi disimpan")
	}
	c.TraMaster()
	return newBill

}
func (c *Console) traReadForm() {
	fmt.Println(strings.Repeat("=", 30))
	fmt.Println("Enigma Laundry - Transaction")
	fmt.Println(strings.Repeat("=", 30))

	var customerID string
	fmt.Print("Masukan ID Customer: ")
	fmt.Scanln(&customerID)

	bill, err := c.traUsecase.FindOne(customerID)
	util.CheckErr(err)

	c.showBillDetail(bill)
}

func (c *Console) showBillDetail(bill *model.Bill) {
	fmt.Printf("Tanggal Transaksi : %v\n", bill.BillDate)
	fmt.Printf("Tanggal Selesai   : %v\n", bill.FinishDate)
	fmt.Printf("Nama Karyawan     : %s\n", bill.Employee.Name)
	fmt.Printf("Nama Customer     : %s\n", bill.Customer.Name)
	fmt.Printf("Nomor Telepon     : %s\n", bill.Customer.PhoneNumber)
	fmt.Println(strings.Repeat("-", 80))
	fmt.Printf("%-10s| %-20s| %-6s| %-20s| %-10s| %-10s\n", "ID", "Nama Produk", "Qty", "Satuan", "Harga", "Total Harga")
	fmt.Println(strings.Repeat("-", 80))
	for _, item := range bill.Items {
		fmt.Printf("%-10s| %-20s| %-6d| %-20s| %-10s| %-10d\n", item.Id, item.Product.Name, item.Qty, item.Product.Uom.Name, item.Product.Price, item.ProductPrice)
	}
	fmt.Println(strings.Repeat("-", 80))
	// fmt.Printf("%76s %-10d\n", "Total Harga : ", bill.TotalPrice())
	c.TraMaster()
}

func (c *Console) Run() {
	c.MainMenuForm()
	for {
		var selectedMenu string
		fmt.Scanln(&selectedMenu)
		switch selectedMenu {
		case "1":
			c.UomMaster()
		case "2":
			c.ProMaster()
		case "3":
			c.EmpMaster()
		case "4":
			c.CusMaster()
		case "5":
			c.TraMaster()
		case "0":
			os.Exit(0)
		default:
			fmt.Println("Menu tidak tersedia")
		}
	}
}

func NewConsole() *Console {
	// konfigurasi
	config := config.NewConfig()

	// koneksi ke database
	dbManager := db.NewDbManager(config)

	// repo
	uomRepo := repository.NewUomRepository(dbManager.ConnectDb())
	uomUsecase := usecase.NewUomUsecase(uomRepo)

	empRepo := repository.NewEmpRepository(dbManager.ConnectDb())
	empUsecase := usecase.NewEmpUsecase(empRepo)

	cusRepo := repository.NewCusRepository(dbManager.ConnectDb())
	cusUsecase := usecase.NewCusUsecase(cusRepo)

	proRepo := repository.NewProRepository(dbManager.ConnectDb())
	proUsecase := usecase.NewProUsecase(proRepo)

	traRepo := repository.NewTraRepository(dbManager.ConnectDb())
	traUsecase := usecase.NewTraUsecase(traRepo)

	return &Console{
		uomUsecase,
		cusUsecase,
		empUsecase,
		proUsecase,
		traUsecase,
	}

}
