package types

import (
	"fmt"
)

type Channel byte

const (
	ChannUnknown  Channel = 0 // 位置渠道
	ChannelMobile Channel = 1 // 手机号
)

func (c Channel) IsValid() bool {
	switch c {
	case ChannelMobile:
		return true
	default:
		return false
	}
}

func ParseChannel(c int) (Channel, error) {
	ch := Channel(c)
	if !ch.IsValid() {
		return ChannUnknown, fmt.Errorf("invalid channel: %v", ch)
	}
	return ch, nil
}

func MustParseChannel(c int) Channel {
	ch, err := ParseChannel(c)
	if err != nil {
		return ChannUnknown
	}
	return ch
}

func AllChannels() []Channel {
	return []Channel{
		ChannelMobile,
	}
}
