package db

type VideoStatus string

const (
	VideoStatusDownloaded VideoStatus = "downloaded"
	VideoStatusPending    VideoStatus = "pending"
)

type Video struct {
	VideoId   string
	ChannelId string
	FileName  string
	Status    VideoStatus
}

type ChannelStatus string

const (
	ChannelStatusArchiving  ChannelStatus = "archiving"
	ChannelStatusMonitoring ChannelStatus = "monitoring"
)

type Channel struct {
	ChannelId string
	Status    ChannelStatus
	LastCheck int
}
