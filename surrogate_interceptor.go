package goflow

//委托代理拦截器
type SurrogateInterceptor struct {
}

func (p *SurrogateInterceptor) GetName() string {
	return "SurrogateInterceptor"
}

func (p *SurrogateInterceptor) Intercept(execution *Execution) {
	for _, task := range execution.Tasks {
		if actors, _ := task.GetTaskActors(); actors != nil {
			for _, actor := range actors {
				agent := GetSurrogate(actor.ActorId, execution.Process.Name)
				if agent != actor.ActorId {
					AddTaskActor(task.Id, task.PerformType, agent)
				}
			}
		}

	}
}

func (p *SurrogateInterceptor) Clone() IInterceptor {
	return &SurrogateInterceptor{}
}
