package main

import (
	clip "github.com/Paulooo0/2clip/pkg/2clip"
	"github.com/Paulooo0/2clip/pkg/2clip/auth"
)

func main() {
	clip.RootCmd.AddCommand(clip.AddCmd)
	clip.AddCmdFlags()

	clip.RootCmd.AddCommand(clip.GetCmd)

	clip.RootCmd.AddCommand(clip.RemoveCmd)

	clip.RootCmd.AddCommand(clip.ListKeysCmd)

	clip.RootCmd.AddCommand(auth.AuthCmd)
	auth.AuthCmdFlags()

	clip.RootCmd.Execute()
}
