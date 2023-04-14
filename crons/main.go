package cron
import (
	"os"
	"os/signal"
	"syscall"
	"github.com/robfig/cron/v3"
	. "github.com/Nozimjon2023/GolangCharts/pkg"
)

func Hourly() { 
  cron := cron.New()
  cron.AddFunc("@hourly", RUN)
  go cron.Start()
  defer cron.Stop()
  sig := make(chan os.Signal, 1)
  signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
  <-sig
}
func Second() {
	cron1 := cron.New()
	defer cron1.Stop()
	cron1.AddFunc("* * * * *", RUN)
	go cron1.Start()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
  }