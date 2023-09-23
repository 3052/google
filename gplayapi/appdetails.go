package gplayapi

import "154.pages.dev/google/gplayapi/gpproto"

type (
   App struct {
      PackageName        string
      AppInfo            *AppInfo
      CategoryImage      *gpproto.Image
      CategoryID         int
      CategoryName       string
      Changes            string
      ContainsAds        bool
      CoverImage         *gpproto.Image
      Description        string
      DeveloperName      string
      DisplayName        string
      DownloadString     string
      EarlyAccess        bool
      IconImage          *gpproto.Image
      InstantAppLink     string
      IsFree             bool
      IsSystem           bool
      LiveStreamUrl      string
      OfferDetails       map[string]string
      OfferType          int32
      Price              string
      PromotionStreamUrl string
      Screenshots        []*gpproto.Image
      ShareUrl           string
      ShortDescription   string
      Size               int64
      TargetSdk          int
      TestingProgram     *TestingProgram
      UpdatedOn          string
      VersionCode        int
      VersionName        string
      Video              *gpproto.Image
   }

   AppInfo struct {
      AppInfoMap map[string]string
   }

   TestingProgram struct {
      Image                    *gpproto.Image
      DisplayName              string
      Email                    string
      IsAvailable              bool
      isSubscribed             bool
      IsSubscribedAndInstalled bool
   }

   TestingProgramStatus struct {
      Subscribed   bool
      Unsubscribed bool
   }
)
