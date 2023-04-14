package main

import (
	"github.com/Nozimjon2023/GolangCharts"
	"os"
	"os/signal"
	"syscall"
	"github.com/robfig/cron/v3"
)

func main() { //точка входа модуля
  // pkg.Run()
  cron := cron.New()
  cron.AddFunc("@hourly", pkg.Run)// @hourly pkg.Run() вызывается каждый час
  go cron.Start()
  defer cron.Stop()
  sig := make(chan os.Signal, 1)
  signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
  <-sig
  
}