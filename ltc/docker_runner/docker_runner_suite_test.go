package docker_runner_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestAppRunning(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DockerRunner Suite")
}