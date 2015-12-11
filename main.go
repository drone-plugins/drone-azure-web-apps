package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/drone-plugins/drone-git-push/repo"
	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

var (
	build     string
	buildDate string
)

func main() {
	fmt.Printf("Drone Azure Web Apps Plugin built at %s\n", buildDate)

	workspace := drone.Workspace{}
	repo := drone.Repo{}
	build := drone.Build{}
	vargs := Params{}

	plugin.Param("workspace", &workspace)
	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	if len(vargs.Username) == 0 {
		fmt.Println("Please provide an username")

		os.Exit(1)
		return
	}

	if len(vargs.Password) == 0 {
		fmt.Println("Please provide a password")

		os.Exit(1)
		return
	}

	if len(vargs.Site) == 0 {
		vargs.Site = repo.Name
	}

	if len(vargs.Slot) == 0 {
		vargs.Slot = vargs.Site
	}

	err := run(&workspace, &build, &vargs)

	if err != nil {
		fmt.Println(err)

		os.Exit(1)
		return
	}
}

func run(workspace *drone.Workspace, build *drone.Build, vargs *Params) error {
	repo.GlobalName(build).Run()
	repo.GlobalUser(build).Run()

	defer func() {
		execute(
			repo.RemoteRemove(
				"deploy"),
			workspace)
	}()

	cmd := repo.RemoteAdd(
		"deploy",
		remote(vargs))

	if err := execute(cmd, workspace); err != nil {
		return err
	}

	if vargs.Commit {
		if err := execute(repo.ForceAdd(), workspace); err != nil {
			return err
		}

		if err := execute(repo.ForceCommit(), workspace); err != nil {
			return err
		}
	}

	cmd = repo.RemotePush(
		"deploy",
		"master",
		vargs.Force)

	if err := execute(cmd, workspace); err != nil {
		return err
	}

	return nil
}

func remote(vargs *Params) string {
	return fmt.Sprintf(
		"https://%s:%s@%s.scm.azurewebsites.net:443/%s.git",
		vargs.Username,
		vargs.Password,
		vargs.Slot,
		vargs.Site)
}

func execute(cmd *exec.Cmd, workspace *drone.Workspace) error {
	trace(cmd)

	cmd.Dir = workspace.Path
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	return cmd.Run()
}

func trace(cmd *exec.Cmd) {
	fmt.Println("$", strings.Join(cmd.Args, " "))
}
