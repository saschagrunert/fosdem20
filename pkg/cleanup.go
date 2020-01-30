package pkg

import (
	. "github.com/saschagrunert/demo"
	"github.com/urfave/cli/v2"
)

func Cleanup(ctx *cli.Context) error {
	_ = Ensure(
		"podman pod rm -fa",
		"podman rm -fa",
		"sudo podman pod rm -fa",
		"sudo podman rm -fa",
	)
	return nil
}
