package main

import (
	"fmt"
	"os"
	"time"

	//"github.com/jinzhu/gorm"
	//_ "github.com/lib/pq" // https://stackoverflow.com/questions/52789531/how-do-i-solve-panic-sql-unknown-driver-postgres-forgotten-import

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/pkg/errors"
)

type Customer struct {
	Name      string    `gorm:"column:name;primary_key"`
	Age       int       `gorm:"column:age"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	// TODO: Explore the DeletedAt field...
	//DeletedAt time.Time `gorm:"column:deleted_at"`
	//DeletedAt *time.Time `sql:"index"`
	//DeletedAt time.Time `sql:"index"`
	//DeletedAt time.Time
	//DeletedAt gorm.DeletedAt `gorm:"index"`
}

func main() {
	fmt.Println("Gorm")
	fmt.Println()

	host := "localhost"
	port := "5432"
	user := "johan"
	dbname := "mydb"
	password := "LinuxVM123"
	if len(os.Args) == 6 {
		host = os.Args[1]
		port = os.Args[2]
		user = os.Args[3]
		dbname = os.Args[4]
		password = os.Args[5]
	}

	fmt.Println("Connecting...")
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password)
	fmt.Printf("*** dbURI: %s\n", dbURI)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		handleError("Failed to connect to database.", err)
	}
	defer func() {
		// TODO: Drop table?
		fmt.Println("Closing connection...")
		// There is no "Close" method in Gorm 2 - need to call teh Close method of sql.DB
		sqlDB, err := db.DB()
		if err == nil {
			sqlDB.Close()
		}
	}()
	fmt.Printf("Connected to db: %v\n", db)
	fmt.Println()

	fmt.Println("Creating tables...")
	db.AutoMigrate(&Customer{})
	fmt.Println()

	fmt.Println("Inserting rows...")
	// Gorm functions typically return a value of type *gorm.DB,
	// which contains an Error field that needs to be checked.
	result := db.Create(&Customer{Name: "Alice", Age: 42})
	fmt.Printf("Result: %#v\n", result)
	fmt.Printf("Inserted %d rows\n", result.RowsAffected)
	if result.Error != nil {
		handleError("Failed to insert row.", err)
	}
	// A more idiomatic way to check for errors...
	err = db.Create(&Customer{Name: "Bob", Age: 64}).Error
	if err != nil {
		handleError("Failed to insert row.", err)
	}
	// Even more idiomatic? (I prefer the one above)
	if err = db.Create(&Customer{Name: "Carol", Age: 16}).Error; err != nil {
		handleError("Failed to insert row.", err)
	}
	fmt.Println()

	fmt.Println("Selecting first row...")
	firstCustomer := Customer{}
	result = db.First(&firstCustomer)
	fmt.Printf("Result: %#v\n", result)
	err = result.Error
	if err != nil {
		handleError("Failed to select first row.", err)
	}
	fmt.Printf("Found %d rows\n", result.RowsAffected)
	fmt.Printf("First row: %v\n", firstCustomer)
	fmt.Println()

	fmt.Println("Selecting last row...")
	lastCustomer := Customer{}
	result = db.Last(&lastCustomer)
	fmt.Printf("Result: %#v\n", result)
	err = result.Error
	if err != nil {
		handleError("Failed to select last row.", err)
	}
	fmt.Printf("Found %d rows\n", result.RowsAffected)
	fmt.Printf("Last row: %v\n", lastCustomer)
	fmt.Println()

	fmt.Println("Selecting specific row...")
	name := "Bob"
	customer := Customer{}
	result = db.Model(customer).Where("name = ?", name).Find(&customer)
	fmt.Printf("Result: %#v\n", result)
	err = result.Error
	if err != nil {
		handleError("Failed to select specific row.", err)
	}
	fmt.Printf("Found %d rows\n", result.RowsAffected)
	fmt.Printf("Customer: %v\n", customer)
	fmt.Println()

	fmt.Println("Updating...")
	customer.Age = 65
	result = db.Save(&customer)
	fmt.Printf("Result: %#v\n", result)
	fmt.Printf("Updated %d rows\n", result.RowsAffected)
	if result.Error != nil {
		handleError("Failed to update row.", err)
	}

	// Save will insert a new row if it doesn't exist yet...
	customer.Name = "Dave"
	customer.Age = 32
	result = db.Save(&customer)
	fmt.Printf("Result: %#v\n", result)
	fmt.Printf("Updated %d rows\n", result.RowsAffected)
	if result.Error != nil {
		handleError("Failed to update row.", err)
	}
	fmt.Println()

	fmt.Println("Deleting...")
	customer.Name = "Carol"
	err = delete(db, &customer)
	if err != nil {
		handleError("Failed to delete 'Carol'", err)
	}
	fmt.Println()

	fmt.Println("Deleting in a DB transaction...")
	err = db.Transaction(func(tx *gorm.DB) error {
		// This should work because "Alice" still exists
		fmt.Println("Deleting 'Alice'... (should succeed)")
		customer.Name = "Alice"
		// Note that we pass the tx - not the db
		err = delete(tx, &customer)
		if err != nil {
			return err
		}
		// This should fail because "Edgar" does not exust
		fmt.Println("Deleting 'Edgar'... (should fail, rolling back delete of 'Alice')")
		customer.Name = "Edgar"
		err = delete(tx, &customer)
		if err != nil {
			// Returning an error will rollback the delete of "Alice"
			fmt.Println("Rollback...")
			return err
		}
		// Returning nil will commit the transaction, ie. both the deletes
		fmt.Println("Commit.")
		return nil
	})
	if err != nil {
		//handleError("Failed to delete both 'Alice' and 'Edgar'", err)
		fmt.Printf("Failed to delete both 'Alice' and 'Edgar' Error: %v\n", err)
		fmt.Println("Continuing as error is expected...")
	}
	fmt.Println()

	fmt.Println("Selecting all rows...")
	customers := []Customer{}
	result = db.Find(&customers)
	fmt.Printf("Result: %#v\n", result)
	err = result.Error
	if err != nil {
		handleError("Failed to select all rows.", err)
	}
	fmt.Printf("Found %d rows\n", result.RowsAffected)
	for i, customer := range customers {
		fmt.Printf("  Customer #%d: %v\n", i, customer)
	}
	fmt.Println()

	fmt.Println("Dropping table...")
	// TODO: No "Close" ,ethod in Gorm 2?
	//db.DropTable(&Customer{})
	fmt.Println()

	fmt.Println("Done.")
}

// TableName, if present, is used by Gorm to determine the database table name.
// Otherwise it will derive table name "customers" from the type name "Customer",
// or table name "shopping_baskets" from the type name "ShoppingBasket"
func (a *Customer) TableName() string {
	return "gorm_customers"
}

func handleError(message string, err error) {
	fmt.Print(message)
	fmt.Printf(" Error: %v\n", err)
	os.Exit(1)
}

// Using a separate function because Gorm's Delete function does not return an error
// https://stackoverflow.com/questions/67154864/how-to-handle-gorm-error-at-delete-function
func delete(db *gorm.DB, customer *Customer) error {
	db = db.Delete(&customer)
	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected < 1 {
		return errors.New("row cannot be deleted because it doesn't exist")
	}
	return nil
}
