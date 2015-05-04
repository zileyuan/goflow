package goflow

import (
	"fmt"
	"time"

	"github.com/lunny/log"
)

//流程服务
type ProcessService struct {
	ProcessCache map[string]*Process
	NameCache    map[string]string
}

//初始化服务对象
func (p *ProcessService) InitProcessService() {
	p.ProcessCache = make(map[string]*Process)
	p.NameCache = make(map[string]string)
}

//缓存Process
func (p *ProcessService) Cache(process *Process) {

	processName := process.Name + DEFAULT_SEPARATOR + IntToStr(process.Version)
	delete(p.ProcessCache, processName)

	if process.Model == nil {
		processModel := &ProcessModel{}
		processModel.BuildRelationship([]byte(process.Content))
		process.SetModel(processModel)
	}

	processName = process.Name + DEFAULT_SEPARATOR + IntToStr(process.Version)
	p.ProcessCache[processName] = process
	p.NameCache[process.Id] = processName
}

//部署Process
func (p *ProcessService) Deploy(input []byte, creator string) string {

	processModel := &ProcessModel{}
	processModel.BuildRelationship(input)

	ver := -1
	oldProcess, _ := GetLatestProcess(processModel.Name)
	if oldProcess == nil {
		ver = oldProcess.Version
	}

	process := &Process{
		Id:         NewUUID(),
		State:      FS_ACTIVITY,
		Content:    string(input),
		Creator:    creator,
		CreateTime: time.Now(),
		Version:    ver + 1,
	}
	process.SetModel(processModel)
	p.Cache(process)
	Save(process, process.Id)

	return process.Id
}

//重新部署Process
func (p *ProcessService) ReDeploy(id string, input []byte) {
	process := new(Process)
	success, err := process.GetProcessById(id)
	if err != nil {
		log.Errorf("error to get process by id %v", err)
		panic(fmt.Errorf("error to get process by id!"))
	}

	if success {
		process.Content = string(input)
		p.Cache(process)
		Update(process, process.Id)
	} else {
		log.Infof("fail to get process by id %v", err)
	}
}

//卸载部署
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

//根据ID得到Process
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

//根据名称、版本得到Process
func (p *ProcessService) GetProcessByVersion(name string, version int) *Process {
	var dbProcess *Process
	if version == -1 {
		dbProcess, _ := GetLatestProcess(name)
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
