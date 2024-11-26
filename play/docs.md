# Overview

package `play`

## Index

- [Constants](#constants)
- [Variables](#variables)
- [Types](#types)
  - [type Apk](#type-apk)
    - [func (a Apk) Field1() string](#func-apk-field1)
    - [func (a Apk) Url() string](#func-apk-url)
  - [type Delivery](#type-delivery)
    - [func (d Delivery) Apk() func() (Apk, bool)](#func-delivery-apk)
    - [func (d Delivery) Obb() func() (Obb, bool)](#func-delivery-obb)
    - [func (d Delivery) Url() string](#func-delivery-url)
  - [type Details](#type-details)
    - [func (d Details) Downloads() (uint64, bool)](#func-details-downloads)
    - [func (d Details) Name() (string, bool)](#func-details-name)
    - [func (d Details) String() string](#func-details-string)
  - [type GoogleAuth](#type-googleauth)
    - [func (g GoogleAuth) Acquire(checkin \*GoogleCheckin, id string) error](#func-googleauth-acquire)
    - [func (g GoogleAuth) Delivery(
  check \*GoogleCheckin, app \*StoreApp, single bool,
) (\*Delivery, error)](#func-googleauth-delivery)
    - [func (g GoogleAuth) Details(
  check \*GoogleCheckin, doc string, single bool,
) (\*Details, error)](#func-googleauth-details)
  - [type GoogleCheckin](#type-googlecheckin)
    - [func (GoogleCheckin) Marshal(device \*GoogleDevice) ([]byte, error)](#func-googlecheckin-marshal)
    - [func (g \*GoogleCheckin) Unmarshal(data []byte) error](#func-googlecheckin-unmarshal)
  - [type GoogleDevice](#type-googledevice)
    - [func (g \*GoogleDevice) Sync(check \*GoogleCheckin) error](#func-googledevice-sync)
  - [type GoogleToken](#type-googletoken)
    - [func (g \*GoogleToken) Auth() (\*GoogleAuth, error)](#func-googletoken-auth)
    - [func (GoogleToken) Marshal(token string) ([]byte, error)](#func-googletoken-marshal)
    - [func (g \*GoogleToken) Unmarshal(data []byte) error](#func-googletoken-unmarshal)
  - [type Obb](#type-obb)
    - [func (o Obb) Field1() uint64](#func-obb-field1)
    - [func (o Obb) Url() string](#func-obb-url)
  - [type StoreApp](#type-storeapp)
    - [func (s \*StoreApp) Apk(value string) string](#func-storeapp-apk)
    - [func (s \*StoreApp) Obb(value uint64) string](#func-storeapp-obb)
  - [type Values](#type-values)
    - [func (v Values) Set(query string) error](#func-values-set)
- [Source files](#source-files)

## Constants

com.roku.web.trc

```go
const Leanback = "android.software.leanback"
```

## Variables

developer.android.com/ndk/guides/abis

```go
var Abis = []string{

  "x86",
  "x86_64",

  "armeabi-v7a",

  "arm64-v8a",
}
```

```go
var Device = GoogleDevice{
  Feature: []string{

    "android.hardware.location.gps",

    "android.software.midi",

    "android.hardware.camera.front",

    "android.hardware.camera.flash",

    "android.hardware.microphone",

    "android.software.device_admin",

    "android.hardware.touchscreen",
    "android.hardware.wifi",

    "android.hardware.sensor.accelerometer",

    "android.hardware.camera",
    "android.hardware.location",
    "android.hardware.screen.portrait",

    "android.hardware.screen.landscape",

    "android.hardware.location.network",

    "android.hardware.bluetooth",
    "android.hardware.bluetooth_le",
    "android.hardware.camera.autofocus",
    "android.hardware.usb.host",

    "android.hardware.sensor.compass",

    "android.hardware.telephony",
  },
  Library: []string{

    "org.apache.http.legacy",

    "android.test.runner",
  },
  Texture: []string{

    "GL_OES_compressed_ETC1_RGB8_texture",

    "GL_KHR_texture_compression_astc_ldr",
  },
}
```

## Types

### type [Apk](./delivery.go#L22)

```go
type Apk struct {
  Message protobuf.Message
}
```

### func (Apk) [Field1](./delivery.go#L17)

```go
func (a Apk) Field1() string
```

### func (Apk) [Url](./delivery.go#L12)

```go
func (a Apk) Url() string
```

### type [Delivery](./delivery.go#L31)

```go
type Delivery struct {
  Message protobuf.Message
}
```

### func (Delivery) [Apk](./delivery.go#L43)

```go
func (d Delivery) Apk() func() (Apk, bool)
```

### func (Delivery) [Obb](./delivery.go#L35)

```go
func (d Delivery) Obb() func() (Obb, bool)
```

### func (Delivery) [Url](./delivery.go#L26)

```go
func (d Delivery) Url() string
```

### type [Details](./details.go#L147)

```go
type Details struct {
  Message protobuf.Message
}
```

### func (Details) [Downloads](./details.go#L31)

```go
func (d Details) Downloads() (uint64, bool)
```

### func (Details) [Name](./details.go#L38)

```go
func (d Details) Name() (string, bool)
```

### func (Details) [String](./details.go#L92)

```go
func (d Details) String() string
```

### type [GoogleAuth](./auth.go#L11)

```go
type GoogleAuth struct {
  Values Values
}
```

### func (GoogleAuth) [Acquire](./acquire.go#L11)

```go
func (g GoogleAuth) Acquire(checkin *GoogleCheckin, id string) error
```

### func (GoogleAuth) [Delivery](./delivery.go#L51)

```go
func (g GoogleAuth) Delivery(
  check *GoogleCheckin, app *StoreApp, single bool,
) (*Delivery, error)
```

### func (GoogleAuth) [Details](./details.go#L151)

```go
func (g GoogleAuth) Details(
  check *GoogleCheckin, doc string, single bool,
) (*Details, error)
```

### type [GoogleCheckin](./checkin.go#L62)

```go
type GoogleCheckin struct {
  Message protobuf.Message
}
```

### func (GoogleCheckin) [Marshal](./checkin.go#L11)

```go
func (GoogleCheckin) Marshal(device *GoogleDevice) ([]byte, error)
```

### func (\*GoogleCheckin) [Unmarshal](./checkin.go#L66)

```go
func (g *GoogleCheckin) Unmarshal(data []byte) error
```

### type [GoogleDevice](./play.go#L108)

```go
type GoogleDevice struct {
  Abi     string
  Feature []string
  Library []string
  Texture []string
}
```

### func (\*GoogleDevice) [Sync](./sync.go#L10)

```go
func (g *GoogleDevice) Sync(check *GoogleCheckin) error
```

### type [GoogleToken](./auth.go#L23)

```go
type GoogleToken struct {
  Values Values
}
```

### func (\*GoogleToken) [Auth](./auth.go#L52)

```go
func (g *GoogleToken) Auth() (*GoogleAuth, error)
```

### func (GoogleToken) [Marshal](./auth.go#L27)

```go
func (GoogleToken) Marshal(token string) ([]byte, error)
```

### func (\*GoogleToken) [Unmarshal](./auth.go#L47)

```go
func (g *GoogleToken) Unmarshal(data []byte) error
```

### type [Obb](./delivery.go#L105)

```go
type Obb struct {
  Message protobuf.Message
}
```

### func (Obb) [Field1](./delivery.go#L100)

```go
func (o Obb) Field1() uint64
```

### func (Obb) [Url](./delivery.go#L95)

```go
func (o Obb) Url() string
```

### type [StoreApp](./play.go#L146)

```go
type StoreApp struct {
  Id      string
  Version uint64
}
```

play.google.com/store/apps/details?id=com.google.android.apps.youtube.unplugged

### func (\*StoreApp) [Apk](./play.go#L119)

```go
func (s *StoreApp) Apk(value string) string
```

### func (\*StoreApp) [Obb](./play.go#L131)

```go
func (s *StoreApp) Obb(value uint64) string
```

### type [Values](./auth.go#L89)

```go
type Values map[string]string
```

### func (Values) [Set](./auth.go#L79)

```go
func (v Values) Set(query string) error
```

## Source files

[acquire.go](./acquire.go)
[auth.go](./auth.go)
[checkin.go](./checkin.go)
[delivery.go](./delivery.go)
[details.go](./details.go)
[play.go](./play.go)
[sync.go](./sync.go)
