package pkg

import (
	. "github.com/saschagrunert/demo"
)

func Mount() *Run {
	r := NewRun(
		"Mount",
		"This demo shows how to mount the container storage",
		"of a running container",
	)

	r.Step(S(
		"Let’s create a container which tries to access a file",
	), S(
		"sudo podman run --name cat-test -d alpine ",
		"sh -c 'while true; do cat test; sleep 2; done'",
	))

	r.Step(S(
		"The file should not be available, which can be verified via the logs",
	), S(
		"sudo podman logs -t cat-test",
	))

	r.Step(S(
		"Now the containers file-system can be mounted",
	), S(
		"sudo podman mount cat-test",
	))

	r.Step(S(
		"And we can write the test file to it",
	), S(
		"echo hello world | sudo tee -a $(sudo podman mount cat-test)/test",
	))

	r.Step(S(
		"Now we can verify via the logs that the container has",
		"access to the file ",
	), S(
		"sudo podman logs -t cat-test",
	))

	r.Step(S(
		"Afterwards we can unmount the directory again",
	), S(
		"sudo podman unmount cat-test",
	))

	return r
}
