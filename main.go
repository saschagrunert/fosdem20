package main

import (
	. "github.com/saschagrunert/demo"
	"github.com/saschagrunert/fosdem20/pkg"
	"github.com/urfave/cli/v2"
)

func main() {
	d := New()
	d.Name = "fosdem20"
	d.HideVersion = true
	d.Usage = "Demo material used for the Podman talk at FOSDEM 2020"
	d.Authors = []*cli.Author{{
		Name: "Sascha Grunert", Email: "sgrunert@suse.com",
	}}
	d.Setup(pkg.Cleanup)
	d.Cleanup(pkg.Cleanup)

	d.Add(pkg.Basic(), "basics", "run the basics of podman demo")
	d.Add(pkg.Isolation(), "isolation", "run the isolation demo")
	d.Add(pkg.Pods(), "pods", "run the pod management demo")
	d.Add(pkg.Security(), "security", "run the security demo")
	d.Add(pkg.Networking(), "networking", "run the networking demo")
	d.Add(pkg.Mount(), "mount", "run the mount container storage demo")
	d.Add(pkg.Images(), "images", "run the container images demo")
	d.Add(pkg.Systemd(), "systemd", "run the systemd demo")
	d.Add(pkg.Registries(), "registries", "run the registries demo")
	d.Add(pkg.RegistryMirrors(), "registry-mirror", "run the registry mirror demo")
	d.Run()
}
