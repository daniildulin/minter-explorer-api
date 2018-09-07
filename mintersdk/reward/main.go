package reward

import (
	"github.com/daniildulin/explorer-api/mintersdk/converter"
	"math/big"
)

// Total blocks for reward
const totalBlocksCount = 44512766

// Total blocks for reward with extra 1ns
const totalBlocksCountWithPlus = 44512784

// Max reward
const maxReward = 111

type MinterReward struct {
}

type MinterRewardError struct {
	message string
}

func (e *MinterRewardError) Error() string {
	return e.message
}

// Get reward by the block number in PIP
func Get(blockNumber uint) (*big.Int, error) {

	if blockNumber <= 0 {
		return big.NewInt(0), &MinterRewardError{message: "Block number should be greater than 0"}
	}

	if blockNumber > totalBlocksCountWithPlus {
		return converter.ConvertValue(big.NewInt(0), "pip"), nil
	}

	if blockNumber > totalBlocksCount {
		return converter.ConvertValue(big.NewInt(1), "pip"), nil
	}

	reward := formula(blockNumber)

	if reward > maxReward {
		return converter.ConvertValue(big.NewInt(maxReward), "pip"), nil
	} else {
		return converter.ConvertValue(big.NewInt(reward), "pip"), nil
	}
}

// Calculate reward by formula
func formula(blockNumber uint) int64 {

	reward := (maxReward*(totalBlocksCount-blockNumber))/totalBlocksCount + 1

	if blockNumber <= totalBlocksCount*50/100 {
		reward = reward * 15 / 10
	}

	return int64(reward)
}
