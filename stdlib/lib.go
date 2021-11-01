package stdlib

import (
	"fmt"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/bcs"
	"github.com/novifinancial/serde-reflection/serde-generate/runtime/golang/serde"
	diemtypes "github.com/starcoinorg/starcoin-go/types"
)

// Structured representation of a call into a known Move transaction script (legacy).
type ScriptCall interface {
	isScriptCall()
}

// Structured representation of a call into a known Move script function.
type ScriptFunctionCall interface {
	isScriptFunctionCall()
}

// Init a  NFTGallery for accept NFT<NFTMeta, NFTBody>
type ScriptFunctionCall__Accept struct {
	NftMeta diemtypes.TypeTag
	NftBody diemtypes.TypeTag
}

func (*ScriptFunctionCall__Accept) isScriptFunctionCall() {}

type ScriptFunctionCall__AcceptToken struct {
	TokenType diemtypes.TypeTag
}

func (*ScriptFunctionCall__AcceptToken) isScriptFunctionCall() {}

// Batch transfer token to others.
type ScriptFunctionCall__BatchPeerToPeer struct {
	TokenType     diemtypes.TypeTag
	Payeees       diemtypes.VecAccountAddress
	PayeeAuthKeys diemtypes.VecBytes
	Amounts       diemtypes.VecU128
}

func (*ScriptFunctionCall__BatchPeerToPeer) isScriptFunctionCall() {}

// Batch transfer token to others.
type ScriptFunctionCall__BatchPeerToPeerV2 struct {
	TokenType diemtypes.TypeTag
	Payeees   diemtypes.VecAccountAddress
	Amounts   diemtypes.VecU128
}

func (*ScriptFunctionCall__BatchPeerToPeerV2) isScriptFunctionCall() {}

// Cancel current upgrade plan, the `sender` must have `UpgradePlanCapability`.
type ScriptFunctionCall__CancelUpgradePlan struct {
}

func (*ScriptFunctionCall__CancelUpgradePlan) isScriptFunctionCall() {}

type ScriptFunctionCall__CastVote struct {
	Token           diemtypes.TypeTag
	ActionT         diemtypes.TypeTag
	ProposerAddress diemtypes.AccountAddress
	ProposalId      uint64
	Agree           bool
	Votes           serde.Uint128
}

func (*ScriptFunctionCall__CastVote) isScriptFunctionCall() {}

type ScriptFunctionCall__ConvertTwoPhaseUpgradeToTwoPhaseUpgradeV2 struct {
	PackageAddress diemtypes.AccountAddress
}

func (*ScriptFunctionCall__ConvertTwoPhaseUpgradeToTwoPhaseUpgradeV2) isScriptFunctionCall() {}

type ScriptFunctionCall__CreateAccountWithInitialAmount struct {
	TokenType     diemtypes.TypeTag
	FreshAddress  diemtypes.AccountAddress
	AuthKey       []byte
	InitialAmount serde.Uint128
}

func (*ScriptFunctionCall__CreateAccountWithInitialAmount) isScriptFunctionCall() {}

type ScriptFunctionCall__CreateAccountWithInitialAmountV2 struct {
	TokenType     diemtypes.TypeTag
	FreshAddress  diemtypes.AccountAddress
	InitialAmount serde.Uint128
}

func (*ScriptFunctionCall__CreateAccountWithInitialAmountV2) isScriptFunctionCall() {}

// Destroy empty IdentifierNFT
type ScriptFunctionCall__DestroyEmpty struct {
	NftMeta diemtypes.TypeTag
	NftBody diemtypes.TypeTag
}

func (*ScriptFunctionCall__DestroyEmpty) isScriptFunctionCall() {}

// remove terminated proposal from proposer
type ScriptFunctionCall__DestroyTerminatedProposal struct {
	TokenT          diemtypes.TypeTag
	ActionT         diemtypes.TypeTag
	ProposerAddress diemtypes.AccountAddress
	ProposalId      uint64
}

func (*ScriptFunctionCall__DestroyTerminatedProposal) isScriptFunctionCall() {}

// Disable account's auto-accept-token feature.
// The script function is reenterable.
type ScriptFunctionCall__DisableAutoAcceptToken struct {
}

func (*ScriptFunctionCall__DisableAutoAcceptToken) isScriptFunctionCall() {}

type ScriptFunctionCall__EmptyScript struct {
}

func (*ScriptFunctionCall__EmptyScript) isScriptFunctionCall() {}

// Enable account's auto-accept-token feature.
// The script function is reenterable.
type ScriptFunctionCall__EnableAutoAcceptToken struct {
}

func (*ScriptFunctionCall__EnableAutoAcceptToken) isScriptFunctionCall() {}

// Once the proposal is agreed, anyone can call the method to make the proposal happen.
type ScriptFunctionCall__Execute struct {
	TokenT          diemtypes.TypeTag
	ProposerAddress diemtypes.AccountAddress
	ProposalId      uint64
}

func (*ScriptFunctionCall__Execute) isScriptFunctionCall() {}

// Execute module upgrade plan propose by submit module upgrade plan, the propose must been agreed, and anyone can execute this function.
type ScriptFunctionCall__ExecuteModuleUpgradePlanPropose struct {
	Token           diemtypes.TypeTag
	ProposerAddress diemtypes.AccountAddress
	ProposalId      uint64
}

func (*ScriptFunctionCall__ExecuteModuleUpgradePlanPropose) isScriptFunctionCall() {}

type ScriptFunctionCall__ExecuteOnChainConfigProposal struct {
	ConfigT    diemtypes.TypeTag
	ProposalId uint64
}

func (*ScriptFunctionCall__ExecuteOnChainConfigProposal) isScriptFunctionCall() {}

type ScriptFunctionCall__ExecuteOnChainConfigProposalV2 struct {
	TokenType       diemtypes.TypeTag
	ConfigT         diemtypes.TypeTag
	ProposerAddress diemtypes.AccountAddress
	ProposalId      uint64
}

func (*ScriptFunctionCall__ExecuteOnChainConfigProposalV2) isScriptFunctionCall() {}

type ScriptFunctionCall__ExecuteWithdrawProposal struct {
	TokenT          diemtypes.TypeTag
	ProposerAddress diemtypes.AccountAddress
	ProposalId      uint64
}

func (*ScriptFunctionCall__ExecuteWithdrawProposal) isScriptFunctionCall() {}

// Let user change their vote during the voting time.
type ScriptFunctionCall__FlipVote struct {
	TokenT          diemtypes.TypeTag
	ActionT         diemtypes.TypeTag
	ProposerAddress diemtypes.AccountAddress
	ProposalId      uint64
}

func (*ScriptFunctionCall__FlipVote) isScriptFunctionCall() {}

type ScriptFunctionCall__InitDataSource struct {
	OracleT   diemtypes.TypeTag
	InitValue serde.Uint128
}

func (*ScriptFunctionCall__InitDataSource) isScriptFunctionCall() {}

type ScriptFunctionCall__Initialize struct {
	StdlibVersion                uint64
	RewardDelay                  uint64
	PreMineStcAmount             serde.Uint128
	TimeMintStcAmount            serde.Uint128
	TimeMintStcPeriod            uint64
	ParentHash                   []byte
	AssociationAuthKey           []byte
	GenesisAuthKey               []byte
	ChainId                      uint8
	GenesisTimestamp             uint64
	UncleRateTarget              uint64
	EpochBlockCount              uint64
	BaseBlockTimeTarget          uint64
	BaseBlockDifficultyWindow    uint64
	BaseRewardPerBlock           serde.Uint128
	BaseRewardPerUnclePercent    uint64
	MinBlockTimeTarget           uint64
	MaxBlockTimeTarget           uint64
	BaseMaxUnclesPerBlock        uint64
	BaseBlockGasLimit            uint64
	Strategy                     uint8
	ScriptAllowed                bool
	ModulePublishingAllowed      bool
	InstructionSchedule          []byte
	NativeSchedule               []byte
	GlobalMemoryPerByteCost      uint64
	GlobalMemoryPerByteWriteCost uint64
	MinTransactionGasUnits       uint64
	LargeTransactionCutoff       uint64
	InstrinsicGasPerByte         uint64
	MaximumNumberOfGasUnits      uint64
	MinPricePerGasUnit           uint64
	MaxPricePerGasUnit           uint64
	MaxTransactionSizeInBytes    uint64
	GasUnitScalingFactor         uint64
	DefaultAccountSize           uint64
	VotingDelay                  uint64
	VotingPeriod                 uint64
	VotingQuorumRate             uint8
	MinActionDelay               uint64
	TransactionTimeout           uint64
}

func (*ScriptFunctionCall__Initialize) isScriptFunctionCall() {}

type ScriptFunctionCall__InitializeV2 struct {
	StdlibVersion                uint64
	RewardDelay                  uint64
	TotalStcAmount               serde.Uint128
	PreMineStcAmount             serde.Uint128
	TimeMintStcAmount            serde.Uint128
	TimeMintStcPeriod            uint64
	ParentHash                   []byte
	AssociationAuthKey           []byte
	GenesisAuthKey               []byte
	ChainId                      uint8
	GenesisTimestamp             uint64
	UncleRateTarget              uint64
	EpochBlockCount              uint64
	BaseBlockTimeTarget          uint64
	BaseBlockDifficultyWindow    uint64
	BaseRewardPerBlock           serde.Uint128
	BaseRewardPerUnclePercent    uint64
	MinBlockTimeTarget           uint64
	MaxBlockTimeTarget           uint64
	BaseMaxUnclesPerBlock        uint64
	BaseBlockGasLimit            uint64
	Strategy                     uint8
	ScriptAllowed                bool
	ModulePublishingAllowed      bool
	InstructionSchedule          []byte
	NativeSchedule               []byte
	GlobalMemoryPerByteCost      uint64
	GlobalMemoryPerByteWriteCost uint64
	MinTransactionGasUnits       uint64
	LargeTransactionCutoff       uint64
	InstrinsicGasPerByte         uint64
	MaximumNumberOfGasUnits      uint64
	MinPricePerGasUnit           uint64
	MaxPricePerGasUnit           uint64
	MaxTransactionSizeInBytes    uint64
	GasUnitScalingFactor         uint64
	DefaultAccountSize           uint64
	VotingDelay                  uint64
	VotingPeriod                 uint64
	VotingQuorumRate             uint8
	MinActionDelay               uint64
	TransactionTimeout           uint64
}

func (*ScriptFunctionCall__InitializeV2) isScriptFunctionCall() {}

type ScriptFunctionCall__Mint struct {
	Index       uint64
	MerkleProof diemtypes.VecBytes
}

func (*ScriptFunctionCall__Mint) isScriptFunctionCall() {}

type ScriptFunctionCall__DummyTokenMint struct {
	Amount serde.Uint128
}

func (*ScriptFunctionCall__DummyTokenMint) isScriptFunctionCall() {}

type ScriptFunctionCall__PeerToPeer struct {
	TokenType    diemtypes.TypeTag
	Payee        diemtypes.AccountAddress
	PayeeAuthKey []byte
	Amount       serde.Uint128
}

func (*ScriptFunctionCall__PeerToPeer) isScriptFunctionCall() {}

type ScriptFunctionCall__PeerToPeerBatch struct {
	TokenType     diemtypes.TypeTag
	Payeees       []byte
	PayeeAuthKeys []byte
	Amount        serde.Uint128
}

func (*ScriptFunctionCall__PeerToPeerBatch) isScriptFunctionCall() {}

type ScriptFunctionCall__PeerToPeerV2 struct {
	TokenType diemtypes.TypeTag
	Payee     diemtypes.AccountAddress
	Amount    serde.Uint128
}

func (*ScriptFunctionCall__PeerToPeerV2) isScriptFunctionCall() {}

type ScriptFunctionCall__PeerToPeerWithMetadata struct {
	TokenType    diemtypes.TypeTag
	Payee        diemtypes.AccountAddress
	PayeeAuthKey []byte
	Amount       serde.Uint128
	Metadata     []byte
}

func (*ScriptFunctionCall__PeerToPeerWithMetadata) isScriptFunctionCall() {}

type ScriptFunctionCall__PeerToPeerWithMetadataV2 struct {
	TokenType diemtypes.TypeTag
	Payee     diemtypes.AccountAddress
	Amount    serde.Uint128
	Metadata  []byte
}

func (*ScriptFunctionCall__PeerToPeerWithMetadataV2) isScriptFunctionCall() {}

// Entrypoint for the proposal.
type ScriptFunctionCall__Propose struct {
	TokenT           diemtypes.TypeTag
	VotingDelay      uint64
	VotingPeriod     uint64
	VotingQuorumRate uint8
	MinActionDelay   uint64
	ExecDelay        uint64
}

func (*ScriptFunctionCall__Propose) isScriptFunctionCall() {}

type ScriptFunctionCall__ProposeModuleUpgradeV2 struct {
	Token         diemtypes.TypeTag
	ModuleAddress diemtypes.AccountAddress
	PackageHash   []byte
	Version       uint64
	ExecDelay     uint64
	Enforced      bool
}

func (*ScriptFunctionCall__ProposeModuleUpgradeV2) isScriptFunctionCall() {}

type ScriptFunctionCall__ProposeUpdateConsensusConfig struct {
	UncleRateTarget           uint64
	BaseBlockTimeTarget       uint64
	BaseRewardPerBlock        serde.Uint128
	BaseRewardPerUnclePercent uint64
	EpochBlockCount           uint64
	BaseBlockDifficultyWindow uint64
	MinBlockTimeTarget        uint64
	MaxBlockTimeTarget        uint64
	BaseMaxUnclesPerBlock     uint64
	BaseBlockGasLimit         uint64
	Strategy                  uint8
	ExecDelay                 uint64
}

func (*ScriptFunctionCall__ProposeUpdateConsensusConfig) isScriptFunctionCall() {}

type ScriptFunctionCall__ProposeUpdateMoveLanguageVersion struct {
	NewVersion uint64
	ExecDelay  uint64
}

func (*ScriptFunctionCall__ProposeUpdateMoveLanguageVersion) isScriptFunctionCall() {}

type ScriptFunctionCall__ProposeUpdateRewardConfig struct {
	RewardDelay uint64
	ExecDelay   uint64
}

func (*ScriptFunctionCall__ProposeUpdateRewardConfig) isScriptFunctionCall() {}

type ScriptFunctionCall__ProposeUpdateTxnPublishOption struct {
	ScriptAllowed           bool
	ModulePublishingAllowed bool
	ExecDelay               uint64
}

func (*ScriptFunctionCall__ProposeUpdateTxnPublishOption) isScriptFunctionCall() {}

type ScriptFunctionCall__ProposeUpdateTxnTimeoutConfig struct {
	DurationSeconds uint64
	ExecDelay       uint64
}

func (*ScriptFunctionCall__ProposeUpdateTxnTimeoutConfig) isScriptFunctionCall() {}

type ScriptFunctionCall__ProposeUpdateVmConfig struct {
	InstructionSchedule          []byte
	NativeSchedule               []byte
	GlobalMemoryPerByteCost      uint64
	GlobalMemoryPerByteWriteCost uint64
	MinTransactionGasUnits       uint64
	LargeTransactionCutoff       uint64
	InstrinsicGasPerByte         uint64
	MaximumNumberOfGasUnits      uint64
	MinPricePerGasUnit           uint64
	MaxPricePerGasUnit           uint64
	MaxTransactionSizeInBytes    uint64
	GasUnitScalingFactor         uint64
	DefaultAccountSize           uint64
	ExecDelay                    uint64
}

func (*ScriptFunctionCall__ProposeUpdateVmConfig) isScriptFunctionCall() {}

type ScriptFunctionCall__ProposeWithdraw struct {
	TokenT    diemtypes.TypeTag
	Receiver  diemtypes.AccountAddress
	Amount    serde.Uint128
	Period    uint64
	ExecDelay uint64
}

func (*ScriptFunctionCall__ProposeWithdraw) isScriptFunctionCall() {}

// queue agreed proposal to execute.
type ScriptFunctionCall__QueueProposalAction struct {
	TokenT          diemtypes.TypeTag
	ActionT         diemtypes.TypeTag
	ProposerAddress diemtypes.AccountAddress
	ProposalId      uint64
}

func (*ScriptFunctionCall__QueueProposalAction) isScriptFunctionCall() {}

type ScriptFunctionCall__RegisterOracle struct {
	OracleT   diemtypes.TypeTag
	Precision uint8
}

func (*ScriptFunctionCall__RegisterOracle) isScriptFunctionCall() {}

// revoke all votes on a proposal
type ScriptFunctionCall__RevokeVote struct {
	Token           diemtypes.TypeTag
	Action          diemtypes.TypeTag
	ProposerAddress diemtypes.AccountAddress
	ProposalId      uint64
}

func (*ScriptFunctionCall__RevokeVote) isScriptFunctionCall() {}

// revoke some votes on a proposal
type ScriptFunctionCall__RevokeVoteOfPower struct {
	Token           diemtypes.TypeTag
	Action          diemtypes.TypeTag
	ProposerAddress diemtypes.AccountAddress
	ProposalId      uint64
	Power           serde.Uint128
}

func (*ScriptFunctionCall__RevokeVoteOfPower) isScriptFunctionCall() {}

type ScriptFunctionCall__RotateAuthenticationKey struct {
	NewKey []byte
}

func (*ScriptFunctionCall__RotateAuthenticationKey) isScriptFunctionCall() {}

// a alias of execute_module_upgrade_plan_propose, will deprecated in the future.
type ScriptFunctionCall__SubmitModuleUpgradePlan struct {
	Token           diemtypes.TypeTag
	ProposerAddress diemtypes.AccountAddress
	ProposalId      uint64
}

func (*ScriptFunctionCall__SubmitModuleUpgradePlan) isScriptFunctionCall() {}

// Directly submit a upgrade plan, the `sender`'s module upgrade plan must been PackageTxnManager::STRATEGY_TWO_PHASE and have UpgradePlanCapability
type ScriptFunctionCall__SubmitUpgradePlan struct {
	PackageHash []byte
	Version     uint64
	Enforced    bool
}

func (*ScriptFunctionCall__SubmitUpgradePlan) isScriptFunctionCall() {}

// association account should call this script after upgrade from v2 to v3.
type ScriptFunctionCall__TakeLinearWithdrawCapability struct {
}

func (*ScriptFunctionCall__TakeLinearWithdrawCapability) isScriptFunctionCall() {}

// Take Offer and put to signer's Collection<Offered>.
type ScriptFunctionCall__TakeOffer struct {
	Offered      diemtypes.TypeTag
	OfferAddress diemtypes.AccountAddress
}

func (*ScriptFunctionCall__TakeOffer) isScriptFunctionCall() {}

// Transfer NFT<NFTMeta, NFTBody> with `id` from `sender` to `receiver`
type ScriptFunctionCall__Transfer struct {
	NftMeta  diemtypes.TypeTag
	NftBody  diemtypes.TypeTag
	Id       uint64
	Receiver diemtypes.AccountAddress
}

func (*ScriptFunctionCall__Transfer) isScriptFunctionCall() {}

type ScriptFunctionCall__UnstakeVote struct {
	Token           diemtypes.TypeTag
	Action          diemtypes.TypeTag
	ProposerAddress diemtypes.AccountAddress
	ProposalId      uint64
}

func (*ScriptFunctionCall__UnstakeVote) isScriptFunctionCall() {}

type ScriptFunctionCall__Update struct {
	OracleT diemtypes.TypeTag
	Value   serde.Uint128
}

func (*ScriptFunctionCall__Update) isScriptFunctionCall() {}

// Update `sender`'s module upgrade strategy to `strategy`
type ScriptFunctionCall__UpdateModuleUpgradeStrategy struct {
	Strategy uint8
}

func (*ScriptFunctionCall__UpdateModuleUpgradeStrategy) isScriptFunctionCall() {}

// Stdlib upgrade script from v2 to v3
type ScriptFunctionCall__UpgradeFromV2ToV3 struct {
	TotalStcAmount serde.Uint128
}

func (*ScriptFunctionCall__UpgradeFromV2ToV3) isScriptFunctionCall() {}

type ScriptFunctionCall__UpgradeFromV5ToV6 struct {
}

func (*ScriptFunctionCall__UpgradeFromV5ToV6) isScriptFunctionCall() {}

type ScriptFunctionCall__UpgradeFromV6ToV7 struct {
}

func (*ScriptFunctionCall__UpgradeFromV6ToV7) isScriptFunctionCall() {}

type ScriptFunctionCall__UpgradeFromV7ToV8 struct {
}

func (*ScriptFunctionCall__UpgradeFromV7ToV8) isScriptFunctionCall() {}

type ScriptFunctionCall__WithdrawAndSplitLtWithdrawCap struct {
	TokenT     diemtypes.TypeTag
	ForAddress diemtypes.AccountAddress
	Amount     serde.Uint128
	LockPeriod uint64
}

func (*ScriptFunctionCall__WithdrawAndSplitLtWithdrawCap) isScriptFunctionCall() {}

type ScriptFunctionCall__WithdrawTokenWithLinearWithdrawCapability struct {
	TokenT diemtypes.TypeTag
}

func (*ScriptFunctionCall__WithdrawTokenWithLinearWithdrawCapability) isScriptFunctionCall() {}

// Build a Diem `TransactionPayload` from a structured object `ScriptFunctionCall`.
func EncodeScriptFunction(call ScriptFunctionCall) diemtypes.TransactionPayload {
	switch call := call.(type) {
	case *ScriptFunctionCall__Accept:
		return EncodeAcceptScriptFunction(call.NftMeta, call.NftBody)
	case *ScriptFunctionCall__AcceptToken:
		return EncodeAcceptTokenScriptFunction(call.TokenType)
	case *ScriptFunctionCall__BatchPeerToPeer:
		return EncodeBatchPeerToPeerScriptFunction(call.TokenType, call.Payeees, call.PayeeAuthKeys, call.Amounts)
	case *ScriptFunctionCall__BatchPeerToPeerV2:
		return EncodeBatchPeerToPeerV2ScriptFunction(call.TokenType, call.Payeees, call.Amounts)
	case *ScriptFunctionCall__CancelUpgradePlan:
		return EncodeCancelUpgradePlanScriptFunction()
	case *ScriptFunctionCall__CastVote:
		return EncodeCastVoteScriptFunction(call.Token, call.ActionT, call.ProposerAddress, call.ProposalId, call.Agree, call.Votes)
	case *ScriptFunctionCall__ConvertTwoPhaseUpgradeToTwoPhaseUpgradeV2:
		return EncodeConvertTwoPhaseUpgradeToTwoPhaseUpgradeV2ScriptFunction(call.PackageAddress)
	case *ScriptFunctionCall__CreateAccountWithInitialAmount:
		return EncodeCreateAccountWithInitialAmountScriptFunction(call.TokenType, call.FreshAddress, call.AuthKey, call.InitialAmount)
	case *ScriptFunctionCall__CreateAccountWithInitialAmountV2:
		return EncodeCreateAccountWithInitialAmountV2ScriptFunction(call.TokenType, call.FreshAddress, call.InitialAmount)
	case *ScriptFunctionCall__DestroyEmpty:
		return EncodeDestroyEmptyScriptFunction(call.NftMeta, call.NftBody)
	case *ScriptFunctionCall__DestroyTerminatedProposal:
		return EncodeDestroyTerminatedProposalScriptFunction(call.TokenT, call.ActionT, call.ProposerAddress, call.ProposalId)
	case *ScriptFunctionCall__DisableAutoAcceptToken:
		return EncodeDisableAutoAcceptTokenScriptFunction()
	case *ScriptFunctionCall__EmptyScript:
		return EncodeEmptyScriptScriptFunction()
	case *ScriptFunctionCall__EnableAutoAcceptToken:
		return EncodeEnableAutoAcceptTokenScriptFunction()
	case *ScriptFunctionCall__Execute:
		return EncodeExecuteScriptFunction(call.TokenT, call.ProposerAddress, call.ProposalId)
	case *ScriptFunctionCall__ExecuteModuleUpgradePlanPropose:
		return EncodeExecuteModuleUpgradePlanProposeScriptFunction(call.Token, call.ProposerAddress, call.ProposalId)
	case *ScriptFunctionCall__ExecuteOnChainConfigProposal:
		return EncodeExecuteOnChainConfigProposalScriptFunction(call.ConfigT, call.ProposalId)
	case *ScriptFunctionCall__ExecuteOnChainConfigProposalV2:
		return EncodeExecuteOnChainConfigProposalV2ScriptFunction(call.TokenType, call.ConfigT, call.ProposerAddress, call.ProposalId)
	case *ScriptFunctionCall__ExecuteWithdrawProposal:
		return EncodeExecuteWithdrawProposalScriptFunction(call.TokenT, call.ProposerAddress, call.ProposalId)
	case *ScriptFunctionCall__FlipVote:
		return EncodeFlipVoteScriptFunction(call.TokenT, call.ActionT, call.ProposerAddress, call.ProposalId)
	case *ScriptFunctionCall__InitDataSource:
		return EncodeInitDataSourceScriptFunction(call.OracleT, call.InitValue)
	case *ScriptFunctionCall__Initialize:
		return EncodeInitializeScriptFunction(call.StdlibVersion, call.RewardDelay, call.PreMineStcAmount, call.TimeMintStcAmount, call.TimeMintStcPeriod, call.ParentHash, call.AssociationAuthKey, call.GenesisAuthKey, call.ChainId, call.GenesisTimestamp, call.UncleRateTarget, call.EpochBlockCount, call.BaseBlockTimeTarget, call.BaseBlockDifficultyWindow, call.BaseRewardPerBlock, call.BaseRewardPerUnclePercent, call.MinBlockTimeTarget, call.MaxBlockTimeTarget, call.BaseMaxUnclesPerBlock, call.BaseBlockGasLimit, call.Strategy, call.ScriptAllowed, call.ModulePublishingAllowed, call.InstructionSchedule, call.NativeSchedule, call.GlobalMemoryPerByteCost, call.GlobalMemoryPerByteWriteCost, call.MinTransactionGasUnits, call.LargeTransactionCutoff, call.InstrinsicGasPerByte, call.MaximumNumberOfGasUnits, call.MinPricePerGasUnit, call.MaxPricePerGasUnit, call.MaxTransactionSizeInBytes, call.GasUnitScalingFactor, call.DefaultAccountSize, call.VotingDelay, call.VotingPeriod, call.VotingQuorumRate, call.MinActionDelay, call.TransactionTimeout)
	case *ScriptFunctionCall__InitializeV2:
		return EncodeInitializeV2ScriptFunction(call.StdlibVersion, call.RewardDelay, call.TotalStcAmount, call.PreMineStcAmount, call.TimeMintStcAmount, call.TimeMintStcPeriod, call.ParentHash, call.AssociationAuthKey, call.GenesisAuthKey, call.ChainId, call.GenesisTimestamp, call.UncleRateTarget, call.EpochBlockCount, call.BaseBlockTimeTarget, call.BaseBlockDifficultyWindow, call.BaseRewardPerBlock, call.BaseRewardPerUnclePercent, call.MinBlockTimeTarget, call.MaxBlockTimeTarget, call.BaseMaxUnclesPerBlock, call.BaseBlockGasLimit, call.Strategy, call.ScriptAllowed, call.ModulePublishingAllowed, call.InstructionSchedule, call.NativeSchedule, call.GlobalMemoryPerByteCost, call.GlobalMemoryPerByteWriteCost, call.MinTransactionGasUnits, call.LargeTransactionCutoff, call.InstrinsicGasPerByte, call.MaximumNumberOfGasUnits, call.MinPricePerGasUnit, call.MaxPricePerGasUnit, call.MaxTransactionSizeInBytes, call.GasUnitScalingFactor, call.DefaultAccountSize, call.VotingDelay, call.VotingPeriod, call.VotingQuorumRate, call.MinActionDelay, call.TransactionTimeout)
	case *ScriptFunctionCall__Mint:
		return EncodeMintScriptFunction(call.Index, call.MerkleProof)
	case *ScriptFunctionCall__DummyTokenMint:
		return EncodeDummyTokenMintScriptFunction(call.Amount)
	case *ScriptFunctionCall__PeerToPeer:
		return EncodePeerToPeerScriptFunction(call.TokenType, call.Payee, call.PayeeAuthKey, call.Amount)
	case *ScriptFunctionCall__PeerToPeerBatch:
		return EncodePeerToPeerBatchScriptFunction(call.TokenType, call.Payeees, call.PayeeAuthKeys, call.Amount)
	case *ScriptFunctionCall__PeerToPeerV2:
		return EncodePeerToPeerV2ScriptFunction(call.TokenType, call.Payee, call.Amount)
	case *ScriptFunctionCall__PeerToPeerWithMetadata:
		return EncodePeerToPeerWithMetadataScriptFunction(call.TokenType, call.Payee, call.PayeeAuthKey, call.Amount, call.Metadata)
	case *ScriptFunctionCall__PeerToPeerWithMetadataV2:
		return EncodePeerToPeerWithMetadataV2ScriptFunction(call.TokenType, call.Payee, call.Amount, call.Metadata)
	case *ScriptFunctionCall__Propose:
		return EncodeProposeScriptFunction(call.TokenT, call.VotingDelay, call.VotingPeriod, call.VotingQuorumRate, call.MinActionDelay, call.ExecDelay)
	case *ScriptFunctionCall__ProposeModuleUpgradeV2:
		return EncodeProposeModuleUpgradeV2ScriptFunction(call.Token, call.ModuleAddress, call.PackageHash, call.Version, call.ExecDelay, call.Enforced)
	case *ScriptFunctionCall__ProposeUpdateConsensusConfig:
		return EncodeProposeUpdateConsensusConfigScriptFunction(call.UncleRateTarget, call.BaseBlockTimeTarget, call.BaseRewardPerBlock, call.BaseRewardPerUnclePercent, call.EpochBlockCount, call.BaseBlockDifficultyWindow, call.MinBlockTimeTarget, call.MaxBlockTimeTarget, call.BaseMaxUnclesPerBlock, call.BaseBlockGasLimit, call.Strategy, call.ExecDelay)
	case *ScriptFunctionCall__ProposeUpdateMoveLanguageVersion:
		return EncodeProposeUpdateMoveLanguageVersionScriptFunction(call.NewVersion, call.ExecDelay)
	case *ScriptFunctionCall__ProposeUpdateRewardConfig:
		return EncodeProposeUpdateRewardConfigScriptFunction(call.RewardDelay, call.ExecDelay)
	case *ScriptFunctionCall__ProposeUpdateTxnPublishOption:
		return EncodeProposeUpdateTxnPublishOptionScriptFunction(call.ScriptAllowed, call.ModulePublishingAllowed, call.ExecDelay)
	case *ScriptFunctionCall__ProposeUpdateTxnTimeoutConfig:
		return EncodeProposeUpdateTxnTimeoutConfigScriptFunction(call.DurationSeconds, call.ExecDelay)
	case *ScriptFunctionCall__ProposeUpdateVmConfig:
		return EncodeProposeUpdateVmConfigScriptFunction(call.InstructionSchedule, call.NativeSchedule, call.GlobalMemoryPerByteCost, call.GlobalMemoryPerByteWriteCost, call.MinTransactionGasUnits, call.LargeTransactionCutoff, call.InstrinsicGasPerByte, call.MaximumNumberOfGasUnits, call.MinPricePerGasUnit, call.MaxPricePerGasUnit, call.MaxTransactionSizeInBytes, call.GasUnitScalingFactor, call.DefaultAccountSize, call.ExecDelay)
	case *ScriptFunctionCall__ProposeWithdraw:
		return EncodeProposeWithdrawScriptFunction(call.TokenT, call.Receiver, call.Amount, call.Period, call.ExecDelay)
	case *ScriptFunctionCall__QueueProposalAction:
		return EncodeQueueProposalActionScriptFunction(call.TokenT, call.ActionT, call.ProposerAddress, call.ProposalId)
	case *ScriptFunctionCall__RegisterOracle:
		return EncodeRegisterOracleScriptFunction(call.OracleT, call.Precision)
	case *ScriptFunctionCall__RevokeVote:
		return EncodeRevokeVoteScriptFunction(call.Token, call.Action, call.ProposerAddress, call.ProposalId)
	case *ScriptFunctionCall__RevokeVoteOfPower:
		return EncodeRevokeVoteOfPowerScriptFunction(call.Token, call.Action, call.ProposerAddress, call.ProposalId, call.Power)
	case *ScriptFunctionCall__RotateAuthenticationKey:
		return EncodeRotateAuthenticationKeyScriptFunction(call.NewKey)
	case *ScriptFunctionCall__SubmitModuleUpgradePlan:
		return EncodeSubmitModuleUpgradePlanScriptFunction(call.Token, call.ProposerAddress, call.ProposalId)
	case *ScriptFunctionCall__SubmitUpgradePlan:
		return EncodeSubmitUpgradePlanScriptFunction(call.PackageHash, call.Version, call.Enforced)
	case *ScriptFunctionCall__TakeLinearWithdrawCapability:
		return EncodeTakeLinearWithdrawCapabilityScriptFunction()
	case *ScriptFunctionCall__TakeOffer:
		return EncodeTakeOfferScriptFunction(call.Offered, call.OfferAddress)
	case *ScriptFunctionCall__Transfer:
		return EncodeTransferScriptFunction(call.NftMeta, call.NftBody, call.Id, call.Receiver)
	case *ScriptFunctionCall__UnstakeVote:
		return EncodeUnstakeVoteScriptFunction(call.Token, call.Action, call.ProposerAddress, call.ProposalId)
	case *ScriptFunctionCall__Update:
		return EncodeUpdateScriptFunction(call.OracleT, call.Value)
	case *ScriptFunctionCall__UpdateModuleUpgradeStrategy:
		return EncodeUpdateModuleUpgradeStrategyScriptFunction(call.Strategy)
	case *ScriptFunctionCall__UpgradeFromV2ToV3:
		return EncodeUpgradeFromV2ToV3ScriptFunction(call.TotalStcAmount)
	case *ScriptFunctionCall__UpgradeFromV5ToV6:
		return EncodeUpgradeFromV5ToV6ScriptFunction()
	case *ScriptFunctionCall__UpgradeFromV6ToV7:
		return EncodeUpgradeFromV6ToV7ScriptFunction()
	case *ScriptFunctionCall__UpgradeFromV7ToV8:
		return EncodeUpgradeFromV7ToV8ScriptFunction()
	case *ScriptFunctionCall__WithdrawAndSplitLtWithdrawCap:
		return EncodeWithdrawAndSplitLtWithdrawCapScriptFunction(call.TokenT, call.ForAddress, call.Amount, call.LockPeriod)
	case *ScriptFunctionCall__WithdrawTokenWithLinearWithdrawCapability:
		return EncodeWithdrawTokenWithLinearWithdrawCapabilityScriptFunction(call.TokenT)
	}
	panic("unreachable")
}

// Try to recognize a Diem `Script` and convert it into a structured object `ScriptCall`.
func DecodeScript(script *diemtypes.Script) (ScriptCall, error) {
	if helper := script_decoder_map[string(script.Code)]; helper != nil {
		val, err := helper(script)
		return val, err
	} else {
		return nil, fmt.Errorf("Unknown script bytecode: %s", string(script.Code))
	}
}

// Try to recognize a Diem `TransactionPayload` and convert it into a structured object `ScriptFunctionCall`.
func DecodeScriptFunctionPayload(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := script.(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if helper := script_function_decoder_map[string(script.Value.Module.Name)+string(script.Value.Function)]; helper != nil {
			val, err := helper(script)
			return val, err
		} else {
			return nil, fmt.Errorf("Unknown script function: %s::%s", script.Value.Module.Name, script.Value.Function)
		}
	default:
		return nil, fmt.Errorf("Unknown transaction payload encountered when decoding")
	}
}

// Init a  NFTGallery for accept NFT<NFTMeta, NFTBody>
func EncodeAcceptScriptFunction(nft_meta diemtypes.TypeTag, nft_body diemtypes.TypeTag) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "NFTGalleryScripts"},
			Function: "accept",
			TyArgs:   []diemtypes.TypeTag{nft_meta, nft_body},
			Args:     [][]byte{},
		},
	}
}

func EncodeAcceptTokenScriptFunction(token_type diemtypes.TypeTag) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "Account"},
			Function: "accept_token",
			TyArgs:   []diemtypes.TypeTag{token_type},
			Args:     [][]byte{},
		},
	}
}

// Batch transfer token to others.
func EncodeBatchPeerToPeerScriptFunction(token_type diemtypes.TypeTag, payeees diemtypes.VecAccountAddress, _payee_auth_keys diemtypes.VecBytes, amounts diemtypes.VecU128) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "TransferScripts"},
			Function: "batch_peer_to_peer",
			TyArgs:   []diemtypes.TypeTag{token_type},
			Args:     [][]byte{encode_vecaccountaddress_argument(payeees), encode_vecbytes_argument(_payee_auth_keys), encode_vecu128_argument(amounts)},
		},
	}
}

// Batch transfer token to others.
func EncodeBatchPeerToPeerV2ScriptFunction(token_type diemtypes.TypeTag, payeees diemtypes.VecAccountAddress, amounts diemtypes.VecU128) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "TransferScripts"},
			Function: "batch_peer_to_peer_v2",
			TyArgs:   []diemtypes.TypeTag{token_type},
			Args:     [][]byte{encode_vecaccountaddress_argument(payeees), encode_vecu128_argument(amounts)},
		},
	}
}

// Cancel current upgrade plan, the `sender` must have `UpgradePlanCapability`.
func EncodeCancelUpgradePlanScriptFunction() diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "ModuleUpgradeScripts"},
			Function: "cancel_upgrade_plan",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{},
		},
	}
}

func EncodeCastVoteScriptFunction(token diemtypes.TypeTag, action_t diemtypes.TypeTag, proposer_address diemtypes.AccountAddress, proposal_id uint64, agree bool, votes serde.Uint128) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "DaoVoteScripts"},
			Function: "cast_vote",
			TyArgs:   []diemtypes.TypeTag{token, action_t},
			Args:     [][]byte{encode_address_argument(proposer_address), encode_u64_argument(proposal_id), encode_bool_argument(agree), encode_u128_argument(votes)},
		},
	}
}

func EncodeConvertTwoPhaseUpgradeToTwoPhaseUpgradeV2ScriptFunction(package_address diemtypes.AccountAddress) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "PackageTxnManager"},
			Function: "convert_TwoPhaseUpgrade_to_TwoPhaseUpgradeV2",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{encode_address_argument(package_address)},
		},
	}
}

func EncodeCreateAccountWithInitialAmountScriptFunction(token_type diemtypes.TypeTag, fresh_address diemtypes.AccountAddress, _auth_key []byte, initial_amount serde.Uint128) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "Account"},
			Function: "create_account_with_initial_amount",
			TyArgs:   []diemtypes.TypeTag{token_type},
			Args:     [][]byte{encode_address_argument(fresh_address), encode_u8vector_argument(_auth_key), encode_u128_argument(initial_amount)},
		},
	}
}

func EncodeCreateAccountWithInitialAmountV2ScriptFunction(token_type diemtypes.TypeTag, fresh_address diemtypes.AccountAddress, initial_amount serde.Uint128) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "Account"},
			Function: "create_account_with_initial_amount_v2",
			TyArgs:   []diemtypes.TypeTag{token_type},
			Args:     [][]byte{encode_address_argument(fresh_address), encode_u128_argument(initial_amount)},
		},
	}
}

// Destroy empty IdentifierNFT
func EncodeDestroyEmptyScriptFunction(nft_meta diemtypes.TypeTag, nft_body diemtypes.TypeTag) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "IdentifierNFTScripts"},
			Function: "destroy_empty",
			TyArgs:   []diemtypes.TypeTag{nft_meta, nft_body},
			Args:     [][]byte{},
		},
	}
}

// remove terminated proposal from proposer
func EncodeDestroyTerminatedProposalScriptFunction(token_t diemtypes.TypeTag, action_t diemtypes.TypeTag, proposer_address diemtypes.AccountAddress, proposal_id uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "Dao"},
			Function: "destroy_terminated_proposal",
			TyArgs:   []diemtypes.TypeTag{token_t, action_t},
			Args:     [][]byte{encode_address_argument(proposer_address), encode_u64_argument(proposal_id)},
		},
	}
}

// Disable account's auto-accept-token feature.
// The script function is reenterable.
func EncodeDisableAutoAcceptTokenScriptFunction() diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "AccountScripts"},
			Function: "disable_auto_accept_token",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{},
		},
	}
}

func EncodeEmptyScriptScriptFunction() diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "EmptyScripts"},
			Function: "empty_script",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{},
		},
	}
}

// Enable account's auto-accept-token feature.
// The script function is reenterable.
func EncodeEnableAutoAcceptTokenScriptFunction() diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "AccountScripts"},
			Function: "enable_auto_accept_token",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{},
		},
	}
}

// Once the proposal is agreed, anyone can call the method to make the proposal happen.
func EncodeExecuteScriptFunction(token_t diemtypes.TypeTag, proposer_address diemtypes.AccountAddress, proposal_id uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "ModifyDaoConfigProposal"},
			Function: "execute",
			TyArgs:   []diemtypes.TypeTag{token_t},
			Args:     [][]byte{encode_address_argument(proposer_address), encode_u64_argument(proposal_id)},
		},
	}
}

// Execute module upgrade plan propose by submit module upgrade plan, the propose must been agreed, and anyone can execute this function.
func EncodeExecuteModuleUpgradePlanProposeScriptFunction(token diemtypes.TypeTag, proposer_address diemtypes.AccountAddress, proposal_id uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "ModuleUpgradeScripts"},
			Function: "execute_module_upgrade_plan_propose",
			TyArgs:   []diemtypes.TypeTag{token},
			Args:     [][]byte{encode_address_argument(proposer_address), encode_u64_argument(proposal_id)},
		},
	}
}

func EncodeExecuteOnChainConfigProposalScriptFunction(config_t diemtypes.TypeTag, proposal_id uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "OnChainConfigScripts"},
			Function: "execute_on_chain_config_proposal",
			TyArgs:   []diemtypes.TypeTag{config_t},
			Args:     [][]byte{encode_u64_argument(proposal_id)},
		},
	}
}

func EncodeExecuteOnChainConfigProposalV2ScriptFunction(token_type diemtypes.TypeTag, config_t diemtypes.TypeTag, proposer_address diemtypes.AccountAddress, proposal_id uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "OnChainConfigScripts"},
			Function: "execute_on_chain_config_proposal_v2",
			TyArgs:   []diemtypes.TypeTag{token_type, config_t},
			Args:     [][]byte{encode_address_argument(proposer_address), encode_u64_argument(proposal_id)},
		},
	}
}

func EncodeExecuteWithdrawProposalScriptFunction(token_t diemtypes.TypeTag, proposer_address diemtypes.AccountAddress, proposal_id uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "TreasuryScripts"},
			Function: "execute_withdraw_proposal",
			TyArgs:   []diemtypes.TypeTag{token_t},
			Args:     [][]byte{encode_address_argument(proposer_address), encode_u64_argument(proposal_id)},
		},
	}
}

// Let user change their vote during the voting time.
func EncodeFlipVoteScriptFunction(token_t diemtypes.TypeTag, action_t diemtypes.TypeTag, proposer_address diemtypes.AccountAddress, proposal_id uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "DaoVoteScripts"},
			Function: "flip_vote",
			TyArgs:   []diemtypes.TypeTag{token_t, action_t},
			Args:     [][]byte{encode_address_argument(proposer_address), encode_u64_argument(proposal_id)},
		},
	}
}

func EncodeInitDataSourceScriptFunction(oracle_t diemtypes.TypeTag, init_value serde.Uint128) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "PriceOracleScripts"},
			Function: "init_data_source",
			TyArgs:   []diemtypes.TypeTag{oracle_t},
			Args:     [][]byte{encode_u128_argument(init_value)},
		},
	}
}

func EncodeInitializeScriptFunction(stdlib_version uint64, reward_delay uint64, pre_mine_stc_amount serde.Uint128, time_mint_stc_amount serde.Uint128, time_mint_stc_period uint64, parent_hash []byte, association_auth_key []byte, genesis_auth_key []byte, chain_id uint8, genesis_timestamp uint64, uncle_rate_target uint64, epoch_block_count uint64, base_block_time_target uint64, base_block_difficulty_window uint64, base_reward_per_block serde.Uint128, base_reward_per_uncle_percent uint64, min_block_time_target uint64, max_block_time_target uint64, base_max_uncles_per_block uint64, base_block_gas_limit uint64, strategy uint8, script_allowed bool, module_publishing_allowed bool, instruction_schedule []byte, native_schedule []byte, global_memory_per_byte_cost uint64, global_memory_per_byte_write_cost uint64, min_transaction_gas_units uint64, large_transaction_cutoff uint64, instrinsic_gas_per_byte uint64, maximum_number_of_gas_units uint64, min_price_per_gas_unit uint64, max_price_per_gas_unit uint64, max_transaction_size_in_bytes uint64, gas_unit_scaling_factor uint64, default_account_size uint64, voting_delay uint64, voting_period uint64, voting_quorum_rate uint8, min_action_delay uint64, transaction_timeout uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "Genesis"},
			Function: "initialize",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{encode_u64_argument(stdlib_version), encode_u64_argument(reward_delay), encode_u128_argument(pre_mine_stc_amount), encode_u128_argument(time_mint_stc_amount), encode_u64_argument(time_mint_stc_period), encode_u8vector_argument(parent_hash), encode_u8vector_argument(association_auth_key), encode_u8vector_argument(genesis_auth_key), encode_u8_argument(chain_id), encode_u64_argument(genesis_timestamp), encode_u64_argument(uncle_rate_target), encode_u64_argument(epoch_block_count), encode_u64_argument(base_block_time_target), encode_u64_argument(base_block_difficulty_window), encode_u128_argument(base_reward_per_block), encode_u64_argument(base_reward_per_uncle_percent), encode_u64_argument(min_block_time_target), encode_u64_argument(max_block_time_target), encode_u64_argument(base_max_uncles_per_block), encode_u64_argument(base_block_gas_limit), encode_u8_argument(strategy), encode_bool_argument(script_allowed), encode_bool_argument(module_publishing_allowed), encode_u8vector_argument(instruction_schedule), encode_u8vector_argument(native_schedule), encode_u64_argument(global_memory_per_byte_cost), encode_u64_argument(global_memory_per_byte_write_cost), encode_u64_argument(min_transaction_gas_units), encode_u64_argument(large_transaction_cutoff), encode_u64_argument(instrinsic_gas_per_byte), encode_u64_argument(maximum_number_of_gas_units), encode_u64_argument(min_price_per_gas_unit), encode_u64_argument(max_price_per_gas_unit), encode_u64_argument(max_transaction_size_in_bytes), encode_u64_argument(gas_unit_scaling_factor), encode_u64_argument(default_account_size), encode_u64_argument(voting_delay), encode_u64_argument(voting_period), encode_u8_argument(voting_quorum_rate), encode_u64_argument(min_action_delay), encode_u64_argument(transaction_timeout)},
		},
	}
}

func EncodeInitializeV2ScriptFunction(stdlib_version uint64, reward_delay uint64, total_stc_amount serde.Uint128, pre_mine_stc_amount serde.Uint128, time_mint_stc_amount serde.Uint128, time_mint_stc_period uint64, parent_hash []byte, association_auth_key []byte, genesis_auth_key []byte, chain_id uint8, genesis_timestamp uint64, uncle_rate_target uint64, epoch_block_count uint64, base_block_time_target uint64, base_block_difficulty_window uint64, base_reward_per_block serde.Uint128, base_reward_per_uncle_percent uint64, min_block_time_target uint64, max_block_time_target uint64, base_max_uncles_per_block uint64, base_block_gas_limit uint64, strategy uint8, script_allowed bool, module_publishing_allowed bool, instruction_schedule []byte, native_schedule []byte, global_memory_per_byte_cost uint64, global_memory_per_byte_write_cost uint64, min_transaction_gas_units uint64, large_transaction_cutoff uint64, instrinsic_gas_per_byte uint64, maximum_number_of_gas_units uint64, min_price_per_gas_unit uint64, max_price_per_gas_unit uint64, max_transaction_size_in_bytes uint64, gas_unit_scaling_factor uint64, default_account_size uint64, voting_delay uint64, voting_period uint64, voting_quorum_rate uint8, min_action_delay uint64, transaction_timeout uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "Genesis"},
			Function: "initialize_v2",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{encode_u64_argument(stdlib_version), encode_u64_argument(reward_delay), encode_u128_argument(total_stc_amount), encode_u128_argument(pre_mine_stc_amount), encode_u128_argument(time_mint_stc_amount), encode_u64_argument(time_mint_stc_period), encode_u8vector_argument(parent_hash), encode_u8vector_argument(association_auth_key), encode_u8vector_argument(genesis_auth_key), encode_u8_argument(chain_id), encode_u64_argument(genesis_timestamp), encode_u64_argument(uncle_rate_target), encode_u64_argument(epoch_block_count), encode_u64_argument(base_block_time_target), encode_u64_argument(base_block_difficulty_window), encode_u128_argument(base_reward_per_block), encode_u64_argument(base_reward_per_uncle_percent), encode_u64_argument(min_block_time_target), encode_u64_argument(max_block_time_target), encode_u64_argument(base_max_uncles_per_block), encode_u64_argument(base_block_gas_limit), encode_u8_argument(strategy), encode_bool_argument(script_allowed), encode_bool_argument(module_publishing_allowed), encode_u8vector_argument(instruction_schedule), encode_u8vector_argument(native_schedule), encode_u64_argument(global_memory_per_byte_cost), encode_u64_argument(global_memory_per_byte_write_cost), encode_u64_argument(min_transaction_gas_units), encode_u64_argument(large_transaction_cutoff), encode_u64_argument(instrinsic_gas_per_byte), encode_u64_argument(maximum_number_of_gas_units), encode_u64_argument(min_price_per_gas_unit), encode_u64_argument(max_price_per_gas_unit), encode_u64_argument(max_transaction_size_in_bytes), encode_u64_argument(gas_unit_scaling_factor), encode_u64_argument(default_account_size), encode_u64_argument(voting_delay), encode_u64_argument(voting_period), encode_u8_argument(voting_quorum_rate), encode_u64_argument(min_action_delay), encode_u64_argument(transaction_timeout)},
		},
	}
}

// Mint a GenesisNFT
func EncodeMintScriptFunction(index uint64, merkle_proof diemtypes.VecBytes) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "GenesisNFTScripts"},
			Function: "mint",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{encode_u64_argument(index), encode_vecbytes_argument(merkle_proof)},
		},
	}
}

func EncodeDummyTokenMintScriptFunction(amount serde.Uint128) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "DummyTokenScripts"},
			Function: "mint",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{encode_u128_argument(amount)},
		},
	}
}

func EncodePeerToPeerScriptFunction(token_type diemtypes.TypeTag, payee diemtypes.AccountAddress, _payee_auth_key []byte, amount serde.Uint128) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "TransferScripts"},
			Function: "peer_to_peer",
			TyArgs:   []diemtypes.TypeTag{token_type},
			Args:     [][]byte{encode_address_argument(payee), encode_u8vector_argument(_payee_auth_key), encode_u128_argument(amount)},
		},
	}
}

func EncodePeerToPeerBatchScriptFunction(token_type diemtypes.TypeTag, _payeees []byte, _payee_auth_keys []byte, _amount serde.Uint128) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "TransferScripts"},
			Function: "peer_to_peer_batch",
			TyArgs:   []diemtypes.TypeTag{token_type},
			Args:     [][]byte{encode_u8vector_argument(_payeees), encode_u8vector_argument(_payee_auth_keys), encode_u128_argument(_amount)},
		},
	}
}

func EncodePeerToPeerV2ScriptFunction(token_type diemtypes.TypeTag, payee diemtypes.AccountAddress, amount serde.Uint128) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "TransferScripts"},
			Function: "peer_to_peer_v2",
			TyArgs:   []diemtypes.TypeTag{token_type},
			Args:     [][]byte{encode_address_argument(payee), encode_u128_argument(amount)},
		},
	}
}

func EncodePeerToPeerWithMetadataScriptFunction(token_type diemtypes.TypeTag, payee diemtypes.AccountAddress, _payee_auth_key []byte, amount serde.Uint128, metadata []byte) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "TransferScripts"},
			Function: "peer_to_peer_with_metadata",
			TyArgs:   []diemtypes.TypeTag{token_type},
			Args:     [][]byte{encode_address_argument(payee), encode_u8vector_argument(_payee_auth_key), encode_u128_argument(amount), encode_u8vector_argument(metadata)},
		},
	}
}

func EncodePeerToPeerWithMetadataV2ScriptFunction(token_type diemtypes.TypeTag, payee diemtypes.AccountAddress, amount serde.Uint128, metadata []byte) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "TransferScripts"},
			Function: "peer_to_peer_with_metadata_v2",
			TyArgs:   []diemtypes.TypeTag{token_type},
			Args:     [][]byte{encode_address_argument(payee), encode_u128_argument(amount), encode_u8vector_argument(metadata)},
		},
	}
}

// Entrypoint for the proposal.
func EncodeProposeScriptFunction(token_t diemtypes.TypeTag, voting_delay uint64, voting_period uint64, voting_quorum_rate uint8, min_action_delay uint64, exec_delay uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "ModifyDaoConfigProposal"},
			Function: "propose",
			TyArgs:   []diemtypes.TypeTag{token_t},
			Args:     [][]byte{encode_u64_argument(voting_delay), encode_u64_argument(voting_period), encode_u8_argument(voting_quorum_rate), encode_u64_argument(min_action_delay), encode_u64_argument(exec_delay)},
		},
	}
}

func EncodeProposeModuleUpgradeV2ScriptFunction(token diemtypes.TypeTag, module_address diemtypes.AccountAddress, package_hash []byte, version uint64, exec_delay uint64, enforced bool) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "ModuleUpgradeScripts"},
			Function: "propose_module_upgrade_v2",
			TyArgs:   []diemtypes.TypeTag{token},
			Args:     [][]byte{encode_address_argument(module_address), encode_u8vector_argument(package_hash), encode_u64_argument(version), encode_u64_argument(exec_delay), encode_bool_argument(enforced)},
		},
	}
}

func EncodeProposeUpdateConsensusConfigScriptFunction(uncle_rate_target uint64, base_block_time_target uint64, base_reward_per_block serde.Uint128, base_reward_per_uncle_percent uint64, epoch_block_count uint64, base_block_difficulty_window uint64, min_block_time_target uint64, max_block_time_target uint64, base_max_uncles_per_block uint64, base_block_gas_limit uint64, strategy uint8, exec_delay uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "OnChainConfigScripts"},
			Function: "propose_update_consensus_config",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{encode_u64_argument(uncle_rate_target), encode_u64_argument(base_block_time_target), encode_u128_argument(base_reward_per_block), encode_u64_argument(base_reward_per_uncle_percent), encode_u64_argument(epoch_block_count), encode_u64_argument(base_block_difficulty_window), encode_u64_argument(min_block_time_target), encode_u64_argument(max_block_time_target), encode_u64_argument(base_max_uncles_per_block), encode_u64_argument(base_block_gas_limit), encode_u8_argument(strategy), encode_u64_argument(exec_delay)},
		},
	}
}

func EncodeProposeUpdateMoveLanguageVersionScriptFunction(new_version uint64, exec_delay uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "OnChainConfigScripts"},
			Function: "propose_update_move_language_version",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{encode_u64_argument(new_version), encode_u64_argument(exec_delay)},
		},
	}
}

func EncodeProposeUpdateRewardConfigScriptFunction(reward_delay uint64, exec_delay uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "OnChainConfigScripts"},
			Function: "propose_update_reward_config",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{encode_u64_argument(reward_delay), encode_u64_argument(exec_delay)},
		},
	}
}

func EncodeProposeUpdateTxnPublishOptionScriptFunction(script_allowed bool, module_publishing_allowed bool, exec_delay uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "OnChainConfigScripts"},
			Function: "propose_update_txn_publish_option",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{encode_bool_argument(script_allowed), encode_bool_argument(module_publishing_allowed), encode_u64_argument(exec_delay)},
		},
	}
}

func EncodeProposeUpdateTxnTimeoutConfigScriptFunction(duration_seconds uint64, exec_delay uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "OnChainConfigScripts"},
			Function: "propose_update_txn_timeout_config",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{encode_u64_argument(duration_seconds), encode_u64_argument(exec_delay)},
		},
	}
}

func EncodeProposeUpdateVmConfigScriptFunction(instruction_schedule []byte, native_schedule []byte, global_memory_per_byte_cost uint64, global_memory_per_byte_write_cost uint64, min_transaction_gas_units uint64, large_transaction_cutoff uint64, instrinsic_gas_per_byte uint64, maximum_number_of_gas_units uint64, min_price_per_gas_unit uint64, max_price_per_gas_unit uint64, max_transaction_size_in_bytes uint64, gas_unit_scaling_factor uint64, default_account_size uint64, exec_delay uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "OnChainConfigScripts"},
			Function: "propose_update_vm_config",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{encode_u8vector_argument(instruction_schedule), encode_u8vector_argument(native_schedule), encode_u64_argument(global_memory_per_byte_cost), encode_u64_argument(global_memory_per_byte_write_cost), encode_u64_argument(min_transaction_gas_units), encode_u64_argument(large_transaction_cutoff), encode_u64_argument(instrinsic_gas_per_byte), encode_u64_argument(maximum_number_of_gas_units), encode_u64_argument(min_price_per_gas_unit), encode_u64_argument(max_price_per_gas_unit), encode_u64_argument(max_transaction_size_in_bytes), encode_u64_argument(gas_unit_scaling_factor), encode_u64_argument(default_account_size), encode_u64_argument(exec_delay)},
		},
	}
}

func EncodeProposeWithdrawScriptFunction(token_t diemtypes.TypeTag, receiver diemtypes.AccountAddress, amount serde.Uint128, period uint64, exec_delay uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "TreasuryScripts"},
			Function: "propose_withdraw",
			TyArgs:   []diemtypes.TypeTag{token_t},
			Args:     [][]byte{encode_address_argument(receiver), encode_u128_argument(amount), encode_u64_argument(period), encode_u64_argument(exec_delay)},
		},
	}
}

// queue agreed proposal to execute.
func EncodeQueueProposalActionScriptFunction(token_t diemtypes.TypeTag, action_t diemtypes.TypeTag, proposer_address diemtypes.AccountAddress, proposal_id uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "Dao"},
			Function: "queue_proposal_action",
			TyArgs:   []diemtypes.TypeTag{token_t, action_t},
			Args:     [][]byte{encode_address_argument(proposer_address), encode_u64_argument(proposal_id)},
		},
	}
}

func EncodeRegisterOracleScriptFunction(oracle_t diemtypes.TypeTag, precision uint8) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "PriceOracleScripts"},
			Function: "register_oracle",
			TyArgs:   []diemtypes.TypeTag{oracle_t},
			Args:     [][]byte{encode_u8_argument(precision)},
		},
	}
}

// revoke all votes on a proposal
func EncodeRevokeVoteScriptFunction(token diemtypes.TypeTag, action diemtypes.TypeTag, proposer_address diemtypes.AccountAddress, proposal_id uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "DaoVoteScripts"},
			Function: "revoke_vote",
			TyArgs:   []diemtypes.TypeTag{token, action},
			Args:     [][]byte{encode_address_argument(proposer_address), encode_u64_argument(proposal_id)},
		},
	}
}

// revoke some votes on a proposal
func EncodeRevokeVoteOfPowerScriptFunction(token diemtypes.TypeTag, action diemtypes.TypeTag, proposer_address diemtypes.AccountAddress, proposal_id uint64, power serde.Uint128) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "DaoVoteScripts"},
			Function: "revoke_vote_of_power",
			TyArgs:   []diemtypes.TypeTag{token, action},
			Args:     [][]byte{encode_address_argument(proposer_address), encode_u64_argument(proposal_id), encode_u128_argument(power)},
		},
	}
}

func EncodeRotateAuthenticationKeyScriptFunction(new_key []byte) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "Account"},
			Function: "rotate_authentication_key",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{encode_u8vector_argument(new_key)},
		},
	}
}

// a alias of execute_module_upgrade_plan_propose, will deprecated in the future.
func EncodeSubmitModuleUpgradePlanScriptFunction(token diemtypes.TypeTag, proposer_address diemtypes.AccountAddress, proposal_id uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "ModuleUpgradeScripts"},
			Function: "submit_module_upgrade_plan",
			TyArgs:   []diemtypes.TypeTag{token},
			Args:     [][]byte{encode_address_argument(proposer_address), encode_u64_argument(proposal_id)},
		},
	}
}

// Directly submit a upgrade plan, the `sender`'s module upgrade plan must been PackageTxnManager::STRATEGY_TWO_PHASE and have UpgradePlanCapability
func EncodeSubmitUpgradePlanScriptFunction(package_hash []byte, version uint64, enforced bool) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "ModuleUpgradeScripts"},
			Function: "submit_upgrade_plan",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{encode_u8vector_argument(package_hash), encode_u64_argument(version), encode_bool_argument(enforced)},
		},
	}
}

// association account should call this script after upgrade from v2 to v3.
func EncodeTakeLinearWithdrawCapabilityScriptFunction() diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "StdlibUpgradeScripts"},
			Function: "take_linear_withdraw_capability",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{},
		},
	}
}

// Take Offer and put to signer's Collection<Offered>.
func EncodeTakeOfferScriptFunction(offered diemtypes.TypeTag, offer_address diemtypes.AccountAddress) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "Offer"},
			Function: "take_offer",
			TyArgs:   []diemtypes.TypeTag{offered},
			Args:     [][]byte{encode_address_argument(offer_address)},
		},
	}
}

// Transfer NFT<NFTMeta, NFTBody> with `id` from `sender` to `receiver`
func EncodeTransferScriptFunction(nft_meta diemtypes.TypeTag, nft_body diemtypes.TypeTag, id uint64, receiver diemtypes.AccountAddress) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "NFTGalleryScripts"},
			Function: "transfer",
			TyArgs:   []diemtypes.TypeTag{nft_meta, nft_body},
			Args:     [][]byte{encode_u64_argument(id), encode_address_argument(receiver)},
		},
	}
}

func EncodeUnstakeVoteScriptFunction(token diemtypes.TypeTag, action diemtypes.TypeTag, proposer_address diemtypes.AccountAddress, proposal_id uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "DaoVoteScripts"},
			Function: "unstake_vote",
			TyArgs:   []diemtypes.TypeTag{token, action},
			Args:     [][]byte{encode_address_argument(proposer_address), encode_u64_argument(proposal_id)},
		},
	}
}

func EncodeUpdateScriptFunction(oracle_t diemtypes.TypeTag, value serde.Uint128) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "PriceOracleScripts"},
			Function: "update",
			TyArgs:   []diemtypes.TypeTag{oracle_t},
			Args:     [][]byte{encode_u128_argument(value)},
		},
	}
}

// Update `sender`'s module upgrade strategy to `strategy`
func EncodeUpdateModuleUpgradeStrategyScriptFunction(strategy uint8) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "ModuleUpgradeScripts"},
			Function: "update_module_upgrade_strategy",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{encode_u8_argument(strategy)},
		},
	}
}

// Stdlib upgrade script from v2 to v3
func EncodeUpgradeFromV2ToV3ScriptFunction(total_stc_amount serde.Uint128) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "StdlibUpgradeScripts"},
			Function: "upgrade_from_v2_to_v3",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{encode_u128_argument(total_stc_amount)},
		},
	}
}

func EncodeUpgradeFromV5ToV6ScriptFunction() diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "StdlibUpgradeScripts"},
			Function: "upgrade_from_v5_to_v6",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{},
		},
	}
}

func EncodeUpgradeFromV6ToV7ScriptFunction() diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "StdlibUpgradeScripts"},
			Function: "upgrade_from_v6_to_v7",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{},
		},
	}
}

func EncodeUpgradeFromV7ToV8ScriptFunction() diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "StdlibUpgradeScripts"},
			Function: "upgrade_from_v7_to_v8",
			TyArgs:   []diemtypes.TypeTag{},
			Args:     [][]byte{},
		},
	}
}

func EncodeWithdrawAndSplitLtWithdrawCapScriptFunction(token_t diemtypes.TypeTag, for_address diemtypes.AccountAddress, amount serde.Uint128, lock_period uint64) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "TreasuryScripts"},
			Function: "withdraw_and_split_lt_withdraw_cap",
			TyArgs:   []diemtypes.TypeTag{token_t},
			Args:     [][]byte{encode_address_argument(for_address), encode_u128_argument(amount), encode_u64_argument(lock_period)},
		},
	}
}

func EncodeWithdrawTokenWithLinearWithdrawCapabilityScriptFunction(token_t diemtypes.TypeTag) diemtypes.TransactionPayload {
	return &diemtypes.TransactionPayload__ScriptFunction{
		diemtypes.ScriptFunction{
			Module:   diemtypes.ModuleId{Address: [16]uint8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, Name: "TreasuryScripts"},
			Function: "withdraw_token_with_linear_withdraw_capability",
			TyArgs:   []diemtypes.TypeTag{token_t},
			Args:     [][]byte{},
		},
	}
}

func decode_accept_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 2 {
			return nil, fmt.Errorf("Was expecting 2 type arguments")
		}
		if len(script.Value.Args) < 0 {
			return nil, fmt.Errorf("Was expecting 0 regular arguments")
		}
		var call ScriptFunctionCall__Accept
		call.NftMeta = script.Value.TyArgs[0]
		call.NftBody = script.Value.TyArgs[1]
		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_accept_token_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 0 {
			return nil, fmt.Errorf("Was expecting 0 regular arguments")
		}
		var call ScriptFunctionCall__AcceptToken
		call.TokenType = script.Value.TyArgs[0]
		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_batch_peer_to_peer_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 3 {
			return nil, fmt.Errorf("Was expecting 3 regular arguments")
		}
		var call ScriptFunctionCall__BatchPeerToPeer
		call.TokenType = script.Value.TyArgs[0]

		if val, err := diemtypes.BcsDeserializeVecAccountAddress(script.Value.Args[0]); err == nil {
			call.Payeees = val
		} else {
			return nil, err
		}

		if val, err := diemtypes.BcsDeserializeVecBytes(script.Value.Args[1]); err == nil {
			call.PayeeAuthKeys = val
		} else {
			return nil, err
		}

		if val, err := diemtypes.BcsDeserializeVecU128(script.Value.Args[2]); err == nil {
			call.Amounts = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_batch_peer_to_peer_v2_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 2 {
			return nil, fmt.Errorf("Was expecting 2 regular arguments")
		}
		var call ScriptFunctionCall__BatchPeerToPeerV2
		call.TokenType = script.Value.TyArgs[0]

		if val, err := diemtypes.BcsDeserializeVecAccountAddress(script.Value.Args[0]); err == nil {
			call.Payeees = val
		} else {
			return nil, err
		}

		if val, err := diemtypes.BcsDeserializeVecU128(script.Value.Args[1]); err == nil {
			call.Amounts = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_cancel_upgrade_plan_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 0 {
			return nil, fmt.Errorf("Was expecting 0 regular arguments")
		}
		var call ScriptFunctionCall__CancelUpgradePlan
		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_cast_vote_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 2 {
			return nil, fmt.Errorf("Was expecting 2 type arguments")
		}
		if len(script.Value.Args) < 4 {
			return nil, fmt.Errorf("Was expecting 4 regular arguments")
		}
		var call ScriptFunctionCall__CastVote
		call.Token = script.Value.TyArgs[0]
		call.ActionT = script.Value.TyArgs[1]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.ProposerAddress = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU64(); err == nil {
			call.ProposalId = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[2]).DeserializeBool(); err == nil {
			call.Agree = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[3]).DeserializeU128(); err == nil {
			call.Votes = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_convert_TwoPhaseUpgrade_to_TwoPhaseUpgradeV2_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 1 {
			return nil, fmt.Errorf("Was expecting 1 regular arguments")
		}
		var call ScriptFunctionCall__ConvertTwoPhaseUpgradeToTwoPhaseUpgradeV2

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.PackageAddress = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_create_account_with_initial_amount_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 3 {
			return nil, fmt.Errorf("Was expecting 3 regular arguments")
		}
		var call ScriptFunctionCall__CreateAccountWithInitialAmount
		call.TokenType = script.Value.TyArgs[0]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.FreshAddress = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeBytes(); err == nil {
			call.AuthKey = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[2]).DeserializeU128(); err == nil {
			call.InitialAmount = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_create_account_with_initial_amount_v2_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 2 {
			return nil, fmt.Errorf("Was expecting 2 regular arguments")
		}
		var call ScriptFunctionCall__CreateAccountWithInitialAmountV2
		call.TokenType = script.Value.TyArgs[0]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.FreshAddress = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU128(); err == nil {
			call.InitialAmount = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_destroy_empty_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 2 {
			return nil, fmt.Errorf("Was expecting 2 type arguments")
		}
		if len(script.Value.Args) < 0 {
			return nil, fmt.Errorf("Was expecting 0 regular arguments")
		}
		var call ScriptFunctionCall__DestroyEmpty
		call.NftMeta = script.Value.TyArgs[0]
		call.NftBody = script.Value.TyArgs[1]
		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_destroy_terminated_proposal_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 2 {
			return nil, fmt.Errorf("Was expecting 2 type arguments")
		}
		if len(script.Value.Args) < 2 {
			return nil, fmt.Errorf("Was expecting 2 regular arguments")
		}
		var call ScriptFunctionCall__DestroyTerminatedProposal
		call.TokenT = script.Value.TyArgs[0]
		call.ActionT = script.Value.TyArgs[1]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.ProposerAddress = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU64(); err == nil {
			call.ProposalId = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_disable_auto_accept_token_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 0 {
			return nil, fmt.Errorf("Was expecting 0 regular arguments")
		}
		var call ScriptFunctionCall__DisableAutoAcceptToken
		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_empty_script_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 0 {
			return nil, fmt.Errorf("Was expecting 0 regular arguments")
		}
		var call ScriptFunctionCall__EmptyScript
		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_enable_auto_accept_token_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 0 {
			return nil, fmt.Errorf("Was expecting 0 regular arguments")
		}
		var call ScriptFunctionCall__EnableAutoAcceptToken
		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_execute_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 2 {
			return nil, fmt.Errorf("Was expecting 2 regular arguments")
		}
		var call ScriptFunctionCall__Execute
		call.TokenT = script.Value.TyArgs[0]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.ProposerAddress = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU64(); err == nil {
			call.ProposalId = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_execute_module_upgrade_plan_propose_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 2 {
			return nil, fmt.Errorf("Was expecting 2 regular arguments")
		}
		var call ScriptFunctionCall__ExecuteModuleUpgradePlanPropose
		call.Token = script.Value.TyArgs[0]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.ProposerAddress = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU64(); err == nil {
			call.ProposalId = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_execute_on_chain_config_proposal_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 1 {
			return nil, fmt.Errorf("Was expecting 1 regular arguments")
		}
		var call ScriptFunctionCall__ExecuteOnChainConfigProposal
		call.ConfigT = script.Value.TyArgs[0]

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeU64(); err == nil {
			call.ProposalId = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_execute_on_chain_config_proposal_v2_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 2 {
			return nil, fmt.Errorf("Was expecting 2 type arguments")
		}
		if len(script.Value.Args) < 2 {
			return nil, fmt.Errorf("Was expecting 2 regular arguments")
		}
		var call ScriptFunctionCall__ExecuteOnChainConfigProposalV2
		call.TokenType = script.Value.TyArgs[0]
		call.ConfigT = script.Value.TyArgs[1]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.ProposerAddress = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU64(); err == nil {
			call.ProposalId = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_execute_withdraw_proposal_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 2 {
			return nil, fmt.Errorf("Was expecting 2 regular arguments")
		}
		var call ScriptFunctionCall__ExecuteWithdrawProposal
		call.TokenT = script.Value.TyArgs[0]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.ProposerAddress = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU64(); err == nil {
			call.ProposalId = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_flip_vote_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 2 {
			return nil, fmt.Errorf("Was expecting 2 type arguments")
		}
		if len(script.Value.Args) < 2 {
			return nil, fmt.Errorf("Was expecting 2 regular arguments")
		}
		var call ScriptFunctionCall__FlipVote
		call.TokenT = script.Value.TyArgs[0]
		call.ActionT = script.Value.TyArgs[1]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.ProposerAddress = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU64(); err == nil {
			call.ProposalId = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_init_data_source_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 1 {
			return nil, fmt.Errorf("Was expecting 1 regular arguments")
		}
		var call ScriptFunctionCall__InitDataSource
		call.OracleT = script.Value.TyArgs[0]

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeU128(); err == nil {
			call.InitValue = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_initialize_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 41 {
			return nil, fmt.Errorf("Was expecting 41 regular arguments")
		}
		var call ScriptFunctionCall__Initialize

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeU64(); err == nil {
			call.StdlibVersion = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU64(); err == nil {
			call.RewardDelay = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[2]).DeserializeU128(); err == nil {
			call.PreMineStcAmount = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[3]).DeserializeU128(); err == nil {
			call.TimeMintStcAmount = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[4]).DeserializeU64(); err == nil {
			call.TimeMintStcPeriod = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[5]).DeserializeBytes(); err == nil {
			call.ParentHash = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[6]).DeserializeBytes(); err == nil {
			call.AssociationAuthKey = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[7]).DeserializeBytes(); err == nil {
			call.GenesisAuthKey = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[8]).DeserializeU8(); err == nil {
			call.ChainId = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[9]).DeserializeU64(); err == nil {
			call.GenesisTimestamp = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[10]).DeserializeU64(); err == nil {
			call.UncleRateTarget = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[11]).DeserializeU64(); err == nil {
			call.EpochBlockCount = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[12]).DeserializeU64(); err == nil {
			call.BaseBlockTimeTarget = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[13]).DeserializeU64(); err == nil {
			call.BaseBlockDifficultyWindow = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[14]).DeserializeU128(); err == nil {
			call.BaseRewardPerBlock = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[15]).DeserializeU64(); err == nil {
			call.BaseRewardPerUnclePercent = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[16]).DeserializeU64(); err == nil {
			call.MinBlockTimeTarget = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[17]).DeserializeU64(); err == nil {
			call.MaxBlockTimeTarget = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[18]).DeserializeU64(); err == nil {
			call.BaseMaxUnclesPerBlock = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[19]).DeserializeU64(); err == nil {
			call.BaseBlockGasLimit = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[20]).DeserializeU8(); err == nil {
			call.Strategy = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[21]).DeserializeBool(); err == nil {
			call.ScriptAllowed = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[22]).DeserializeBool(); err == nil {
			call.ModulePublishingAllowed = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[23]).DeserializeBytes(); err == nil {
			call.InstructionSchedule = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[24]).DeserializeBytes(); err == nil {
			call.NativeSchedule = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[25]).DeserializeU64(); err == nil {
			call.GlobalMemoryPerByteCost = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[26]).DeserializeU64(); err == nil {
			call.GlobalMemoryPerByteWriteCost = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[27]).DeserializeU64(); err == nil {
			call.MinTransactionGasUnits = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[28]).DeserializeU64(); err == nil {
			call.LargeTransactionCutoff = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[29]).DeserializeU64(); err == nil {
			call.InstrinsicGasPerByte = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[30]).DeserializeU64(); err == nil {
			call.MaximumNumberOfGasUnits = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[31]).DeserializeU64(); err == nil {
			call.MinPricePerGasUnit = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[32]).DeserializeU64(); err == nil {
			call.MaxPricePerGasUnit = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[33]).DeserializeU64(); err == nil {
			call.MaxTransactionSizeInBytes = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[34]).DeserializeU64(); err == nil {
			call.GasUnitScalingFactor = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[35]).DeserializeU64(); err == nil {
			call.DefaultAccountSize = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[36]).DeserializeU64(); err == nil {
			call.VotingDelay = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[37]).DeserializeU64(); err == nil {
			call.VotingPeriod = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[38]).DeserializeU8(); err == nil {
			call.VotingQuorumRate = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[39]).DeserializeU64(); err == nil {
			call.MinActionDelay = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[40]).DeserializeU64(); err == nil {
			call.TransactionTimeout = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_initialize_v2_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 42 {
			return nil, fmt.Errorf("Was expecting 42 regular arguments")
		}
		var call ScriptFunctionCall__InitializeV2

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeU64(); err == nil {
			call.StdlibVersion = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU64(); err == nil {
			call.RewardDelay = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[2]).DeserializeU128(); err == nil {
			call.TotalStcAmount = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[3]).DeserializeU128(); err == nil {
			call.PreMineStcAmount = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[4]).DeserializeU128(); err == nil {
			call.TimeMintStcAmount = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[5]).DeserializeU64(); err == nil {
			call.TimeMintStcPeriod = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[6]).DeserializeBytes(); err == nil {
			call.ParentHash = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[7]).DeserializeBytes(); err == nil {
			call.AssociationAuthKey = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[8]).DeserializeBytes(); err == nil {
			call.GenesisAuthKey = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[9]).DeserializeU8(); err == nil {
			call.ChainId = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[10]).DeserializeU64(); err == nil {
			call.GenesisTimestamp = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[11]).DeserializeU64(); err == nil {
			call.UncleRateTarget = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[12]).DeserializeU64(); err == nil {
			call.EpochBlockCount = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[13]).DeserializeU64(); err == nil {
			call.BaseBlockTimeTarget = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[14]).DeserializeU64(); err == nil {
			call.BaseBlockDifficultyWindow = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[15]).DeserializeU128(); err == nil {
			call.BaseRewardPerBlock = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[16]).DeserializeU64(); err == nil {
			call.BaseRewardPerUnclePercent = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[17]).DeserializeU64(); err == nil {
			call.MinBlockTimeTarget = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[18]).DeserializeU64(); err == nil {
			call.MaxBlockTimeTarget = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[19]).DeserializeU64(); err == nil {
			call.BaseMaxUnclesPerBlock = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[20]).DeserializeU64(); err == nil {
			call.BaseBlockGasLimit = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[21]).DeserializeU8(); err == nil {
			call.Strategy = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[22]).DeserializeBool(); err == nil {
			call.ScriptAllowed = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[23]).DeserializeBool(); err == nil {
			call.ModulePublishingAllowed = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[24]).DeserializeBytes(); err == nil {
			call.InstructionSchedule = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[25]).DeserializeBytes(); err == nil {
			call.NativeSchedule = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[26]).DeserializeU64(); err == nil {
			call.GlobalMemoryPerByteCost = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[27]).DeserializeU64(); err == nil {
			call.GlobalMemoryPerByteWriteCost = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[28]).DeserializeU64(); err == nil {
			call.MinTransactionGasUnits = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[29]).DeserializeU64(); err == nil {
			call.LargeTransactionCutoff = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[30]).DeserializeU64(); err == nil {
			call.InstrinsicGasPerByte = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[31]).DeserializeU64(); err == nil {
			call.MaximumNumberOfGasUnits = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[32]).DeserializeU64(); err == nil {
			call.MinPricePerGasUnit = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[33]).DeserializeU64(); err == nil {
			call.MaxPricePerGasUnit = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[34]).DeserializeU64(); err == nil {
			call.MaxTransactionSizeInBytes = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[35]).DeserializeU64(); err == nil {
			call.GasUnitScalingFactor = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[36]).DeserializeU64(); err == nil {
			call.DefaultAccountSize = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[37]).DeserializeU64(); err == nil {
			call.VotingDelay = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[38]).DeserializeU64(); err == nil {
			call.VotingPeriod = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[39]).DeserializeU8(); err == nil {
			call.VotingQuorumRate = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[40]).DeserializeU64(); err == nil {
			call.MinActionDelay = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[41]).DeserializeU64(); err == nil {
			call.TransactionTimeout = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_mint_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 2 {
			return nil, fmt.Errorf("Was expecting 2 regular arguments")
		}
		var call ScriptFunctionCall__Mint

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeU64(); err == nil {
			call.Index = val
		} else {
			return nil, err
		}

		if val, err := diemtypes.BcsDeserializeVecBytes(script.Value.Args[1]); err == nil {
			call.MerkleProof = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_dummytoken_mint_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 1 {
			return nil, fmt.Errorf("Was expecting 1 regular arguments")
		}
		var call ScriptFunctionCall__DummyTokenMint

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeU128(); err == nil {
			call.Amount = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_peer_to_peer_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 3 {
			return nil, fmt.Errorf("Was expecting 3 regular arguments")
		}
		var call ScriptFunctionCall__PeerToPeer
		call.TokenType = script.Value.TyArgs[0]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.Payee = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeBytes(); err == nil {
			call.PayeeAuthKey = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[2]).DeserializeU128(); err == nil {
			call.Amount = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_peer_to_peer_batch_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 3 {
			return nil, fmt.Errorf("Was expecting 3 regular arguments")
		}
		var call ScriptFunctionCall__PeerToPeerBatch
		call.TokenType = script.Value.TyArgs[0]

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeBytes(); err == nil {
			call.Payeees = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeBytes(); err == nil {
			call.PayeeAuthKeys = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[2]).DeserializeU128(); err == nil {
			call.Amount = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_peer_to_peer_v2_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 2 {
			return nil, fmt.Errorf("Was expecting 2 regular arguments")
		}
		var call ScriptFunctionCall__PeerToPeerV2
		call.TokenType = script.Value.TyArgs[0]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.Payee = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU128(); err == nil {
			call.Amount = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_peer_to_peer_with_metadata_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 4 {
			return nil, fmt.Errorf("Was expecting 4 regular arguments")
		}
		var call ScriptFunctionCall__PeerToPeerWithMetadata
		call.TokenType = script.Value.TyArgs[0]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.Payee = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeBytes(); err == nil {
			call.PayeeAuthKey = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[2]).DeserializeU128(); err == nil {
			call.Amount = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[3]).DeserializeBytes(); err == nil {
			call.Metadata = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_peer_to_peer_with_metadata_v2_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 3 {
			return nil, fmt.Errorf("Was expecting 3 regular arguments")
		}
		var call ScriptFunctionCall__PeerToPeerWithMetadataV2
		call.TokenType = script.Value.TyArgs[0]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.Payee = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU128(); err == nil {
			call.Amount = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[2]).DeserializeBytes(); err == nil {
			call.Metadata = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_propose_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 5 {
			return nil, fmt.Errorf("Was expecting 5 regular arguments")
		}
		var call ScriptFunctionCall__Propose
		call.TokenT = script.Value.TyArgs[0]

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeU64(); err == nil {
			call.VotingDelay = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU64(); err == nil {
			call.VotingPeriod = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[2]).DeserializeU8(); err == nil {
			call.VotingQuorumRate = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[3]).DeserializeU64(); err == nil {
			call.MinActionDelay = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[4]).DeserializeU64(); err == nil {
			call.ExecDelay = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_propose_module_upgrade_v2_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 5 {
			return nil, fmt.Errorf("Was expecting 5 regular arguments")
		}
		var call ScriptFunctionCall__ProposeModuleUpgradeV2
		call.Token = script.Value.TyArgs[0]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.ModuleAddress = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeBytes(); err == nil {
			call.PackageHash = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[2]).DeserializeU64(); err == nil {
			call.Version = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[3]).DeserializeU64(); err == nil {
			call.ExecDelay = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[4]).DeserializeBool(); err == nil {
			call.Enforced = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_propose_update_consensus_config_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 12 {
			return nil, fmt.Errorf("Was expecting 12 regular arguments")
		}
		var call ScriptFunctionCall__ProposeUpdateConsensusConfig

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeU64(); err == nil {
			call.UncleRateTarget = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU64(); err == nil {
			call.BaseBlockTimeTarget = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[2]).DeserializeU128(); err == nil {
			call.BaseRewardPerBlock = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[3]).DeserializeU64(); err == nil {
			call.BaseRewardPerUnclePercent = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[4]).DeserializeU64(); err == nil {
			call.EpochBlockCount = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[5]).DeserializeU64(); err == nil {
			call.BaseBlockDifficultyWindow = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[6]).DeserializeU64(); err == nil {
			call.MinBlockTimeTarget = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[7]).DeserializeU64(); err == nil {
			call.MaxBlockTimeTarget = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[8]).DeserializeU64(); err == nil {
			call.BaseMaxUnclesPerBlock = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[9]).DeserializeU64(); err == nil {
			call.BaseBlockGasLimit = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[10]).DeserializeU8(); err == nil {
			call.Strategy = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[11]).DeserializeU64(); err == nil {
			call.ExecDelay = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_propose_update_move_language_version_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 2 {
			return nil, fmt.Errorf("Was expecting 2 regular arguments")
		}
		var call ScriptFunctionCall__ProposeUpdateMoveLanguageVersion

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeU64(); err == nil {
			call.NewVersion = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU64(); err == nil {
			call.ExecDelay = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_propose_update_reward_config_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 2 {
			return nil, fmt.Errorf("Was expecting 2 regular arguments")
		}
		var call ScriptFunctionCall__ProposeUpdateRewardConfig

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeU64(); err == nil {
			call.RewardDelay = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU64(); err == nil {
			call.ExecDelay = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_propose_update_txn_publish_option_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 3 {
			return nil, fmt.Errorf("Was expecting 3 regular arguments")
		}
		var call ScriptFunctionCall__ProposeUpdateTxnPublishOption

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeBool(); err == nil {
			call.ScriptAllowed = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeBool(); err == nil {
			call.ModulePublishingAllowed = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[2]).DeserializeU64(); err == nil {
			call.ExecDelay = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_propose_update_txn_timeout_config_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 2 {
			return nil, fmt.Errorf("Was expecting 2 regular arguments")
		}
		var call ScriptFunctionCall__ProposeUpdateTxnTimeoutConfig

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeU64(); err == nil {
			call.DurationSeconds = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU64(); err == nil {
			call.ExecDelay = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_propose_update_vm_config_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 14 {
			return nil, fmt.Errorf("Was expecting 14 regular arguments")
		}
		var call ScriptFunctionCall__ProposeUpdateVmConfig

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeBytes(); err == nil {
			call.InstructionSchedule = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeBytes(); err == nil {
			call.NativeSchedule = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[2]).DeserializeU64(); err == nil {
			call.GlobalMemoryPerByteCost = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[3]).DeserializeU64(); err == nil {
			call.GlobalMemoryPerByteWriteCost = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[4]).DeserializeU64(); err == nil {
			call.MinTransactionGasUnits = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[5]).DeserializeU64(); err == nil {
			call.LargeTransactionCutoff = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[6]).DeserializeU64(); err == nil {
			call.InstrinsicGasPerByte = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[7]).DeserializeU64(); err == nil {
			call.MaximumNumberOfGasUnits = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[8]).DeserializeU64(); err == nil {
			call.MinPricePerGasUnit = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[9]).DeserializeU64(); err == nil {
			call.MaxPricePerGasUnit = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[10]).DeserializeU64(); err == nil {
			call.MaxTransactionSizeInBytes = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[11]).DeserializeU64(); err == nil {
			call.GasUnitScalingFactor = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[12]).DeserializeU64(); err == nil {
			call.DefaultAccountSize = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[13]).DeserializeU64(); err == nil {
			call.ExecDelay = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_propose_withdraw_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 4 {
			return nil, fmt.Errorf("Was expecting 4 regular arguments")
		}
		var call ScriptFunctionCall__ProposeWithdraw
		call.TokenT = script.Value.TyArgs[0]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.Receiver = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU128(); err == nil {
			call.Amount = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[2]).DeserializeU64(); err == nil {
			call.Period = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[3]).DeserializeU64(); err == nil {
			call.ExecDelay = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_queue_proposal_action_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 2 {
			return nil, fmt.Errorf("Was expecting 2 type arguments")
		}
		if len(script.Value.Args) < 2 {
			return nil, fmt.Errorf("Was expecting 2 regular arguments")
		}
		var call ScriptFunctionCall__QueueProposalAction
		call.TokenT = script.Value.TyArgs[0]
		call.ActionT = script.Value.TyArgs[1]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.ProposerAddress = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU64(); err == nil {
			call.ProposalId = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_register_oracle_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 1 {
			return nil, fmt.Errorf("Was expecting 1 regular arguments")
		}
		var call ScriptFunctionCall__RegisterOracle
		call.OracleT = script.Value.TyArgs[0]

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeU8(); err == nil {
			call.Precision = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_revoke_vote_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 2 {
			return nil, fmt.Errorf("Was expecting 2 type arguments")
		}
		if len(script.Value.Args) < 2 {
			return nil, fmt.Errorf("Was expecting 2 regular arguments")
		}
		var call ScriptFunctionCall__RevokeVote
		call.Token = script.Value.TyArgs[0]
		call.Action = script.Value.TyArgs[1]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.ProposerAddress = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU64(); err == nil {
			call.ProposalId = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_revoke_vote_of_power_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 2 {
			return nil, fmt.Errorf("Was expecting 2 type arguments")
		}
		if len(script.Value.Args) < 3 {
			return nil, fmt.Errorf("Was expecting 3 regular arguments")
		}
		var call ScriptFunctionCall__RevokeVoteOfPower
		call.Token = script.Value.TyArgs[0]
		call.Action = script.Value.TyArgs[1]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.ProposerAddress = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU64(); err == nil {
			call.ProposalId = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[2]).DeserializeU128(); err == nil {
			call.Power = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_rotate_authentication_key_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 1 {
			return nil, fmt.Errorf("Was expecting 1 regular arguments")
		}
		var call ScriptFunctionCall__RotateAuthenticationKey

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeBytes(); err == nil {
			call.NewKey = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_submit_module_upgrade_plan_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 2 {
			return nil, fmt.Errorf("Was expecting 2 regular arguments")
		}
		var call ScriptFunctionCall__SubmitModuleUpgradePlan
		call.Token = script.Value.TyArgs[0]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.ProposerAddress = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU64(); err == nil {
			call.ProposalId = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_submit_upgrade_plan_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 3 {
			return nil, fmt.Errorf("Was expecting 3 regular arguments")
		}
		var call ScriptFunctionCall__SubmitUpgradePlan

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeBytes(); err == nil {
			call.PackageHash = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU64(); err == nil {
			call.Version = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[2]).DeserializeBool(); err == nil {
			call.Enforced = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_take_linear_withdraw_capability_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 0 {
			return nil, fmt.Errorf("Was expecting 0 regular arguments")
		}
		var call ScriptFunctionCall__TakeLinearWithdrawCapability
		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_take_offer_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 1 {
			return nil, fmt.Errorf("Was expecting 1 regular arguments")
		}
		var call ScriptFunctionCall__TakeOffer
		call.Offered = script.Value.TyArgs[0]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.OfferAddress = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_transfer_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 2 {
			return nil, fmt.Errorf("Was expecting 2 type arguments")
		}
		if len(script.Value.Args) < 2 {
			return nil, fmt.Errorf("Was expecting 2 regular arguments")
		}
		var call ScriptFunctionCall__Transfer
		call.NftMeta = script.Value.TyArgs[0]
		call.NftBody = script.Value.TyArgs[1]

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeU64(); err == nil {
			call.Id = val
		} else {
			return nil, err
		}

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[1]); err == nil {
			call.Receiver = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_unstake_vote_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 2 {
			return nil, fmt.Errorf("Was expecting 2 type arguments")
		}
		if len(script.Value.Args) < 2 {
			return nil, fmt.Errorf("Was expecting 2 regular arguments")
		}
		var call ScriptFunctionCall__UnstakeVote
		call.Token = script.Value.TyArgs[0]
		call.Action = script.Value.TyArgs[1]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.ProposerAddress = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU64(); err == nil {
			call.ProposalId = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_update_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 1 {
			return nil, fmt.Errorf("Was expecting 1 regular arguments")
		}
		var call ScriptFunctionCall__Update
		call.OracleT = script.Value.TyArgs[0]

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeU128(); err == nil {
			call.Value = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_update_module_upgrade_strategy_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 1 {
			return nil, fmt.Errorf("Was expecting 1 regular arguments")
		}
		var call ScriptFunctionCall__UpdateModuleUpgradeStrategy

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeU8(); err == nil {
			call.Strategy = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_upgrade_from_v2_to_v3_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 1 {
			return nil, fmt.Errorf("Was expecting 1 regular arguments")
		}
		var call ScriptFunctionCall__UpgradeFromV2ToV3

		if val, err := bcs.NewDeserializer(script.Value.Args[0]).DeserializeU128(); err == nil {
			call.TotalStcAmount = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_upgrade_from_v5_to_v6_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 0 {
			return nil, fmt.Errorf("Was expecting 0 regular arguments")
		}
		var call ScriptFunctionCall__UpgradeFromV5ToV6
		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_upgrade_from_v6_to_v7_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 0 {
			return nil, fmt.Errorf("Was expecting 0 regular arguments")
		}
		var call ScriptFunctionCall__UpgradeFromV6ToV7
		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_upgrade_from_v7_to_v8_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 0 {
			return nil, fmt.Errorf("Was expecting 0 type arguments")
		}
		if len(script.Value.Args) < 0 {
			return nil, fmt.Errorf("Was expecting 0 regular arguments")
		}
		var call ScriptFunctionCall__UpgradeFromV7ToV8
		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_withdraw_and_split_lt_withdraw_cap_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 3 {
			return nil, fmt.Errorf("Was expecting 3 regular arguments")
		}
		var call ScriptFunctionCall__WithdrawAndSplitLtWithdrawCap
		call.TokenT = script.Value.TyArgs[0]

		if val, err := diemtypes.BcsDeserializeAccountAddress(script.Value.Args[0]); err == nil {
			call.ForAddress = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[1]).DeserializeU128(); err == nil {
			call.Amount = val
		} else {
			return nil, err
		}

		if val, err := bcs.NewDeserializer(script.Value.Args[2]).DeserializeU64(); err == nil {
			call.LockPeriod = val
		} else {
			return nil, err
		}

		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

func decode_withdraw_token_with_linear_withdraw_capability_script_function(script diemtypes.TransactionPayload) (ScriptFunctionCall, error) {
	switch script := interface{}(script).(type) {
	case *diemtypes.TransactionPayload__ScriptFunction:
		if len(script.Value.TyArgs) < 1 {
			return nil, fmt.Errorf("Was expecting 1 type arguments")
		}
		if len(script.Value.Args) < 0 {
			return nil, fmt.Errorf("Was expecting 0 regular arguments")
		}
		var call ScriptFunctionCall__WithdrawTokenWithLinearWithdrawCapability
		call.TokenT = script.Value.TyArgs[0]
		return &call, nil
	default:
		return nil, fmt.Errorf("Unexpected TransactionPayload encountered when decoding a script function")
	}
}

var script_decoder_map = map[string]func(*diemtypes.Script) (ScriptCall, error){}

var script_function_decoder_map = map[string]func(diemtypes.TransactionPayload) (ScriptFunctionCall, error){
	"NFTGalleryScriptsaccept":                                       decode_accept_script_function,
	"Accountaccept_token":                                           decode_accept_token_script_function,
	"TransferScriptsbatch_peer_to_peer":                             decode_batch_peer_to_peer_script_function,
	"TransferScriptsbatch_peer_to_peer_v2":                          decode_batch_peer_to_peer_v2_script_function,
	"ModuleUpgradeScriptscancel_upgrade_plan":                       decode_cancel_upgrade_plan_script_function,
	"DaoVoteScriptscast_vote":                                       decode_cast_vote_script_function,
	"PackageTxnManagerconvert_TwoPhaseUpgrade_to_TwoPhaseUpgradeV2": decode_convert_TwoPhaseUpgrade_to_TwoPhaseUpgradeV2_script_function,
	"Accountcreate_account_with_initial_amount":                     decode_create_account_with_initial_amount_script_function,
	"Accountcreate_account_with_initial_amount_v2":                  decode_create_account_with_initial_amount_v2_script_function,
	"IdentifierNFTScriptsdestroy_empty":                             decode_destroy_empty_script_function,
	"Daodestroy_terminated_proposal":                                decode_destroy_terminated_proposal_script_function,
	"AccountScriptsdisable_auto_accept_token":                       decode_disable_auto_accept_token_script_function,
	"EmptyScriptsempty_script":                                      decode_empty_script_script_function,
	"AccountScriptsenable_auto_accept_token":                        decode_enable_auto_accept_token_script_function,
	"ModifyDaoConfigProposalexecute":                                decode_execute_script_function,
	"ModuleUpgradeScriptsexecute_module_upgrade_plan_propose":       decode_execute_module_upgrade_plan_propose_script_function,
	"OnChainConfigScriptsexecute_on_chain_config_proposal":          decode_execute_on_chain_config_proposal_script_function,
	"OnChainConfigScriptsexecute_on_chain_config_proposal_v2":       decode_execute_on_chain_config_proposal_v2_script_function,
	"TreasuryScriptsexecute_withdraw_proposal":                      decode_execute_withdraw_proposal_script_function,
	"DaoVoteScriptsflip_vote":                                       decode_flip_vote_script_function,
	"PriceOracleScriptsinit_data_source":                            decode_init_data_source_script_function,
	"Genesisinitialize":                                             decode_initialize_script_function,
	"Genesisinitialize_v2":                                          decode_initialize_v2_script_function,
	"GenesisNFTScriptsmint":                                         decode_mint_script_function,
	"DummyTokenScriptsmint":                                         decode_mint_script_function,
	"TransferScriptspeer_to_peer":                                   decode_peer_to_peer_script_function,
	"TransferScriptspeer_to_peer_batch":                             decode_peer_to_peer_batch_script_function,
	"TransferScriptspeer_to_peer_v2":                                decode_peer_to_peer_v2_script_function,
	"TransferScriptspeer_to_peer_with_metadata":                     decode_peer_to_peer_with_metadata_script_function,
	"TransferScriptspeer_to_peer_with_metadata_v2":                  decode_peer_to_peer_with_metadata_v2_script_function,
	"ModifyDaoConfigProposalpropose":                                decode_propose_script_function,
	"ModuleUpgradeScriptspropose_module_upgrade_v2":                 decode_propose_module_upgrade_v2_script_function,
	"OnChainConfigScriptspropose_update_consensus_config":           decode_propose_update_consensus_config_script_function,
	"OnChainConfigScriptspropose_update_move_language_version":      decode_propose_update_move_language_version_script_function,
	"OnChainConfigScriptspropose_update_reward_config":              decode_propose_update_reward_config_script_function,
	"OnChainConfigScriptspropose_update_txn_publish_option":         decode_propose_update_txn_publish_option_script_function,
	"OnChainConfigScriptspropose_update_txn_timeout_config":         decode_propose_update_txn_timeout_config_script_function,
	"OnChainConfigScriptspropose_update_vm_config":                  decode_propose_update_vm_config_script_function,
	"TreasuryScriptspropose_withdraw":                               decode_propose_withdraw_script_function,
	"Daoqueue_proposal_action":                                      decode_queue_proposal_action_script_function,
	"PriceOracleScriptsregister_oracle":                             decode_register_oracle_script_function,
	"DaoVoteScriptsrevoke_vote":                                     decode_revoke_vote_script_function,
	"DaoVoteScriptsrevoke_vote_of_power":                            decode_revoke_vote_of_power_script_function,
	"Accountrotate_authentication_key":                              decode_rotate_authentication_key_script_function,
	"ModuleUpgradeScriptssubmit_module_upgrade_plan":                decode_submit_module_upgrade_plan_script_function,
	"ModuleUpgradeScriptssubmit_upgrade_plan":                       decode_submit_upgrade_plan_script_function,
	"StdlibUpgradeScriptstake_linear_withdraw_capability":           decode_take_linear_withdraw_capability_script_function,
	"Offertake_offer":                                               decode_take_offer_script_function,
	"NFTGalleryScriptstransfer":                                     decode_transfer_script_function,
	"DaoVoteScriptsunstake_vote":                                    decode_unstake_vote_script_function,
	"PriceOracleScriptsupdate":                                      decode_update_script_function,
	"ModuleUpgradeScriptsupdate_module_upgrade_strategy":            decode_update_module_upgrade_strategy_script_function,
	"StdlibUpgradeScriptsupgrade_from_v2_to_v3":                     decode_upgrade_from_v2_to_v3_script_function,
	"StdlibUpgradeScriptsupgrade_from_v5_to_v6":                     decode_upgrade_from_v5_to_v6_script_function,
	"StdlibUpgradeScriptsupgrade_from_v6_to_v7":                     decode_upgrade_from_v6_to_v7_script_function,
	"StdlibUpgradeScriptsupgrade_from_v7_to_v8":                     decode_upgrade_from_v7_to_v8_script_function,
	"TreasuryScriptswithdraw_and_split_lt_withdraw_cap":             decode_withdraw_and_split_lt_withdraw_cap_script_function,
	"TreasuryScriptswithdraw_token_with_linear_withdraw_capability": decode_withdraw_token_with_linear_withdraw_capability_script_function,
}

func encode_bool_argument(arg bool) []byte {

	s := bcs.NewSerializer()
	if err := s.SerializeBool(arg); err == nil {
		return s.GetBytes()
	}

	panic("Unable to serialize argument of type bool")
}

func encode_u8_argument(arg uint8) []byte {

	s := bcs.NewSerializer()
	if err := s.SerializeU8(arg); err == nil {
		return s.GetBytes()
	}

	panic("Unable to serialize argument of type u8")
}

func encode_u64_argument(arg uint64) []byte {

	s := bcs.NewSerializer()
	if err := s.SerializeU64(arg); err == nil {
		return s.GetBytes()
	}

	panic("Unable to serialize argument of type u64")
}

func encode_u128_argument(arg serde.Uint128) []byte {

	s := bcs.NewSerializer()
	if err := s.SerializeU128(arg); err == nil {
		return s.GetBytes()
	}

	panic("Unable to serialize argument of type u128")
}

func encode_address_argument(arg diemtypes.AccountAddress) []byte {

	if val, err := arg.BcsSerialize(); err == nil {
		{
			return val
		}
	}

	panic("Unable to serialize argument of type address")
}

func encode_u8vector_argument(arg []byte) []byte {

	s := bcs.NewSerializer()
	if err := s.SerializeBytes(arg); err == nil {
		return s.GetBytes()
	}

	panic("Unable to serialize argument of type u8vector")
}

func encode_vecu128_argument(arg diemtypes.VecU128) []byte {

	if val, err := arg.BcsSerialize(); err == nil {
		{
			return val
		}
	}

	panic("Unable to serialize argument of type vecu128")
}

func encode_vecaccountaddress_argument(arg diemtypes.VecAccountAddress) []byte {

	if val, err := arg.BcsSerialize(); err == nil {
		{
			return val
		}
	}

	panic("Unable to serialize argument of type vecaccountaddress")
}

func encode_vecbytes_argument(arg diemtypes.VecBytes) []byte {

	if val, err := arg.BcsSerialize(); err == nil {
		{
			return val
		}
	}

	panic("Unable to serialize argument of type vecbytes")
}

func decode_bool_argument(arg diemtypes.TransactionArgument) (value bool, err error) {
	if arg, ok := arg.(*diemtypes.TransactionArgument__Bool); ok {
		value = bool(*arg)
	} else {
		err = fmt.Errorf("Was expecting a Bool argument")
	}
	return
}

func decode_u8_argument(arg diemtypes.TransactionArgument) (value uint8, err error) {
	if arg, ok := arg.(*diemtypes.TransactionArgument__U8); ok {
		value = uint8(*arg)
	} else {
		err = fmt.Errorf("Was expecting a U8 argument")
	}
	return
}

func decode_u64_argument(arg diemtypes.TransactionArgument) (value uint64, err error) {
	if arg, ok := arg.(*diemtypes.TransactionArgument__U64); ok {
		value = uint64(*arg)
	} else {
		err = fmt.Errorf("Was expecting a U64 argument")
	}
	return
}

func decode_u128_argument(arg diemtypes.TransactionArgument) (value serde.Uint128, err error) {
	if arg, ok := arg.(*diemtypes.TransactionArgument__U128); ok {
		value = serde.Uint128(*arg)
	} else {
		err = fmt.Errorf("Was expecting a U128 argument")
	}
	return
}

func decode_address_argument(arg diemtypes.TransactionArgument) (value diemtypes.AccountAddress, err error) {
	if arg, ok := arg.(*diemtypes.TransactionArgument__Address); ok {
		value = arg.Value
	} else {
		err = fmt.Errorf("Was expecting a Address argument")
	}
	return
}

func decode_u8vector_argument(arg diemtypes.TransactionArgument) (value []byte, err error) {
	if arg, ok := arg.(*diemtypes.TransactionArgument__U8Vector); ok {
		value = []byte(*arg)
	} else {
		err = fmt.Errorf("Was expecting a U8Vector argument")
	}
	return
}

func decode_vecu128_argument(arg diemtypes.TransactionArgument) (value diemtypes.VecU128, err error) {
	if arg, ok := arg.(*diemtypes.TransactionArgument__VecU128); ok {
		value = arg.Value
	} else {
		err = fmt.Errorf("Was expecting a VecU128 argument")
	}
	return
}

func decode_vecaccountaddress_argument(arg diemtypes.TransactionArgument) (value diemtypes.VecAccountAddress, err error) {
	if arg, ok := arg.(*diemtypes.TransactionArgument__VecAccountAddress); ok {
		value = arg.Value
	} else {
		err = fmt.Errorf("Was expecting a VecAccountAddress argument")
	}
	return
}

func decode_vecbytes_argument(arg diemtypes.TransactionArgument) (value diemtypes.VecBytes, err error) {
	if arg, ok := arg.(*diemtypes.TransactionArgument__VecBytes); ok {
		value = arg.Value
	} else {
		err = fmt.Errorf("Was expecting a VecBytes argument")
	}
	return
}
