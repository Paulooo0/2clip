package main

import (
	clip "github.com/Paulooo0/2clip/pkg/2clip"
	"github.com/Paulooo0/2clip/pkg/2clip/add"
	"github.com/Paulooo0/2clip/pkg/2clip/auth"
	"github.com/Paulooo0/2clip/pkg/2clip/get"
	"github.com/Paulooo0/2clip/pkg/2clip/list"
	"github.com/Paulooo0/2clip/pkg/2clip/remove"
)

func main() {
	clip.RootCmd.AddCommand(add.AddCmd)
	add.AddCmdFlags()

	clip.RootCmd.AddCommand(get.GetCmd)

	clip.RootCmd.AddCommand(remove.RemoveCmd)

	clip.RootCmd.AddCommand(list.ListKeysCmd)

	clip.RootCmd.AddCommand(auth.AuthCmd)
	auth.AuthCmdFlags()

	clip.RootCmd.Execute()
}
