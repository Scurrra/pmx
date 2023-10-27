package cmd

import (
	"strings"

	"github.com/Scurrra/pmx/pmxconfig"
	gh_license "github.com/Shresht7/gh-license/api"
	. "github.com/eminarican/safetypes"
	"github.com/sa-/slicefunk"
	"github.com/spf13/cobra"
	"golang.org/x/mod/semver"
)

var (
	name        string
	authors     []string
	projectType string // pmxconfig.ProjectType
	version     string
	license     string
	description string
	repository  string
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize new documentation.",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if noInteract {
			return pmxconfig.Configure(
				name,
				authors,
				pmxconfig.ProjectType(projectType),
				version,
				license,
				description,
				repository,
			)
		}

		if len(args) > 0 && args[0] != "." {
			name = args[0]
		} else if !cmd.Flags().Lookup("name").Changed {
			name = promptGetInput(
				promptContent{
					"Name of the project is not provided",
					"Name of the project: ",
					None[string](),
				},
			)
		}

		if !cmd.Flags().Lookup("project_type").Changed {
			projectType = promptGetSelect(
				promptContent{
					"Type of the project is not provided",
					"Type of the project: ",
					None[string](),
				},
				[]pmxconfig.ProjectType{pmxconfig.LIB, pmxconfig.BIN},
			)
		}

		if !cmd.Flags().Lookup("authors").Changed {
			authors = strings.Split(promptGetInput(
				promptContent{
					"Authors of the project is not provided",
					"Authors of the project: ",
					Some[string](""),
				},
			), ",")
		}

		if !cmd.Flags().Lookup("version").Changed {
			version = promptGetInput(
				promptContent{
					"Version of the project is not provided",
					"Version of the project: ",
					Some[string]("v0.0.1"),
				},
			)
			for ; !semver.IsValid(version); version = promptGetInput(
				promptContent{
					"Version of the project is not provided",
					"Version of the project: ",
					Some[string]("v0.0.1"),
				},
			) {
			}
		}

		if !cmd.Flags().Lookup("license").Changed {
			license = promptGetSelect(
				promptContent{
					"Type of the project is not provided",
					"Type of the project: ",
					None[string](),
				},
				slicefunk.Map[gh_license.LicenseSimple, string](licenses, func(ls gh_license.LicenseSimple) string {
					return ls.Spdx_id
				}),
			)
		}

		if !cmd.Flags().Lookup("description").Changed {
			description = promptGetInput(
				promptContent{
					"Description of the project is not provided",
					"Description of the project: ",
					Some[string](""),
				},
			)
		}

		if !cmd.Flags().Lookup("repository").Changed {
			repository = promptGetInput(
				promptContent{
					"Repository url of the project is not provided",
					"Repository url of the project: ",
					Some[string](""),
				},
			)
		}

		return pmxconfig.Configure(
			name,
			authors,
			pmxconfig.ProjectType(projectType),
			version,
			license,
			description,
			repository,
		)
	},
}
