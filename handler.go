package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

//SwitchModel 切换下一个模型
func SwitchModel(w http.ResponseWriter, req *http.Request) {
	ids, ok := req.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		log.Printf("err")
		//error handling
	}
	id, _ := strconv.Atoi(ids[0])
	modelList := GetList()
	newId := (id + 1) % len(modelList.Models)
	newMsgId := (id + 1) % len(modelList.Messages)
	w.Header().Set("Content-Type", "application/json")
	var resp ModelChangeResp
	resp.Model.ID = newId
	resp.Model.Message = modelList.Messages[newMsgId]
	respJson, _ := json.Marshal(resp)
	w.Write(respJson)
}

func RandomTextures(w http.ResponseWriter, req *http.Request) {
	ids, ok := req.URL.Query()["id"]
	if !ok || len(ids) < 1 {
		log.Printf("err")
		//error handling
	}
	id := ids[0]
	strs := strings.Split(id, "-")
	modelId, _ := strconv.Atoi(strs[0])
	textureId, _ := strconv.Atoi(strs[1])
	modelList := GetList()
	w.Header().Set("Content-Type", "application/json")
	var resp TextureChangeResp

	switch modelList.Models[modelId].(type) {
	case string:
		resp.Textures.ID = textureId
	case []interface{}:
		length := len(modelList.Models[modelId].([]interface{}))
		resp.Textures.ID = rand.Intn(length)
		for resp.Textures.ID == textureId {
			resp.Textures.ID = rand.Intn(length)
		}

	}
	respJson, _ := json.Marshal(resp)
	w.Write(respJson)
}

//GetModel - `/get/?id=1-23` 获取 分组 1 的 第 23 号 皮肤
func GetModel(w http.ResponseWriter, req *http.Request) {
	ids, ok := req.URL.Query()["id"]
	var id string
	if !ok || len(ids) < 1 {
		//log.Printf("err")
		//error handling
		w.Write([]byte("error"))
		return
	} else {
		id = ids[0]
	}
	strs := strings.Split(id, "-")
	modelId, _ := strconv.Atoi(strs[0])
	textureId, _ := strconv.Atoi(strs[1])
	modelList := GetList()
	var path string
	switch modelList.Models[modelId].(type) {
	case string:
		path = modelList.Models[modelId].(string)
	case []interface{}:
		path = modelList.Models[modelId].([]interface{})[textureId].(string)

	}

	model := GetModelData(path)

	resp, _ := json.Marshal(model)
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
}
