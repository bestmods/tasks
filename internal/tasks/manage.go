package tasks

import (
	"bestmods/tasks/internal/config"
	"fmt"

	"github.com/gamemann/Rust-Auto-Wipe/pkg/chttp"
)

func Exec(task config.Task, debug int) {
	data, _, err := chttp.SendHTTPReq(task.URL, task.Auth, task.Method, nil)

	if err != nil {
		fmt.Printf("Error executing task '%s' with method '%s' and auth '%s'.\n", task.URL, task.Method, task.Auth)
	}

	if debug > 1 {
		fmt.Printf("Task '%s' with method '%s' executed.\n", task.URL, task.Method)
	}

	if debug > 2 {
		fmt.Println(data)
	}
}
