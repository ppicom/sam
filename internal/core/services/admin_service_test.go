package services

import (
	os_mocks "github.com/pjover/sam/internal/core/os/mocks"
	ports_mocks "github.com/pjover/sam/internal/core/ports/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func Test_CreateDirectory(t *testing.T) {
	mockedTimeManager := new(os_mocks.TimeManager)
	mockedTimeManager.On("Now").Return(time.Date(2021, time.October, 31, 21, 14, 0, 0, time.UTC))
	mockedFileManager := new(os_mocks.FileManager)
	mockedFileManager.On("Exists", mock.Anything).Return(false, nil)
	mockedFileManager.On("CreateDirectory", mock.Anything).Return(nil)
	mockedConfigService := new(ports_mocks.ConfigService)
	mockedConfigService.On("Get", "dirs.home").Return("/fake/dir")
	mockedConfigService.On("Set", mock.Anything, mock.Anything).Return(nil)
	mockedExecManager := new(os_mocks.ExecManager)
	mockedExecManager.On("BrowseTo", mock.Anything).Return(nil)

	sut := NewAdminService(mockedConfigService, mockedTimeManager, mockedFileManager, mockedExecManager)

	t.Run("Should return current month", func(t *testing.T) {
		msg, err := sut.CreateDirectory(false, false)
		assert.Equal(t, "Creat el directori /fake/dir/211000-Factures del mes d'Octubre", msg)
		assert.Equal(t, nil, err)
	})

	t.Run("Should return previous month", func(t *testing.T) {
		msg, err := sut.CreateDirectory(true, false)
		assert.Equal(t, "Creat el directori /fake/dir/210900-Factures del mes de Setembre", msg)
		assert.Equal(t, nil, err)
	})

	t.Run("Should return next month", func(t *testing.T) {
		msg, err := sut.CreateDirectory(false, true)
		assert.Equal(t, "Creat el directori /fake/dir/211100-Factures del mes de Novembre", msg)
		assert.Equal(t, nil, err)
	})
}

func Test_Backup_ok(t *testing.T) {
	mockedTimeManager := new(os_mocks.TimeManager)
	mockedTimeManager.On("Now").Return(time.Date(2021, time.October, 31, 21, 14, 0, 0, time.UTC))
	mockedFileManager := new(os_mocks.FileManager)
	mockedFileManager.On("GetTempDirectory").Return("/tmp/sam", nil)
	mockedFileManager.On("Exists", "/fake/dir").Return(true, nil)
	mockedFileManager.On("Zip", "/fake/dir/211031-Backup.zip", mock.Anything).Return(nil)
	mockedConfigService := new(ports_mocks.ConfigService)
	mockedConfigService.On("Get", "dirs.backup").Return("/fake/dir")
	mockedExecManager := new(os_mocks.ExecManager)
	mockedExecManager.On("Run", "mongoexport", "--db=hobbit", "--collection=consumption", "--out=/tmp/sam/consumption.json").Return(nil)
	mockedExecManager.On("Run", "mongoexport", "--db=hobbit", "--collection=customer", "--out=/tmp/sam/customer.json").Return(nil)
	mockedExecManager.On("Run", "mongoexport", "--db=hobbit", "--collection=invoice", "--out=/tmp/sam/invoice.json").Return(nil)
	mockedExecManager.On("Run", "mongoexport", "--db=hobbit", "--collection=product", "--out=/tmp/sam/product.json").Return(nil)
	mockedExecManager.On("Run", "mongoexport", "--db=hobbit", "--collection=sequence", "--out=/tmp/sam/sequence.json").Return(nil)

	sut := NewAdminService(mockedConfigService, mockedTimeManager, mockedFileManager, mockedExecManager)

	t.Run("Should return message with right filename", func(t *testing.T) {
		msg, err := sut.Backup()
		assert.Equal(t, "Completada la còpia de seguretat de la base de dades a /fake/dir/211031-Backup.zip ...", msg)
		assert.Equal(t, nil, err)
	})
}