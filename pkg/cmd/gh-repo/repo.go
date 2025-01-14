package repo

import (
	"github.com/MakeNowJust/heredoc"
	repoBrowseCmd "github.com/abdfnx/gh/pkg/cmd/gh-repo/browse"
	repoCloneCmd "github.com/abdfnx/gh/pkg/cmd/gh-repo/clone"
	repoCreateCmd "github.com/abdfnx/gh/pkg/cmd/gh-repo/create"
	creditsCmd "github.com/abdfnx/gh/pkg/cmd/gh-repo/credits"
	repoForkCmd "github.com/abdfnx/gh/pkg/cmd/gh-repo/fork"
	gardenCmd "github.com/abdfnx/gh/pkg/cmd/gh-repo/garden"
	repoListCmd "github.com/abdfnx/gh/pkg/cmd/gh-repo/list"
	repoSyncCmd "github.com/abdfnx/gh/pkg/cmd/gh-repo/sync"
	repoViewCmd "github.com/abdfnx/gh/pkg/cmd/gh-repo/view"
	"github.com/abdfnx/gh/pkg/cmdutil"
	git_config "github.com/david-tomson/tran-git"
	"github.com/abdfnx/gh/utils"
	"github.com/spf13/cobra"

	"github.com/abdfnx/gh/pkg/cmd/factory"
)

func NewCmdRepo(f *cmdutil.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gh-repo <command>",
		Short: "Create, clone, fork, and view repositories.",
		Long:  `Work with GitHub repositories`,
		Example: heredoc.Doc(`
			tran gh-repo create
			tran gh-repo clone moby/moby
		`),
		Annotations: map[string]string{
			"help:arguments": heredoc.Doc(`
				A repository can be supplied as an argument in any of the following formats:
				- "OWNER/REPO"
				- by URL, e.g. "https://github.com/OWNER/REPO"
			`),
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			username := git_config.GitConfig()

			if username == ":username" {
				utils.AuthMessage()

			} else {
				cmd.Help()
			}

			return nil
		},
	}

	repoResolvingCmdFactory := *f
	repoResolvingCmdFactory.BaseRepo = factory.SmartBaseRepoFunc(f)

	cmd.AddCommand(repoViewCmd.NewCmdView(f, nil))
	cmd.AddCommand(repoForkCmd.NewCmdFork(f, nil))
	cmd.AddCommand(repoCloneCmd.NewCmdClone(f, nil))
	cmd.AddCommand(repoCreateCmd.NewCmdCreate(f, nil))
	cmd.AddCommand(repoListCmd.NewCmdList(f, nil))
	cmd.AddCommand(repoSyncCmd.NewCmdSync(f, nil))
	cmd.AddCommand(creditsCmd.NewCmdRepoCredits(f, nil))
	cmd.AddCommand(gardenCmd.NewCmdGarden(f, nil))
	cmd.AddCommand(repoBrowseCmd.NewCmdBrowse(&repoResolvingCmdFactory, nil))

	return cmd
}
