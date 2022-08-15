package e2e

import (
	"github.com/pjover/sam/internal/adapters/cfg"
	"github.com/pjover/sam/internal/adapters/os"
	"github.com/pjover/sam/internal/domain/ports"
	"github.com/pjover/sam/internal/domain/services/billing"
	"github.com/pjover/sam/test/fakes"
)

func InjectDependencies() ports.CommandManager {

	configService := cfg.NewConfigService() // TODO Fake it!
	dbService := fakes.NewE2eDbService()
	osService := os.NewOsService() // TODO Fake it!

	billingService := billing.NewBillingService(configService, dbService, osService)

	return NewE2eCommandManager(billingService)
}
