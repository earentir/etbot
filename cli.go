package main

import cli "github.com/jawher/mow.cli"

func cliTBStart(cmd *cli.Cmd) {
	cmd.Spec = "[-s]"

	TBSAllowOutput := cmd.BoolOpt("S", false, "Allow the Bot to output to twitch")
	_ = TBSAllowOutput
}
