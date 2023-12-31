package keeper_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/unigrid-project/cosmos-gridnode/testutil/keeper"
	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/keeper"
	"github.com/unigrid-project/cosmos-gridnode/x/gridnode/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.GridnodeKeeper(t)
	return keeper.NewMsgServerImpl(*k), ctx
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}
