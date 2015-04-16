package core

import (
	"encoding/xml"
	"goflow/define"
	"goflow/entity"
	"goflow/model"
	"io/ioutil"
	"time"
)

type ProcessService struct {
	processCache map[string]model.ProcessModel
	nameCache    map[string]string
}

func (p *ProcessService) deploy(xmlProcess string, creator string) string {
	content, err := ioutil.ReadFile(xmlProcess)
	if err != nil {
		panic(err)
	}

	var pm model.ProcessModel
	err = xml.Unmarshal(content, &pm)

	if err != nil {
		panic(err)
	}

	process := &entity.Process{
		Name:        pm.Name,
		DisplayName: pm.DisplayName,
		InstanceUrl: pm.InstanceUrl,
		State:       define.PS_ACTIVITY,
		Content:     string(content),
		Creator:     creator,
		CreateTime:  time.Now(),
		Model:       pm,
	}

	process.Save()

	return process.Id

}
