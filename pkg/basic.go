package pkg

import (
	. "github.com/saschagrunert/demo"
)

func Basic() *Run {
	r := NewRun(
		"Basics",
		"The first demo demonstrates the basic usage of podman",
	)

	r.Step(S(
		"Running a container in podman is a trivial task.",
		"Specified container images will be automatically downloaded",
		"in the same way as we would expect it from different",
		"container runtimes.",
	), S(
		"podman run alpine echo hello world",
	))

	r.Step(S(
		"We can utilize mainly all command line parameters we already",
		"know from docker",
	), S(
		"podman run --name container -d alpine sleep infinity",
	))

	r.Step(S(
		"Now the container should be running in the background",
	), S(
		"podman ps",
	))

	r.Step(S(
		"And we can exec that container, too",
	), S(
		"podman exec container ps aux",
	))

	r.Step(S(
		"There is no deamon running for podman, which means that some",
		"background process has to ensure the lifetime of the container.",
		"This is the job of the container monitoring tool `conmon`",
	), S(
		"ps x | grep conmon",
	))

	r.Step(S(
		"We can see that conmon interacts with runc to keep track of the",
		"containers lifecycle.",
	), S(
		"runc list -f json | jq .",
	))

	r.Step(S(
		"If we now kill the containers process, then conmon reports to podman",
		"that the container has exited",
	), S(
		"pkill -9 sleep",
	))

	r.Step(nil, S(
		"podman ps -a",
	))

	return r
}
