package goflow

import (
	"fmt"
	"testing"

	"github.com/Knetic/govaluate"
)

//测试表达式
func TestExpression(t *testing.T) {
	fmt.Printf("--- Start TestExpression ---")
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

	fmt.Printf("--- End TestExpression ---")
}

//测试参与方式ALL
func TestActorAll(t *testing.T) {
	fmt.Printf("--- Start TestActorAll ---")
	bytes := LoadXML("res/actorall.xml")
	engine := NewEngine()
	processId := engine.Deploy(bytes, "")
	args := map[string]interface{}{
		"task1.operator": []string{"1", "2"},
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %s", order.Id)

	fmt.Printf("--- End TestActorAll ---")
}

//测试分叉和合并
func TestForkJoin(t *testing.T) {
	fmt.Printf("--- Start TestForkJoin ---")
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
	fmt.Printf("--- End TestForkJoin ---")
}

//测试决策1
func TestDecision1(t *testing.T) {
	fmt.Printf("--- Start TestDecision1 ---")
	bytes := LoadXML("res/decision1.xml")
	engine := NewEngine()
	processId := engine.Deploy(bytes, "")

	args := map[string]interface{}{
		"task2.operator": []string{"1"},
		"content":        250.0,
	}
	order := engine.StartInstanceById(processId, "2", args)
	t.Logf("OrderId %s", order.Id)
	fmt.Printf("--- End TestDecision1 ---")
}

//测试决策2
func TestDecision2(t *testing.T) {
	fmt.Printf("--- Start TestDecision2 ---")
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
	fmt.Printf("--- End TestDecision2 ---")
}

//简单测试
func TestSimple(t *testing.T) {
	fmt.Printf("--- Start TestSimple ---")
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
	fmt.Printf("--- End TestSimple ---")
}

//测试协办流程
func TestAssist(t *testing.T) {
	fmt.Printf("--- Start TestAssist ---")
	bytes := LoadXML("res/assist.xml")
	engine := NewEngine()
	engine.Deploy(bytes, "")
	args := map[string]interface{}{
		"content": 250.0,
	}
	order := engine.StartInstanceByName("assist", -1, "2", args)
	tasks := GetActiveTasksByOrderId(order.Id)
	for _, task := range tasks {
		CreateNewTask(task.Id, TT_ASSIST, "test")
	}
	fmt.Printf("--- End TestAssist ---")
}

//测试子流程1
func TestSubProcess1(t *testing.T) {
	fmt.Printf("--- Start TestSubProcess1 ---")
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
	fmt.Printf("--- End TestSubProcess1 ---")
}

//测试子流程2
func TestSubProcess2(t *testing.T) {
	fmt.Printf("--- Start TestSubProcess2 ---")
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
	fmt.Printf("--- End TestSubProcess2 ---")
}

//测试小组
func TestGroup(t *testing.T) {
	fmt.Printf("--- Start TestGroup ---")
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
	fmt.Printf("--- End TestGroup ---")
}

//测试权限
func TestRight(t *testing.T) {
	fmt.Printf("--- Start TestRight ---")
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
	fmt.Printf("--- End TestRight ---")
}

//测试任务提取
func TestTake(t *testing.T) {
	fmt.Printf("--- Start TestTake ---")
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
	fmt.Printf("--- End TestTake ---")
}
