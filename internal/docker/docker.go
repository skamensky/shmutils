package docker

import (
	"fmt"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/skamensky/shmutils/pkg/command"
)

func GetDockerContainers() (string, error) {
	res := command.New("docker", "container", "ls").WithTreatStderrAsErr(true).Run()
	return res.Stdout.String(), res.Err
}

func AttachToDockerContainer(containerId string) error {
	getCommand := func(shell string) string {
		return fmt.Sprintf("docker exec -it %s %s", containerId, shell)
	}
	cmd := command.New("bash", "-c", getCommand("bash")).WithTreatStderrAsErr(true).WithAttachToTerminal(true)
	res := cmd.Run()
	if res.Err != nil {
		fmt.Println("Couldn't run bash, trying sh")
		cmd = command.New("bash", "-c", getCommand("sh")).WithTreatStderrAsErr(true).WithAttachToTerminal(true)
		res = cmd.Run()
		if res.Err != nil {
			return res.Err
		}
	}
	fmt.Println(res.Stderr.String())
	return nil
}

func ShellIntoContainer() {
	containers, err := GetDockerContainers()

	if err != nil {
		panic(err)
	}
	if len(strings.Split(containers, "\n")) < 1 {
		fmt.Println("No containers found")
	}
	chosenItem := ""
	userChoseFirstRow := false
	for {
		label := "Select container"
		if userChoseFirstRow {
			label = "Select container (you can't select the first row)"
		}
		prompt := promptui.Select{
			Label: label,
			Items: strings.Split(containers, "\n"),
		}
		_, item, err := prompt.Run()
		if err != nil {
			panic(err)
		}
		if !strings.Contains(item, "CONTAINER ID") {
			chosenItem = item
			break
		} else {
			userChoseFirstRow = true
		}
	}

	items := strings.Split(chosenItem, " ")
	err = AttachToDockerContainer(items[0])
	if err != nil {
		panic(err)
	}
}
