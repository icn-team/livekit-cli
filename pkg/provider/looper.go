package provider

import (
	"github.com/icn-team/webrtc/v3"

	"github.com/livekit/protocol/livekit"
	lksdk "github.com/icn-team/server-sdk-go"
)

type VideoLooper interface {
	lksdk.SampleProvider
	Codec() webrtc.RTPCodecCapability
	ToLayer(quality livekit.VideoQuality) *livekit.VideoLayer
}
