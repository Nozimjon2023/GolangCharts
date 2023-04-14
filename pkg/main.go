package pkg

import (

	"os"
	"os/signal"
	"syscall"
	 "github.com/robfig/cron"
	
)

func RUN() {
 
  cron1 := cron.New()
  defer cron1.Stop()
  cron1.AddFunc("* * * * *", A)
  go cron1.Start()
  
  
  sig := make(chan os.Signal, 1)
  signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
  <-sig
}