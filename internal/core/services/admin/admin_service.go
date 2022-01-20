package admin

import (
	"fmt"
	"github.com/pjover/sam/internal/core/ports"
	"github.com/pjover/sam/internal/core/services/lang"
	"log"
	"path"
	"time"
)

type adminService struct {
	configService ports.ConfigService
	osService     ports.OsService
	langService   lang.LangService
}

func NewAdminService(configService ports.ConfigService, osService ports.OsService, langService lang.LangService) ports.AdminService {
	service := adminService{
		configService: configService,
		osService:     osService,
		langService:   langService,
	}
	msg, err := service.CreateWorkingDirectory()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
	return service
}

func (a adminService) Backup() (string, error) {
	fmt.Println("Fent la còpia de seguretat de la base de dades ...")

	tmpDirPath, err := a.osService.GetTempDirectory()
	if err != nil {
		return "", err
	}

	var strSlice = []string{"consumption", "customer", "invoice", "product", "sequence"}
	var files = []string{}
	for _, value := range strSlice {
		fileName := fmt.Sprintf("%s.json", value)
		filePath := path.Join(tmpDirPath, fileName)
		err := a.osService.RunCommand(
			"mongoexport",
			"--db=hobbit",
			fmt.Sprintf("--collection=%s", value),
			fmt.Sprintf("--out=%s", filePath),
		)
		files = append(files, filePath)
		if err != nil {
			return "", err
		}
	}

	zipFilePath, err := a.getZipFilePath()
	if err != nil {
		return "", err
	}

	err = a.osService.CreateZipFile(zipFilePath, files)
	if err != nil {
		return "", err
	}

	return fmt.Sprint("Completada la còpia de seguretat de la base de dades a ", zipFilePath, " ..."), nil
}

func (a adminService) getZipFilePath() (string, error) {
	backupDir := a.configService.Get("dirs.backup")
	exists, err := a.osService.ItemExists(backupDir)
	if err != nil {
		return "", err
	}
	if !exists {
		return "", fmt.Errorf("el directori %s no existeix", backupDir)
	}
	dateStr := a.osService.Now().Format("060102")
	backupFileName := fmt.Sprintf("%s-Backup.zip", dateStr)
	backupFilePath := path.Join(backupDir, backupFileName)
	return backupFilePath, nil
}

func (a adminService) CreateWorkingDirectory() (string, error) {
	workingTime := a.getWorkingTime()
	yearMonth := workingTime.Format("2006-01")
	dirName := a.langService.WorkingDir(workingTime)

	dirPath := path.Join(a.configService.Get("dirs.home"), dirName)
	exists, err := a.osService.ItemExists(dirPath)
	if err != nil {
		return "", err
	}
	if exists {
		_ = a.updateConfig(yearMonth, dirName)
		return fmt.Sprint("Treballant al directori ", dirPath), nil
	}

	if err := a.osService.CreateDirectory(dirPath); err != nil {
		return "", err
	}

	if err := a.updateConfig(yearMonth, dirName); err != nil {
		return "", err
	}

	if err := a.osService.OpenUrlInBrowser(dirPath); err != nil {
		return "", err
	}

	return fmt.Sprint("Creat el directori ", dirPath), nil
}

func (a adminService) getWorkingTime() time.Time {
	var t = a.osService.Now()
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
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
