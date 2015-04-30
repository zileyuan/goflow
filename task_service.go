package goflow

import "strings"

type TaskService struct {
}

func (p *TaskService) RemoveTaskActor(taskId string, actors ...string) {
	task := &Task{}
	task.GetTaskById(taskId)
	if len(actors) > 0 && task.TaskType == TT_MAJOR {
		v := JsonToMap(task.Variable)
		oldActors := strings.Split(v[DEFAULT_KEY_ACTOR].(string), ",")
		for _, actor := range actors {
			for k, s := range oldActors {
				if s == actor {
					oldActors = StringsRemoveAtIndex(oldActors, k)
					break
				}
			}
		}
		v[DEFAULT_KEY_ACTOR] = oldActors
		task.Variable = MapToJson(v)
		Update(task, task.Id)
	}
}
