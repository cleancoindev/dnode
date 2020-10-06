package simulator

import (
	"encoding/csv"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

type SimReportCSVWriter struct {
	writer *csv.Writer
}

type CSVWriterClose func()

var Headers = []string{
	"BlockHeight",
	"BlockTime",
	"SimDuration",
	"Validators: Bonded",
	"Validators: Unbonding",
	"Validators: Unbonded",
	"Staking: Bonded",
	"Staking: NotBonded",
	"Staking: LPs",
	"Staking: ActiveRedelegations",
	"Mint: MinInflation",
	"Mint: MaxInflation",
	"Mint: AnnualProvision",
	"Mint: BlocksPerYear",
	"Dist: FoundationPool",
	"Dist: PTreasuryPool",
	"Dist: LiquidityPPool",
	"Dist: HARP",
	"Supply: Total [main]",
	"Supply: Total [staking]",
	"Supply: Total [LP]",
	"Stats: Bonded/TotalSupply [bonding]",
	"Stats: Bonded/TotalSupply [LPs]",
	"Counters: Bonding: Delegations",
	"Counters: Bonding: Redelegations",
	"Counters: Bonding: Undelegations",
	"Counters: LP: Delegations",
	"Counters: LP: Redelegations",
	"Counters: LP: Undelegations",
	"Counters: Rewards",
	"Counters: RewardsCollected",
	"Counters: Commissions",
	"Counters: CommissionsCollected",
	"Counters: LockedRewards",
}

func NewSimReportCSVWriter(t *testing.T, filePath string) (*SimReportCSVWriter, CSVWriterClose) {
	file, err := os.Create(filePath)
	require.Nil(t, err)

	closeFn := func() {
		file.Close()
	}

	writer := csv.NewWriter(file)
	err = writer.Write(Headers)
	require.Nil(t, err)

	return &SimReportCSVWriter{
		writer: writer,
	}, closeFn
}

func (w *SimReportCSVWriter) Write(item SimReportItem) {
	defer w.writer.Flush()

	data := []string{
		strconv.FormatInt(item.BlockHeight, 10),
		item.BlockTime.Format("02.01.2006T15:04:05"),
		FormatDuration(item.SimulationDur),
		strconv.Itoa(item.ValidatorsBonded),
		strconv.Itoa(item.ValidatorsUnbonding),
		strconv.Itoa(item.ValidatorsUnbonded),
		item.StakingBonded.String(),
		item.StakingNotBonded.String(),
		item.StakingLPs.String(),
		strconv.Itoa(item.RedelegationsInProcess),
		item.MintMinInflation.String(),
		item.MintMaxInflation.String(),
		item.MintAnnualProvisions.String(),
		strconv.FormatUint(item.MintBlocksPerYear, 10),
		item.DistFoundationPool.String(),
		item.DistPublicTreasuryPool.String(),
		item.DistLiquidityProvidersPool.String(),
		item.DistHARP.String(),
		item.SupplyTotalMain.String(),
		item.SupplyTotalStaking.String(),
		item.SupplyTotalLP.String(),
		item.StatsBondedRatio.String(),
		item.StatsLPRatio.String(),
		strconv.FormatInt(item.Counters.BDelegations, 10),
		strconv.FormatInt(item.Counters.BRedelegations, 10),
		strconv.FormatInt(item.Counters.BUndelegations, 10),
		strconv.FormatInt(item.Counters.LPDelegations, 10),
		strconv.FormatInt(item.Counters.LPRedelegations, 10),
		strconv.FormatInt(item.Counters.LPUndelegations, 10),
		strconv.FormatInt(item.Counters.Rewards, 10),
		item.Counters.RewardsCollected.String(),
		strconv.FormatInt(item.Counters.Commissions, 10),
		item.Counters.CommissionsCollected.String(),
		strconv.FormatInt(item.Counters.LockedRewards, 10),
	}

	_ = w.writer.Write(data)
}
