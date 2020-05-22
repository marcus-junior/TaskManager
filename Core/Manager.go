package Core

import (
	"fmt"
	"github.com/taskManager/Entity"
	"os"
	"os/exec"
	"strings"
)

type Manager struct{
	Output strings.Builder
}


func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (m Manager) Run(taskManager Entity.TaskManager){

	m.Output = strings.Builder{}
	m.Output.WriteString("------------------------------------------------------------------------------------------------")
	m.Output.WriteString("\n" + taskManager.Name + "\n")
	m.Output.WriteString("------------------------------------------------------------------------------------------------")

	for i  := 0; i < len(taskManager.Tasks); i++ {

		taskTemp := taskManager.Tasks[i]
		cmd := exec.Command(taskManager.Executor.Path, taskManager.Executor.Parameter, taskTemp.Command)
		cmd.Stdout = &taskTemp.Out
		cmd.Stderr = &taskTemp.Stderr
		cmd.Run()

		m.Output.WriteString("\n")
		m.Output.WriteString("------------------------------------------------------------------------------------------------\n")
		m.Output.WriteString("Command =>   " + taskTemp.Label            + "\n")
		m.Output.WriteString("------------------------------------------------------------------------------------------------\n")
		m.Output.WriteString("Success :: \n" + taskTemp.Out.String()     + "\n")
		m.Output.WriteString("------------------------------------------------------------------------------------------------\n")
		m.Output.WriteString("Error ::   \n"  + taskTemp.Stderr.String() + "\n")
		m.Output.WriteString("------------------------------------------------------------------------------------------------")

	}

	file, err := os.Create(taskManager.FileOutput)
	check(err)

	result, err := file.WriteString(m.Output.String())
	check(err)

	if result > 0 {
		fmt.Println("Processamento Finalizado com sucesso!")
	}
}