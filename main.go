package main

import (
	"bestmods/tasks/internal/config"
	"bestmods/tasks/internal/tasks"
	"flag"
	"fmt"
	"os"

	cron "github.com/robfig/cron/v3"
)

const VERSION = "1.0.0"

func main() {
	var list bool
	var version bool
	var help bool

	// Handle command line parameters.
	flag.BoolVar(&list, "list", false, "Lists all tasks.")
	flag.BoolVar(&version, "version", false, "Shows the current version.")
	flag.BoolVar(&help, "help", false, "Shows the help menu.")

	configFile := flag.String("cfg", "/etc/bestmods-tasks/tasks.conf", "The path to the config file.")

	flag.Parse()

	// Check verison flag.
	if version {
		fmt.Println(VERSION)

		os.Exit(0)
	}

	// Check help flag.
	if help {
		fmt.Printf("./tasks --cfg <cfgFile> --list --version --help\n")
		fmt.Printf("\t--cfg => Path to config file. Default path is /etc/bestmods-tasks/tasks.conf.\n")
		fmt.Printf("\t--list => Lists configuration file.\n")
		fmt.Printf("\t--version => Prints the current version.\n")
		fmt.Printf("\t--help => Prints the help menu.\n\n")

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
	sch := cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)))

	for _, task := range cfg.Tasks {
		_, err = sch.AddFunc(task.CronStr, func() {
			tasks.Exec(task, cfg.Debug)
		})

		if err != nil {
			fmt.Printf("Error initiating task\n")
			fmt.Println(err)

			continue
		}

		if cfg.Debug > 0 {
			fmt.Printf("Task '%s' ('%s') initiated.\n", task.URL, task.Method)
		}
	}

	sch.Start()

	// Start blocking.
	select {}
}
