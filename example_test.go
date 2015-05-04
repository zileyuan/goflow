package goflow

import (
	"testing"

	"github.com/Knetic/govaluate"
)

//测试表达式
func TestExpression(t *testing.T) {
	expression1, _ := govaluate.NewEvaluableExpression("content")
	parameters1 := make(map[string]interface{})
	parameters1["content"] = "toTask1"
	next1, _ := expression1.Evaluate(parameters1)
	t.Logf("next1 %v", next1)

	expression2, _ := govaluate.NewEvaluableExpression("content == 200")
	parameters2 := make(map[string]interface{})
	parameters2["content"] = 200.0
	next2, _ := expression2.Evaluate(parameters2)
	t.Logf("next2 %v", next2)

	expression3, _ := govaluate.NewEvaluableExpression("content > 200")
	parameters3 := make(map[string]interface{})
	parameters3["content"] = 200.0
	next3, _ := expression3.Evaluate(parameters3)
	t.Logf("next3 %v", next3)

	expression4, _ := govaluate.NewEvaluableExpression("content < 200")
	parameters4 := make(map[string]interface{})
	parameters4["content"] = 200.0
	next4, _ := expression4.Evaluate(parameters4)
	t.Logf("next4 %v", next4)
}

//测试参与方式ALL
func TestActorAll(t *testing.T) {
	bytes := LoadXML("res/actorall.xml")
	engine := NewEngine()
	processId := engine.Deploy(bytes, "")
	args := map[string]interface{}{
		"task1.operator": []string{"1", "2"},
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %s", order.Id)
}

//测试分叉和合并
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
	tasks := GetActiveTasksByOrderId(order.Id)
	for _, task := range tasks {
		engine.ExecuteByTaskId(task.Id, "1", args)
	}
}

//测试决策1
func TestDecision1(t *testing.T) {
	bytes := LoadXML("res/decision1.xml")
	engine := NewEngine()
	processId := engine.Deploy(bytes, "")

	args := map[string]interface{}{
		"task2.operator": []string{"1"},
		"content":        250.0,
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %s", order.Id)
}

//测试决策2
func TestDecision2(t *testing.T) {
	bytes := LoadXML("res/decision2.xml")
	engine := NewEngine()
	processId := engine.Deploy(bytes, "")

	args := map[string]interface{}{
		"task1.operator": []string{"1"},
		"task2.operator": []string{"1"},
		"task3.operator": []string{"1"},
		"content":        250.0,
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %s", order.Id)
}

//简单测试
func TestSimple(t *testing.T) {
	bytes := LoadXML("res/simple.xml")
	engine := NewEngine()
	processId := engine.Deploy(bytes, "")
	args := map[string]interface{}{
		"task1.operator": []string{"1"},
		"task2.operator": []string{"1"},
		"task3.operator": []string{"1"},
		"content":        250.0,
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %s", order.Id)
}

//测试参与方式-辅助
func TestAssist(t *testing.T) {
	bytes := LoadXML("res/assist.xml")
	engine := NewEngine()
	engine.Deploy(bytes, "")
	order := engine.StartInstanceByName("assist", -1, "2", nil)
	tasks := GetActiveTasksByOrderId(order.Id)
	for _, task := range tasks {
		CreateNewTask(task.Id, TT_ASSIST, "test")
	}
}

//测试子流程1
func TestSubProcess1(t *testing.T) {
	engine := NewEngine()
	bytes := LoadXML("res/subprocess.child.xml")
	processId := engine.Deploy(bytes, "")
	bytes = LoadXML("res/subprocess.sp1.xml")
	processId = engine.Deploy(bytes, "")

	args := map[string]interface{}{
		"task1.operator": []string{"1"},
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %s", order.Id)

	tasks := GetActiveTasksByOrderId(order.Id)
	for _, task := range tasks {
		engine.ExecuteByTaskId(task.Id, "1", args)
	}
}

//测试子流程2
func TestSubProcess2(t *testing.T) {
	engine := NewEngine()
	bytes := LoadXML("res/subprocess.child.xml")
	processId := engine.Deploy(bytes, "")
	bytes = LoadXML("res/subprocess.sp2.xml")
	processId = engine.Deploy(bytes, "")

	args := map[string]interface{}{
		"task1.operator": []string{"1"},
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %s", order.Id)

	tasks := GetActiveTasksByOrderId(order.Id)
	for _, task := range tasks {
		engine.ExecuteByTaskId(task.Id, "1", args)
	}
}

//测试小组
func TestGroup(t *testing.T) {
	bytes := LoadXML("res/group.xml")
	engine := NewEngine()
	engine.Deploy(bytes, "")
	args := map[string]interface{}{
		"task1.operator": []string{"role1"},
	}
	order := engine.StartInstanceByName("group", 0, "2", args)
	t.Logf("OrderId %s", order.Id)
	tasks := GetActiveTasksByOrderId(order.Id)
	for _, task := range tasks {
		engine.ExecuteByTaskId(task.Id, "test1", args)
	}
}

//测试权限
func TestRight(t *testing.T) {
	bytes := LoadXML("res/right.xml")
	engine := NewEngine()
	processId := engine.Deploy(bytes, "")
	args := map[string]interface{}{
		"task1.operator": []string{"2"},
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %s", order.Id)
	tasks := GetActiveTasksByOrderId(order.Id)
	for _, task := range tasks {
		engine.ExecuteByTaskId(task.Id, string(ER_ADMIN), args)
	}
}

//测试任务提取
func TestTake(t *testing.T) {
	bytes := LoadXML("res/take.xml")
	engine := NewEngine()
	processId := engine.Deploy(bytes, "")
	args := map[string]interface{}{
		"task1.operator": []string{"1"},
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %s", order.Id)
	tasks := GetActiveTasksByOrderId(order.Id)
	for _, task := range tasks {
		TakeTask(task.Id, "1")
	}
}
