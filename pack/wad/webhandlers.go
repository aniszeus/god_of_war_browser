package wad

import (
	"bytes"
	"fmt"
	"net/http"
	"reflect"

	"github.com/mogaika/god_of_war_browser/webutils"
)

func (wad *Wad) WebHandlerForNodeByTagId(w http.ResponseWriter, tagId TagId) error {
	tag := wad.GetTagById(tagId)
	node := wad.GetNodeById(tag.NodeId)
	data, serverId, err := wad.GetInstanceFromNode(node.Id)
	if err == nil {
		type Result struct {
			Tag      *Tag
			Data     interface{}
			ServerId uint32
		}
		val, err := data.Marshal(wad.GetNodeResourceByTagId(node.Tag.Id))
		if err != nil {
			return fmt.Errorf("Error marshaling node %d from %s: %v", tagId, wad.Name(), err.(error))
		} else {
			webutils.WriteJson(w, &Result{Tag: node.Tag, Data: val, ServerId: serverId})
		}
	} else {
		return fmt.Errorf("File %s-%d[%s] parsing error: %v", wad.Name(), node.Tag.Id, node.Tag.Name, err)
	}
	return nil
}

func (wad *Wad) WebHandlerDumpTagData(w http.ResponseWriter, id TagId) {
	tag := wad.GetTagById(id)
	webutils.WriteFile(w, bytes.NewBuffer(tag.Data), tag.Name)
}

func (wad *Wad) WebHandlerCallResourceHttpAction(w http.ResponseWriter, r *http.Request, id TagId, action string) error {
	if inst, _, err := wad.GetInstanceFromTag(id); err == nil {
		rt := reflect.TypeOf(inst)
		method, has := rt.MethodByName("HttpAction")
		if !has {
			return fmt.Errorf("Error: %s has not func SubfileGetter", rt.Name())
		} else {
			method.Func.Call([]reflect.Value{
				reflect.ValueOf(inst),
				reflect.ValueOf(wad.GetNodeResourceByTagId(id)),
				reflect.ValueOf(w),
				reflect.ValueOf(r),
				reflect.ValueOf(action),
			}[:])
			return nil
		}
	} else {
		return fmt.Errorf("File %d instance getting error: %v", id, err)
	}
}
