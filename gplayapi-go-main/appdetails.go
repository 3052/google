package gplayapi

import (
	"net/http"

	"github.com/Juby210/gplayapi-go/gpproto"
)

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

func (client *GooglePlayClient) GetAppDetails(packageName string) (*App, error) {
	r, _ := http.NewRequest("GET", UrlDetails+"?doc="+packageName, nil)
	payload, err := client.doAuthedReq(r)
	if err != nil {
		return nil, err
	}
	if payload == nil {
		return nil, ErrNilPayload
	}
	return buildApp(payload.DetailsResponse), nil
}

func buildApp(res *gpproto.DetailsResponse) *App {
	return buildAppFromItem(res.Item)
}

func buildAppFromItem(item *gpproto.Item) *App {
	details := item.Details.AppDetails
	app := &App{
		PackageName:      *item.Id,
		CategoryID:       int(item.GetCategoryId()),
		DisplayName:      item.GetTitle(),
		Description:      item.GetDescriptionHtml(),
		ShortDescription: item.GetPromotionalDescription(),
		ShareUrl:         item.GetShareUrl(),
		VersionName:      details.GetVersionString(),
		VersionCode:      int(details.GetVersionCode()),
		CategoryName:     details.GetCategoryName(),
		Size:             details.GetInfoDownloadSize(),
		DownloadString:   details.GetDownloadLabelAbbreviated(),
		Changes:          details.GetRecentChangesHtml(),
		ContainsAds:      details.InstallNotes != nil,
		EarlyAccess:      details.EarlyAccessInfo != nil,
		DeveloperName:    details.GetDeveloperName(),
		TargetSdk:        int(details.GetTargetSdkVersion()),
		UpdatedOn:        details.GetInfoUpdatedOn(),
	}

	if len(item.Offer) > 0 {
		offer := item.Offer[0]
		app.OfferType = offer.GetOfferType()
		app.IsFree = offer.GetMicros() == 0
		app.Price = offer.GetFormattedAmount()
	}

	if app.DeveloperName == "" {
		app.DeveloperName = item.GetCreator()
	}

	if details.InstantLink != nil {
		app.InstantAppLink = details.GetInstantLink()
	}

	parseAppInfo(app, item)
	parseStreamUrls(app, item)
	parseImages(app, item)

	parseTestingProgram(app, details)

	return app
}

func parseAppInfo(app *App, item *gpproto.Item) {
	if item.AppInfo != nil {
		app.AppInfo = &AppInfo{map[string]string{}}
		for _, s := range item.AppInfo.Section {
			if s.Label != nil && s.Container != nil && s.GetContainer().Description != nil {
				app.AppInfo.AppInfoMap[s.GetLabel()] = s.GetContainer().GetDescription()
			}
		}
	}
}

func parseStreamUrls(app *App, item *gpproto.Item) {
	if item.Annotations != nil {
		app.LiveStreamUrl = item.Annotations.GetLiveStreamUrl()
		app.PromotionStreamUrl = item.Annotations.GetPromotionStreamUrl()
	}
}

func parseImages(app *App, item *gpproto.Item) {
	for _, image := range item.Image {
		switch image.GetImageType() {
		case ImageTypeCategoryIcon:
			app.CategoryImage = image
		case ImageTypeAppIcon:
			app.IconImage = image
		case ImageTypeYoutubeVideoThumbnail:
			app.Video = image
		case ImageTypePlayStorePageBackground:
			app.CoverImage = image
		case ImageTypeAppScreenshot:
			app.Screenshots = append(app.Screenshots, image)
		}
	}

	if len(app.Screenshots) == 0 {
		if item.Annotations != nil && item.Annotations.SectionImage != nil {
			for _, imageContainer := range item.Annotations.SectionImage.ImageContainer {
				app.Screenshots = append(app.Screenshots, imageContainer.Image)
			}
		}
	}
}

func parseTestingProgram(app *App, details *gpproto.AppDetails) {
	if details.TestingProgramInfo != nil {
		testingProgram := details.TestingProgramInfo
		app.TestingProgram = &TestingProgram{
			Image:                    testingProgram.Image,
			DisplayName:              testingProgram.GetDisplayName(),
			Email:                    testingProgram.GetEmail(),
			IsAvailable:              true,
			isSubscribed:             testingProgram.GetSubscribed(),
			IsSubscribedAndInstalled: testingProgram.GetSubscribedAndInstalled(),
		}
	}
}
