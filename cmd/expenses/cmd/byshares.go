// Copyright (c) 2022 Alvaro Silva <alvaro.fagner@gmail.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package cmd

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/anvari1313/splitwise.go"
	"github.com/spf13/cobra"
	"github.com/splitwise-batch/pkg/common/entity"
	"github.com/splitwise-batch/pkg/security"
)

var (
	userShares []string
	paidBy     uint64
)

// bysharesCmd represents the byshares command
var bysharesCmd = &cobra.Command{
	Use:   "byshares",
	Short: "Creates an expense by shares",
	RunE:  byshares,
}

func byshares(cmd *cobra.Command, args []string) error {
	shares, err := parseUserShares(userShares)
	if err != nil {
		return err
	}

	auth, err := security.Authenticate(credsFile)
	if err != nil {
		return err
	}

	client := splitwise.NewClient(auth)

	expenses, err := ReadExpenses(ExpensesFile)
	if err != nil {
		return err
	}

	for _, expense := range expenses {
		amounts, err := splitAmount(expense.Cost, shares)
		if err != nil {
			return err
		}

		resp, err := client.CreateExpenseByShare(context.Background(), expense, amounts)
		if err != nil {
			return fmt.Errorf("could not create expense %+v, error=%s", expense, err.Error())
		}

		fmt.Printf("Cost: %s, Amounts: %+v\n", expense.Cost, amounts)
		fmt.Println("Expense created:", resp)
	}

	return nil
}

func splitAmount(cost string, shares []entity.Share) ([]splitwise.UserShare, error) {
	costFloat, err := strconv.ParseFloat(cost, 64)
	if err != nil {
		return nil, err
	}

	var partsTotal float64
	for _, share := range shares {
		partsTotal += share.Share
	}

	var result []splitwise.UserShare
	for _, share := range shares {
		var paid float64
		if share.UserID == paidBy {
			paid = costFloat
		} else {
			paid = 0
		}

		owed := costFloat * share.Share / partsTotal
		result = append(result, splitwise.UserShare{
			UserID:    share.UserID,
			PaidShare: fmt.Sprintf("%.2f", paid),
			OwedShare: fmt.Sprintf("%.2f", owed),
		})

	}

	return result, nil
}

func parseUserShares(userShares []string) ([]entity.Share, error) {
	var result []entity.Share
	for _, share := range userShares {
		s := strings.Split(share, ":")
		if len(s) != 2 {
			return nil, fmt.Errorf("can't parse user share %s", share)
		}

		userID, err := strconv.Atoi(s[0])
		if err != nil {
			return nil, fmt.Errorf("can't parse userID from %s", s[0])
		}

		share, err := strconv.ParseFloat(s[1], 64)
		if err != nil {
			return nil, fmt.Errorf("can't parse paid share from %s", s[1])
		}

		result = append(result, entity.Share{
			UserID: uint64(userID),
			Share:  share,
		})
	}

	return result, nil
}

func init() {
	rootCmd.AddCommand(bysharesCmd)

	bysharesCmd.PersistentFlags().StringSliceVar(
		&userShares,
		"user-share",
		[]string{},
		"the pair user and respective share in the format userID:share. Can be passed many times. Requires at least two.",
	)
	bysharesCmd.PersistentFlags().Uint64Var(&paidBy, "paid-by", 0, "user who paid for the expense")
}
