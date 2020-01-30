package pkg

import (
	. "github.com/saschagrunert/demo"
)

func Isolation() *Run {
	r := NewRun(
		"Isolation",
		"This demo shows the isolation between containers and the host",
	)

	r.Step(S(
		"With podman we have the chance to change the isolation between",
		"the running containers and their host.",
		"For example, we could remove the PID namespace isolation",
	), S(
		"podman run --pid=host alpine ps aux | head -20",
	))

	r.Step(S(
		"Or we could let two container share their namespace.",
		"For this we have to create a start container.",
	), S(
		"podman run --name first -d alpine sleep infinity",
	))

	r.Step(S(
		"Now we can start another container and specifying to join the PID",
		"namespace with the first one.",
	), S(
		"podman run --pid container:first --name second -d alpine sleep infinity",
	))

	r.Step(S(
		"We now have two running containers",
	), S(
		"podman ps",
	))

	r.Step(S(
		"But they share the same PID namespaces",
	), S(
		"podman exec -it first ps aux",
	))

	r.Step(S(
		"Podman is able to provide further namespace information",
	), S(
		"podman ps --ns",
	))

	return r
}
