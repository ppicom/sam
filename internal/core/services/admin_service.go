package services

import (
	"fmt"
	"github.com/pjover/sam/internal/core/os"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/pjover/sam/internal/translate"
	"path"
	"time"
)

type adminService struct {
	configService ports.ConfigService
	timeManager   os.TimeManager
	fileManager   os.FileManager
	execManager   os.ExecManager
}

func NewAdminService(configService ports.ConfigService, timeManager os.TimeManager, fileManager os.FileManager, execManager os.ExecManager) ports.AdminService {
	return adminService{
		configService: configService,
		timeManager:   timeManager,
		fileManager:   fileManager,
		execManager:   execManager,
	}
}

func (a adminService) Backup() (string, error) {
	fmt.Println("Fent la còpia de seguretat de la base de dades ...")

	dateStr := a.timeManager.Now().Format("060102")
	backupDirPath, err := a.getBackupDirPath()
	if err != nil {
		return "", err
	}

	tmpDirName := fmt.Sprintf("%s-Backup", dateStr)
	tmpDirPath := path.Join(backupDirPath, tmpDirName)
	if err := a.fileManager.CreateDirectory(tmpDirPath); err != nil {
		return "", err
	}

	if err := a.fileManager.ChangeToDirectory(tmpDirPath); err != nil {
		return "", err
	}

	var strSlice = []string{"consumption", "customer", "invoice", "product", "sequence"}
	for _, value := range strSlice {
		if err := a.execManager.Run(
			"mongoexport",
			"--db=hobbit_prod",
			fmt.Sprintf("--collection=%s", value),
			fmt.Sprintf("--out=%s.json", value),
		); err != nil {
			return "", err
		}
	}

	if err := a.fileManager.ChangeToDirectory(backupDirPath); err != nil {
		return "", err
	}

	zipFileName := fmt.Sprintf("%s.zip", tmpDirName)
	if err := a.execManager.Run(
		"zip",
		"-r",
		zipFileName,
		fmt.Sprintf("%s/", tmpDirName),
	); err != nil {
		return "", err
	}

	if err := a.fileManager.RemoveDirectory(tmpDirPath); err != nil {
		return "", err
	}

	filePath := path.Join(backupDirPath, zipFileName)
	return fmt.Sprint("Completada la còpia de seguretat de la base de dades a ", filePath, " ..."), nil
}

func (a adminService) getBackupDirPath() (string, error) {
	backupDirPath := a.configService.Get("dirs.backup")
	exists, err := a.fileManager.Exists(backupDirPath)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", fmt.Errorf("El directori ", backupDirPath, " no existeix")
	}
	return backupDirPath, nil
}

func (a adminService) CreateDirectory(previousMonth bool, nextMonth bool) (string, error) {
	workingTime := a.getWorkingTime(previousMonth, nextMonth)
	yearMonth := workingTime.Format("2006-01")
	dirName := translate.WorkingDir(workingTime)

	dirPath := path.Join(a.configService.Get("dirs.home"), dirName)
	exists, err := a.fileManager.Exists(dirPath)
	if err != nil {
		return "", err
	}
	if exists {
		return fmt.Sprintf("El directori %s ja existeix", dirPath), nil
	}

	if err := a.fileManager.CreateDirectory(dirPath); err != nil {
		return "", err
	}

	if err := a.updateConfig(yearMonth, dirName); err != nil {
		return "", err
	}

	if err := a.execManager.BrowseTo(dirPath); err != nil {
		return "", err
	}

	return fmt.Sprint("Creat el directori ", dirPath), nil
}

func (a adminService) getWorkingTime(previousMonth bool, nextMonth bool) time.Time {
	var t = a.timeManager.Now()
	var workingTime = time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)

	if previousMonth {
		workingTime = workingTime.AddDate(0, -1, 0)
	} else if nextMonth {
		workingTime = workingTime.AddDate(0, 1, 0)
	}
	return workingTime
}

func (a adminService) updateConfig(yearMonth string, dirName string) error {
	if err := a.configService.Set("yearMonth", yearMonth); err != nil {
		return err
	}
	if err := a.configService.Set("dirs.current", dirName); err != nil {
		return err
	}
	return nil
}
