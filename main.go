package main

import (
	_ "broadcast_back_end/internal/packed"
	"github.com/gogf/gf/v2/os/gctx"

	"broadcast_back_end/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
