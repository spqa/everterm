/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spqa/everterm/pkg/client"
	"github.com/spqa/everterm/pkg/config"
	"log"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Config Everterm",
	// Long: ``,
	Run: configureHandler,
}

func configureHandler(cmd *cobra.Command, args []string) {
	prompt := promptui.Prompt{
		Label: "Enter API key",
	}
	result, err := prompt.Run()
	if err != nil {
		log.Fatalln("Prompt failed")
	}
	err = validateApiKey(result)
	if err != nil {
		log.Fatal("Invalid API key", err)
	}
	if result != "" {
		config.SetConfigAndSave(config.ApiKey, result)
	}
}

func validateApiKey(input string) error {
	if input == "" {
		return nil
	}
	userStore, err := client.GetUserStoreClient()
	if err != nil {
		return err
	}
	user, err := userStore.GetUser(context.Background(), input)
	if err != nil {
		return err
	}
	log.Printf("User token validate success: userId %d", user.ID)
	return nil
}

func init() {
	rootCmd.AddCommand(configureCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configureCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configureCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
