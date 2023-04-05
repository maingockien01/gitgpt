package cli

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gitgpt/pck/git"
	"gitgpt/pck/gptclient"

	"github.com/spf13/cobra"
)

// CLI
var (
	// RootCmd is the root command for the application.
	RootCmd = &cobra.Command{
		Use:   "gitgpt",
		Short: "gitgpt is a command line tool for generating git commit messages.",
		Long: `gitgpt is a command line tool for generating git commit messages.
It uses OpenAI's GPT-3 to generate meaningful git commit messages.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Main logic
			main()
		},
	}
)

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	// Get secret key from env
	apiKey := os.Getenv("OPENAI_API_KEY")

	// Get client
	client := gptclient.NewClient(apiKey)

	// Get project and ticket
	// projectTicket, err := git.GetProjectTicket()

	// if err != nil {
	// 	panic(err)
	// }

	// Get git diff
	gitDiff, err := git.GetGitDiff()

	if err != nil {
		panic(err)
	}

	// Get git log
	gitLog, err := git.GetGitLog()

	if err != nil {
		panic(err)
	}

	// Get branch name
	branchName, err := git.GetBranchName()

	if err != nil {
		panic(err)
	}

	// Train model
	ctx := context.Background()
	trainMsg := "Here is the histories of git commit messages: \n" + gitLog + "\n" + "Here is the current branch name: \n" + branchName + "\n" + "Here is the git diff: \n" + gitDiff + "\n" + "Suggest a short and meaningful git commit message:"
	commitMsg := OTHER
	commitMsgs := []string{}

	promptMsg := trainMsg
	// Get completion
	for commitMsg == OTHER {

		conmitMsgs, err := client.GetCompletion(ctx, promptMsg)

		if err != nil {
			panic(err)
		}

		prompt := PromptSelect{
			Label: "Select messages",
			Items: gptclient.ParseMessages(conmitMsgs),
		}
		commitMsg = PromptGetSelect(prompt)

		commitMsgs = append(commitMsgs, commitMsg)

		promptMsg = "List of git messages suggested by GPT-3: \n" + strings.Join(commitMsgs, "\n") + "\n" + trainMsg

	}
	// Copy commit message to clipboard

	commitCmd := exec.Command("pbcopy")
	commitCmd.Stdin = strings.NewReader(fmt.Sprintf("git commit -m \"%s\"", commitMsg))
	commitCmd.Run()
}
