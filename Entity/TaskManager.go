package Entity

type TaskManager struct {
	Name       string   `json:"name"`
	FileOutput string   `json:"fileOutput"`
	Executor   Executor `json:"Executor"`
	Tasks    []Task     `json:"Tasks"`
}