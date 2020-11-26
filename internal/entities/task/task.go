package task

type Task struct {
	ID         int    `json:"id"`
	Task       string `json:"task"`
	RemindTime string `json:"remindtime"`
	UUID       string `json:"uuid"`
	Reminder   string `json:"reminder"`
}
