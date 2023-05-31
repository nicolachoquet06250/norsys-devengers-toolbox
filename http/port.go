package http

import (
	"norsys-devengers-toolbox/arrays"
	"norsys-devengers-toolbox/cli"
	"norsys-devengers-toolbox/environement"
	"norsys-devengers-toolbox/numbers"
	"regexp"
	"strconv"
	"strings"
)

func ChooseUnusedPort() int {
	out := cli.ExeCmd("netstat -nao")

	lines := strings.Split(string(out), environement.BackLine())

	var listeningPorts []int

	for _, l := range lines {
		var re = regexp.MustCompile(`^ {2}TCP {4}\d{0,3}\.\d{0,3}\.\d{0,3}\.\d{0,3}:(?P<local_port>80\d+) +\d{0,3}\.\d{0,3}\.\d{0,3}\.\d{0,3}:(?P<remote_port>\d+) +(LISTENING) +\d+$`)

		for _, match := range re.FindAllString(l, -1) {
			matches := re.FindStringSubmatch(match)

			localPort, _ := strconv.Atoi(matches[re.SubexpIndex("local_port")])

			listeningPorts = append(listeningPorts, localPort)
		}
	}

	localPort := numbers.RandomNumber(8000, 8099)

	chosenPort := localPort

	exists, _ := arrays.InArray(chosenPort, listeningPorts)
	if exists == true {
		return ChooseUnusedPort()
	} else {
		environement.ChosenPort = localPort
		return environement.ChosenPort
	}
}
