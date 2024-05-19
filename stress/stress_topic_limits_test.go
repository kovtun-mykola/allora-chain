package stress_test

import (
	"context"
	"os"
	"testing"

	chain_test "github.com/allora-network/allora-chain/stress/chain"
	"github.com/stretchr/testify/require"
)

func SetupTopicLimitsTest(t *testing.T) TestMetadata {
	ret := TestMetadata{}
	ret.t = t
	var err error
	ret.ctx = context.Background()
	// userHomeDir, _ := os.UserHomeDir()
	// home := filepath.Join(userHomeDir, ".allorad")
	node, err := chain_test.NewNode(
		t,
		chain_test.NodeConfig{
			NodeRPCAddress: "http://localhost:26657",
			AlloraHomeDir:  "./devnet/genesis",
		},
	)
	require.NoError(t, err)
	ret.n = node
	return ret
}

func TestStressTestTopicLimitsSuite(t *testing.T) {
	if _, isIntegration := os.LookupEnv("STRESS_TEST_TOPIC_LIMIT"); isIntegration == false {
		t.Skip("Skipping Stress Test topic limit unless explicitly enabled")
	}

	t.Log(">>> Setting up connection to local node <<<")
	m := Setup(t)

	t.Log(">>> Test Making Topic Creation Limits <<<")
	CreateTopicLoop(m)
}
