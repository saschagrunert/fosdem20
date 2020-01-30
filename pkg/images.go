package pkg

import (
	. "github.com/saschagrunert/demo"
)

func Images() *Run {
	r := NewRun(
		"Images",
		"This demo shows how podman manages container images",
	)

	r.Step(S(
		"Surely podman is able to build container images",
	), S(
		"echo FROM scratch > Containerfile",
	))
	r.Step(nil, S("podman build ."))

	r.Step(S(
		"Podman uses buildah under the hood to build containers.",
		"This means we can use mostly all features of buildah in",
		"podman as well, like the Containerfile preprocessor:",
	), S(
		`cat <<EOT > Containerfile.in
FROM alpine
#include "./Containerfile-clang"
EOT`,
	))

	r.Step(nil, S("echo RUN apk add clang > Containerfile-clang"))
	r.Step(nil, S("podman build -t clang -f Containerfile.in ."))

	r.Step(S(
		"Podman is now able to show an image tree, which",
		"contains useful image information",
	), S(
		"podman image tree clang",
	))

	return r
}
