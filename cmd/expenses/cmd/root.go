// Copyright (c) 2022 Alvaro Silva <alvaro.fagner@gmail.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/anvari1313/splitwise.go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile      string
	credsFile    string
	ExpensesFile string
	currencyCode string
	groupID      uint32
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "expenses",
	Short: "A CLI application for automating the creation of Splitwise expenses in batches",
	Long:  "expenses byshares --from-file expenses.csv --user-share user_one:2 --user-share user_two:1",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func ReadExpenses(filePath string) ([]splitwise.Expense, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	fileLines := readLines(f)

	f.Close()

	var expenses []splitwise.Expense
	for _, line := range fileLines {
		expense, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, *expense)
	}

	return expenses, nil
}

func readLines(f *os.File) []string {
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines
}

func parseLine(line string) (*splitwise.Expense, error) {
	s := strings.Split(line, ";")
	if len(s) != 3 {
		return nil, fmt.Errorf("Line must have the entries: date,name,amount. Current line: %s", line)
	}

	// TODO: Support different date/time layouts.
	dateLayout := "2006-01-02"
	date, err := time.Parse(dateLayout, s[0])
	if err != nil {
		return nil, err
	}

	name := s[1]

	cost := strings.Replace(s[2], ",", "", -1)

	return &splitwise.Expense{
		Cost:         cost,
		Description:  name,
		Date:         date.String(),
		CurrencyCode: currencyCode,
		GroupId:      groupID,
	}, nil
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.createexpenses.yaml)")
	rootCmd.PersistentFlags().StringVar(&credsFile, "credentials", ".credentials.json",
		"path to file containing the SplitWise authentication key (default is .credentials.json)")
	// TODO Add flags for making it possible to create a single expense with no need of the csv file.
	//      The csv file then should be optional and not be marked as required anymore.
	rootCmd.PersistentFlags().StringVar(&ExpensesFile, "from-file", "", "path to file containing the expenses")
	rootCmd.MarkFlagRequired("from-file")
	rootCmd.PersistentFlags().StringVar(&currencyCode, "currency-code", "USD", "A currency code. Must be in the list from get_currencies")
	rootCmd.PersistentFlags().Uint32Var(&groupID, "group-id", 0, "user group id where the expense will be added to")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".createexpenses" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".createexpenses")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
