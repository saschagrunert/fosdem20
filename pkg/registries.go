package pkg

import (
	"io/ioutil"

	. "github.com/saschagrunert/demo"
)

const r1 = `# An array of host[:port] registries to try when
# pulling an unqualified image, in order.
unqualified-search-registries = ["docker.io", "quay.io"]

[[registry]]
prefix = "localhost"
location = "docker.io/library"
blocked = false
insecure = false
`

const (
	f = "registries.conf"
	k = "/etc/containers/" + f
)

func Registries() *Run {
	r := NewRun(
		"Registry Configurations",
		"This demo shows how to configure registries with podman",
	)

	r.Setup(func() error {
		if err := ioutil.WriteFile(f, []byte(r1), 0o644); err != nil {
			return err
		}

		return Ensure(
			"[ ! -f "+k+".bak ] && sudo mv "+k+" "+k+".bak || true",
			"sudo cp "+f+" "+k,
			"rm "+f,
		)
	})

	r.Step(S(
		"Podman supports multiple registry configuration syntaxes.",
		"From now on we focus on the latest version, which comes with the",
		"highest set of features. The default configuration can be found",
		"at "+k,
	), S(
		"grep -B2 unqualified-search-registries "+k,
	))

	r.Step(S(
		"The `unqualified-search-registries` allows us to pull images without",
		"prepending a registry prefix",
	), S(
		"podman pull hello-world",
	))

	r.Step(S(
		"A single registry can be specified within a [[registry]] entry",
	), S(
		`grep -A4 '^\[\[registry\]\]' `+k,
	))

	r.Step(S(
		"We have been rewritten the docker library to localhost.",
		"Now it is possible to pull via localhost",
	), S(
		"podman --log-level=debug pull localhost/alpine 2>&1 | grep rewritten",
	))

	return r
}
