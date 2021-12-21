package ports

import "github.com/pjover/sam/internal/core/model"

type DbService interface {
	FindCustomer(code int) (model.Customer, error)
	FindChild(code int) (model.Child, error)
	FindInvoice(code string) (model.Invoice, error)
	FindProduct(code string) (model.Product, error)
	FindAllProducts() ([]model.Product, error)
	FindInvoicesByYearMonth(yearMonth string) ([]model.Invoice, error)
	FindInvoicesByCustomer(customerCode int) ([]model.Invoice, error)
	FindInvoicesByCustomerAndYearMonth(customerCode int, yearMonth string) ([]model.Invoice, error)
	FindActiveCustomers() ([]model.Customer, error)
	FindActiveChildren() ([]model.Child, error)
}
