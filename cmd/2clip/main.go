package main

import (
	clip "github.com/Paulooo0/2clip/pkg/2clip"
	"github.com/Paulooo0/2clip/pkg/2clip/util"
)

func main() {

	util.RootCmd.AddCommand(clip.AddCmd)

	util.RootCmd.AddCommand(clip.GetCmd)

	util.RootCmd.AddCommand(clip.RemoveCmd)

	util.RootCmd.AddCommand(clip.ListKeysCmd)
	clip.AddCmdFlags()

	util.RootCmd.Execute()
}
