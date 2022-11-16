package cli

import (
	"github.com/spf13/cobra"
	rpctypes "github.com/tendermint/tendermint/rpc/core/types"
)

func NewWorkerHelper(cmd *cobra.Command, blockEvent rpctypes.ResultEvent) error {
	return nil
}
