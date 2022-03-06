package di

import (
	"github.com/pjover/sam/internal/adapters/cli"
	billingCli "github.com/pjover/sam/internal/adapters/cli/billing"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/billing"
)

func billingServiceDI(cfgService ports.ConfigService, dbService ports.DbService, osService ports.OsService, cmdManager cli.CmdManager) {
	billingService := billing.NewBillingService(cfgService, dbService, osService)
	cmdManager.AddCommand(billingCli.NewInsertConsumptionsCmd(billingService))
	cmdManager.AddCommand(billingCli.NewBillConsumptionsCmd(billingService))
	cmdManager.AddCommand(billingCli.NewRectifyConsumptionsCmd(billingService))
}
