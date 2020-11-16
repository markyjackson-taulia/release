package CheckPrerequisites

import (
	"fmt"
	"os"
	"os/exec"
)

func CheckPrerequisites() error {
	docker_version := exec.Command("docker", "-v")
	docker_version_Err := system.RunOSCommand(docker_version, logger)
	if docker_version_Err != nil {
		return errors.Wrapf(docker_version_Err, "Attempting to get docker version")
	}
}
