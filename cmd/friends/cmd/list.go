// Copyright 2022 vscode
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cmd

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/anvari1313/splitwise.go"
	"github.com/spf13/cobra"
	"github.com/splitwise-batch/pkg/security"
)

var credsFile string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List friends for the current user",
	Run:   list,
}

func list(cmd *cobra.Command, args []string) {
	auth, err := security.Authenticate(credsFile)
	if err != nil {
		fmt.Println(err)
	}

	client := splitwise.NewClient(auth)

	friends, err := client.Friends(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	// TODO: Option to save to file
	j, _ := json.MarshalIndent(friends, "", "    ")
	fmt.Println(string(j))
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.PersistentFlags().StringVar(&credsFile, "credentials", ".credentials.json",
		"path to file containing the SplitWise authentication key (default is .credentials.json)")
}
