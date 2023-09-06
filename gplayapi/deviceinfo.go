package gplayapi

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Juby210/gplayapi-go/gpproto"
)

type (
	DeviceInfo struct {
		_Build           *DeviceBuildInfo
		_SimOperator     string
		_Platforms       []string
		_OtaInstalled    bool
		_CellOperator    string
		_Roaming         string
		_TimeZone        string
		_TouchScreen     int32
		_Keyboard        int32
		_Navigation      int32
		_ScreenLayout    int32
		_Screen          *DeviceInfoScreen
		_GLVersion       int32
		_GLExtensions    []string
		_SharedLibraries []string
		_Features        []string
		_Locales         []string
	}

	DeviceBuildInfo struct {
		*gpproto.AndroidBuildProto
		_VersionRelease int
	}

	DeviceInfoScreen struct {
		_Density int32
		_Width   int32
		_Height  int32
	}
)

func (i *DeviceInfo) GetAuthUserAgent() string {
	return fmt.Sprintf("GoogleAuth/1.4 %s %s", i._Build.GetDevice(), i._Build.GetId())
}

func (i *DeviceInfo) GetUserAgent() string {
	params := []string{
		"api=3",
		"versionCode=81582300",
		"sdk=28",
		"device=" + i._Build.GetDevice(),
		"hardware=" + i._Build.GetDevice(),
		"product=" + i._Build.GetProduct(),
		"platformVersionRelease=" + strconv.Itoa(i._Build._VersionRelease),
		"model=" + i._Build.GetModel(),
		"buildId=" + i._Build.GetId(),
		"isWideScreen=0",
		"supportedAbis=" + strings.Join(i._Platforms, ";"),
	}
	return "Android-Finsky/15.8.23-all [0] [PR] 259261889 (" + strings.Join(params, ",") + ")"
}

func (i *DeviceInfo) GenerateAndroidCheckInRequest() *gpproto.AndroidCheckinRequest {
	var int0 int64 = 0
	timestamp := time.Now().Unix()
	i._Build.Timestamp = &timestamp

	return &gpproto.AndroidCheckinRequest{
		Id: &int0,
		Checkin: &gpproto.AndroidCheckinProto{
			Build:           i._Build.AndroidBuildProto,
			LastCheckinMsec: &int0,
			CellOperator:    &i._CellOperator,
			SimOperator:     &i._SimOperator,
			Roaming:         &i._Roaming,
			UserNumber:      ptrInt32(0),
		},
		Locale:              ptrStr("en_GB"),
		TimeZone:            &i._TimeZone,
		Version:             ptrInt32(3),
		DeviceConfiguration: i.GetDeviceConfigProto(),
		Fragment:            ptrInt32(0),
	}
}

func (i *DeviceInfo) GetDeviceConfigProto() *gpproto.DeviceConfigurationProto {
	var mem int64 = 8589935000
	return &gpproto.DeviceConfigurationProto{
		TouchScreen:            &i._TouchScreen,
		Keyboard:               &i._Keyboard,
		Navigation:             &i._Navigation,
		ScreenLayout:           &i._ScreenLayout,
		HasHardKeyboard:        ptrBool(false),
		HasFiveWayNavigation:   ptrBool(false),
		ScreenDensity:          &i._Screen._Density,
		GlEsVersion:            &i._GLVersion,
		SystemSharedLibrary:    i._SharedLibraries,
		SystemAvailableFeature: i._Features,
		NativePlatform:         i._Platforms,
		ScreenWidth:            &i._Screen._Width,
		ScreenHeight:           &i._Screen._Height,
		SystemSupportedLocale:  i._Locales,
		GlExtension:            i._GLExtensions,
		DeviceClass:            ptrInt32(0),
		LowRamDevice:           ptrInt32(0),
		TotalMemoryBytes:       &mem,
		MaxNumOf_CPUCores:      ptrInt32(8),
		DeviceFeature:          i.GetDeviceFeatures(),
	}
}

func (i *DeviceInfo) GetDeviceFeatures() (ret []*gpproto.DeviceFeature) {
	var int0 int32 = 0
	for _, f := range i._Features {
		name := f
		ret = append(ret, &gpproto.DeviceFeature{Name: &name, Value: &int0})
	}
	return ret
}
