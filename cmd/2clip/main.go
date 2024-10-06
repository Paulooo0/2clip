package main

import (
	clip "github.com/Paulooo0/2clip/internal/2clip"
	"github.com/Paulooo0/2clip/internal/2clip/add"
	"github.com/Paulooo0/2clip/internal/2clip/auth"
	"github.com/Paulooo0/2clip/internal/2clip/get"
	"github.com/Paulooo0/2clip/internal/2clip/list"
	"github.com/Paulooo0/2clip/internal/2clip/remove"
)

func main() {
	clip.RootCmd.AddCommand(add.AddCmd)
	add.AddCmdFlags()

	clip.RootCmd.AddCommand(get.GetCmd)
	get.GetCmdFlags()

	clip.RootCmd.AddCommand(remove.RemoveCmd)
	remove.RemoveCmdFlags()

	clip.RootCmd.AddCommand(list.ListKeysCmd)

	clip.RootCmd.AddCommand(auth.AuthCmd)
	auth.AuthCmdFlags()

	clip.RootCmd.Execute()
}
