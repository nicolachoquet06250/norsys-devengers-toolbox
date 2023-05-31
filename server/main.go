package server

import (
	"log"
	"net/http"
	httpPort "norsys-devengers-toolbox/http"
	"norsys-devengers-toolbox/router"
	"strconv"
	"strings"
)

func Process(registerRoutes bool, logger *log.Logger) {
	port := httpPort.ChooseUnusedPort()

	logger.Println("using port :", port)

	if registerRoutes {
		router.Routes()
	}

	portStr := strconv.FormatInt(int64(port), 10)
	portSuffix := ":" + portStr

	if registerRoutes {
		logger.Println("server opened on http://localhost" + portSuffix)
	}

	err := http.ListenAndServe(portSuffix, nil)

	if err != nil {
		if strings.Contains(err.Error(), "invalid port") {
			logger.Println(err.Error())
			Process(false, logger)
		} else {
			logger.Fatal(err)
		}
	}
}
