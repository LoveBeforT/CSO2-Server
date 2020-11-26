package typestruct

import "sync"

type (
	//分区服务器，管理所拥有的频道
	ChannelServer struct {
		ServerIndex  uint8
		ServerStatus uint8
		ServerType   uint8
		ServerName   []byte
		ChannelCount uint8
		Channels     []*ChannelInfo
	}

	//主服务器，管理各个分区
	ServerManager struct {
		ServerNum uint8
		Servers   []*ChannelServer
	}

	//频道信息，隶属于分区服务器,用于请求服务器和请求频道
	ChannelInfo struct {
		ChannelID     uint8
		ChannelName   []byte
		Unk00         uint16
		Unk01         uint16
		Unk02         uint8
		ChannelType   uint8
		ChannelStatus uint8
		NextRoomID    uint8
		RoomNum       uint16
		Rooms         map[uint16]*Room
		RoomNums      map[uint8]uint16

		ChannelMutex *sync.Mutex
	}
)

const (
	MAXCHANNELNUM       = 16
	MAXSERVERNUM        = 15
	MAXCHANNELROOMNUM   = 0xFF
	MAXROOMNUM          = 0xFFFF
	DefalutServerName   = "한국 서버"
	DefalutChannelName1 = "채널 1"
	DefalutChannelName2 = "채널 2"

	//貌似非3以外的都被客户端认为是战队频道
	ChannelServerTypeNormal = 1
	ChannelServerTypeTeam   = 3

	//ChannelType
	ChannelTypeFree         = 0
	ChannelTypeNovice       = 1
	ChannelTypeNoviceLowKAD = 2
	ChannelTypeClan         = 3
	ChannelTypeBigCity      = 4
	ChannelTypeActive       = 5 //only open for active

	ChannelStatusBusy   = 0
	ChannelStatusNormal = 1
)
