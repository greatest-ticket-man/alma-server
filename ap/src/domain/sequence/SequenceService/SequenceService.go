package SequenceService

import (
	"alma-server/ap/src/repository/user/UserSequenceRepository"
	"context"
)

// NextReserveSeq 予約のシーケンス番号
func NextReserveSeq(ctx context.Context, eventID string) uint64 {
	return UserSequenceRepository.Next(ctx, eventID, UserSequenceRepository.ReserveKey)
}
