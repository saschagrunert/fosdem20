package pkg

import (
	"os"

	. "github.com/saschagrunert/demo"
)

func Pods() *Run {
	r := NewRun(
		"Pods",
		"This demo shows basic pod management in podman",
	)

	r.Setup(func() error {
		_ = os.RemoveAll("pod.yml")
		return nil
	})

	r.Step(S(
		"Pods are a concept from Kubernetes, where multiple containers",
		"share a set of common resources like the network interface.",
		"Podman supports pods as well. Let’s create one",
	), S(
		"podman pod create --name my-pod",
	))

	r.Step(S(
		"A new infra container has been created which is the base for the pod",
	), S(
		"podman pod ps",
	))

	r.Step(S(
		"Now it is possible to add a container to the pod",
	), S(
		"podman run --pod my-pod -d alpine sleep infinity",
	))

	r.Step(S(
		"Pods per default do not share the PID namespace.",
		"They share the network interface and IPC namespace,",
		"which makes it possible to communicate via the same localhost",
	), S(
		"podman run --pod my-pod -d nginx:alpine",
	))

	r.Step(S(
		"The webserver is now reachable via localhost, too",
	), S(
		"podman run --pod my-pod alpine wget -q -O- 127.0.0.1",
	))

	r.Step(S(
		"It is also possible to generate a Kubernetes manifest",
		"from the pod",
	), S(
		"podman generate kube -f pod.yml my-pod",
	))
	r.Step(nil, S("cat pod.yml"))

	r.Step(S(
		"The manifest can now be used to play around in Kubernetes.",
		"This works also the other way around. We can import the",
		"pod.yml back into podman as well. Let’s ensure that no workload",
		"is running",
	), S(
		"podman pod rm -fa",
	))
	r.Step(nil, S("podman ps -a"))

	r.Step(S(
		"Now we can utilize the `play kube` subcommand to re-create everything",
		"we deleted in the previous step",
	), S(
		"podman play kube pod.yml",
	))

	r.Step(S(
		"Everything should be now up and running again",
	), S(
		"podman ps -a",
	))

	return r
}
