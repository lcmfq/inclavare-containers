// +build linux

package main

import (
	"github.com/alibaba/inclavare-containers/shim/runtime/v2/rune/v2"
	"github.com/containerd/containerd/runtime/v2/shim"
)

func main() {
	shim.Run("io.containerd.rune.v2", v2.New)
}
