package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/appleboy/drone-git-push/repo"
)

type (
	Author struct {
		Name  string
		Email string
	}

	Commit struct {
		Author Author
	}

	Repo struct {
		Name string
	}

	Config struct {
		Username string
		Password string
		Site     string
		Slot     string
		Force    bool
		Commit   bool
	}

	Plugin struct {
		Commit Commit
		Repo   Repo
		Config Config
	}
)

func (p Plugin) Exec() error {
	if p.Config.Site == "" {
		p.Config.Site = p.Repo.Name
	}

	if p.Config.Slot == "" {
		p.Config.Slot = p.Repo.Name
	}

	if err := p.execute(repo.GlobalName(p.Commit.Author.Name)); err != nil {
		return err
	}

	if err := p.execute(repo.GlobalUser(p.Commit.Author.Email)); err != nil {
		return err
	}

	if err := p.execute(repo.RemoteAdd("deploy", p.remote())); err != nil {
		return err
	}

	defer p.execute(repo.RemoteRemove("deploy")) // nolint: errcheck

	if p.Config.Commit {
		if err := p.execute(repo.ForceAdd()); err != nil {
			return err
		}

		if err := p.execute(repo.ForceCommit("", false)); err != nil {
			return err
		}
	}

	return p.execute(repo.RemotePush("deploy", "master", p.Config.Force, false))
}

func (p Plugin) execute(cmd *exec.Cmd) error {
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	p.trace(cmd)
	return cmd.Run()
}

func (p Plugin) trace(cmd *exec.Cmd) {
	fmt.Println("$", strings.Join(cmd.Args, " "))
}

func (p Plugin) remote() string {
	return fmt.Sprintf(
		"https://%s:%s@%s.scm.azurewebsites.net:443/%s.git",
		p.Config.Username,
		p.Config.Password,
		p.Config.Slot,
		p.Config.Site,
	)
}
