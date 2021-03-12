package messari

// GetAllAssetsResp is struct which holds the response shape of the "/assets" api call
type GetAllAssetsResp struct {
	Status Status  `json:"status,omitempty"`
	Data   []Asset `json:"data,omitempty"`
}

// GetAssetResp is struct which holds repsonse shape of "/assets/:key" api call
type GetAssetResp struct {
	Status Status        `json:"status,omitempty"`
	Data   AssetMetaData `json:"data,omitempty"`
}

// GetAssetMetricsResp is struct which holds response shape of "/assets/:key/metrics" api call
type GetAssetMetricsResp struct {
	Status Status               `json:"status,omitempty"`
	Data   AssetMetricsMetadata `json:"data,omitempty"`
}

// AssetMetricsMetadata struct holds Metrics and AssetMetaData fields
type AssetMetricsMetadata struct {
	AssetMetaData
	Metrics
}

// AssetMetaData is struct which holds fields of an Asset's basic metadata
type AssetMetaData struct {
	ID     string `json:"id,omitempty"`
	Symbol string `json:"symbol,omitempty"`
	Name   string `json:"name,omitempty"`
	Slug   string `json:"slug,omitempty"`
}

// Asset is a struct representing a cryptocurrency asset in Messari API
type Asset struct {
	ID                  string  `json:"id,omitempty"`
	Symbol              string  `json:"symbol,omitempty"`
	Name                string  `json:"name,omitempty"`
	Slug                string  `json:"slug,omitempty"`
	InternalTempAgoraID *string `json:"_internal_temp_agora_id,omitempty"`
	Metrics             Metrics `json:"metrics,omitempty"`
	Profile             Profile `json:"profile,omitempty"`
}

// Metrics struct is an Asset's collection of metrics
type Metrics struct {
	MarketData             MarketData          `json:"market_data,omitempty"`
	Marketcap              Marketcap           `json:"marketcap,omitempty"`
	Supply                 MetricsSupply       `json:"supply,omitempty"`
	BlockchainStats24Hours map[string]*float64 `json:"blockchain_stats_24_hours,omitempty"`
	MarketDataLiquidity    MarketDataLiquidity `json:"market_data_liquidity,omitempty"`
	AllTimeHigh            AllTimeHigh         `json:"all_time_high,omitempty"`
	CycleLow               CycleLow            `json:"cycle_low,omitempty"`
	TokenSaleStats         TokenSaleStats      `json:"token_sale_stats,omitempty"`
	StakingStats           StakingStats        `json:"staking_stats,omitempty"`
	MiningStats            MiningStats         `json:"mining_stats,omitempty"`
	DeveloperActivity      DeveloperActivity   `json:"developer_activity,omitempty"`
	RoiData                map[string]*float64 `json:"roi_data,omitempty"`
	RoiByYear              map[string]*float64 `json:"roi_by_year,omitempty"`
	RiskMetrics            RiskMetrics         `json:"risk_metrics,omitempty"`
	MiscData               MiscData            `json:"misc_data,omitempty"`
	LendRates              map[string]float64  `json:"lend_rates,omitempty"`
	BorrowRates            *BorrowRates        `json:"borrow_rates,omitempty"`
	LoanData               LoanData            `json:"loan_data,omitempty"`
	Reddit                 Reddit              `json:"reddit,omitempty"`
	OnChainData            map[string]*float64 `json:"on_chain_data,omitempty"`
	ExchangeFlows          ExchangeFlows       `json:"exchange_flows,omitempty"`
	AlertMessages          interface{}         `json:"alert_messages,omitempty"`
}

// AllTimeHigh struct holds data around an Asset's all time high metric
type AllTimeHigh struct {
	Price             *float64 `json:"price,omitempty"`
	At                *string  `json:"at,omitempty"`
	DaysSince         *float64 `json:"days_since,omitempty"`
	PercentDown       *float64 `json:"percent_down,omitempty"`
	BreakevenMultiple *float64 `json:"breakeven_multiple,omitempty"`
}

// BorrowRates struct holds data for borrowing rates metric of an Asset
type BorrowRates struct {
	Coinlist     *float64 `json:"coinlist,omitempty"`
	Aavestable   *float64 `json:"aavestable,omitempty"`
	Aavevariable *float64 `json:"aavevariable,omitempty"`
	Compound     *float64 `json:"compound,omitempty"`
	Dydx         *float64 `json:"dydx,omitempty"`
	Nuo          *float64 `json:"nuo,omitempty"`
	Nexo         *float64 `json:"nexo,omitempty"`
}

// CycleLow struct holds an Asset's metric around its low cycle
type CycleLow struct {
	Price     *float64 `json:"price,omitempty"`
	At        *string  `json:"at,omitempty"`
	PercentUp *float64 `json:"percent_up,omitempty"`
	DaysSince *float64 `json:"days_since,omitempty"`
}

// DeveloperActivity struct holds data for an Asset's development metrics
type DeveloperActivity struct {
	Stars                   *float64 `json:"stars,omitempty"`
	Watchers                *float64 `json:"watchers,omitempty"`
	CommitsLast3Months      *float64 `json:"commits_last_3_months,omitempty"`
	CommitsLast1Year        *float64 `json:"commits_last_1_year,omitempty"`
	LinesAddedLast3Months   *float64 `json:"lines_added_last_3_months,omitempty"`
	LinesAddedLast1Year     *float64 `json:"lines_added_last_1_year,omitempty"`
	LinesDeletedLast3Months *float64 `json:"lines_deleted_last_3_months,omitempty"`
	LinesDeletedLast1Year   *float64 `json:"lines_deleted_last_1_year,omitempty"`
}

// ExchangeFlows struct is an Asset's metric around exchange flows
type ExchangeFlows struct {
	SupplyExchangeUsd                   *float64 `json:"supply_exchange_usd,omitempty"`
	FlowInExchangeNativeUnitsInclusive  *float64 `json:"flow_in_exchange_native_units_inclusive,omitempty"`
	FlowInExchangeUsdInclusive          *float64 `json:"flow_in_exchange_usd_inclusive,omitempty"`
	FlowInExchangeNativeUnits           *float64 `json:"flow_in_exchange_native_units,omitempty"`
	FlowInExchangeUsd                   *float64 `json:"flow_in_exchange_usd,omitempty"`
	FlowOutExchangeNativeUnitsInclusive *float64 `json:"flow_out_exchange_native_units_inclusive,omitempty"`
	FlowOutExchangeUsdInclusive         *float64 `json:"flow_out_exchange_usd_inclusive,omitempty"`
	FlowOutExchangeNativeUnits          *float64 `json:"flow_out_exchange_native_units,omitempty"`
	FlowOutExchangeUsd                  *float64 `json:"flow_out_exchange_usd,omitempty"`
}

// LoanData struct holds information around an Asset's loaning metric
type LoanData struct {
	OriginatedLast24HoursUsd           *float64 `json:"originated_last_24_hours_usd,omitempty"`
	OutstandingDebtUsd                 *float64 `json:"outstanding_debt_usd,omitempty"`
	RepaidLast24HoursUsd               *float64 `json:"repaid_last_24_hours_usd,omitempty"`
	CollateralizedLast24HoursUsd       *float64 `json:"collateralized_last_24_hours_usd,omitempty"`
	CollateralLiquidatedLast24HoursUsd *float64 `json:"collateral_liquidated_last_24_hours_usd,omitempty"`
}

// MarketData struct holds information on an Asset's MarketData
type MarketData struct {
	PriceUsd                               float64       `json:"price_usd,omitempty"`
	PriceBtc                               float64       `json:"price_btc,omitempty"`
	PriceEth                               float64       `json:"price_eth,omitempty"`
	VolumeLast24Hours                      float64       `json:"volume_last_24_hours,omitempty"`
	RealVolumeLast24Hours                  float64       `json:"real_volume_last_24_hours,omitempty"`
	VolumeLast24HoursOverstatementMultiple float64       `json:"volume_last_24_hours_overstatement_multiple,omitempty"`
	PercentChangeUsdLast1Hour              *float64      `json:"percent_change_usd_last_1_hour,omitempty"`
	PercentChangeUsdLast24Hours            float64       `json:"percent_change_usd_last_24_hours,omitempty"`
	PercentChangeBtcLast24Hours            float64       `json:"percent_change_btc_last_24_hours,omitempty"`
	PercentChangeEthLast24Hours            float64       `json:"percent_change_eth_last_24_hours,omitempty"`
	OhlcvLast1Hour                         OHLCVLastHour `json:"ohlcv_last_1_hour,omitempty"`
	OhlcvLast24Hour                        OHLCVLastHour `json:"ohlcv_last_24_hour,omitempty"`
	LastTradeAt                            string        `json:"last_trade_at,omitempty"`
}

// OHLCVLastHour struct holds an Asset's open, high, low, close and volume data
type OHLCVLastHour struct {
	Open   float64 `json:"open,omitempty"`
	High   float64 `json:"high,omitempty"`
	Low    float64 `json:"low,omitempty"`
	Close  float64 `json:"close,omitempty"`
	Volume float64 `json:"volume,omitempty"`
}

// MarketDataLiquidity struct holds data around an Asset's market liquidity
type MarketDataLiquidity struct {
	ClearingPricesToSell interface{} `json:"clearing_prices_to_sell,omitempty"`
	Marketcap            interface{} `json:"marketcap,omitempty"`
	AssetBidDepth        interface{} `json:"asset_bid_depth,omitempty"`
	UsdBidDepth          interface{} `json:"usd_bid_depth,omitempty"`
	UpdatedAt            string      `json:"updated_at,omitempty"`
}

// Marketcap struct holds data around an Asset's market cap
type Marketcap struct {
	MarketcapDominancePercent        float64  `json:"marketcap_dominance_percent,omitempty"`
	CurrentMarketcapUsd              float64  `json:"current_marketcap_usd,omitempty"`
	Y2050MarketcapUsd                *float64 `json:"y_2050_marketcap_usd,omitempty"`
	YPlus10MarketcapUsd              *float64 `json:"y_plus10_marketcap_usd,omitempty"`
	LiquidMarketcapUsd               *float64 `json:"liquid_marketcap_usd,omitempty"`
	RealizedMarketcapUsd             *float64 `json:"realized_marketcap_usd,omitempty"`
	VolumeTurnoverLast24HoursPercent *float64 `json:"volume_turnover_last_24_hours_percent,omitempty"`
}

// MiningStats struct holds data around an Asset's mining
type MiningStats struct {
	MiningAlgo                 *string  `json:"mining_algo,omitempty"`
	NetworkHashRate            *string  `json:"network_hash_rate,omitempty"`
	AvailableOnNicehashPercent *float64 `json:"available_on_nicehash_percent,omitempty"`
	The1HourAttackCost         *float64 `json:"1_hour_attack_cost,omitempty"`
	The24HoursAttackCost       *float64 `json:"24_hours_attack_cost,omitempty"`
	AttackAppeal               *float64 `json:"attack_appeal,omitempty"`
	MiningRevenueNative        *float64 `json:"mining_revenue_native,omitempty"`
	MiningRevenueUsd           *float64 `json:"mining_revenue_usd,omitempty"`
	AverageDifficulty          *float64 `json:"average_difficulty,omitempty"`
}

// MiscData struct is a struct holding miscellaneous data of an Asset
type MiscData struct {
	PrivateMarketPriceUsd              *float64 `json:"private_market_price_usd,omitempty"`
	VladimirClubCost                   *float64 `json:"vladimir_club_cost,omitempty"`
	BtcCurrentNormalizedSupplyPriceUsd *float64 `json:"btc_current_normalized_supply_price_usd,omitempty"`
	BtcY2050NormalizedSupplyPriceUsd   *float64 `json:"btc_y2050_normalized_supply_price_usd,omitempty"`
	AssetCreatedAt                     *string  `json:"asset_created_at,omitempty"`
	AssetAgeDays                       *float64 `json:"asset_age_days,omitempty"`
	Categories                         []string `json:"categories,omitempty"`
	Sectors                            []string `json:"sectors,omitempty"`
	Tags                               []string `json:"tags,omitempty"`
}

// Reddit struct holds data for an Asset's reddit details
type Reddit struct {
	ActiveUserCount *float64 `json:"active_user_count,omitempty"`
	Subscribers     *float64 `json:"subscribers,omitempty"`
}

// RiskMetrics struct ..
type RiskMetrics struct {
	SharpeRatios    SharpeRatios    `json:"sharpe_ratios,omitempty"`
	VolatilityStats VolatilityStats `json:"volatility_stats,omitempty"`
}

// SharpeRatios struct ..
type SharpeRatios struct {
	Last30Days float64  `json:"last_30_days,omitempty"`
	Last90Days *float64 `json:"last_90_days,omitempty"`
	Last1Year  *float64 `json:"last_1_year,omitempty"`
	Last3Years *float64 `json:"last_3_years,omitempty"`
}

// VolatilityStats struct holds data for an Asset's volatility (in it's RiskMetrics)
type VolatilityStats struct {
	VolatilityLast30Days float64  `json:"volatility_last_30_days,omitempty"`
	VolatilityLast90Days *float64 `json:"volatility_last_90_days,omitempty"`
	VolatilityLast1Year  *float64 `json:"volatility_last_1_year,omitempty"`
	VolatilityLast3Years *float64 `json:"volatility_last_3_years,omitempty"`
}

// StakingStats struct is stats around an Asset's staking
type StakingStats struct {
	StakingYieldPercent     *float64 `json:"staking_yield_percent,omitempty"`
	StakingType             *string  `json:"staking_type,omitempty"`
	StakingMinimum          *float64 `json:"staking_minimum,omitempty"`
	TokensStaked            *float64 `json:"tokens_staked,omitempty"`
	TokensStakedPercent     *float64 `json:"tokens_staked_percent,omitempty"`
	RealStakingYieldPercent *float64 `json:"real_staking_yield_percent,omitempty"`
}

// MetricsSupply struct ..
type MetricsSupply struct {
	Y2050                  *float64 `json:"y_2050,omitempty"`
	YPlus10                *float64 `json:"y_plus10,omitempty"`
	Liquid                 *float64 `json:"liquid,omitempty"`
	Circulating            float64  `json:"circulating,omitempty"`
	Y2050IssuedPercent     *float64 `json:"y_2050_issued_percent,omitempty"`
	AnnualInflationPercent *float64 `json:"annual_inflation_percent,omitempty"`
	StockToFlow            *float64 `json:"stock_to_flow,omitempty"`
	YPlus10IssuedPercent   *float64 `json:"y_plus10_issued_percent,omitempty"`
}

// TokenSaleStats struct holds data for a token's sale
type TokenSaleStats struct {
	SaleProceedsUsd        *float64    `json:"sale_proceeds_usd,omitempty"`
	SaleStartDate          *string     `json:"sale_start_date,omitempty"`
	SaleEndDate            *string     `json:"sale_end_date,omitempty"`
	RoiSinceSaleUsdPercent *float64    `json:"roi_since_sale_usd_percent,omitempty"`
	RoiSinceSaleBtcPercent *float64    `json:"roi_since_sale_btc_percent,omitempty"`
	RoiSinceSaleEthPercent interface{} `json:"roi_since_sale_eth_percent,omitempty"`
}

// Profile struct holds data for an Asset's profile
type Profile struct {
	General      ProfileGeneral `json:"general,omitempty"`
	Contributors Entities       `json:"contributors,omitempty"`
	Advisors     Entities       `json:"advisors,omitempty"`
	Investors    Entities       `json:"investors,omitempty"`
	Ecosystem    Ecosystem      `json:"ecosystem,omitempty"`
	Economics    Economics      `json:"economics,omitempty"`
	Technology   Technology     `json:"technology,omitempty"`
	Governance   Governance     `json:"governance,omitempty"`
	Metadata     Metadata       `json:"metadata,omitempty"`
}

// Entities struct represents list of Individuals and Organizations
type Entities struct {
	Individuals   []Individual   `json:"individuals,omitempty"`
	Organizations []Organization `json:"organizations,omitempty"`
}

// Individual struct holds information about an Individual
type Individual struct {
	Slug        string  `json:"slug,omitempty"`
	FirstName   string  `json:"first_name,omitempty"`
	LastName    string  `json:"last_name,omitempty"`
	Title       string  `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	AvatarURL   *string `json:"avatar_url,omitempty"`
}

// Organization struct holds information about an Organization
type Organization struct {
	Slug        string  `json:"slug,omitempty"`
	Name        string  `json:"name,omitempty"`
	Logo        *string `json:"logo,omitempty"`
	Description *string `json:"description,omitempty"`
}

// Economics struct holds information about an Asset's economics
type Economics struct {
	Token                Token                `json:"token,omitempty"`
	Launch               Launch               `json:"launch,omitempty"`
	ConsensusAndEmission ConsensusAndEmission `json:"consensus_and_emission,omitempty"`
	NativeTreasury       NativeTreasury       `json:"native_treasury,omitempty"`
}

// ConsensusAndEmission struct ..
type ConsensusAndEmission struct {
	Supply    ConsensusAndEmissionSupply `json:"supply,omitempty"`
	Consensus Consensus                  `json:"consensus,omitempty"`
}

// Consensus struct ..
type Consensus struct {
	ConsensusDetails          *string  `json:"consensus_details,omitempty"`
	GeneralConsensusMechanism *string  `json:"general_consensus_mechanism,omitempty"`
	PreciseConsensusMechanism *string  `json:"precise_consensus_mechanism,omitempty"`
	TargetedBlockTime         *float64 `json:"targeted_block_time,omitempty"`
	BlockReward               *float64 `json:"block_reward,omitempty"`
	MiningAlgorithm           *string  `json:"mining_algorithm,omitempty"`
	NextHalvingDate           *string  `json:"next_halving_date,omitempty"`
	IsVictimOf51PercentAttack *bool    `json:"is_victim_of_51_percent_attack,omitempty"`
}

// ConsensusAndEmissionSupply struct ..
type ConsensusAndEmissionSupply struct {
	SupplyCurveDetails  *string              `json:"supply_curve_details,omitempty"`
	GeneralEmissionType *GeneralEmissionType `json:"general_emission_type,omitempty"`
	PreciseEmissionType *string              `json:"precise_emission_type,omitempty"`
	IsCappedSupply      *bool                `json:"is_capped_supply,omitempty"`
	MaxSupply           *float64             `json:"max_supply,omitempty"`
}

// Launch struct holds information about the launch of an Asset
type Launch struct {
	General             LaunchGeneral       `json:"general,omitempty"`
	Fundraising         Fundraising         `json:"fundraising,omitempty"`
	InitialDistribution InitialDistribution `json:"initial_distribution,omitempty"`
}

// Fundraising struct holds information on a Asset's fundraising efforts
type Fundraising struct {
	SalesRounds                 []SalesRound                 `json:"sales_rounds,omitempty"`
	SalesDocuments              []BlockExplorer              `json:"sales_documents,omitempty"`
	SalesTreasuryAccounts       []interface{}                `json:"sales_treasury_accounts,omitempty"`
	TreasuryPolicies            []interface{}                `json:"treasury_policies,omitempty"`
	ProjectedUseOfSalesProceeds []ProjectedUseOfSalesProceed `json:"projected_use_of_sales_proceeds,omitempty"`
}

// ProjectedUseOfSalesProceed struct ..
type ProjectedUseOfSalesProceed struct {
	Category           *string  `json:"category,omitempty"`
	AmountInPercentage *float64 `json:"amount_in_percentage,omitempty"`
}

// BlockExplorer struct holds information on an Asset's block explorer
type BlockExplorer struct {
	Name string `json:"name,omitempty"`
	Link string `json:"link,omitempty"`
}

// SalesRound struct ..
type SalesRound struct {
	Title                        string      `json:"title,omitempty"`
	StartDate                    *string     `json:"start_date,omitempty"`
	Type                         *string     `json:"type,omitempty"`
	Details                      *string     `json:"details,omitempty"`
	EndDate                      *string     `json:"end_date,omitempty"`
	NativeTokensAllocated        *float64    `json:"native_tokens_allocated,omitempty"`
	AssetCollected               *string     `json:"asset_collected,omitempty"`
	PricePerTokenInAsset         *float64    `json:"price_per_token_in_asset,omitempty"`
	EquivalentPricePerTokenInUsd *float64    `json:"equivalent_price_per_token_in_usd,omitempty"`
	AmountCollectedInAsset       *float64    `json:"amount_collected_in_asset,omitempty"`
	AmountCollectedInUsd         *float64    `json:"amount_collected_in_usd,omitempty"`
	IsKycRequired                interface{} `json:"is_kyc_required,omitempty"`
	RestrictedJurisdictions      []string    `json:"restricted_jurisdictions,omitempty"`
}

// LaunchGeneral struct ..
type LaunchGeneral struct {
	LaunchStyle   *string `json:"launch_style,omitempty"`
	LaunchDetails *string `json:"launch_details,omitempty"`
}

// InitialDistribution struct holds info on an Asset's initial token distribution metrics
type InitialDistribution struct {
	InitialSupply            *float64                 `json:"initial_supply,omitempty"`
	InitialSupplyRepartition InitialSupplyRepartition `json:"initial_supply_repartition,omitempty"`
	TokenDistributionDate    *string                  `json:"token_distribution_date,omitempty"`
	GenesisBlockDate         *string                  `json:"genesis_block_date,omitempty"`
}

// InitialSupplyRepartition struct ..
type InitialSupplyRepartition struct {
	AllocatedToInvestorsPercentage                 *float64 `json:"allocated_to_investors_percentage,omitempty"`
	AllocatedToOrganizationOrFoundersPercentage    *float64 `json:"allocated_to_organization_or_founders_percentage,omitempty"`
	AllocatedToPreminedRewardsOrAirdropsPercentage *float64 `json:"allocated_to_premined_rewards_or_airdrops_percentage,omitempty"`
}

// NativeTreasury struct ..
type NativeTreasury struct {
	Accounts             []interface{} `json:"accounts,omitempty"`
	TreasuryUsageDetails interface{}   `json:"treasury_usage_details,omitempty"`
}

// Token struct holds information on a token
type Token struct {
	TokenName         *string         `json:"token_name,omitempty"`
	TokenType         *TokenType      `json:"token_type,omitempty"`
	TokenAddress      *string         `json:"token_address,omitempty"`
	BlockExplorers    []BlockExplorer `json:"block_explorers,omitempty"`
	Multitoken        []interface{}   `json:"multitoken,omitempty"`
	TokenUsage        *string         `json:"token_usage,omitempty"`
	TokenUsageDetails *string         `json:"token_usage_details,omitempty"`
}

// Ecosystem struct holds information on an Asset's ecosystem
type Ecosystem struct {
	Assets        []EcosystemAsset `json:"assets,omitempty"`
	Organizations []Organization   `json:"organizations,omitempty"`
}

// EcosystemAsset struct is a struct that hold details of an Asset within an Asset's ecosystem
type EcosystemAsset struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// ProfileGeneral struct ..
type ProfileGeneral struct {
	Overview   GeneralOverview `json:"overview,omitempty"`
	Background Background      `json:"background,omitempty"`
	Roadmap    []Roadmap       `json:"roadmap,omitempty"`
	Regulation Regulation      `json:"regulation,omitempty"`
}

// Background struct ..
type Background struct {
	BackgroundDetails    *string        `json:"background_details,omitempty"`
	IssuingOrganizations []Organization `json:"issuing_organizations,omitempty"`
}

// GeneralOverview struct ..
type GeneralOverview struct {
	IsVerified     *bool           `json:"is_verified,omitempty"`
	Tagline        *string         `json:"tagline,omitempty"`
	Category       *string         `json:"category,omitempty"`
	Sector         *string         `json:"sector,omitempty"`
	Tags           *string         `json:"tags,omitempty"`
	ProjectDetails *string         `json:"project_details,omitempty"`
	OfficialLinks  []BlockExplorer `json:"official_links,omitempty"`
}

// Regulation struct ..
type Regulation struct {
	RegulatoryDetails *string  `json:"regulatory_details,omitempty"`
	SfarScore         *float64 `json:"sfar_score,omitempty"`
	SfarSummary       *string  `json:"sfar_summary,omitempty"`
}

// Roadmap struct ..
type Roadmap struct {
	Title   string  `json:"title,omitempty"`
	Date    *string `json:"date,omitempty"`
	Type    *string `json:"type,omitempty"`
	Details *string `json:"details,omitempty"`
}

// Governance struct ..
type Governance struct {
	GovernanceDetails *string           `json:"governance_details,omitempty"`
	OnchainGovernance OnchainGovernance `json:"onchain_governance,omitempty"`
	Grants            []interface{}     `json:"grants,omitempty"`
}

// OnchainGovernance struct ..
type OnchainGovernance struct {
	OnchainGovernanceType    *string `json:"onchain_governance_type,omitempty"`
	OnchainGovernanceDetails *string `json:"onchain_governance_details,omitempty"`
	IsTreasuryDecentralized  *bool   `json:"is_treasury_decentralized,omitempty"`
}

// Metadata struct ..
type Metadata struct {
	UpdatedAt string `json:"updated_at,omitempty"`
}

// Technology struct ..
type Technology struct {
	Overview TechnologyOverview `json:"overview,omitempty"`
	Security Security           `json:"security,omitempty"`
}

// TechnologyOverview struct ..
type TechnologyOverview struct {
	TechnologyDetails  *string            `json:"technology_details,omitempty"`
	ClientRepositories []ClientRepository `json:"client_repositories,omitempty"`
}

// ClientRepository struct ..
type ClientRepository struct {
	Name        string  `json:"name,omitempty"`
	Link        string  `json:"link,omitempty"`
	LicenseType *string `json:"license_type,omitempty"`
}

// Security struct ..
type Security struct {
	Audits                          []Roadmap `json:"audits,omitempty"`
	KnownExploitsAndVulnerabilities []Roadmap `json:"known_exploits_and_vulnerabilities,omitempty"`
}

// Status struct ..
type Status struct {
	Elapsed   float64 `json:"elapsed,omitempty"`
	Timestamp string  `json:"timestamp,omitempty"`
}

// GeneralEmissionType is a string which dictate's an Asset's emission type
type GeneralEmissionType string

// A const type of GeneralEmissionType
const (
	BurnMint     GeneralEmissionType = "Burn & Mint"
	Deflationary GeneralEmissionType = "Deflationary"
	FixedSupply  GeneralEmissionType = "Fixed Supply"
	Inflationary GeneralEmissionType = "Inflationary"
)

// TokenType is a string which represents type of a token
type TokenType string

// A const type which represents a kind of TokenType
const (
	ERC20OmniTRC20 TokenType = "ERC-20, Omni, TRC-20"
	Erc20          TokenType = "ERC-20"
	Native         TokenType = "Native"
)
