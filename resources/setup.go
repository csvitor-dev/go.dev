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
	"react-email": "npm run export:email",
}

func Prepare(args []string) error {
	var err error

	if len(args) > 0 {
		err = prepareOnly(args...)
	} else {
		err = prepareAll()
	}
	return err
}

func prepareAll() error {
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
		log.Printf("'%s' resource loaded successfully!", resource)
		os.Chdir("..")
	}
	log.Println("All resources loaded successfully!")

	return os.Chdir("..")
}

func prepareOnly(resources ...string) error {
	err := os.Chdir("./resources")

	if err != nil {
		return fmt.Errorf("resources: fail to prepare resources (%v)", err)
	}

	for _, resource := range resources {
		command, exists := avaliableResources[resource]

		if !exists {
			return fmt.Errorf("resources: the resource '%s' do not exists", resource)
		}

		if err = os.Chdir(fmt.Sprintf("./%s", resource)); err != nil {
			return fmt.Errorf("resources: fail to prepare resource from %s", resource)
		}
		args := strings.Split(command, " ")
		_, err = exec.Command(args[0], args[1:]...).Output()

		if err != nil {
			return fmt.Errorf("resources: fail on execute '%s' (%v)", command, err)
		}
		log.Printf("'%s' resource loaded successfully!", resource)
		os.Chdir("..")
	}
	log.Printf("Only resources [%s] was load successfully!\n", strings.Join(resources, ", "))

	return os.Chdir("..")
}
