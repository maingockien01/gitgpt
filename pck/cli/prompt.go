package cli

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
)

// CLI prompt
type PromptSelect struct {
	Label string
	Items []string
}

const OTHER = "Other"

func PromptGetSelect(pc PromptSelect) string {
	items := pc.Items
	var result string
	var err error

	// Add "Other" option
	items = append(items, OTHER)

	prompt := promptui.Select{
		Label: pc.Label,
		Items: items,
	}

	_, result, err = prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}
