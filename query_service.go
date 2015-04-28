package goflow

type QueryService struct {
}

func (p *QueryService) GetActiveTasks(orderId string) []*Task {
	task := &Task{}
	tasks, _ := task.GetActiveTasks(orderId)
	return tasks
}
