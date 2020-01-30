package pkg

import (
	. "github.com/saschagrunert/demo"
)

func Security() *Run {
	r := NewRun(
		"Security",
		"This demo shows the container security features of podman",
	)

	r.StepCanFail(S(
		"Podman allows us to limit the Linux capabilities a container",
		"has. For example we could drop all of them",
	), S(
		"podman run --cap-drop=all alpine ping 127.0.0.1",
	))

	r.Step(S(
		"Or we add just the single capability we need",
	), S(
		"podman run --cap-drop=all --cap-add=net_raw alpine ping -c3 127.0.0.1",
	))

	r.Step(S(
		"We could also use SECCOMP to limit the container in a more",
		"powerful way. For that we have to define a SECCOMP profile JSON",
	), S(
		`cat <<EOT > profile.json
{
  "defaultAction": "SCMP_ACT_ALLOW",
  "architectures": ["SCMP_ARCH_X86_64"],
  "syscalls": [ { "names": ["mkdir"], "action": "SCMP_ACT_ERRNO" } ]
}
EOT`,
	))

	r.StepCanFail(S(
		"Now podman can utilize the profile to limit the `mkdir(2)` syscall",
	), S(
		"podman run --security-opt seccomp=profile.json alpine mkdir test",
	))

	r.Step(S(
		"Podman supports AppArmor and SELinux as well.",
		"For example if we load a simple 'deny-write' AppArmor profile",
	), S(
		`cat <<EOT > deny-write
#include <tunables/global>

profile deny-write flags=(attach_disconnected) {
  #include <abstractions/base>

  file,

  # Deny all file writes.
  deny /** w,
}
EOT`,
	))

	r.Step(S(
		"The profile has to be loaded before running podman",
	), S(
		"sudo apparmor_parser -r deny-write",
	))

	r.StepCanFail(S(
		"Now podman can use the profile in the same way as for SECCOMP",
	), S(
		"sudo podman run --security-opt apparmor=deny-write alpine mkdir test",
	))

	r.Step(S(
		"It is always recommended to run container based workloads as non root.",
		"This reduces the possible attack surface. It’s fairly easy to switch",
		"to a different user in podman via the `--user` flag.",
	), S(
		"podman run --user guest alpine whoami",
	))

	r.StepCanFail(S(
		"This user should not be able to modify the root file system of the container",
	), S(
		"podman run --user guest alpine touch test",
	))

	return r
}
