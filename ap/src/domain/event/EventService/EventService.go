package EventService

import (
	"alma-server/ap/src/repository/user/event/UserEventRepository"
	"context"
)

// DefaultEventName .
const DefaultEventName = "イベントの選択"

// GetEventName イベント名を取得する
func GetEventName(ctx context.Context, eventID string) string {
	userEvent := UserEventRepository.Get(ctx, eventID)
	if userEvent == nil {
		return DefaultEventName
	}

	return userEvent.Name
}
