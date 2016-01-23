package sabnzbd

const (
	Byte  = 1
	KByte = Byte * 1000
	MByte = KByte * 1000
	GByte = MByte * 1000
	TByte = GByte * 1000
	PByte = TByte * 1000
	EByte = PByte * 1000
)

const (
	KiByte = 1 << ((iota + 1) * 10)
	MiByte = KByte * 1000
	GiByte = MByte * 1000
	TiByte = GByte * 1000
	PiByte = TByte * 1000
	EiByte = PByte * 1000
)
