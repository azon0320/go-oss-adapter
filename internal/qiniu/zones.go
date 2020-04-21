package qiniu

import "github.com/qiniu/api.v7/v7/storage"

const (
	ZoneHuadong  = "east-china"
	ZoneHuabei   = "north-china"
	ZoneHuanan   = "south-china"
	ZoneBeimei   = "north-usa"
	ZoneXinjiapo = "singapore"
)

func GetZoneFromString(zonestr string) *storage.Zone {
	switch zonestr {
	case ZoneHuadong:
		return &storage.ZoneHuadong
	case ZoneHuabei:
		return &storage.ZoneHuabei
	case ZoneHuanan:
		return &storage.ZoneHuanan
	case ZoneBeimei:
		return &storage.ZoneBeimei
	case ZoneXinjiapo:
		return &storage.ZoneXinjiapo
	default:
		return nil
	}
}

func GetZoneFromStringDefault(zonestr string, defzone *storage.Zone) *storage.Zone {
	var zone = GetZoneFromString(zonestr)
	if zone == nil {
		return defzone
	} else {
		return zone
	}
}
