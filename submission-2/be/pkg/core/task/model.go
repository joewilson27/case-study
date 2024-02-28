package task

const _TASK_NOT_FOUND_ = "Task not found"

type TaskAdd struct {
	Name    string `json:"name"`
	Content string `json:"content"`
	Status  string `json:"status"`
}
