package goflow

import "testing"

func TestActorAll(t *testing.T) {
	bytes := LoadXML("res/actorall.xml")
	engine := NewEngine()
	processId := engine.Deploy(bytes, "")
	args := map[string]interface{}{
		"task1.operator": []string{"1", "2"},
	}
	order := engine.StartInstanceById(processId, "2", args)
}

func TestForkJoin(t *testing.T) {
	bytes := LoadXML("res/forkjoin.xml")
	engine := NewEngine()
	processId := engine.Deploy(bytes, "")
	args := map[string]interface{}{
		"task1.operator": []string{"1"},
		"task2.operator": []string{"1"},
		"task3.operator": []string{"1"},
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %s", order.Id)
	tasks := engine.GetActiveTasks(order.Id)
	for _, task := range tasks {
		engine.ExecuteByTaskId(task.Id, "1")
	}
}

func TestDecision1(t *testing.T) {
	bytes := LoadXML("res/decision1.xml")
	engine := NewEngine()
	processId := engine.Deploy(bytes, "")

	args := map[string]interface{}{
		"task2.operator": []string{"1"},
		"content":        250,
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %s", order.Id)
}

func TestDecision2(t *testing.T) {
	bytes := LoadXML("res/decision2.xml")
	engine := NewEngine()
	processId := engine.Deploy(bytes, "")

	args := map[string]interface{}{
		"task1.operator": []string{"1"},
		"task2.operator": []string{"1"},
		"task3.operator": []string{"1"},
		"content":        250,
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %s", order.Id)
}

func TestSimple(t *testing.T) {
	bytes := LoadXML("res/simple.xml")
	engine := NewEngine()
	processId := engine.Deploy(bytes, "")
	args := map[string]interface{}{
		"task1.operator": []string{"1"},
		"task2.operator": []string{"1"},
		"task3.operator": []string{"1"},
		"content":        250,
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %s", order.Id)
}

func TestAssist(t *testing.T) {
	bytes := LoadXML("res/assist.xml")
	engine := NewEngine()
	processId := engine.Deploy(bytes, "")
	order := engine.StartInstanceByName("assist", -1, "2", nil)
	tasks := engine.GetActiveTasks(order.Id)
	for _, task := range tasks {
		engine.CreateNewTask(task.Id, TT_ASSIST, "test")
	}
}

func TestSubProcess1(t *testing.T) {
	engine := NewEngine()
	bytes := LoadXML("res/subprocess.child.xml")
	processId := engine.Deploy(bytes, "")
	bytes := LoadXML("res/subprocess.sp1.xml")
	processId := engine.Deploy(bytes, "")

	args := map[string]interface{}{
		"task1.operator": []string{"1"},
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %s", order.Id)

	tasks := engine.GetActiveTasks(order.Id)
	for _, task := range tasks {
		engine.ExecuteByTaskId(task.Id, "1", args)
	}
}

func TestSubProcess2(t *testing.T) {
	engine := NewEngine()
	bytes := LoadXML("res/subprocess.child.xml")
	processId := engine.Deploy(bytes, "")
	bytes := LoadXML("res/subprocess.sp2.xml")
	processId := engine.Deploy(bytes, "")

	args := map[string]interface{}{
		"task1.operator": []string{"1"},
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %s", order.Id)

	tasks := engine.GetActiveTasks(order.Id)
	for _, task := range tasks {
		engine.ExecuteByTaskId(task.Id, "1", args)
	}
}

func TestGroup(t *testing.T) {
	bytes := LoadXML("res/group.xml")
	engine := NewEngine()
	processId := engine.Deploy(bytes, "")
	args := map[string]interface{}{
		"task1.operator": []string{"role1"},
	}
	order := engine.StartInstanceByName("group", 0, "2", args)
	t.Logf("OrderId %s", order.Id)
	tasks := engine.GetActiveTasks(order.Id)
	for _, task := range tasks {
		engine.ExecuteByTaskId(task.Id, "test1", args)
	}
}

func TestRight(t *testing.T) {
	bytes := LoadXML("res/right.xml")
	engine := NewEngine()
	processId := engine.Deploy(bytes, "")
	args := map[string]interface{}{
		"task1.operator": []string{"2"},
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %s", order.Id)
	tasks := engine.GetActiveTasks(order.Id)
	for _, task := range tasks {
		engine.ExecuteByTaskId(task.Id, ER_ADMIN, args)
	}
}

func TestTake(t *testing.T) {
	bytes := LoadXML("res/take.xml")
	engine := NewEngine()
	processId := engine.Deploy(bytes, "")
	args := map[string]interface{}{
		"task1.operator": []string{"1"},
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %s", order.Id)
	tasks := engine.GetActiveTasks(order.Id)
	for _, task := range tasks {
		engine.Take(task.Id, "1")
	}
}
