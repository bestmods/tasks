package main

import (
	"bestmods/tasks/internal/config"
	"bestmods/tasks/internal/tasks"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/go-co-op/gocron"
)

const VERSION = "1.0.0"

func main() {
	var list bool
	var version bool

	// Handle command line parameters.
	flag.BoolVar(&list, "list", false, "Lists all tasks.")
	flag.BoolVar(&version, "version", false, "Shows the current version.")

	configFile := flag.String("cfg", "/etc/bestmods-tasks/tasks.conf", "The path to the config file.")

	flag.Parse()

	// Check verison flag.
	if version {
		fmt.Println(VERSION)

		os.Exit(0)
	}

	// Load config.
	var cfg config.Config
	err := cfg.LoadConfig(*configFile)

	if err != nil {
		fmt.Println("Error reading config file!")
		fmt.Println(err)

		os.Exit(1)
	}

	if list {
		fmt.Printf("Debug => %d\n", cfg.Debug)
		for i, task := range cfg.Tasks {
			fmt.Printf("Task #%d\n", i)
			fmt.Printf("\tCron String => %s\n", task.CronStr)
			fmt.Printf("\tURL => %s\n", task.URL)
			fmt.Printf("\tAuth => %s\n", task.Auth)
			fmt.Printf("\tMethod => %s\n\n", task.Method)
		}

		os.Exit(0)
	}

	// Setup cron jobs.
	sch := gocron.NewScheduler(time.UTC)

	for _, task := range cfg.Tasks {
		sch.Cron(task.CronStr).Do(func() {
			tasks.Exec(task, cfg.Debug)
		})

		if cfg.Debug > 0 {
			fmt.Printf("Task '%s' ('%s') initiated.\n", task.URL, task.Method)
		}
	}

	// Start blocking.
	sch.StartBlocking()
}
