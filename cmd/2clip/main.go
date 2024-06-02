package main

import (
	clip "github.com/Paulooo0/2clip/pkg/2clip"
)

func main() {

	clip.RootCmd.AddCommand(clip.AddCmd)

	clip.RootCmd.AddCommand(clip.GetCmd)

	clip.RootCmd.AddCommand(clip.RemoveCmd)

	clip.RootCmd.AddCommand(clip.ListKeysCmd)
	clip.AddCmdFlags()

	clip.RootCmd.AddCommand(clip.AuthCmd)

	clip.RootCmd.Execute()
}
