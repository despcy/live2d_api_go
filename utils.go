package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GetList() ModelList {
	// Open our jsonFile
	jsonFile, err := os.Open("model_list.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var list ModelList
	json.Unmarshal(byteValue, &list)
	return list
}

func GetModelData(path string) ModelData {
	// Open our jsonFile
	jsonFile, err := os.Open("./model/" + path + "/index.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var model ModelData
	json.Unmarshal(byteValue, &model)
	changeDataPathInModel(&model, path)
	return model
}

func changeDataPathInModel(data *ModelData, path string) {
	baseDir := "../live2Dmodel/" + path + "/"
	//model path
	data.Model = baseDir + data.Model
	//texture path
	if data.Pose != "" {
		data.Pose = baseDir + data.Pose
	}
	if data.Physics != "" {
		data.Physics = baseDir + data.Physics
	}
	for i, _ := range data.Textures {
		data.Textures[i] = baseDir + data.Textures[i]
	}
	//motions path
	for i, _ := range data.Motions.Idle {
		data.Motions.Idle[i].File = baseDir + data.Motions.Idle[i].File
	}
	for i, _ := range data.Motions.Sleepy {
		data.Motions.Sleepy[i].File = baseDir + data.Motions.Sleepy[i].File
	}
	for i, _ := range data.Motions.FlickHead {
		data.Motions.FlickHead[i].File = baseDir + data.Motions.FlickHead[i].File
	}
	for i, _ := range data.Motions.TapBody {
		data.Motions.TapBody[i].File = baseDir + data.Motions.TapBody[i].File
		data.Motions.TapBody[i].Sound = baseDir + data.Motions.TapBody[i].Sound
	}
	for i, _ := range data.Expressions {
		data.Expressions[i].File = baseDir + data.Expressions[i].File
	}
	//
}
