package pkg

import (
	. "github.com/saschagrunert/demo"
)

func Networking() *Run {
	r := NewRun(
		"Networking",
		"This demo shows how networking is operating in podman",
	)

	r.StepCanFail(S(
		"Podman uses the Container Networking Interface (CNI) project",
		"to provide networking capablities",
	), S(
		"sudo podman network ls",
	))

	r.StepCanFail(S(
		"The configuration of the network plugins will be done in a system-wide",
		"directory, like /etc/cni/net.d",
	), S(
		"jq . /etc/cni/net.d/87-podman-bridge.conflist",
	))

	r.StepCanFail(S(
		"Podman chooses the default CNI network based on its configuration",
		"in /etc/containers/libpod.conf",
	), S(
		"grep -B8 cni_default_network /etc/containers/libpod.conf",
	))

	r.StepCanFail(S(
		"The CNI bridge plugin will ensure that we get the correct IP",
		"addresses assigned, let’s checkout how this looks like",
	), S(
		"sudo podman run --name container -d alpine sleep infinity",
	))
	r.StepCanFail(nil, S("sudo podman exec container ip addr"))
	r.StepCanFail(nil, S("ip netns"))
	r.StepCanFail(nil, S("ip link list type veth"))

	r.StepCanFail(S(
		"When running podman in rootless mode, then the project slirp4netns",
		"provides a user-mode networking via a TAP device.",
	), S(
		"podman run alpine ip addr",
	))

	r.StepCanFail(S(
		"It is also possible to run a container in host network mode",
	), S(
		"podman run --net host alpine ip addr",
	))

	r.StepCanFail(S(
		"Or join two network namespaces together… But for those use-cases we",
		"usually have pods, right?",
	), nil)

	return r
}
