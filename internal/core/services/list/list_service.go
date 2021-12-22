package list

import (
	"bytes"
	"github.com/pjover/sam/internal/core/model"
	"github.com/pjover/sam/internal/core/ports"
)

type listService struct {
	dbService ports.DbService
}

func NewListService(dbService ports.DbService) ports.ListService {
	return listService{
		dbService: dbService,
	}
}

func (l listService) ListCustomerInvoices(customerCode int) (string, error) {
	invoices, err := l.dbService.FindInvoicesByCustomer(customerCode)
	if err != nil {
		return "", err
	}
	return listInvoices(invoices)
}

func (l listService) ListCustomerYearMonthInvoices(customerCode int, yearMonth string) (string, error) {
	invoices, err := l.dbService.FindInvoicesByCustomerAndYearMonth(customerCode, yearMonth)
	if err != nil {
		return "", err
	}
	return listInvoices(invoices)
}

func (l listService) ListYearMonthInvoices(yearMonth string) (string, error) {
	invoices, err := l.dbService.FindInvoicesByYearMonth(yearMonth)
	if err != nil {
		return "", err
	}
	return listInvoices(invoices)
}

func listInvoices(invoices []model.Invoice) (string, error) {
	var buffer bytes.Buffer
	for _, invoice := range invoices {
		buffer.WriteString(invoice.String() + "\n")
	}
	return buffer.String(), nil
}

func (l listService) ListProducts() (string, error) {
	products, err := l.dbService.FindAllProducts()
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	for _, product := range products {
		buffer.WriteString(product.String() + "\n")
	}
	return buffer.String(), nil
}

func (l listService) ListCustomers() (string, error) {
	customers, err := l.dbService.FindActiveCustomers()
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	for _, customer := range customers {
		buffer.WriteString(customer.String() + "\n")
	}
	return buffer.String(), nil
}

func (l listService) ListChildren() (string, error) {
	children, err := l.dbService.FindActiveChildren()
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	for _, child := range children {
		buffer.WriteString(child.String() + "\n")
	}
	return buffer.String(), nil
}

func (l listService) ListMails(group string, groupByLanguage bool) (string, error) {
	customers, err := l.dbService.FindActiveCustomers()
	if err != nil {
		return "", err
	}

	if groupByLanguage {
		return l.listMailsByLanguage(customers)
	} else if group == "ALL" {
		return l.listAllMails(customers)
	} else {
		return l.listMailsByGroup(group, customers)
	}
}

func (l listService) listMailsByLanguage(customers []model.Customer) (string, error) {
	var caBuffer, esBuffer bytes.Buffer
	caBuffer.WriteString("CA:\n")
	esBuffer.WriteString("ES:\n")
	for _, customer := range customers {
		if customer.Language == "CA" {
			caBuffer.WriteString(customer.InvoiceHolder.Mail() + ", ")
		} else {
			esBuffer.WriteString(customer.InvoiceHolder.Mail() + ", ")
		}
	}
	return caBuffer.String() + "\n" + esBuffer.String(), nil
}

func (l listService) listAllMails(customers []model.Customer) (string, error) {
	var buffer bytes.Buffer
	for _, customer := range customers {
		buffer.WriteString(customer.InvoiceHolder.Mail() + ", ")
	}
	return buffer.String(), nil
}

func (l listService) listMailsByGroup(group string, customers []model.Customer) (string, error) {
	var buffer bytes.Buffer
	buffer.WriteString(group + ":\n")
	for _, customer := range customers {
		var in bool
		for _, child := range customer.Children {
			if child.Group == group {
				in = true
				break
			}
		}
		if in {
			buffer.WriteString(customer.InvoiceHolder.Mail() + ", ")
		}
	}
	return buffer.String(), nil
}
