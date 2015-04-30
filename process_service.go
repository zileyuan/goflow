package goflow

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/lunny/log"
)

type ProcessService struct {
	ProcessCache map[string]*Process
	NameCache    map[string]string
}

func (p *ProcessService) InitProcessService() {
	p.ProcessCache = make(map[string]*Process)
	p.NameCache = make(map[string]string)
}

func (p *ProcessService) Cache(process *Process) {

	processName := process.Name + DEFAULT_SEPARATOR + IntToStr(process.Version)
	delete(p.ProcessCache, processName)

	var pm ProcessModel
	err := xml.Unmarshal(process.Content, &pm)

	if err != nil {
		log.Errorf("error to unmarshal xml %v", err)
		panic(fmt.Errorf("error to unmarshal xml!"))
	}
	process.SetModel(&pm)

	processName = process.Name + DEFAULT_SEPARATOR + IntToStr(process.Version)
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
	Save(process, process.Id)

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
		Update(process, process.Id)
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
		Update(process, process.Id)
	} else {
		log.Infof("fail to get process by id %v", err)
	}
}

func (p *ProcessService) GetProcessById(id string) *Process {
	processName := p.NameCache[id]
	process := p.ProcessCache[processName]

	if process == nil {
		process = &Process{}
		process.GetProcessById(id)
		p.Cache(process)
	}
	return process
}

func (p *ProcessService) GetProcessByVersion(name string, version int) *Process {
	dbProcess := &Process{}
	if version == -1 {
		dbProcess, _ := dbProcess.GetLatestProcess(name)
		if dbProcess == nil {
			return nil
		}
	}
	processName := name + DEFAULT_SEPARATOR + IntToStr(dbProcess.Version)
	process := p.ProcessCache[processName]
	if process == nil {
		process = dbProcess
		p.Cache(process)
	}
	return process
}
