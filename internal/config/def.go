package config

type Task struct {
	CronStr string `json:"cronstr"`
	URL     string `json:"url"`
	Auth    string `json:"auth"`
	Method  string `json:"method"`
}

type Config struct {
	Tasks []Task `json:"tasks"`
}
