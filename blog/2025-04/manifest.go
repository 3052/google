package main

import "fmt"

func main() {
   fmt.Println("hello world")
}

type Manifest struct {
   XMLName                   xml.Name `xml:"manifest"`
   Text                      string   `xml:",chardata"`
   Android                   string   `xml:"android,attr"`
   VersionCode               string   `xml:"versionCode,attr"`
   VersionName               string   `xml:"versionName,attr"`
   CompileSdkVersion         string   `xml:"compileSdkVersion,attr"`
   CompileSdkVersionCodename string   `xml:"compileSdkVersionCodename,attr"`
   RequiredSplitTypes        string   `xml:"requiredSplitTypes,attr"`
   SplitTypes                string   `xml:"splitTypes,attr"`
   Package                   string   `xml:"package,attr"`
   PlatformBuildVersionCode  string   `xml:"platformBuildVersionCode,attr"`
   PlatformBuildVersionName  string   `xml:"platformBuildVersionName,attr"`
   UsesSdk                   struct {
      Text             string `xml:",chardata"`
      MinSdkVersion    string `xml:"minSdkVersion,attr"`
      TargetSdkVersion string `xml:"targetSdkVersion,attr"`
   } `xml:"uses-sdk"`
   UsesPermission []struct {
      Text string `xml:",chardata"`
      Name string `xml:"name,attr"`
   } `xml:"uses-permission"`
   UsesFeature struct {
      Text     string `xml:",chardata"`
      Name     string `xml:"name,attr"`
      Required string `xml:"required,attr"`
   } `xml:"uses-feature"`
   Queries struct {
      Text   string `xml:",chardata"`
      Intent []struct {
         Text   string `xml:",chardata"`
         Action struct {
            Text string `xml:",chardata"`
            Name string `xml:"name,attr"`
         } `xml:"action"`
         Category struct {
            Text string `xml:",chardata"`
            Name string `xml:"name,attr"`
         } `xml:"category"`
         Data struct {
            Text   string `xml:",chardata"`
            Scheme string `xml:"scheme,attr"`
         } `xml:"data"`
      } `xml:"intent"`
      Package struct {
         Text string `xml:",chardata"`
         Name string `xml:"name,attr"`
      } `xml:"package"`
   } `xml:"queries"`
   Permission struct {
      Text            string `xml:",chardata"`
      Name            string `xml:"name,attr"`
      ProtectionLevel string `xml:"protectionLevel,attr"`
   } `xml:"permission"`
   SupportsScreens struct {
      Text         string `xml:",chardata"`
      AnyDensity   string `xml:"anyDensity,attr"`
      SmallScreens string `xml:"smallScreens,attr"`
      LargeScreens string `xml:"largeScreens,attr"`
   } `xml:"supports-screens"`
   Application struct {
      Text                  string `xml:",chardata"`
      Theme                 string `xml:"theme,attr"`
      Label                 string `xml:"label,attr"`
      Icon                  string `xml:"icon,attr"`
      Name                  string `xml:"name,attr"`
      AllowTaskReparenting  string `xml:"allowTaskReparenting,attr"`
      AllowBackup           string `xml:"allowBackup,attr"`
      LargeHeap             string `xml:"largeHeap,attr"`
      SupportsRtl           string `xml:"supportsRtl,attr"`
      ExtractNativeLibs     string `xml:"extractNativeLibs,attr"`
      FullBackupContent     string `xml:"fullBackupContent,attr"`
      ResizeableActivity    string `xml:"resizeableActivity,attr"`
      NetworkSecurityConfig string `xml:"networkSecurityConfig,attr"`
      RoundIcon             string `xml:"roundIcon,attr"`
      AppComponentFactory   string `xml:"appComponentFactory,attr"`
      DataExtractionRules   string `xml:"dataExtractionRules,attr"`
      MetaData              []struct {
         Text     string `xml:",chardata"`
         Name     string `xml:"name,attr"`
         Value    string `xml:"value,attr"`
         Resource string `xml:"resource,attr"`
      } `xml:"meta-data"`
      Activity []struct {
         Text                  string `xml:",chardata"`
         Theme                 string `xml:"theme,attr"`
         Name                  string `xml:"name,attr"`
         Exported              string `xml:"exported,attr"`
         ScreenOrientation     string `xml:"screenOrientation,attr"`
         ConfigChanges         string `xml:"configChanges,attr"`
         ExcludeFromRecents    string `xml:"excludeFromRecents,attr"`
         LaunchMode            string `xml:"launchMode,attr"`
         NoHistory             string `xml:"noHistory,attr"`
         Label                 string `xml:"label,attr"`
         WindowSoftInputMode   string `xml:"windowSoftInputMode,attr"`
         ParentActivityName    string `xml:"parentActivityName,attr"`
         TaskAffinity          string `xml:"taskAffinity,attr"`
         AutoRemoveFromRecents string `xml:"autoRemoveFromRecents,attr"`
         Enabled               string `xml:"enabled,attr"`
         FitsSystemWindows     string `xml:"fitsSystemWindows,attr"`
         IntentFilter          []struct {
            Text       string `xml:",chardata"`
            Label      string `xml:"label,attr"`
            AutoVerify string `xml:"autoVerify,attr"`
            Action     struct {
               Text string `xml:",chardata"`
               Name string `xml:"name,attr"`
            } `xml:"action"`
            Category []struct {
               Text string `xml:",chardata"`
               Name string `xml:"name,attr"`
            } `xml:"category"`
            Data []struct {
               Text        string `xml:",chardata"`
               Scheme      string `xml:"scheme,attr"`
               Host        string `xml:"host,attr"`
               PathPrefix  string `xml:"pathPrefix,attr"`
               PathPattern string `xml:"pathPattern,attr"`
            } `xml:"data"`
         } `xml:"intent-filter"`
      } `xml:"activity"`
      Service []struct {
         Text                     string `xml:",chardata"`
         Name                     string `xml:"name,attr"`
         Exported                 string `xml:"exported,attr"`
         Label                    string `xml:"label,attr"`
         Permission               string `xml:"permission,attr"`
         Description              string `xml:"description,attr"`
         CanRetrieveWindowContent string `xml:"canRetrieveWindowContent,attr"`
         ForegroundServiceType    string `xml:"foregroundServiceType,attr"`
         DirectBootAware          string `xml:"directBootAware,attr"`
         Enabled                  string `xml:"enabled,attr"`
         VisibleToInstantApps     string `xml:"visibleToInstantApps,attr"`
         IntentFilter             struct {
            Text     string `xml:",chardata"`
            Priority string `xml:"priority,attr"`
            Action   struct {
               Text string `xml:",chardata"`
               Name string `xml:"name,attr"`
            } `xml:"action"`
            Category struct {
               Text string `xml:",chardata"`
               Name string `xml:"name,attr"`
            } `xml:"category"`
         } `xml:"intent-filter"`
         MetaData []struct {
            Text     string `xml:",chardata"`
            Name     string `xml:"name,attr"`
            Resource string `xml:"resource,attr"`
            Value    string `xml:"value,attr"`
         } `xml:"meta-data"`
      } `xml:"service"`
      Receiver []struct {
         Text            string `xml:",chardata"`
         Name            string `xml:"name,attr"`
         Exported        string `xml:"exported,attr"`
         Permission      string `xml:"permission,attr"`
         Enabled         string `xml:"enabled,attr"`
         DirectBootAware string `xml:"directBootAware,attr"`
         IntentFilter    []struct {
            Text   string `xml:",chardata"`
            Action []struct {
               Text string `xml:",chardata"`
               Name string `xml:"name,attr"`
            } `xml:"action"`
         } `xml:"intent-filter"`
         MetaData struct {
            Text  string `xml:",chardata"`
            Name  string `xml:"name,attr"`
            Value string `xml:"value,attr"`
         } `xml:"meta-data"`
      } `xml:"receiver"`
      Property struct {
         Text     string `xml:",chardata"`
         Name     string `xml:"name,attr"`
         Resource string `xml:"resource,attr"`
      } `xml:"property"`
      Provider []struct {
         Text            string `xml:",chardata"`
         Name            string `xml:"name,attr"`
         Exported        string `xml:"exported,attr"`
         Authorities     string `xml:"authorities,attr"`
         InitOrder       string `xml:"initOrder,attr"`
         DirectBootAware string `xml:"directBootAware,attr"`
      } `xml:"provider"`
      UsesLibrary []struct {
         Text     string `xml:",chardata"`
         Name     string `xml:"name,attr"`
         Required string `xml:"required,attr"`
      } `xml:"uses-library"`
   } `xml:"application"`
} 
