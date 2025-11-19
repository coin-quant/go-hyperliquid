package hyperliquid

import (
	"fmt"
)

type ActiveAssetDataSubscriptionParams struct {
	User string
	Coin string
}

func (w *WebsocketClient) ActiveAssetData(
	params ActiveAssetDataSubscriptionParams,
	callback func(ActiveAssetData, error),
) (*Subscription, error) {
	remotePayload := remoteActiveAssetDataSubscriptionPayload{
		Type: ChannelActiveAssetData,
		User: params.User,
		Coin: params.Coin,
	}

	return w.subscribe(remotePayload, func(msg any) {
		activeAssetData, ok := msg.(ActiveAssetData)
		if !ok {
			callback(ActiveAssetData{}, fmt.Errorf("SubscribeToActiveAssetData invalid message type"))
			return
		}

		callback(activeAssetData, nil)
	})
}
