package integrations

import (
	"os"
	"path/filepath"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

func CronRun() {
	job := cron.New()

	job.AddFunc("0 0 1 * *", deleteOldFilesFromSystem)

	job.Start()
}

func deleteOldFilesFromSystem() {
	threeMonthsAgo := time.Now().AddDate(0, -3, 0)

	filepath.Walk(viper.GetString("test"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.ModTime().Before(threeMonthsAgo) {
			err := os.Remove(path)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
