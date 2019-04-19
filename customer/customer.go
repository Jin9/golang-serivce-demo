package customer

import (
	"database/sql"
	"fmt"

	"gitlab.com/chinnawat.w/golang-service-demo/model"

	// pq is a postgres driver
	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "changeme"
	dbname   = "profile"
)

const (
	insertCustomer              = `INSERT INTO users.customers (token, name, age, email, phone) VALUES ($1, $2, $3, $4, $5);`
	findTokenByToken            = `SELECT customer.token FROM users.customers as customer where customer.token='$1';`
	findCustomerDetailByToken   = `SELECT customer.name, customer.age, customer.email, customer.phone FROM users.customers as customer where customer.token=$1`
	findAllCustomer             = `SELECT customer.token, customer.name, customer.age, customer.email, customer.phone FROM users.customers as customer order by customer.id asc`
	findIDByToken               = `SELECT customer.id FROM users.customers as customer where customer.token=$1`
	updateCustomerDetailByToken = `UPDATE users.customers SET name=$2, age=$3, email=$4, phone=$5 WHERE token=$1;`
	updateTokenByID             = `UPDATE users.customers SET token=$2 WHERE id=$1;`
	deleteCustomerByToken       = `DELETE FROM users.customers as customer where customer.token=$1;`
)

func connectDB() (db *sql.DB, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// CheckDuplicateToken is a function to check duplicate token before create new member detail
func CheckDuplicateToken(token string) (err error) {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Query(findTokenByToken, token)
	if err != nil {
		return err
	}

	return nil
}

// InsertCustomer is used to insert new detail to db
func InsertCustomer(token string, customer *model.Customer) (err error) {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStatement := insertCustomer

	_, err = db.Exec(sqlStatement, token, customer.Name, customer.Age, customer.Email, customer.Phone)
	if err != nil {
		return err
	}

	return nil
}

func mapCustomerDetail(row *sql.Row) (customer *model.Customer, err error) {
	var name, email, phone string
	var age int

	switch err := row.Scan(&name, &age, &email, &phone); err {
	case sql.ErrNoRows:
		return nil, sql.ErrNoRows
	case nil:
		customer = model.NewCustomer(name, age, email, phone)
		break
	default:
		return nil, err
	}

	return customer, nil
}

// FindCustomerDetailByToken is used to find customer detail by token
func FindCustomerDetailByToken(token string) (customer *model.Customer, err error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow(findCustomerDetailByToken, token)
	customer, err = mapCustomerDetail(row)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func mapAllCustomersDetail(rows *sql.Rows) (customers []*model.CustomerDetail, err error) {
	for rows.Next() {
		var token, name, email, phone string
		var age int
		err = rows.Scan(&token, &name, &age, &email, &phone)
		if err != nil {
			return nil, err
		}
		customer := model.NewCustomerDetail(token, model.NewCustomer(name, age, email, phone))
		customers = append(customers, customer)
	}

	return customers, nil
}

// FindAllCustomerDetail is used to show all customers detail
func FindAllCustomerDetail() (customers []*model.CustomerDetail, err error) {
	db, err := connectDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(findAllCustomer)
	if err != nil {
		return nil, err
	}

	customers, err = mapAllCustomersDetail(rows)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

// UpdateCustomerDetail is used for update Customer data
func UpdateCustomerDetail(token string, customer *model.Customer) (err error) {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStatement := updateCustomerDetailByToken
	_, err = db.Exec(sqlStatement, token, customer.Name, customer.Age, customer.Email, customer.Phone)
	if err != nil {
		return err
	}
	return nil
}

func mapCustomerID(row *sql.Row) (id uint64, err error) {

	switch err := row.Scan(&id); err {
	case sql.ErrNoRows:
		return 0, sql.ErrNoRows
	case nil:
		break
	default:
		return 0, err
	}

	return id, nil
}

func findCustomerIDByToken(token string) (id uint64, err error) {
	db, err := connectDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	row := db.QueryRow(findIDByToken, token)
	id, err = mapCustomerID(row)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UpdateCustomerToken is used for update customer token after update customer detail
func UpdateCustomerToken(existToken string, newToken string) (err error) {
	id, err := findCustomerIDByToken(existToken)
	if err != nil {
		return err
	}

	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStatement := updateTokenByID
	_, err = db.Exec(sqlStatement, id, newToken)
	if err != nil {
		return err
	}

	return nil
}

// DeleteCustomerByToken is used for delete customer
func DeleteCustomerByToken(token string) (err error) {
	db, err := connectDB()
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStatement := deleteCustomerByToken
	_, err = db.Exec(sqlStatement, token)
	if err != nil {
		return err
	}

	return nil
}
