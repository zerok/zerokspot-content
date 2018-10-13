package algoliasearch

type DeleteTaskRes struct {
	DeletedAt string `json:"deletedAt"`
	TaskID    int    `json:"taskID"`
}

type UpdateTaskRes struct {
	TaskID    int    `json:"taskID"`
	UpdatedAt string `json:"updatedAt"`
}

type UpdateTaskWithIDRes struct {
	UpdateTaskRes
	ID string `json:"id"`
}

type TaskStatusRes struct {
	Status      string `json:"status"`
	PendingTask bool   `json:"pendingTask"`
}
