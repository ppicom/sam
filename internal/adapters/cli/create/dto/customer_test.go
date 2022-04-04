package dto

import (
	"github.com/pjover/sam/internal/domain"
	"github.com/pjover/sam/internal/domain/model"
	"github.com/pjover/sam/test/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCustomerToModel(t *testing.T) {
	tests := []struct {
		name     string
		customer Customer
		want     model.Customer
	}{
		{
			name: "customer DTO to model",
			customer: Customer{
				Id:     148,
				Active: true,
				Children: []Child{
					{
						Id:            1850,
						Name:          "Laura",
						Surname:       "Llull",
						SecondSurname: "Bibiloni",
						BirthDate:     "2019-05-25",
						Group:         "EI_1",
						Active:        true,
					},
					{
						Id:            1851,
						Name:          "Aina",
						Surname:       "Llull",
						SecondSurname: "Bibiloni",
						TaxID:         "60235657Z",
						BirthDate:     "2019-05-25",
						Group:         "EI_1",
						Active:        true,
					},
				},
				Adults: []Adult{
					{
						Name:          "Cara",
						Surname:       "Santamaria",
						SecondSurname: "Novella",
						TaxID:         "36361882D",
						Role:          "MOTHER",
						Address: Address{
							Street:  "Carrer Ucraïna 2022, 1st",
							ZipCode: "07007",
							City:    "Palma",
							State:   "Illes Balears",
						},
						Email:            "cara@sgu.org",
						MobilePhone:      "654321098",
						HomePhone:        "987654321",
						GrandMotherPhone: "685698789",
						GrandParentPhone: "658785478",
						WorkPhone:        "987525444",
						BirthDate:        test_data.TestDate.Format(domain.YearMonthDayLayout),
						Nationality:      "US",
					},
					{
						Name:          "Bob",
						Surname:       "Novella",
						SecondSurname: "Sagan",
						TaxID:         "71032204Q",
						Role:          "FATHER",
						Address: Address{
							Street:  "Carrer Ucraïna 2022, 1st",
							ZipCode: "07007",
							City:    "Palma",
							State:   "Illes Balears",
						},
						Email:            "bob@sgu.org",
						MobilePhone:      "654321097",
						HomePhone:        "987654322",
						GrandMotherPhone: "685698788",
						GrandParentPhone: "658785477",
						WorkPhone:        "987525446",
						BirthDate:        test_data.TestDate.Format(domain.YearMonthDayLayout),
						Nationality:      "UK",
					},
				},
				InvoiceHolder: InvoiceHolder{
					Name:  "Cara Santamaria Novella",
					TaxID: "36361882D",
					Address: Address{
						Street:  "Carrer Ucraïna 2022, 1st",
						ZipCode: "07007",
						City:    "Palma",
						State:   "Illes Balears",
					},
					Email:       "cara@sgu.org",
					PaymentType: "BANK_DIRECT_DEBIT",
					Iban:        "ES2830668859978258529057",
				},
				Note:     "Nota del client",
				Language: "CA",
			},
			want: test_data.Customer148,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CustomerToModel(tt.customer)
			assert.Equal(t, tt.want, got)
		})
	}
}
