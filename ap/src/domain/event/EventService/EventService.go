package EventService

import (
	"alma-server/ap/src/repository/user/event/UserEventRepository"
	"context"
)

// DefaultEventName .
const DefaultEventName = "イベントの選択"

// GetEventName イベント名を取得する
func GetEventName(ctx context.Context, eventID string) string {
	userEvent := GetEvent(ctx, eventID)
	if userEvent == nil {
		return DefaultEventName
	}

	return userEvent.Name
}

// GetEvent イベントを取得する
func GetEvent(ctx context.Context, eventID string) *UserEventRepository.UserEvent {
	return UserEventRepository.Get(ctx, eventID)
}
