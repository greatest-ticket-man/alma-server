package ticketconst

// PriceType 価格タイプ
type PriceType int32

const (
	// Adv Advance Tickets 前売り券
	Adv PriceType = 1

	// Dos Doors 当日券
	Dos PriceType = 2

	// Invite Invitation 招待
	Invite PriceType = 3
)
