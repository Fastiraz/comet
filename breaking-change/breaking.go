package breakchange

import (
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/huh"
)

type Order struct {
	BreakingChange string
}

func IsBreakingChange() bool {
	var order Order
	accessible, _ := strconv.ParseBool(os.Getenv("ACCESSIBLE"))

	form := huh.NewForm(
		// huh.NewGroup(huh.NewNote().
		// 	Title("Breaking Change").
		// 	Description("You are about to choose if your commit contains breaking changes.\n\nA breaking change commit is a commit that introduces a breaking change.\n\nIf so, choose `yes` else `no`?\n\n").
		// 	Next(true).
		// 	NextLabel("Next"),
		// ),

		huh.NewGroup(
			huh.NewSelect[string]().
				Options(
					huh.NewOptions(
						"Yes",
						"No",
					)...).
				Title("Is this a breaking change?").
				Description("Choose if this commit introduces a breaking change.").
				Value(&order.BreakingChange),
		),
	).WithAccessible(accessible)

	err := form.Run()
	if err != nil {
		fmt.Println("Uh oh:", err)
		os.Exit(1)
	}

	return order.BreakingChange == "Yes"
}
