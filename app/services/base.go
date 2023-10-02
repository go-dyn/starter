package services

import "encoding/json"

type IBase interface {
	MergeTo(from any, to any) error
}

type BaseImpl struct {
}

func NewBaseImpl() *BaseImpl {
	return &BaseImpl{}
}

func (r *BaseImpl) MergeTo(from any, to any) (err error) {
	strByte, _ := json.Marshal(from)
	err = json.Unmarshal(strByte, to)
	return
}
