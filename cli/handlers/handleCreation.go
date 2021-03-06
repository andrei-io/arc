package handlers

import (
	"arc/project"
	"arc/util"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var (
	projectLocation string
	projectChoices  []string
	qsCreate        []*survey.Question
	tagChoices      []string
	tagLocation     string
	homeLocation    string
)

func init() {
	// This function is ran first in this file
	// so i can set config vars here
	projectLocation = fmt.Sprintf("%v/.project-templates", os.Getenv("HOME"))
	projectChoices, _ = util.GetFolders(projectLocation)
	homeLocation, _ = os.UserHomeDir()
	tagLocation = fmt.Sprintf("%v/dev", homeLocation)
	tagChoices, _ = util.GetFolders(tagLocation)
	// For future use, "None" will serve as a way not to add it to the central store
	tagChoices = append(tagChoices, "None")

	current, _ := os.Getwd()

	preselectedTag := util.GetTag(current, tagChoices)

	qsCreate = []*survey.Question{
		{
			Name:     "name",
			Prompt:   &survey.Input{Message: "Project name:"},
			Validate: util.ValidateName,
		},
		{
			Name: "language",
			Prompt: &survey.Select{
				Message: "Project language:",
				Options: projectChoices,
			},
		},
		{
			Name:     "folder",
			Prompt:   &survey.Input{Message: "Where to create:"},
			Validate: util.ValidateName,
		},
		{
			Name: "Tag",
			Prompt: &survey.Select{
				Message: "Tag: ",
				Options: tagChoices,
				Default: preselectedTag,
			},
		},
		{
			Name: "repo",
			Prompt: &survey.Select{
				Message: "Create git repo?",
				Options: []string{
					"Yes",
					"No",
				},
			},
		},
	}
}

// Type of response
type answerCreate struct {
	Name   string
	Lang   string `survey:"language"`
	Folder string
	Repo   string
	Tag    string
}

func HandleCreation(cmd *cobra.Command, args []string) {
	answers := answerCreate{}

	err := survey.Ask(qsCreate, &answers)

	if err != nil {
		fmt.Printf("Error while asking questions: %s", err.Error())
		return
	}

	var langs []string
	langs = append(langs, answers.Lang)
	project := project.Project{
		Name:     answers.Name,
		Lang:     langs,
		Repo:     answers.Repo == "Yes",
		Location: answers.Folder,
		Tag:      answers.Tag,
	}

	err = project.CreatePoject()
	if err != nil {
		fmt.Printf("Error while creating project:\n\t %s", err.Error())
	}

}
