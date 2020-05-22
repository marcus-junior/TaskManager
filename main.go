package main

import (
	"encoding/json"
	"github.com/taskManager/Core"
	"github.com/taskManager/Entity"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	file, e := ioutil.ReadFile("config.json")
	check(e)

	taskManager := Entity.TaskManager{}
	e = json.Unmarshal([]byte(file), &taskManager)
	check(e)

	manager := Core.Manager{}
	manager.Run(taskManager)
}