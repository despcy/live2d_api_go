package main

/*{
    "models": [
        "GroupName/ModelName",
        [
            "Group1/Model1",
            "Group1/Model2",
            "Group2/Model1"
        ]
    ],
    "messages": [
        "Example 1",
        "Example 2"
    ]
}
*/
type ModelList struct {
	Models   []interface{} `json:"models"`
	Messages []string      `json:"messages"`
}

type ModelChangeResp struct {
	Model struct {
		ID      int    `json:"id"`
		Message string `json:"message"`
	} `json:"model"`
}
type TextureChangeResp struct {
	Textures struct {
		ID int `json:"id"`
	} `json:"textures"`
}
type ModelData struct {
	Version  string   `json:"version"`
	Model    string   `json:"model"`
	Textures []string `json:"textures"`
	Pose     string   `json:"pose"`
	Physics  string   `json:"physics"`
	Layout   struct {
		CenterX float64 `json:"center_x"`
		CenterY float64 `json:"center_y"`
		Width   float64 `json:"width"`
	} `json:"layout"`
	HitAreasCustom struct {
		HeadX []float64 `json:"head_x"`
		HeadY []float64 `json:"head_y"`
		BodyX []float64 `json:"body_x"`
		BodyY []float64 `json:"body_y"`
	} `json:"hit_areas_custom"`
	Motions struct {
		Idle []struct {
			File string `json:"file"`
		} `json:"idle"`
		Sleepy []struct {
			File string `json:"file"`
		} `json:"sleepy"`
		FlickHead []struct {
			File string `json:"file"`
		} `json:"flick_head"`
		TapBody []struct {
			Dialogue int    `json:"dialogue"`
			Sound    string `json:"sound"`
			File     string `json:"file"`
		} `json:"tap_body"`
	} `json:"motions"`

	Expressions []struct {
		Name string `json:"name"`
		File string `json:"file"`
	} `json:"expressions"`
}
