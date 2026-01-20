package polymarket

import (
	"strconv"
	"time"

	"github.com/Eraguzy/PolyNew/internal/storage"
)

type Tag struct {
	ID                  string    `json:"id"`
	Label               string    `json:"label"`
	Slug                string    `json:"slug"`
	ForceShow           bool      `json:"forceShow"`
	PublishedAt         string    `json:"publishedAt"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
	RequiresTranslation bool      `json:"requiresTranslation"`
}

func (s Tag) ToTableTag() (storage.TableTags, error) {
	strid, err := strconv.Atoi(s.ID)
	if err != nil {
		return storage.TableTags{}, err
	}
	return storage.TableTags{
		TagID:   strid,
		Name:    s.Label,
		Tracked: true, // default to false, can be updated later
	}, nil
}

// Market represents a market payload (see test.json -> "type Market struct").
type Market struct {
	ID                           int           `json:"id"`
	Question                     string        `json:"question"`
	ConditionID                  string        `json:"conditionId"`
	Slug                         string        `json:"slug"`
	TwitterCardImage             string        `json:"twitterCardImage"`
	EndDate                      time.Time     `json:"endDate"`
	Category                     string        `json:"category"`
	Liquidity                    string        `json:"liquidity"`
	Image                        string        `json:"image"`
	Icon                         string        `json:"icon"`
	Description                  string        `json:"description"`
	Outcomes                     string        `json:"outcomes"`
	OutcomePrices                string        `json:"outcomePrices"`
	Volume                       string        `json:"volume"`
	Active                       bool          `json:"active"`
	MarketType                   string        `json:"marketType"`
	Closed                       bool          `json:"closed"`
	MarketMakerAddress           string        `json:"marketMakerAddress"`
	UpdatedBy                    int           `json:"updatedBy"`
	CreatedAt                    time.Time     `json:"createdAt"`
	UpdatedAt                    time.Time     `json:"updatedAt"`
	ClosedTime                   string        `json:"closedTime"`
	MailchimpTag                 string        `json:"mailchimpTag"`
	Archived                     bool          `json:"archived"`
	Restricted                   bool          `json:"restricted"`
	VolumeNum                    float64       `json:"volumeNum"`
	LiquidityNum                 float64       `json:"liquidityNum"`
	EndDateISO                   string        `json:"endDateIso"`
	HasReviewedDates             bool          `json:"hasReviewedDates"`
	ReadyForCron                 bool          `json:"readyForCron"`
	Volume24hr                   float64       `json:"volume24hr"`
	Volume1wk                    float64       `json:"volume1wk"`
	Volume1mo                    float64       `json:"volume1mo"`
	Volume1yr                    float64       `json:"volume1yr"`
	ClobTokenIDs                 string        `json:"clobTokenIds"`
	FpmmLive                     bool          `json:"fpmmLive"`
	Volume1wkAmm                 float64       `json:"volume1wkAmm"`
	Volume1moAmm                 float64       `json:"volume1moAmm"`
	Volume1yrAmm                 float64       `json:"volume1yrAmm"`
	Volume1wkClob                float64       `json:"volume1wkClob"`
	Volume1moClob                float64       `json:"volume1moClob"`
	Volume1yrClob                float64       `json:"volume1yrClob"`
	Events                       []MarketEvent `json:"events"`
	Creator                      string        `json:"creator"`
	Ready                        bool          `json:"ready"`
	Funded                       bool          `json:"funded"`
	Cyom                         bool          `json:"cyom"`
	Competitive                  int           `json:"competitive"`
	PagerDutyNotificationEnabled bool          `json:"pagerDutyNotificationEnabled"`
	Approved                     bool          `json:"approved"`
	RewardsMinSize               float64       `json:"rewardsMinSize"`
	RewardsMaxSpread             float64       `json:"rewardsMaxSpread"`
	Spread                       float64       `json:"spread"`
	OneDayPriceChange            float64       `json:"oneDayPriceChange"`
	OneHourPriceChange           float64       `json:"oneHourPriceChange"`
	OneWeekPriceChange           float64       `json:"oneWeekPriceChange"`
	OneMonthPriceChange          float64       `json:"oneMonthPriceChange"`
	OneYearPriceChange           float64       `json:"oneYearPriceChange"`
	LastTradePrice               float64       `json:"lastTradePrice"`
	BestBid                      float64       `json:"bestBid"`
	BestAsk                      float64       `json:"bestAsk"`
	ClearBookOnStart             bool          `json:"clearBookOnStart"`
	ManualActivation             bool          `json:"manualActivation"`
	NegRiskOther                 bool          `json:"negRiskOther"`
	UMAResolutionStatuses        string        `json:"umaResolutionStatuses"`
	PendingDeployment            bool          `json:"pendingDeployment"`
	Deploying                    bool          `json:"deploying"`
	RfqEnabled                   bool          `json:"rfqEnabled"`
	HoldingRewardsEnabled        bool          `json:"holdingRewardsEnabled"`
	FeesEnabled                  bool          `json:"feesEnabled"`
	RequiresTranslation          bool          `json:"requiresTranslation"`
}

// MarketEvent represents the event object inside a Market.
type MarketEvent struct {
	ID                  int       `json:"id"`
	Ticker              string    `json:"ticker"`
	Slug                string    `json:"slug"`
	Title               string    `json:"title"`
	Description         string    `json:"description"`
	StartDate           time.Time `json:"startDate"`
	CreationDate        time.Time `json:"creationDate"`
	EndDate             time.Time `json:"endDate"`
	Image               string    `json:"image"`
	Icon                string    `json:"icon"`
	Active              bool      `json:"active"`
	Closed              bool      `json:"closed"`
	Archived            bool      `json:"archived"`
	Featured            bool      `json:"featured"`
	Restricted          bool      `json:"restricted"`
	Liquidity           float64   `json:"liquidity"`
	Volume              float64   `json:"volume"`
	OpenInterest        float64   `json:"openInterest"`
	SortBy              string    `json:"sortBy"`
	Category            string    `json:"category"`
	PublishedAt         string    `json:"published_at"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
	Competitive         int       `json:"competitive"`
	Volume24hr          float64   `json:"volume24hr"`
	Volume1wk           float64   `json:"volume1wk"`
	Volume1mo           float64   `json:"volume1mo"`
	Volume1yr           float64   `json:"volume1yr"`
	LiquidityAmm        float64   `json:"liquidityAmm"`
	LiquidityClob       float64   `json:"liquidityClob"`
	CommentCount        int       `json:"commentCount"`
	Cyom                bool      `json:"cyom"`
	ClosedTime          time.Time `json:"closedTime"`
	ShowAllOutcomes     bool      `json:"showAllOutcomes"`
	ShowMarketImages    bool      `json:"showMarketImages"`
	EnableNegRisk       bool      `json:"enableNegRisk"`
	NegRiskAugmented    bool      `json:"negRiskAugmented"`
	PendingDeployment   bool      `json:"pendingDeployment"`
	Deploying           bool      `json:"deploying"`
	RequiresTranslation bool      `json:"requiresTranslation"`
}

// Event represents an event payload (see test.json -> "type Event struct").
type Event struct {
	ID                  string        `json:"id"`
	Ticker              string        `json:"ticker"`
	Slug                string        `json:"slug"`
	Title               string        `json:"title"`
	Description         string        `json:"description"`
	ResolutionSource    string        `json:"resolutionSource"`
	StartDate           time.Time     `json:"startDate"`
	CreationDate        time.Time     `json:"creationDate"`
	EndDate             time.Time     `json:"endDate"`
	Image               string        `json:"image"`
	Icon                string        `json:"icon"`
	Active              bool          `json:"active"`
	Closed              bool          `json:"closed"`
	Archived            bool          `json:"archived"`
	New                 bool          `json:"new"`
	Featured            bool          `json:"featured"`
	Restricted          bool          `json:"restricted"`
	Liquidity           float64       `json:"liquidity"`
	Volume              float64       `json:"volume"`
	OpenInterest        float64       `json:"openInterest"`
	SortBy              string        `json:"sortBy"`
	Category            string        `json:"category"`
	PublishedAt         string        `json:"published_at"`
	CreatedAt           time.Time     `json:"createdAt"`
	UpdatedAt           time.Time     `json:"updatedAt"`
	Competitive         float64       `json:"competitive"`
	Volume24hr          float64       `json:"volume24hr"`
	Volume1wk           float64       `json:"volume1wk"`
	Volume1mo           float64       `json:"volume1mo"`
	Volume1yr           float64       `json:"volume1yr"`
	LiquidityAmm        float64       `json:"liquidityAmm"`
	LiquidityClob       float64       `json:"liquidityClob"`
	CommentCount        int           `json:"commentCount"`
	Markets             []EventMarket `json:"markets"`
	Series              []EventSeries `json:"series"`
	Tags                []EventTag    `json:"tags"`
	Cyom                bool          `json:"cyom"`
	ClosedTime          time.Time     `json:"closedTime"`
	ShowAllOutcomes     bool          `json:"showAllOutcomes"`
	ShowMarketImages    bool          `json:"showMarketImages"`
	EnableNegRisk       bool          `json:"enableNegRisk"`
	SeriesSlug          string        `json:"seriesSlug"`
	NegRiskAugmented    bool          `json:"negRiskAugmented"`
	PendingDeployment   bool          `json:"pendingDeployment"`
	Deploying           bool          `json:"deploying"`
	RequiresTranslation bool          `json:"requiresTranslation"`
}

func (e Event) ToTableEvent() (storage.TableEvent, error) {
	strid, err := strconv.Atoi(e.ID)
	if err != nil {
		return storage.TableEvent{}, err
	}
	return storage.TableEvent{
		EventID:   strid,
		CreatedAt: e.CreatedAt,
		Ticker:    e.Ticker,
		ImageURL:  &e.Image,
		Title:     &e.Title,
	}, nil
}

// EventMarket represents a market object inside an Event.
type EventMarket struct {
	ID                           string    `json:"id"`
	Question                     string    `json:"question"`
	ConditionID                  string    `json:"conditionId"`
	Slug                         string    `json:"slug"`
	ResolutionSource             string    `json:"resolutionSource"`
	EndDate                      time.Time `json:"endDate"`
	Category                     string    `json:"category"`
	Liquidity                    string    `json:"liquidity"`
	StartDate                    time.Time `json:"startDate"`
	Fee                          string    `json:"fee"`
	Image                        string    `json:"image"`
	Icon                         string    `json:"icon"`
	Description                  string    `json:"description"`
	Outcomes                     string    `json:"outcomes"`
	OutcomePrices                string    `json:"outcomePrices"`
	Volume                       string    `json:"volume"`
	Active                       bool      `json:"active"`
	MarketType                   string    `json:"marketType"`
	Closed                       bool      `json:"closed"`
	MarketMakerAddress           string    `json:"marketMakerAddress"`
	UpdatedBy                    string    `json:"updatedBy"`
	CreatedAt                    time.Time `json:"createdAt"`
	UpdatedAt                    time.Time `json:"updatedAt"`
	ClosedTime                   string    `json:"closedTime"`
	WideFormat                   bool      `json:"wideFormat"`
	New                          bool      `json:"new"`
	SentDiscord                  bool      `json:"sentDiscord"`
	Featured                     bool      `json:"featured"`
	SubmittedBy                  string    `json:"submitted_by"`
	TwitterCardLocation          string    `json:"twitterCardLocation"`
	TwitterCardLastRefreshed     string    `json:"twitterCardLastRefreshed"`
	Archived                     bool      `json:"archived"`
	ResolvedBy                   string    `json:"resolvedBy"`
	Restricted                   bool      `json:"restricted"`
	VolumeNum                    float64   `json:"volumeNum"`
	LiquidityNum                 float64   `json:"liquidityNum"`
	EndDateISO                   string    `json:"endDateIso"`
	StartDateISO                 string    `json:"startDateIso"`
	HasReviewedDates             bool      `json:"hasReviewedDates"`
	ReadyForCron                 bool      `json:"readyForCron"`
	Volume24hr                   float64   `json:"volume24hr"`
	Volume1wk                    float64   `json:"volume1wk"`
	Volume1mo                    float64   `json:"volume1mo"`
	Volume1yr                    float64   `json:"volume1yr"`
	ClobTokenIDs                 string    `json:"clobTokenIds"`
	FpmmLive                     bool      `json:"fpmmLive"`
	Volume1wkAmm                 float64   `json:"volume1wkAmm"`
	Volume1moAmm                 float64   `json:"volume1moAmm"`
	Volume1yrAmm                 float64   `json:"volume1yrAmm"`
	Volume1wkClob                float64   `json:"volume1wkClob"`
	Volume1moClob                float64   `json:"volume1moClob"`
	Volume1yrClob                float64   `json:"volume1yrClob"`
	Creator                      string    `json:"creator"`
	Ready                        bool      `json:"ready"`
	Funded                       bool      `json:"funded"`
	Cyom                         bool      `json:"cyom"`
	Competitive                  float64   `json:"competitive"`
	PagerDutyNotificationEnabled bool      `json:"pagerDutyNotificationEnabled"`
	Approved                     bool      `json:"approved"`
	RewardsMinSize               float64   `json:"rewardsMinSize"`
	RewardsMaxSpread             float64   `json:"rewardsMaxSpread"`
	Spread                       float64   `json:"spread"`
	OneDayPriceChange            float64   `json:"oneDayPriceChange"`
	OneHourPriceChange           float64   `json:"oneHourPriceChange"`
	OneWeekPriceChange           float64   `json:"oneWeekPriceChange"`
	OneMonthPriceChange          float64   `json:"oneMonthPriceChange"`
	OneYearPriceChange           float64   `json:"oneYearPriceChange"`
	LastTradePrice               float64   `json:"lastTradePrice"`
	BestBid                      float64   `json:"bestBid"`
	BestAsk                      float64   `json:"bestAsk"`
	ClearBookOnStart             bool      `json:"clearBookOnStart"`
	ManualActivation             bool      `json:"manualActivation"`
	NegRiskOther                 bool      `json:"negRiskOther"`
	UMAResolutionStatuses        string    `json:"umaResolutionStatuses"`
	PendingDeployment            bool      `json:"pendingDeployment"`
	Deploying                    bool      `json:"deploying"`
	RfqEnabled                   bool      `json:"rfqEnabled"`
	HoldingRewardsEnabled        bool      `json:"holdingRewardsEnabled"`
	FeesEnabled                  bool      `json:"feesEnabled"`
	RequiresTranslation          bool      `json:"requiresTranslation"`
}

func (em EventMarket) ToTableEventMarket(eventID int) (storage.TableEventMarket, error) {
	strid, err := strconv.Atoi(em.ID)
	if err != nil {
		return storage.TableEventMarket{}, err
	}
	return storage.TableEventMarket{
		EventMarketID: strid,
		EventID:       eventID,
		Question:      em.Question,
		ImageURL:      &em.Image,
	}, nil
}

type EventSeries struct {
	ID                  string    `json:"id"`
	Ticker              string    `json:"ticker"`
	Slug                string    `json:"slug"`
	Title               string    `json:"title"`
	SeriesType          string    `json:"seriesType"`
	Recurrence          string    `json:"recurrence"`
	Image               string    `json:"image"`
	Icon                string    `json:"icon"`
	Layout              string    `json:"layout"`
	Active              bool      `json:"active"`
	Closed              bool      `json:"closed"`
	Archived            bool      `json:"archived"`
	New                 bool      `json:"new"`
	Featured            bool      `json:"featured"`
	Restricted          bool      `json:"restricted"`
	PublishedAt         string    `json:"publishedAt"`
	CreatedBy           string    `json:"createdBy"`
	UpdatedBy           string    `json:"updatedBy"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
	CommentsEnabled     bool      `json:"commentsEnabled"`
	Competitive         float64   `json:"competitive"`
	Volume24hr          float64   `json:"volume24hr"`
	StartDate           time.Time `json:"startDate"`
	CommentCount        int       `json:"commentCount"`
	RequiresTranslation bool      `json:"requiresTranslation"`
}

type EventTag struct {
	ID                  string    `json:"id"`
	Label               string    `json:"label"`
	Slug                string    `json:"slug"`
	ForceShow           bool      `json:"forceShow"`
	UpdatedAt           time.Time `json:"updatedAt"`
	RequiresTranslation bool      `json:"requiresTranslation"`
}
