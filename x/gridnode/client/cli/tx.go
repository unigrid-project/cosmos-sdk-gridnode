package cli

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govutils "github.com/cosmos/cosmos-sdk/x/gov/client/utils"

	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/spf13/cobra"
	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	listSeparator              = ","
	FlagMetadata               = "metadata"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		NewCmdDelegate(),
		NewCmdUnDelegate(),
		NewCmdCastVoteFromGridnode(),
	)

	return cmd
}

func NewCmdDelegate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegate [delegator-address] [amount]",
		Short: "Delegate tokens for gridnode",
		Args:  cobra.ExactArgs(2), // Expecting 2 arguments now
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Retrieve the delegator address and amount from args
			delegatorAddress := args[0]
			amountStr := args[1]

			// Convert the amount string to int64
			amount, err := strconv.ParseInt(amountStr, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid amount: %s", amountStr)
			}

			// Create the message
			msg := types.NewMsgDelegateGridnode(delegatorAddress, amount)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdUnDelegate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "undelegate [delegator-address] [amount]",
		Short: "Undelegate tokens from the gridnode module",
		Args:  cobra.ExactArgs(2), // Expecting 2 arguments now
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Retrieve the delegator address and amount from args
			delegatorAddress := args[0]
			amountStr := args[1]

			// Convert the amount string to int64
			amount, err := strconv.ParseInt(amountStr, 10, 64)
			if err != nil {
				return fmt.Errorf("invalid amount: %s", amountStr)
			}

			// Create the message
			msg := types.NewMsgUndelegateGridnode(delegatorAddress, amount)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewCmdCastVoteFromGridnode() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cast-vote [proposal-id] [option]",
		Short: "Cast a vote from a Gridnode, options: yes/no/no_with_veto/abstain",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			// Get voting address
			from := clientCtx.GetFromAddress()
			fmt.Printf("Vote from: %s\n", from.String())
			// Retrieve the public key associated with the address
			pubKey, err := getPublicKeyFromAddress(clientCtx, from)
			if err != nil {
				return fmt.Errorf("failed to get public key for address %s: %v", from, err)
			}
			if pubKey == nil {
				return fmt.Errorf("public key for address %s is nil", from)
			}
			fmt.Printf("Public key: %s\n", pubKey.String())
			// validate that the proposal id is a uint
			proposalID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("proposal-id %s not a valid int, please input a valid proposal-id", args[0])
			}

			metadata, err := cmd.Flags().GetString(FlagMetadata)
			if err != nil {
				return err
			}

			// Find out which vote option user chose
			byteVoteOption, err := v1.VoteOptionFromString(govutils.NormalizeVoteOption(args[1]))
			if err != nil {
				return err
			}

			// Check if it is GridNode
			isNode := types.IsGridnode(from.String())
			if !isNode {
				return errors.New("address is not a GridNode")
			}

			// Build vote message and run basic validation
			msg := v1.NewMsgVote(from, proposalID, byteVoteOption, metadata)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(FlagMetadata, "", "Specify metadata of the vote")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getPublicKeyFromAddress(cliCtx client.Context, addr sdk.AccAddress) (cryptotypes.PubKey, error) {
	// Query the account information using the address
	queryClient := authtypes.NewQueryClient(cliCtx)
	accountResponse, err := queryClient.Account(
		context.Background(),
		&authtypes.QueryAccountRequest{Address: addr.String()},
	)
	if err != nil {
		return nil, err
	}

	// Decode the account information
	var account authtypes.AccountI
	err = cliCtx.InterfaceRegistry.UnpackAny(accountResponse.Account, &account)
	if err != nil {
		return nil, err
	}

	// Retrieve the public key from the account
	return account.GetPubKey(), nil
}
