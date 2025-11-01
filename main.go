package main

import (
	"StudyTgServer/api"
	"StudyTgServer/config"
)

func main() {
	cfg := config.Load()

	studyApi := api.NewStudyApiServer(cfg.StudyApiHost, cfg.StudyApiPort)

}
