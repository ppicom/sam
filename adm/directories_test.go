package adm

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockedSamClock struct {
	mock.Mock
}

func (m *MockedSamClock) Now() time.Time {
	args := m.Called()
	return args.Get(0).(time.Time)
}

func Test_GetDirConfig(t *testing.T) {
	mockedSamTimes := new(MockedSamClock)
	sut := Directories{Timer: mockedSamTimes}
	mockedSamTimes.On("Now").Return(time.Date(2021, time.October, 31, 21, 14, 0, 0, time.UTC))

	t.Run("Should return current month", func(t *testing.T) {
		yearMonth, dirName := sut.GetDirConfig(false, false)
		assert.Equal(t, "2021-10", yearMonth)
		assert.Equal(t, "211000-Factures del mes d'Octubre", dirName)
	})

	t.Run("Should return previous month", func(t *testing.T) {
		yearMonth, dirName := sut.GetDirConfig(true, false)
		assert.Equal(t, "2021-09", yearMonth)
		assert.Equal(t, "210900-Factures del mes de Setembre", dirName)
	})

	t.Run("Should return next month", func(t *testing.T) {
		yearMonth, dirName := sut.GetDirConfig(false, true)
		assert.Equal(t, "2021-11", yearMonth)
		assert.Equal(t, "211100-Factures del mes de Novembre", dirName)
	})
}
