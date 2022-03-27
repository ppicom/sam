package model

import (
	"github.com/pjover/sam/internal/domain/model/payment_type"
	"time"
)

type Invoice struct {
	Id          string
	CustomerId  int
	Date        time.Time
	YearMonth   YearMonth
	ChildrenIds []int
	Lines       []Line
	PaymentType payment_type.PaymentType
	Note        string
	Emailed     bool
	Printed     bool
	SentToBank  bool
}

type Line struct {
	ProductId     string
	Units         float64
	ProductPrice  float64
	TaxPercentage float64
	ChildId       int
}