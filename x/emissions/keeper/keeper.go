package keeper

import (
	"context"
	"errors"
	"math/big"
	"strings"

	cosmosMath "cosmossdk.io/math"

	"cosmossdk.io/collections"
	"cosmossdk.io/core/address"
	storetypes "cosmossdk.io/core/store"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/allora-network/allora-chain/app/params"
	state "github.com/allora-network/allora-chain/x/emissions"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Uint = cosmosMath.Uint
type Int = cosmosMath.Int

type TOPIC_ID = uint64
type ACC_ADDRESS = string
type TARGET = sdk.AccAddress
type TARGET_STR = string
type DELEGATOR = sdk.AccAddress
type DELEGATOR_STR = string
type WORKERS = string
type REPUTERS = string
type BLOCK_NUMBER = int64
type UNIX_TIMESTAMP = uint64
type REQUEST_ID = string

// Emissions rate Constants
// TODO make these not constants and figure out how they should
// be changeable by governance or some algorithm or whatever
const EPOCH_LENGTH = 5
const EMISSIONS_PER_EPOCH = 1000

type Keeper struct {
	cdc          codec.BinaryCodec
	addressCodec address.Codec

	// State management
	schema     collections.Schema
	params     collections.Item[state.Params]
	authKeeper AccountKeeper
	bankKeeper BankKeeper

	// ############################################
	// #             TOPIC STATE:                 #
	// ############################################

	// the next topic id to be used, equal to the number of topics that have been created
	nextTopicId collections.Sequence
	// every topic that has been created indexed by their topicId starting from 1 (0 is reserved for the root network)
	topics collections.Map[TOPIC_ID, state.Topic]
	// for a topic, what is every worker node that has registered to it?
	topicWorkers collections.KeySet[collections.Pair[TOPIC_ID, sdk.AccAddress]]
	// for a topic, what is every reputer node that has registered to it?
	topicReputers collections.KeySet[collections.Pair[TOPIC_ID, sdk.AccAddress]]
	// for an address, what are all the topics that it's registered for?
	addressTopics collections.Map[sdk.AccAddress, []uint64]

	// ############################################
	// #                STAKING                   #
	// ############################################

	// total sum stake of all topics
	allTopicStakeSum collections.Item[Uint]
	// total sum stake of all stakers on the network
	totalStake collections.Item[Uint]
	// for every topic, how much total stake does that topic have accumulated?
	topicStake collections.Map[TOPIC_ID, Uint]
	// For staking information, we have 3 views of the data we need to record:
	// 1. How many tokens did each delegator (e.g. alice) stake?
	//    How many tokens did alice put into the staking system?
	// 2. For each target (e.g. bob), then for each delegator (e.g. alice),
	//    how many tokens did that delegator stake upon that target? (e.g. alice -> bob)
	// 3. For each target, how much accumulated stake from everybody do they have?
	//    How many tokens have been staked upon bob total?
	// Explanation by Example:
	// Let there be Alice and Bob, existing as nodes in topic0
	//
	// At time t_0, Alice stakes 4 tokens upon herself
	// stakeOwnedByDelegator(Alice) -> +4
	// stakePlacement(Alice, Alice) -> +4
	// stakePlacedUponTarget(Alice) -> +4
	// TopicStake(topic0) -> +4
	// TotalStake across everybody: +4
	//
	// Later, at time t_1, Alice stakes 3 tokens upon Bob
	// stakeOwnedByDelegator(Alice) -> +3
	// stakePlacement(Alice, Bob) -> +3
	// stakePlacedUponTarget(Bob) -> +3
	// TopicStake(topic0) -> +3
	// TotalStake across everybody: +3
	//
	// At time t_2, Alice withdraws Alice's stake in herself and Bob
	// stakeOwnedByDelegator(Alice) -> -4 for Alice, -3 for Bob = -7 total
	// stakePlacement(Alice,Alice) -> -4
	// stakePlacement(Alice, Bob) -> -3
	// stakePlacedUponTarget(Alice) -> -4
	// stakePlacedUponTarget(Bob) -> -3
	// TopicStake(topic0) -> -7
	// TotalStake across everybody: -7
	//
	// TODO: enforcement and unit testing of various invariants from this data structure
	// Invariant 1: stakeOwnedByDelegator = sum of all stakes placed by that delegator
	// Invariant 2: targetStake = sum of all bonds placed upon that target
	// Invariant 3: sum of all targetStake = topicStake for that topic
	//
	// map of (delegator) -> total amount staked by delegator on everyone they've staked upon (including potentially themselves)
	stakeOwnedByDelegator collections.Map[DELEGATOR, Uint]
	// map of (delegator, target) -> amount staked by delegator upon target
	stakePlacement collections.Map[collections.Pair[DELEGATOR, TARGET], Uint]
	// map of (target) -> total amount staked upon target by themselves and all other delegators
	stakePlacedUponTarget collections.Map[TARGET, Uint]
	// map of (delegator) -> removal information for that delegator
	stakeRemovalQueue collections.Map[DELEGATOR, state.StakeRemoval]

	// ############################################
	// #        INFERENCE REQUEST MEMPOOL         #
	// ############################################
	mempool collections.Map[collections.Pair[TOPIC_ID, REQUEST_ID], state.InferenceRequest]
	funds   collections.Map[REQUEST_ID, Uint]

	// ############################################
	// #            MISC GLOBAL STATE:            #
	// ############################################

	// map of Reputer, worker -> weight judged by reputer upon worker node
	weights collections.Map[collections.Triple[TOPIC_ID, sdk.AccAddress, sdk.AccAddress], Uint]

	// map of (topic, worker) -> inference
	inferences collections.Map[collections.Pair[TOPIC_ID, sdk.AccAddress], state.Inference]

	// map of worker id to node data about that worker
	workers collections.Map[sdk.AccAddress, state.OffchainNode]

	// map of reputer id to node data about that reputer
	reputers collections.Map[sdk.AccAddress, state.OffchainNode]

	// the last block the token inflation rewards were updated: int64 same as BlockHeight()
	lastRewardsUpdate collections.Item[BLOCK_NUMBER]

	// map of (topic, timestamp, index) -> Inference
	allInferences collections.Map[collections.Pair[TOPIC_ID, UNIX_TIMESTAMP], state.Inferences]
}

func NewKeeper(
	cdc codec.BinaryCodec,
	addressCodec address.Codec,
	storeService storetypes.KVStoreService,
	ak AccountKeeper,
	bk BankKeeper) Keeper {

	sb := collections.NewSchemaBuilder(storeService)
	k := Keeper{
		cdc:                   cdc,
		addressCodec:          addressCodec,
		params:                collections.NewItem(sb, state.ParamsKey, "params", codec.CollValue[state.Params](cdc)),
		authKeeper:            ak,
		bankKeeper:            bk,
		totalStake:            collections.NewItem(sb, state.TotalStakeKey, "total_stake", UintValue),
		topicStake:            collections.NewMap(sb, state.TopicStakeKey, "topic_stake", collections.Uint64Key, UintValue),
		lastRewardsUpdate:     collections.NewItem(sb, state.LastRewardsUpdateKey, "last_rewards_update", collections.Int64Value),
		nextTopicId:           collections.NewSequence(sb, state.NextTopicIdKey, "next_topic_id"),
		topics:                collections.NewMap(sb, state.TopicsKey, "topics", collections.Uint64Key, codec.CollValue[state.Topic](cdc)),
		topicWorkers:          collections.NewKeySet(sb, state.TopicWorkersKey, "topic_workers", collections.PairKeyCodec(collections.Uint64Key, sdk.AccAddressKey)),
		addressTopics:         collections.NewMap(sb, state.AddressTopicsKey, "address_topics", sdk.AccAddressKey, TopicIdListValue),
		topicReputers:         collections.NewKeySet(sb, state.TopicReputersKey, "topic_reputers", collections.PairKeyCodec(collections.Uint64Key, sdk.AccAddressKey)),
		allTopicStakeSum:      collections.NewItem(sb, state.AllTopicStakeSumKey, "all_topic_stake_sum", UintValue),
		stakeOwnedByDelegator: collections.NewMap(sb, state.DelegatorStakeKey, "delegator_stake", sdk.AccAddressKey, UintValue),
		stakePlacement:        collections.NewMap(sb, state.BondsKey, "bonds", collections.PairKeyCodec(sdk.AccAddressKey, sdk.AccAddressKey), UintValue),
		stakePlacedUponTarget: collections.NewMap(sb, state.TargetStakeKey, "target_stake", sdk.AccAddressKey, UintValue),
		stakeRemovalQueue:     collections.NewMap(sb, state.StakeRemovalQueueKey, "stake_removal_queue", sdk.AccAddressKey, codec.CollValue[state.StakeRemoval](cdc)),
		mempool:               collections.NewMap(sb, state.MempoolKey, "mempool", collections.PairKeyCodec(collections.Uint64Key, collections.StringKey), codec.CollValue[state.InferenceRequest](cdc)),
		funds:                 collections.NewMap(sb, state.FundsKey, "funds", collections.StringKey, UintValue),
		weights:               collections.NewMap(sb, state.WeightsKey, "weights", collections.TripleKeyCodec(collections.Uint64Key, sdk.AccAddressKey, sdk.AccAddressKey), UintValue),
		inferences:            collections.NewMap(sb, state.InferencesKey, "inferences", collections.PairKeyCodec(collections.Uint64Key, sdk.AccAddressKey), codec.CollValue[state.Inference](cdc)),
		workers:               collections.NewMap(sb, state.WorkerNodesKey, "worker_nodes", sdk.AccAddressKey, codec.CollValue[state.OffchainNode](cdc)),
		reputers:              collections.NewMap(sb, state.ReputerNodesKey, "reputer_nodes", sdk.AccAddressKey, codec.CollValue[state.OffchainNode](cdc)),
		allInferences:         collections.NewMap(sb, state.AllInferencesKey, "inferences_all", collections.PairKeyCodec(collections.Uint64Key, collections.Uint64Key), codec.CollValue[state.Inferences](cdc)),
	}

	schema, err := sb.Build()
	if err != nil {
		panic(err)
	}

	k.schema = schema

	return k
}

func (k *Keeper) SetParams(ctx context.Context, params state.Params) error {
	return k.params.Set(ctx, params)
}

func (k *Keeper) GetTopicWeightLastRan(ctx context.Context, topicId TOPIC_ID) (uint64, error) {
	topic, err := k.topics.Get(ctx, topicId)
	if err != nil {
		return 0, err
	}
	ret := topic.WeightLastRan
	return ret, nil
}

func (k *Keeper) GetAllInferences(ctx context.Context, topicId TOPIC_ID, timestamp uint64) (*state.Inferences, error) {
	// pair := collections.Join(topicId, timestamp, index)
	key := collections.Join(topicId, timestamp)
	inferences, err := k.allInferences.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	return &inferences, nil
}

func appendToInferences(inferencesSet *state.Inferences, newInference *state.Inference) state.Inferences {
	if inferencesSet == nil || len(inferencesSet.Inferences) == 0 {
		// If Inferences is nil or empty, create a new one with the new Inference
		return state.Inferences{Inferences: []*state.Inference{newInference}}
	}
	// If Inferences is not empty, append the new Inference to the existing ones
	return state.Inferences{Inferences: append(inferencesSet.Inferences, newInference)}
}

func (k *Keeper) InsertInference(ctx context.Context, topicId TOPIC_ID, timestamp uint64, inference state.Inference) error {
	key := collections.Join(topicId, timestamp)
	inferences_set, err := k.allInferences.Get(ctx, key)
	if err != nil {
		inferences_set = state.Inferences{
			Inferences: []*state.Inference{},
		}
	}
	// inferences_new_set := append(inferences_set.Inferences, &inference)
	inferences_new_set := appendToInferences(&inferences_set, &inference)
	return k.allInferences.Set(ctx, key, inferences_new_set)
}

// Insert a complete set of inferences for a topic/timestamp. Overwrites previous ones.
func (k *Keeper) InsertInferences(ctx context.Context, topicId TOPIC_ID, timestamp uint64, inferences state.Inferences) error {
	key := collections.Join(topicId, timestamp)
	err := k.allInferences.Set(ctx, key, inferences)

	return err
}

func (k *Keeper) GetStakePlacedUponTarget(ctx context.Context, target sdk.AccAddress) (Uint, error) {
	ret, err := k.stakePlacedUponTarget.Get(ctx, target)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return cosmosMath.NewUint(0), nil
		}
		return cosmosMath.Uint{}, err
	}
	return ret, nil
}

func (k *Keeper) SetStakePlacedUponTarget(ctx context.Context, target sdk.AccAddress, stake Uint) error {
	if stake.IsZero() {
		return k.stakePlacedUponTarget.Remove(ctx, target)
	}
	return k.stakePlacedUponTarget.Set(ctx, target, stake)
}

// Returns the last block height at which rewards emissions were updated
func (k *Keeper) GetLastRewardsUpdate(ctx context.Context) (int64, error) {
	lastRewardsUpdate, err := k.lastRewardsUpdate.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return 0, nil
		} else {
			return 0, err
		}
	}
	return lastRewardsUpdate, nil
}

// Set the last block height at which rewards emissions were updated
func (k *Keeper) SetLastRewardsUpdate(ctx context.Context, blockHeight int64) error {
	if blockHeight < 0 {
		return state.ErrBlockHeightNegative
	}
	previousBlockHeight, err := k.lastRewardsUpdate.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			previousBlockHeight = 0
		} else {
			return err
		}
	}
	if blockHeight < previousBlockHeight {
		return state.ErrBlockHeightLessThanPrevious
	}
	return k.lastRewardsUpdate.Set(ctx, blockHeight)
}

// return epoch length
func (k *Keeper) EpochLength() int64 {
	return EPOCH_LENGTH
}

// return how many new coins should be minted for the next emission
func (k *Keeper) CalculateAccumulatedEmissions(ctx context.Context) (cosmosMath.Int, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	blockNumber := sdkCtx.BlockHeight()
	lastRewardsUpdate, err := k.GetLastRewardsUpdate(sdkCtx)
	if err != nil {
		return cosmosMath.Int{}, err
	}
	blocksSinceLastUpdate := blockNumber - lastRewardsUpdate
	// number of epochs that have passed (if more than 1)
	epochsPassed := cosmosMath.NewInt(blocksSinceLastUpdate / EPOCH_LENGTH)
	// get emission amount
	return epochsPassed.Mul(cosmosMath.NewInt(EMISSIONS_PER_EPOCH)), nil
}

// mint new rewards coins to this module account
func (k *Keeper) MintRewardsCoins(ctx context.Context, amount cosmosMath.Int) error {
	coins := sdk.NewCoins(sdk.NewCoin(params.DefaultBondDenom, amount))
	return k.bankKeeper.MintCoins(ctx, state.AlloraStakingModuleName, coins)
}

// for a given topic, returns every reputer node registered to it and their normalized stake
func (k *Keeper) GetReputerNormalizedStake(
	ctx sdk.Context,
	topicId TOPIC_ID,
	topicStake *big.Float) (reputerNormalizedStakeMap map[ACC_ADDRESS]*big.Float, retErr error) {
	reputerNormalizedStakeMap = make(map[ACC_ADDRESS]*big.Float)
	rng := collections.NewPrefixedPairRange[TOPIC_ID, sdk.AccAddress](topicId)
	retErr = nil
	retErr = k.topicReputers.Walk(ctx, rng, func(key collections.Pair[TOPIC_ID, sdk.AccAddress]) (stop bool, err error) {
		reputer := key.K2()
		// Get Stake in each reputer
		reputerTargetStake, err := k.stakePlacedUponTarget.Get(ctx, reputer)
		if err != nil {
			return true, err
		}
		reputerTotalStake := big.NewFloat(0).SetInt(reputerTargetStake.BigInt())

		// How much stake does each reputer have as a percentage of the total stake in the topic?
		reputerNormalizedStake := big.NewFloat(0).Quo(reputerTotalStake, topicStake)
		reputerNormalizedStakeMap[reputer.String()] = reputerNormalizedStake
		return false, nil
	})
	return reputerNormalizedStakeMap, retErr
}

func (k *Keeper) GetWorker(ctx context.Context, worker sdk.AccAddress) (state.OffchainNode, error) {
	return k.workers.Get(ctx, worker)
}

func (k *Keeper) GetReputer(ctx context.Context, reputer sdk.AccAddress) (state.OffchainNode, error) {
	return k.reputers.Get(ctx, reputer)
}

// Gets the total sum of all stake in the network across all topics
func (k *Keeper) GetTotalStake(ctx context.Context) (Uint, error) {
	ret, err := k.totalStake.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return cosmosMath.NewUint(0), nil
		}
		return cosmosMath.Uint{}, err
	}
	return ret, nil
}

// Sets the total sum of all stake in the network across all topics
func (k *Keeper) SetTotalStake(ctx context.Context, totalStake Uint) error {
	// total stake does not have a zero guard because totalStake is allowed to be zero
	// it is initialized to zero at genesis anyways.
	return k.totalStake.Set(ctx, totalStake)
}

// A function that accepts a topicId and returns list of Inferences or error
func (k *Keeper) GetLatestInferencesFromTopic(ctx context.Context, topicId TOPIC_ID) ([]*state.InferenceSetForScoring, error) {
	var inferences []*state.InferenceSetForScoring
	var latestTimestamp, err = k.GetTopicWeightLastRan(ctx, topicId)
	if err != nil {
		latestTimestamp = 0
	}
	rng := collections.
		NewPrefixedPairRange[TOPIC_ID, UNIX_TIMESTAMP](topicId).
		StartInclusive(latestTimestamp).
		Descending()

	iter, err := k.allInferences.Iterate(ctx, rng)
	if err != nil {
		return nil, err
	}
	for ; iter.Valid(); iter.Next() {
		kv, err := iter.KeyValue()
		if err != nil {
			return nil, err
		}
		key := kv.Key
		value := kv.Value
		inferenceSet := &state.InferenceSetForScoring{
			TopicId:    key.K1(),
			Timestamp:  key.K2(),
			Inferences: &value,
		}
		inferences = append(inferences, inferenceSet)
	}
	return inferences, nil
}

// Gets the stake in the network for a given topic
func (k *Keeper) GetTopicStake(ctx context.Context, topicId TOPIC_ID) (Uint, error) {
	ret, err := k.topicStake.Get(ctx, topicId)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return cosmosMath.NewUint(0), nil
		}
		return cosmosMath.Uint{}, err
	}
	return ret, nil
}

// Sets the stake in the network for a given topic
func (k *Keeper) SetTopicStake(ctx context.Context, topicId TOPIC_ID, stake Uint) error {
	if stake.IsZero() {
		return k.topicStake.Remove(ctx, topicId)
	}
	return k.topicStake.Set(ctx, topicId, stake)
}

// GetTopicsByCreator returns a slice of all topics created by a given creator.
func (k *Keeper) GetTopicsByCreator(ctx context.Context, creator string) ([]*state.Topic, error) {
	var topicsByCreator []*state.Topic

	err := k.topics.Walk(ctx, nil, func(id TOPIC_ID, topic state.Topic) (bool, error) {
		if topic.Creator == creator {
			topicsByCreator = append(topicsByCreator, &topic)
		}
		return false, nil // Continue iterating
	})

	if err != nil {
		return nil, err
	}

	return topicsByCreator, nil
}

// AddAddressTopics adds new topics to the address's list of topics, avoiding duplicates.
func (k *Keeper) AddAddressTopics(ctx context.Context, address sdk.AccAddress, newTopics []uint64) error {
	// Get the current list of topics for the address
	currentTopics, err := k.GetRegisteredTopicsIdsByAddress(ctx, address)
	if err != nil {
		return err
	}

	topicSet := make(map[uint64]bool)
	for _, topic := range currentTopics {
		topicSet[topic] = true
	}

	for _, newTopic := range newTopics {
		if _, exists := topicSet[newTopic]; !exists {
			currentTopics = append(currentTopics, newTopic)
		}
	}

	// Set the updated list of topics for the address
	return k.addressTopics.Set(ctx, address, currentTopics)
}

// GetRegisteredTopicsByAddress returns a slice of all topics ids registered by a given address.
func (k *Keeper) GetRegisteredTopicsIdsByAddress(ctx context.Context, address sdk.AccAddress) ([]uint64, error) {
	topics, err := k.addressTopics.Get(ctx, address)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			// Return an empty slice if the address is not found, or handle it differently if needed.
			return []uint64{}, nil
		}
		return nil, err
	}
	return topics, nil
}

// GetRegisteredTopicsIdsByWorkerAddress returns a slice of all topics ids registered by a given worker address.
func (k *Keeper) GetRegisteredTopicsIdsByWorkerAddress(ctx context.Context, address sdk.AccAddress) ([]uint64, error) {
	var topicsByAddress []uint64

	err := k.topicWorkers.Walk(ctx, nil, func(pair collections.Pair[TOPIC_ID, sdk.AccAddress]) (bool, error) {
		if pair.K2().String() == address.String() {
			topicsByAddress = append(topicsByAddress, pair.K1())
		}
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	return topicsByAddress, nil
}

// GetRegisteredTopicsIdsByReputerAddress returns a slice of all topics ids registered by a given reputer address.
func (k *Keeper) GetRegisteredTopicsIdsByReputerAddress(ctx context.Context, address sdk.AccAddress) ([]uint64, error) {
	var topicsByAddress []uint64

	err := k.topicReputers.Walk(ctx, nil, func(pair collections.Pair[TOPIC_ID, sdk.AccAddress]) (bool, error) {
		if pair.K2().String() == address.String() {
			topicsByAddress = append(topicsByAddress, pair.K1())
		}
		return false, nil
	})
	if err != nil {
		return nil, err
	}

	return topicsByAddress, nil
}

func (k *Keeper) IterateAllTopicStake(ctx context.Context) (collections.Iterator[uint64, cosmosMath.Uint], error) {
	rng := collections.Range[uint64]{}
	rng.StartInclusive(0)
	end, err := k.nextTopicId.Peek(ctx)
	if err != nil {
		return collections.Iterator[uint64, cosmosMath.Uint]{}, err
	}
	rng.EndExclusive(end)
	return k.topicStake.Iterate(ctx, &rng)
}

// Runs an arbitrary function for every topic in the network
func (k *Keeper) WalkAllTopicStake(ctx context.Context, walkFunc func(topicId TOPIC_ID, stake Uint) (stop bool, err error)) error {
	rng := collections.Range[uint64]{}
	rng.StartInclusive(0)
	end, err := k.nextTopicId.Peek(ctx)
	if err != nil {
		return err
	}
	rng.EndExclusive(end)
	err = k.topicStake.Walk(ctx, &rng, walkFunc)
	return err
}

// GetStakesForAccount returns the list of stakes for a given account address.
func (k *Keeper) GetStakesForAccount(ctx context.Context, delegator sdk.AccAddress) ([]*state.StakeInfo, error) {
	targets, amounts, err := k.GetAllBondsForDelegator(ctx, delegator)
	if err != nil {
		return nil, err
	}

	stakeInfos := make([]*state.StakeInfo, len(targets))
	for i, target := range targets {
		stakeInfos[i] = &state.StakeInfo{
			Address: target.String(),
			Amount:  amounts[i].String(),
		}
	}

	return stakeInfos, nil
}

// Gets next topic id
func (k *Keeper) IncrementTopicId(ctx context.Context) (TOPIC_ID, error) {
	return k.nextTopicId.Next(ctx)
}

// Sets a topic config on a topicId
func (k *Keeper) SetTopic(ctx context.Context, topicId TOPIC_ID, topic state.Topic) error {
	return k.topics.Set(ctx, topicId, topic)
}

// Checks if a topic exists
func (k *Keeper) TopicExists(ctx context.Context, topicId TOPIC_ID) (bool, error) {
	return k.topics.Has(ctx, topicId)
}

// Returns the number of topics that are active in the network
func (k *Keeper) GetNumTopics(ctx context.Context) (TOPIC_ID, error) {
	return k.nextTopicId.Peek(ctx)
}

// GetActiveTopics returns a slice of all active topics.
func (k *Keeper) GetActiveTopics(ctx context.Context) ([]*state.Topic, error) {
	var activeTopics []*state.Topic
	if err := k.topics.Walk(ctx, nil, func(topicId TOPIC_ID, topic state.Topic) (bool, error) {
		if topic.Active { // Check if the topic is marked as active
			activeTopics = append(activeTopics, &topic)
		}
		return false, nil // Continue the iteration
	}); err != nil {
		return nil, err
	}
	return activeTopics, nil
}

// Add stake adds stake to the system for a given delegator and target
// it adds to existing holdings.
// it places the stake upon target, from delegator, in amount.
// it also updates the total stake for the subnet in question and the total global stake.
// see comments in keeper.go data structures for examples of how the data structure tracking works
func (k *Keeper) AddStake(ctx context.Context, topicsIds []TOPIC_ID, delegator string, target string, stake Uint) error {

	// if stake is zero this function is a no-op
	if stake.IsZero() {
		return state.ErrDoNotSetMapValueToZero
	}

	// update the stake array that tracks how much each delegator has invested in the system total
	delegatorAcc, err := sdk.AccAddressFromBech32(delegator)
	if err != nil {
		return err
	}
	delegatorStake, err := k.stakeOwnedByDelegator.Get(ctx, delegatorAcc)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			delegatorStake = cosmosMath.NewUint(0)
		} else {
			return err
		}
	}

	delegatorStakeNew := delegatorStake.Add(stake)
	if err := k.stakeOwnedByDelegator.Set(ctx, delegatorAcc, delegatorStakeNew); err != nil {
		return err
	}

	// Update the bonds amount, which tracks each individual place
	// each delegator has placed their stake.
	// the sum of all bonds for a delegator should equal the delegatorStake
	// and the sum of all bonds on a target should equal the targetStake
	// set Bond(delegator -> target) = Bond(delegator -> target) + stake
	targetAcc, err := sdk.AccAddressFromBech32(target)
	if err != nil {
		return err
	}
	bondIndex := collections.Join(delegatorAcc, targetAcc)
	bond, err := k.stakePlacement.Get(ctx, bondIndex)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			bond = cosmosMath.NewUint(0)
		} else {
			return err
		}
	}
	bondNew := bond.Add(stake)
	if err := k.stakePlacement.Set(ctx, bondIndex, bondNew); err != nil {
		return err
	}

	// set the targetStake for this target target
	// this is the sum total of all bonds placed upon this target
	// from all different people who have placed stake upon this target
	targetStake, err := k.stakePlacedUponTarget.Get(ctx, targetAcc)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			targetStake = cosmosMath.NewUint(0)
		} else {
			return err
		}
	}
	targetStakeNew := targetStake.Add(stake)
	if err := k.stakePlacedUponTarget.Set(ctx, targetAcc, targetStakeNew); err != nil {
		return err
	}

	// Update the sum topic stake for all topics
	for _, topicId := range topicsIds {
		topicStake, err := k.topicStake.Get(ctx, topicId)
		if err != nil {
			if errors.Is(err, collections.ErrNotFound) {
				topicStake = cosmosMath.NewUint(0)
			} else {
				return err
			}
		}
		topicStakeNew := topicStake.Add(stake)
		if err := k.topicStake.Set(ctx, topicId, topicStakeNew); err != nil {
			return err
		}

		// Update the total stake across all topics
		allTopicStakeSum, err := k.allTopicStakeSum.Get(ctx)
		if err != nil {
			if errors.Is(err, collections.ErrNotFound) {
				allTopicStakeSum = cosmosMath.NewUint(0)
			} else {
				return err
			}
		}
		allTopicStakeSumNew := allTopicStakeSum.Add(stake)
		if err := k.allTopicStakeSum.Set(ctx, allTopicStakeSumNew); err != nil {
			return err
		}
	}

	// Update the total stake across the entire system
	totalStake, err := k.totalStake.Get(ctx)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			totalStake = cosmosMath.NewUint(0)
		} else {
			return err
		}
	}
	totalStakeNew := totalStake.Add(stake)
	if err := k.totalStake.Set(ctx, totalStakeNew); err != nil {
		return err
	}

	return nil
}

// Remove stake from bond updates the various data structures associated
// with removing stake from the system for a given delegator and target
// it removes the stake upon target, from delegator, in amount.
// it also updates the total stake for the topic in question and the total global stake.
// see comments in keeper.go data structures for examples of how the data structure tracking works
func (k *Keeper) RemoveStakeFromBond(
	ctx context.Context,
	topicsIds []TOPIC_ID,
	delegator sdk.AccAddress,
	target sdk.AccAddress,
	stake Uint) error {

	if stake.IsZero() {
		return errors.New("stake must be greater than zero")
	}

	// 1. 2. and 3. make checks and update state for
	// delegatorStake, bonds, and targetStake
	err := k.RemoveStakeFromBondMissingTotalOrTopicStake(ctx, delegator, target, stake)
	if err != nil {
		return err
	}

	// Perform State Updates
	// TODO: make this function prevent partial state updates / do rollbacks if any of the set statements fail
	// not necessary as long as callers are responsible, but it would be nice to have

	// topicStake(topic) = topicStake(topic) - stake
	for _, topic := range topicsIds {
		topicStake, err := k.topicStake.Get(ctx, topic)
		if err != nil {
			return err
		}
		if stake.GT(topicStake) {
			return state.ErrIntegerUnderflowTopicStake
		}

		topicStakeNew := topicStake.Sub(stake)
		if topicStakeNew.IsZero() {
			err = k.topicStake.Remove(ctx, topic)
		} else {
			err = k.topicStake.Set(ctx, topic, topicStakeNew)
		}
		if err != nil {
			return err
		}
	}

	// totalStake = totalStake - stake
	// 4. Check: totalStake >= stake
	totalStake, err := k.totalStake.Get(ctx)
	if err != nil {
		return err
	}
	if stake.GT(totalStake) {
		return state.ErrIntegerUnderflowTotalStake
	}

	// we do write zero here, because totalStake is allowed to be zero
	err = k.totalStake.Set(ctx, totalStake.Sub(stake))
	if err != nil {
		return err
	}

	return nil
}

// Remove stake from bond updates the various data structures associated
// with removing stake from the system for a given delegator and target
// it removes the stake upon target, from delegator, in amount.
// it *DOES NOT* update the total stake for the subnet in question and the total global stake.
// this is used by RemoveAllStake to avoid double counting the topic/total stake removal
func (k *Keeper) RemoveStakeFromBondMissingTotalOrTopicStake(
	ctx context.Context,
	delegator sdk.AccAddress,
	target sdk.AccAddress,
	stake Uint) error {
	// 1. Check: delegatorStake(delegator) >= stake
	delegatorStake, err := k.stakeOwnedByDelegator.Get(ctx, delegator)
	if err != nil {
		return err
	}
	if stake.GT(delegatorStake) {
		return state.ErrIntegerUnderflowDelegator
	}

	// 2. Check: bonds(target, delegator) >= stake
	bond, err := k.stakePlacement.Get(ctx, collections.Join(delegator, target))
	if err != nil {
		return err
	}
	if stake.GT(bond) {
		return state.ErrIntegerUnderflowBonds
	}

	// 3. Check: targetStake(target) >= stake
	targetStake, err := k.stakePlacedUponTarget.Get(ctx, target)
	if err != nil {
		return err
	}
	if stake.GT(targetStake) {
		return state.ErrIntegerUnderflowTarget
	}

	// Perform State Updates
	// TODO: make this function prevent partial state updates / do rollbacks if any of the set statements fail
	// not necessary as long as callers are responsible, but it would be nice to have

	// delegatorStake(delegator) = delegatorStake(delegator) - stake
	delegatorStakeNew := delegatorStake.Sub(stake)
	if delegatorStakeNew.IsZero() {
		err = k.stakeOwnedByDelegator.Remove(ctx, delegator)
	} else {
		err = k.stakeOwnedByDelegator.Set(ctx, delegator, delegatorStakeNew)
	}
	if err != nil {
		return err
	}
	// bonds(target, delegator) = bonds(target, delegator) - stake
	bondNew := bond.Sub(stake)
	if bondNew.IsZero() {
		err = k.stakePlacement.Remove(ctx, collections.Join(delegator, target))
	} else {
		err = k.stakePlacement.Set(ctx, collections.Join(delegator, target), bondNew)
	}
	if err != nil {
		return err
	}

	// targetStake(target) = targetStake(target) - stake
	targetStakeNew := targetStake.Sub(stake)
	if targetStakeNew.IsZero() {
		err = k.stakePlacedUponTarget.Remove(ctx, target)
	} else {
		err = k.stakePlacedUponTarget.Set(ctx, target, targetStakeNew)
	}
	if err != nil {
		return err
	}

	return nil
}

// Used by Modify functions to change stake placements. This function subtracts from the stakePlacement mapping ONLY
// and does not modify any of the other stake mappings e.g. delegatorStake totalStake or topicStake in a system.
func (k *Keeper) SubStakePlacement(ctx context.Context, delegator sdk.AccAddress, target sdk.AccAddress, amount Uint) error {
	bond, err := k.GetBond(ctx, delegator, target)
	if err != nil {
		return err
	}
	if amount.GT(bond) {
		return state.ErrIntegerUnderflowBonds
	}
	bondNew := bond.Sub(amount)
	return k.stakePlacement.Set(ctx, collections.Join(delegator, target), bondNew)
}

// Used by Modify functions to change stake placements. This function adds to the stakePlacement mapping ONLY
// and does not modify any of the other stake mappings e.g. delegatorStake totalStake or topicStake in a system.
func (k *Keeper) AddStakePlacement(ctx context.Context, delegator sdk.AccAddress, target sdk.AccAddress, amount Uint) error {
	bond, err := k.GetBond(ctx, delegator, target)
	if err != nil {
		return err
	}
	bondNew := bond.Add(amount)
	return k.stakePlacement.Set(ctx, collections.Join(delegator, target), bondNew)
}

// Used by Modify functions to change stake placements. This function subtracts from the stakePlacedUponTarget mapping ONLY
// and does not modify any of the other stake mappings e.g. delegatorStake totalStake or topicStake in a system.
func (k *Keeper) SubStakePlacedUponTarget(ctx context.Context, target sdk.AccAddress, amount Uint) error {
	targetStake, err := k.GetStakePlacedUponTarget(ctx, target)
	if err != nil {
		return err
	}
	if amount.GT(targetStake) {
		return state.ErrIntegerUnderflowTarget
	}
	targetStakeNew := targetStake.Sub(amount)
	return k.stakePlacedUponTarget.Set(ctx, target, targetStakeNew)
}

// Used by Modify functions to change stake placements. This function adds to the stakePlacedUponTarget mapping ONLY
// and does not modify any of the other stake mappings e.g. delegatorStake totalStake or topicStake in a system.
func (k *Keeper) AddStakePlacedUponTarget(ctx context.Context, target sdk.AccAddress, amount Uint) error {
	targetStake, err := k.GetStakePlacedUponTarget(ctx, target)
	if err != nil {
		return err
	}
	targetStakeNew := targetStake.Add(amount)
	return k.stakePlacedUponTarget.Set(ctx, target, targetStakeNew)
}

// Add stake into an array of topics
func (k *Keeper) AddStakeToTopics(ctx context.Context, topicsIds []TOPIC_ID, stake Uint) error {
	if stake.IsZero() {
		return state.ErrDoNotSetMapValueToZero
	}

	// Calculate the total stake to be added across all topics
	totalStakeToAdd := stake.Mul(cosmosMath.NewUint(uint64(len(topicsIds))))

	for _, topicId := range topicsIds {
		topicStake, err := k.topicStake.Get(ctx, topicId)
		if err != nil {
			if errors.Is(err, collections.ErrNotFound) {
				topicStake = cosmosMath.NewUint(0)
			} else {
				return err
			}
		}

		topicStakeNew := topicStake.Add(stake)
		if err := k.topicStake.Set(ctx, topicId, topicStakeNew); err != nil {
			return err
		}
	}

	// Update the allTopicStakeSum
	allTopicStakeSum, err := k.allTopicStakeSum.Get(ctx)
	if err != nil {
		return err
	}

	newAllTopicStakeSum := allTopicStakeSum.Add(totalStakeToAdd)
	if err := k.allTopicStakeSum.Set(ctx, newAllTopicStakeSum); err != nil {
		return err
	}

	return nil
}

// Remove stake from an array of topics
func (k *Keeper) RemoveStakeFromTopics(ctx context.Context, topicsIds []TOPIC_ID, stake Uint) error {
	if stake.IsZero() {
		return state.ErrDoNotSetMapValueToZero
	}

	// Calculate the total stake to be removed across all topics
	totalStakeToRemove := stake.Mul(cosmosMath.NewUint(uint64(len(topicsIds))))

	for _, topicId := range topicsIds {
		topicStake, err := k.topicStake.Get(ctx, topicId)
		if err != nil {
			return err // If there's an error, it's not because the topic doesn't exist but some other reason
		}

		if topicStake.LT(stake) {
			return state.ErrCannotRemoveMoreStakeThanStakedInTopic
		}

		topicStakeNew := topicStake.Sub(stake)
		if err := k.topicStake.Set(ctx, topicId, topicStakeNew); err != nil {
			return err
		}
	}

	// Update the allTopicStakeSum
	allTopicStakeSum, err := k.allTopicStakeSum.Get(ctx)
	if err != nil {
		return err
	}

	newAllTopicStakeSum := allTopicStakeSum.Sub(totalStakeToRemove)
	if err := k.allTopicStakeSum.Set(ctx, newAllTopicStakeSum); err != nil {
		return err
	}

	return nil
}

// for a given address, find out how much stake they've put into the system
func (k *Keeper) GetDelegatorStake(ctx context.Context, delegator sdk.AccAddress) (Uint, error) {
	ret, err := k.stakeOwnedByDelegator.Get(ctx, delegator)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return cosmosMath.NewUint(0), nil
		}
		return cosmosMath.Uint{}, err
	}
	return ret, nil
}

// For a given delegator and target, find out how much stake the delegator has placed upon the target
func (k *Keeper) GetBond(ctx context.Context, delegator sdk.AccAddress, target sdk.AccAddress) (Uint, error) {
	ret, err := k.stakePlacement.Get(ctx, collections.Join(delegator, target))
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return cosmosMath.NewUint(0), nil
		}
		return cosmosMath.Uint{}, err
	}
	return ret, nil
}

// For a given delegator, return a map of every target they've placed stake upon, and how much stake they've placed upon them
// O(n) over the number of targets registered.
// since maps of byte array types aren't supported in golang, we instead return two equal length arrays
// where the first array is the targets, and the second array is the amount of stake placed upon them
// indexes in the two arrays correspond to each other
// invariant that len(targets) == len(stakes)
func (k *Keeper) GetAllBondsForDelegator(ctx context.Context, delegator sdk.AccAddress) ([]sdk.AccAddress, []Uint, error) {
	targets := make([]sdk.AccAddress, 0)
	amounts := make([]Uint, 0)
	iter, err := k.stakePlacement.Iterate(ctx, nil)
	if err != nil {
		return nil, nil, err
	}

	// iterate over all keys in stakePlacements
	kvs, err := iter.Keys()
	if err != nil {
		return nil, nil, err
	}
	for _, kv := range kvs {
		d := kv.K1()
		// if the delegator key matches the delegator we're looking for
		if d.Equals(delegator) {
			target := kv.K2()
			amount, err := k.stakePlacement.Get(ctx, kv)
			if err != nil {
				return nil, nil, err
			}
			targets = append(targets, target)
			amounts = append(amounts, amount)
		}
	}
	if len(targets) != len(amounts) {
		return nil, nil, state.ErrIterationLengthDoesNotMatch
	}

	return targets, amounts, nil
}

// For a given topic return the matrix (double map) of all (reputers, workers) -> weight of reputer upon worker
func (k *Keeper) GetWeightsFromTopic(ctx context.Context, topicId TOPIC_ID) (map[REPUTERS]map[WORKERS]*Uint, error) {
	weights := make(map[ACC_ADDRESS]map[ACC_ADDRESS]*Uint)
	rng := collections.NewPrefixedTripleRange[TOPIC_ID, sdk.AccAddress, sdk.AccAddress](topicId)
	iter, err := k.weights.Iterate(ctx, rng)
	if err != nil {
		return nil, err
	}

	kvs, err := iter.Keys()
	if err != nil {
		return nil, err
	}
	for _, kv := range kvs {
		reputer := kv.K2()
		worker := kv.K3()
		reputerWeights := weights[reputer.String()]
		if reputerWeights == nil {
			reputerWeights = make(map[ACC_ADDRESS]*Uint)
			weights[reputer.String()] = reputerWeights
		}
		weight, err := k.weights.Get(ctx, kv)
		if err != nil {
			return nil, err
		}
		weights[reputer.String()][worker.String()] = &weight
	}
	return weights, nil
}

// UpdateTopicInferenceLastRan updates the InferenceLastRan timestamp for a given topic.
func (k *Keeper) UpdateTopicInferenceLastRan(ctx context.Context, topicId TOPIC_ID, lastRanTime uint64) error {
	topic, err := k.topics.Get(ctx, topicId)
	if err != nil {
		return err
	}
	topic.InferenceLastRan = lastRanTime
	return k.topics.Set(ctx, topicId, topic)
}

// UpdateTopicWeightLastRan updates the WeightLastRan timestamp for a given topic.
func (k *Keeper) UpdateTopicWeightLastRan(ctx context.Context, topicId TOPIC_ID, lastRanTime uint64) error {
	topic, err := k.topics.Get(ctx, topicId)
	if err != nil {
		return err
	}
	topic.WeightLastRan = lastRanTime
	return k.topics.Set(ctx, topicId, topic)
}

// check a reputer node is registered
func (k *Keeper) IsReputerRegistered(ctx context.Context, reputer sdk.AccAddress) (bool, error) {
	return k.reputers.Has(ctx, reputer)
}

// Adds a new reputer to the reputer tracking data structures, reputers and topicReputers
func (k *Keeper) InsertReputer(ctx context.Context, topicsIds []TOPIC_ID, reputer sdk.AccAddress, reputerInfo state.OffchainNode) error {
	for _, topicId := range topicsIds {
		topickey := collections.Join[uint64, sdk.AccAddress](topicId, reputer)
		err := k.topicReputers.Set(ctx, topickey)
		if err != nil {
			return err
		}
	}
	err := k.reputers.Set(ctx, reputer, reputerInfo)
	if err != nil {
		return err
	}
	err = k.AddAddressTopics(ctx, reputer, topicsIds)
	if err != nil {
		return err
	}
	return nil
}

// check a worker node is registered
func (k *Keeper) IsWorkerRegistered(ctx context.Context, worker sdk.AccAddress) (bool, error) {
	return k.workers.Has(ctx, worker)
}

// Adds a new worker to the worker tracking data structures, workers and topicWorkers
func (k *Keeper) InsertWorker(ctx context.Context, topicsIds []TOPIC_ID, worker sdk.AccAddress, workerInfo state.OffchainNode) error {
	for _, topicId := range topicsIds {
		topickey := collections.Join[uint64, sdk.AccAddress](topicId, worker)
		err := k.topicWorkers.Set(ctx, topickey)
		if err != nil {
			return err
		}
	}
	err := k.workers.Set(ctx, worker, workerInfo)
	if err != nil {
		return err
	}
	err = k.AddAddressTopics(ctx, worker, topicsIds)
	if err != nil {
		return err
	}
	return nil
}

func (k *Keeper) FindWorkerNodesByOwner(ctx sdk.Context, nodeId string) ([]*state.OffchainNode, error) {
	var nodes []*state.OffchainNode
	var nodeIdParts = strings.Split(nodeId, "|")

	if len(nodeIdParts) < 2 {
		nodeIdParts = append(nodeIdParts, "")
	}

	owner, libp2pkey := nodeIdParts[0], nodeIdParts[1]

	iterator, err := k.workers.Iterate(ctx, nil)
	if err != nil {
		return nil, err
	}

	for ; iterator.Valid(); iterator.Next() {
		node, _ := iterator.Value()
		if node.Owner == owner && len(libp2pkey) == 0 || node.Owner == owner && node.LibP2PKey == libp2pkey {
			nodes = append(nodes, &node)
		}
	}

	return nodes, nil
}

func (k *Keeper) GetWorkerAddressByP2PKey(ctx context.Context, p2pKey string) (sdk.AccAddress, error) {
	iterator, err := k.workers.Iterate(ctx, nil)
	if err != nil {
		return nil, err
	}

	for ; iterator.Valid(); iterator.Next() {
		node, _ := iterator.Value()
		if node.LibP2PKey == p2pKey {
			address, err := sdk.AccAddressFromBech32(node.NodeAddress)
			if err != nil {
				return nil, err
			}

			return address, nil
		}
	}

	return nil, collections.ErrNotFound
}

func (k *Keeper) SetWeight(
	ctx context.Context,
	topicId TOPIC_ID,
	reputer sdk.AccAddress,
	worker sdk.AccAddress,
	weight Uint) error {
	key := collections.Join3(topicId, reputer, worker)
	if weight.IsZero() {
		return k.weights.Remove(ctx, key)
	}
	return k.weights.Set(ctx, key, weight)
}

func (k *Keeper) SetInference(
	ctx context.Context,
	topicID TOPIC_ID,
	worker sdk.AccAddress,
	inference state.Inference) error {
	key := collections.Join(topicID, worker)
	return k.inferences.Set(ctx, key, inference)
}

// for a given delegator, get their stake removal information
func (k *Keeper) GetStakeRemovalQueueForDelegator(ctx context.Context, delegator sdk.AccAddress) (state.StakeRemoval, error) {
	return k.stakeRemovalQueue.Get(ctx, delegator)
}

// For a given delegator, adds their stake removal information to the removal queue for delay waiting
func (k *Keeper) SetStakeRemovalQueueForDelegator(ctx context.Context, delegator sdk.AccAddress, removalInfo state.StakeRemoval) error {
	return k.stakeRemovalQueue.Set(ctx, delegator, removalInfo)
}

func (k *Keeper) AddToMempool(ctx context.Context, request state.InferenceRequest) error {
	requestId, err := request.GetRequestId()
	if err != nil {
		return err
	}
	key := collections.Join(request.TopicId, requestId)
	return k.mempool.Set(ctx, key, request)
}

func (k *Keeper) IsRequestInMempool(ctx context.Context, topicId TOPIC_ID, requestId string) (bool, error) {
	return k.mempool.Has(ctx, collections.Join(topicId, requestId))
}

func (k *Keeper) GetMempoolInferenceRequestById(ctx context.Context, topicId TOPIC_ID, requestId string) (state.InferenceRequest, error) {
	return k.mempool.Get(ctx, collections.Join(topicId, requestId))
}

func (k *Keeper) GetMempoolInferenceRequestsForTopic(ctx context.Context, topicId TOPIC_ID) ([]state.InferenceRequest, error) {
	var ret []state.InferenceRequest = make([]state.InferenceRequest, 0)
	rng := collections.NewPrefixedPairRange[TOPIC_ID, string](topicId)
	iter, err := k.mempool.Iterate(ctx, rng)
	if err != nil {
		return nil, err
	}
	for ; iter.Valid(); iter.Next() {
		value, err := iter.Value()
		if err != nil {
			return nil, err
		}
		ret = append(ret, value)
	}
	return ret, nil
}

func (k *Keeper) GetMempool(ctx context.Context) ([]state.InferenceRequest, error) {
	var ret []state.InferenceRequest = make([]state.InferenceRequest, 0)
	iter, err := k.mempool.Iterate(ctx, nil)
	if err != nil {
		return nil, err
	}
	for ; iter.Valid(); iter.Next() {
		value, err := iter.Value()
		if err != nil {
			return nil, err
		}
		ret = append(ret, value)
	}
	return ret, nil

}

func (k *Keeper) SetFunds(ctx context.Context, requestId string, amount Uint) error {
	if amount.IsZero() {
		return k.funds.Remove(ctx, requestId)
	}
	return k.funds.Set(ctx, requestId, amount)
}

func (k *Keeper) GetFunds(ctx context.Context, requestId string) (Uint, error) {
	return k.funds.Get(ctx, requestId)
}
