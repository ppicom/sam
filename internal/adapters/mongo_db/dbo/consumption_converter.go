package dbo

import "github.com/pjover/sam/internal/domain/model"

func ConvertConsumption(consumption Consumption) model.Consumption {
	return model.Consumption{
		Code:            consumption.Code,
		ChildCode:       consumption.ChildCode,
		ProductID:       consumption.ProductID,
		Units:           Decimal128ToFloat64(consumption.Units),
		YearMonth:       consumption.YearMonth,
		Note:            consumption.Note,
		IsRectification: consumption.IsRectification,
		InvoiceCode:     consumption.InvoiceCode,
	}
}

func ConvertConsumptions(consumptions []Consumption) []model.Consumption {
	var out []model.Consumption
	for _, consumption := range consumptions {
		out = append(out, ConvertConsumption(consumption))
	}
	return out
}

func ConvertConsumptionToDbo(consumption model.Consumption) Consumption {
	return Consumption{
		Code:            consumption.Code,
		ChildCode:       consumption.ChildCode,
		ProductID:       consumption.ProductID,
		Units:           Float64ToDecimal128(consumption.Units),
		YearMonth:       consumption.YearMonth,
		Note:            consumption.Note,
		IsRectification: consumption.IsRectification,
		InvoiceCode:     consumption.InvoiceCode,
	}
}

func ConvertConsumptionsToDbo(consumptions []model.Consumption) []interface{} {
	var out []interface{}
	for _, consumption := range consumptions {
		out = append(out, ConvertConsumptionToDbo(consumption))
	}
	return out
}
