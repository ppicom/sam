// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	model "github.com/pjover/sam/internal/domain/model"
	mock "github.com/stretchr/testify/mock"

	payment_type "github.com/pjover/sam/internal/domain/model/payment_type"

	time "time"
)

// DbService is an autogenerated mock type for the DbService type
type DbService struct {
	mock.Mock
}

// FindActiveChildConsumptions provides a mock function with given fields: id
func (_m *DbService) FindActiveChildConsumptions(id int) ([]model.Consumption, error) {
	ret := _m.Called(id)

	var r0 []model.Consumption
	if rf, ok := ret.Get(0).(func(int) []model.Consumption); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Consumption)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindActiveChildren provides a mock function with given fields:
func (_m *DbService) FindActiveChildren() ([]model.Child, error) {
	ret := _m.Called()

	var r0 []model.Child
	if rf, ok := ret.Get(0).(func() []model.Child); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Child)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindActiveCustomers provides a mock function with given fields:
func (_m *DbService) FindActiveCustomers() ([]model.Customer, error) {
	ret := _m.Called()

	var r0 []model.Customer
	if rf, ok := ret.Get(0).(func() []model.Customer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllActiveConsumptions provides a mock function with given fields:
func (_m *DbService) FindAllActiveConsumptions() ([]model.Consumption, error) {
	ret := _m.Called()

	var r0 []model.Consumption
	if rf, ok := ret.Get(0).(func() []model.Consumption); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Consumption)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllProducts provides a mock function with given fields:
func (_m *DbService) FindAllProducts() ([]model.Product, error) {
	ret := _m.Called()

	var r0 []model.Product
	if rf, ok := ret.Get(0).(func() []model.Product); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAllSequences provides a mock function with given fields:
func (_m *DbService) FindAllSequences() ([]model.Sequence, error) {
	ret := _m.Called()

	var r0 []model.Sequence
	if rf, ok := ret.Get(0).(func() []model.Sequence); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Sequence)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindChangedCustomers provides a mock function with given fields: changedSince
func (_m *DbService) FindChangedCustomers(changedSince time.Time) ([]model.Customer, error) {
	ret := _m.Called(changedSince)

	var r0 []model.Customer
	if rf, ok := ret.Get(0).(func(time.Time) []model.Customer); ok {
		r0 = rf(changedSince)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(time.Time) error); ok {
		r1 = rf(changedSince)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindChild provides a mock function with given fields: id
func (_m *DbService) FindChild(id int) (model.Child, error) {
	ret := _m.Called(id)

	var r0 model.Child
	if rf, ok := ret.Get(0).(func(int) model.Child); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Child)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindCustomer provides a mock function with given fields: id
func (_m *DbService) FindCustomer(id int) (model.Customer, error) {
	ret := _m.Called(id)

	var r0 model.Customer
	if rf, ok := ret.Get(0).(func(int) model.Customer); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Customer)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindInvoice provides a mock function with given fields: id
func (_m *DbService) FindInvoice(id string) (model.Invoice, error) {
	ret := _m.Called(id)

	var r0 model.Invoice
	if rf, ok := ret.Get(0).(func(string) model.Invoice); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Invoice)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindInvoicesByCustomer provides a mock function with given fields: customerId
func (_m *DbService) FindInvoicesByCustomer(customerId int) ([]model.Invoice, error) {
	ret := _m.Called(customerId)

	var r0 []model.Invoice
	if rf, ok := ret.Get(0).(func(int) []model.Invoice); ok {
		r0 = rf(customerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Invoice)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(customerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindInvoicesByCustomerAndYearMonth provides a mock function with given fields: customerId, yearMonth
func (_m *DbService) FindInvoicesByCustomerAndYearMonth(customerId int, yearMonth model.YearMonth) ([]model.Invoice, error) {
	ret := _m.Called(customerId, yearMonth)

	var r0 []model.Invoice
	if rf, ok := ret.Get(0).(func(int, model.YearMonth) []model.Invoice); ok {
		r0 = rf(customerId, yearMonth)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Invoice)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, model.YearMonth) error); ok {
		r1 = rf(customerId, yearMonth)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindInvoicesByYearMonth provides a mock function with given fields: yearMonth
func (_m *DbService) FindInvoicesByYearMonth(yearMonth model.YearMonth) ([]model.Invoice, error) {
	ret := _m.Called(yearMonth)

	var r0 []model.Invoice
	if rf, ok := ret.Get(0).(func(model.YearMonth) []model.Invoice); ok {
		r0 = rf(yearMonth)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Invoice)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.YearMonth) error); ok {
		r1 = rf(yearMonth)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindInvoicesByYearMonthAndPaymentTypeAndSentToBank provides a mock function with given fields: yearMonth, paymentType, sentToBank
func (_m *DbService) FindInvoicesByYearMonthAndPaymentTypeAndSentToBank(yearMonth model.YearMonth, paymentType payment_type.PaymentType, sentToBank bool) ([]model.Invoice, error) {
	ret := _m.Called(yearMonth, paymentType, sentToBank)

	var r0 []model.Invoice
	if rf, ok := ret.Get(0).(func(model.YearMonth, payment_type.PaymentType, bool) []model.Invoice); ok {
		r0 = rf(yearMonth, paymentType, sentToBank)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Invoice)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.YearMonth, payment_type.PaymentType, bool) error); ok {
		r1 = rf(yearMonth, paymentType, sentToBank)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindProduct provides a mock function with given fields: id
func (_m *DbService) FindProduct(id string) (model.Product, error) {
	ret := _m.Called(id)

	var r0 model.Product
	if rf, ok := ret.Get(0).(func(string) model.Product); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(model.Product)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertConsumptions provides a mock function with given fields: consumptions
func (_m *DbService) InsertConsumptions(consumptions []model.Consumption) error {
	ret := _m.Called(consumptions)

	var r0 error
	if rf, ok := ret.Get(0).(func([]model.Consumption) error); ok {
		r0 = rf(consumptions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertInvoices provides a mock function with given fields: invoices
func (_m *DbService) InsertInvoices(invoices []model.Invoice) error {
	ret := _m.Called(invoices)

	var r0 error
	if rf, ok := ret.Get(0).(func([]model.Invoice) error); ok {
		r0 = rf(invoices)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SearchCustomers provides a mock function with given fields: searchText
func (_m *DbService) SearchCustomers(searchText string) ([]model.Customer, error) {
	ret := _m.Called(searchText)

	var r0 []model.Customer
	if rf, ok := ret.Get(0).(func(string) []model.Customer); ok {
		r0 = rf(searchText)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Customer)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(searchText)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateConsumptions provides a mock function with given fields: consumptions
func (_m *DbService) UpdateConsumptions(consumptions []model.Consumption) error {
	ret := _m.Called(consumptions)

	var r0 error
	if rf, ok := ret.Get(0).(func([]model.Consumption) error); ok {
		r0 = rf(consumptions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateSequences provides a mock function with given fields: sequences
func (_m *DbService) UpdateSequences(sequences []model.Sequence) error {
	ret := _m.Called(sequences)

	var r0 error
	if rf, ok := ret.Get(0).(func([]model.Sequence) error); ok {
		r0 = rf(sequences)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
