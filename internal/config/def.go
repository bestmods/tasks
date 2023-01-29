package config

type Task struct {
	CronStr string `json:"cronstr"`
	URL     string `json:"url"`
	Auth    string `json:"auth"`
	Method  string `json:"method"`
}

type Config struct {
	Debug int    `json:"debug"`
	Tasks []Task `json:"tasks"`
}
