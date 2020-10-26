package ticketconst

// PriceType 価格タイプ
type PriceType int32

const (
	// Adv Advance Tickets 前売り券
	Adv PriceType = 1

	// DOS Doors 当日券
	DOS PriceType = 2
)
