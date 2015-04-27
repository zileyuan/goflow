package goflow

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/lunny/log"
)

const (
	DEFAULT_SEPARATOR = "."
)

type ProcessService struct {
	ProcessCache map[string]*Process
	NameCache    map[string]string
}

func (p *ProcessService) Cache(process *Process) {

	processName := process.Name + DEFAULT_SEPARATOR + IntString(process.Version)
	delete(p.ProcessCache, processName)

	var pm ProcessModel
	err := xml.Unmarshal(process.Content, &pm)

	if err != nil {
		log.Errorf("error to unmarshal xml %v", err)
		panic(fmt.Errorf("error to unmarshal xml!"))
	}
	process.SetModel(&pm)

	processName = process.Name + DEFAULT_SEPARATOR + IntString(process.Version)
	p.ProcessCache[processName] = process
	p.NameCache[process.Id] = processName
}

func (p *ProcessService) Deploy(input []byte, creator string) string {

	process := &Process{
		State:      FS_ACTIVITY,
		Content:    input,
		Creator:    creator,
		CreateTime: time.Now(),
	}
	p.Cache(process)
	process.Save()

	return process.Id
}

func (p *ProcessService) ReDeploy(id string, input []byte) {
	process := new(Process)
	success, err := process.GetProcessById(id)
	if err != nil {
		log.Errorf("error to get process by id %v", err)
		panic(fmt.Errorf("error to get process by id!"))
	}

	if success {
		process.Content = input
		p.Cache(process)
		process.Update()
	} else {
		log.Infof("fail to get process by id %v", err)
	}
}

func (p *ProcessService) UnDeploy(id string) {
	process := new(Process)
	success, err := process.GetProcessById(id)
	if err != nil {
		log.Errorf("error to get process by id %v", err)
		panic(fmt.Errorf("error to get process by id!"))
	}

	if success {
		process.State = FS_FINISH
		p.Cache(process)
		process.Update()
	} else {
		log.Infof("fail to get process by id %v", err)
	}
}
