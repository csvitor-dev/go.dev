package resources

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func fail(err error) error {
	message := fmt.Sprintf("fail to prepare stylesheet (%s)", err.Error())
	return fmt.Errorf("resource: %s", message)
}

func PrepareTailwind() error {
	err := os.Chdir("./resources/tailwindcss")

	if err != nil {
		return fail(err)
	}
	_, err = exec.Command("npm", "run", "build:css").Output()

	if err != nil {
		return fail(err)
	}
	log.Println("Tailwind classes loaded!")

	return os.Chdir("../..")
}
