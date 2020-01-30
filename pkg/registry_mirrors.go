package pkg

import (
	"io/ioutil"

	. "github.com/saschagrunert/demo"
)

const r2 = `unqualified-search-registries = ["docker.io"]

[[registry]]
location = "docker.io/library"
mirror = [
 { location = "localhost/mirror-path" },           # doesnt work
 { location = "localhost:5000", insecure = true }, # should work
]
`

func RegistryMirrors() *Run {
	r := NewRun(
		"Registry Mirrors",
		"This demo shows how to configure registries mirrors in podman",
	)

	r.Setup(func() error {
		if err := ioutil.WriteFile(f, []byte(r2), 0o644); err != nil {
			return err
		}
		return Ensure(
			"[ ! -f "+k+".bak ] && sudo mv "+k+" "+k+".bak || true",
			"sudo cp "+f+" "+k,
			"rm "+f,
		)
	})

	r.Step(S(
		"Registry mirrors are especially useful in air-gapped scenarios,",
		"where access to the internet is limited.",
		"A registry mirror can be configured like this",
	), S(
		`grep -A5 '^\[\[registry\]\]' `+k,
	))

	r.Step(S(
		"To let the mirror work, we would have to setup one",
		"For this we use podman to setup a local registry",
	), S(
		"podman run --rm --name=registry -p 5000:5000 -d registry",
	))

	r.Step(S(
		"Now we can transfer our target image into the local registry",
	), S(
		"podman pull hello-world",
	))
	r.Step(nil, S(
		"podman tag hello-world localhost:5000/hello-world",
	))
	r.Step(nil, S(
		"podman push --tls-verify=false localhost:5000/hello-world",
	))

	r.Step(S(
		"If we now pull an image from docker.io, then we first lookup our",
		"configured mirrors.",
	), S(
		"podman --log-level=debug pull hello-world 2>&1 | grep rewritten",
	))

	return r
}
