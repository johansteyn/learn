package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq" // https://stackoverflow.com/questions/52789531/how-do-i-solve-panic-sql-unknown-driver-postgres-forgotten-import
)

type Application struct {
	KittyTenantID       string    `gorm:"column:kitty_tenant_id;primary_key"`
	AcmePartnerTenantID string    `gorm:"column:acme_partner_tenant_id;unique_index:acme_tenant"`
	RedirectURL         string    `gorm:"column:redirect_url"`
	CompanyName         string    `gorm:"column:company_name"`
	Version             string    `gorm:"column:application_version"`
	RequiredVSAVersion  string    `gorm:"column:required_vsa_version"`
	Copyright           string    `gorm:"column:application_copyright"`
	ID                  uint32    `gorm:"column:application_id"`
	KittyServerAddress  string    `gorm:"column:authorization_server"`
	Authorized          bool      `gorm:"column:authorized"`
	TicketSync          bool      `gorm:"column:ticket_sync"`
	CreatedAt           time.Time `gorm:"column:created_at"`
}

type OrganizationLink struct {
	KittyTenantID         string    `gorm:"column:kitty_tenant_id;primary_key"`
	KittyOrganizationID   string    `gorm:"column:kitty_organization_id;primary_key;unique"`
	KittyOrganizationName string    `gorm:"column:kitty_organization_name"`
	AcmePartnerID         string    `gorm:"column:acme_partner_id;primary_key"`
	AcmeCustomerID        string    `gorm:"column:acme_customer_id;primary_key"`
	AcmeCustomerName      string    `gorm:"column:acme_customer_name"`
	ProtectionEdition     string    `gorm:"column:acme_protection_edition"`
	InstallAgents         bool      `gorm:"column:install_agents"`
	ApplyDefaultPlan      bool      `gorm:"column:apply_default_protection_plans"`
	AlertsSynchronization bool      `gorm:"column:alerts_synchronization"`
	Devices               uint      `gorm:"column:devices"`
	Protected             uint      `gorm:"column:protected"`
	ID                    string    `gorm:"column:id;unique"`
	CreatedAt             time.Time `gorm:"column:created_at"`
	UpdatedAt             time.Time `gorm:"column:updated_at"`
	//DeletedAt              time.Time `gorm:"column:deleted_at"`
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
	user := "ci_kitty"
	dbname := "cyber_ci_kitty"
	password := "8f40a8a58506fc51b0d2ce1f30be3d08"
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
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Printf("Failed to connect to database. Error: %v\n", err)
		return
	}
	fmt.Printf("Connected to db: %v\n", db)

	kittyTenantID := "99999999999999999999999999"
	fmt.Println("Selecting application...")
	application := Application{}
	err = db.Model(application).Where("kitty_tenant_id = ?", kittyTenantID).Find(&application).Error
	if err != nil {
		fmt.Printf("Failed to select. Error: %v\n", err)
		return
	}
	fmt.Printf("Selected application: %v\n", application)

	fmt.Println("Selecting organization link...")
	organizationLink := OrganizationLink{}
	err = db.Model(organizationLink).Where("kitty_tenant_id = ?", kittyTenantID).Find(&organizationLink).Error
	if err != nil {
		fmt.Printf("Failed to select. Error: %v\n", err)
		return
	}
	fmt.Printf("Selected organization link: %v\n", organizationLink)

}

// TableName, if present, is used by Gorm to determine the database table name.
// Otherwise it will derive table name "applications" from the type name "Application"
func (a *Application) TableName() string {
	return "kitty_applications"
}

func (o *OrganizationLink) TableName() string {
	return "organization_links"
}
