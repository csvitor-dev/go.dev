package resources

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

var avaliableResources = map[string]string{
	"tailwindcss": "npm run build:css",
	"react":       "npm run build:vite",
}

func Prepare() error {
	err := os.Chdir("./resources")

	if err != nil {
		return fmt.Errorf("resources: fail to prepare resources (%v)", err)
	}

	for resource, command := range avaliableResources {
		if err = os.Chdir(fmt.Sprintf("./%s", resource)); err != nil {
			return fmt.Errorf("resources: fail to prepare resource from %s", resource)
		}
		args := strings.Split(command, " ")
		_, err = exec.Command(args[0], args[1:]...).Output()

		if err != nil {
			return fmt.Errorf("resources: fail on execute '%s' (%v)", command, err)
		}
		os.Chdir("..")
	}
	log.Println("Resources loaded successfully!")

	return os.Chdir("..")
}
