package goflow

import "testing"

func TestSimpleProcess(t *testing.T) {
	bytes := LoadXML("res/simple.xml")
	engine := NewEngine()
	processId := engine.processService.Deploy(bytes)
	args := map[string]interface{}{
		"task1.operator": "1",
		"task2.operator": "1",
		"task3.operator": "1",
		"content", 250,
	}
	order := engine.StartInstanceById(processId, "2", args)
}

func TestAssistProcess(t *testing.T) {
	bytes := LoadXML("res/assist.xml")
	engine := NewEngine()
	processId := engine.processService.Deploy(bytes)
	args := map[string]interface{}{
		"task1.operator": "1",
		"task2.operator": "1",
		"task3.operator": "1",
		"content", 250,
	}
	order := engine.StartInstanceById(processId, "2", args)
}
