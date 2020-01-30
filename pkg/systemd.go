package pkg

import (
	. "github.com/saschagrunert/demo"
)

func Systemd() *Run {
	r := NewRun(
		"Systemd",
		"This demo shows how podman can generate systemd units",
	)

	r.Step(S(
		"Podman is capable of generating systemd units for container",
		"workloads. For example if we start a new nginx server.",
	), S(
		"podman run --name nginx -d nginx:alpine",
	))

	r.Step(S(
		"Then we can use `generate systemd` to create the unit",
	), S(
		"podman generate systemd -f -n nginx",
	))
	r.Step(nil, S("cat container-nginx.service"))

	r.Step(S(
		"Now we can load the unit via systemd",
	), S(
		"cp container-nginx.service ~/.config/systemd/user",
	))
	r.Step(nil, S("systemctl enable --user container-nginx"))
	r.Step(nil, S("podman stop nginx"))
	r.Step(nil, S("podman ps"))
	r.Step(nil, S("systemctl start --user container-nginx"))
	r.Step(nil, S("podman ps"))

	r.Step(S(
		"It is also possible to run systemd inside a container with podman",
		"For this we have to prepare a container image",
	), S(
		`cat <<EOT > Containerfile
FROM opensuse/tumbleweed
RUN zypper in -y curl nginx systemd-sysvinit
RUN systemctl enable nginx
CMD [ "/usr/sbin/init" ]
EOT`,
	))
	r.Step(nil, S("podman build -t systemd ."))

	r.Step(S(
		"No we can run the nginx web server via systemd inside podman",
	), S(
		"podman run --name systemd -d systemd",
	))
	r.Step(nil, S("podman exec -it systemd ps aux"))
	r.Step(nil, S("podman exec -it systemd curl 127.0.0.1"))

	return r
}
