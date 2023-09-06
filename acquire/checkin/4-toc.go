package play

import (
   "errors"
   "net/http"
   "net/url"
)

func (c checkin) toc(v url.Values) error {
   req := new(http.Request)
   req.Header = make(http.Header)
   req.Header["Accept-Language"] = []string{"en-GB"}
   req.Header["Host"] = []string{"android.clients.google.com"}
   req.Header["User-Agent"] = []string{"Android-Finsky/15.8.23-all [0] [PR] 259261889 (api=3,versionCode=81582300,sdk=28,device=sargo,hardware=sargo,product=sargo,platformVersionRelease=9,model=Pixel 3a,buildId=PQ3B.190705.003,isWideScreen=0,supportedAbis=arm64-v8a;armeabi-v7a;armeabi)"}
   req.Header["X-Ad-Id"] = []string{"LawadaMera"}
   req.Header["X-Dfe-Client-Id"] = []string{"am-android-google"}
   req.Header["X-Dfe-Content-Filters"] = []string{""}
   req.Header["X-Dfe-Network-Type"] = []string{"4"}
   req.Header["X-Dfe-Request-Params"] = []string{"timeoutMs=4000"}
   req.Header["X-Dfe-Userlanguages"] = []string{"en_GB"}
   req.Header["X-Limit-Ad-Tracking-Enabled"] = []string{"false"}
   req.ProtoMajor = 1
   req.ProtoMinor = 1
   req.URL = new(url.URL)
   req.URL.Host = "android.clients.google.com"
   req.URL.Path = "/fdfe/toc"
   req.URL.Scheme = "https"
   req.Header["Authorization"] = []string{"Bearer " + v.Get("Auth")}
   {
      id, err := c.id()
      if err != nil {
         return err
      }
      req.Header["X-Dfe-Device-Id"] = []string{id}
   }
   req.Header["X-Dfe-Mccmnc"] = []string{"20815"}
   
   req.Header["X-Dfe-Device-Checkin-Consistency-Token"] = []string{"ABFEt1WnR20Eg_6baiIK6y1fzRMEEzRoE-16eaUI2XMN537Wqcii9WCOSAb4wVAYpgV2QqchXRPHzlqnpyVZ3a77fD1UQw0QdXtvWbxKb2g2aKhxwzyUVqDj0BgHIvcYo2GYEx1gdpTpHkJsOo2feQJO3JCZfd6rvCmyvpeeNp2VxVc78D2mYBxz2MUTWZjZeKe_kHkeWJoyaG6WpIb6jgasPQH1C3YtvgX4WfaH78mALP2Fat8Hr9L9RxzIzE0OBhFgsD_gNW-jTjR08z_Z0id9FZX47RQuM-u44304cM-WqIHPDL1_bE8"}
   req.Header["X-Dfe-Device-Config-Token"] = []string{"CjkaNwoTMzU5NTc4OTg0MTIyNjg5NDc4ORIgChAxNjkzOTYxMzg0NzcxODQ2EgwIqJnfpwYQ8N6F8AI="}
   req.Header["X-Dfe-Encoded-Targets"] = []string{"CAESN/qigQYC2AMBFfUbyA7SM5Ij/CvfBoIDgxHqGP8R3xzIBvoQtBKFDZ4HAY4FrwSVMasHBO0O2Q8akgYRAQECAQO7AQEpKZ0CnwECAwRrAQYBr9PPAoK7sQMBAQMCBAkIDAgBAwEDBAICBAUZEgMEBAMLAQEBBQEBAcYBARYED+cBfS8CHQEKkAEMMxcBIQoUDwYHIjd3DQ4MFk0JWGYZEREYAQOLAYEBFDMIEYMBAgICAgICOxkCD18LGQKEAcgDBIQBAgGLARkYCy8oBTJlBCUocxQn0QUBDkkGxgNZQq0BZSbeAmIDgAEBOgGtAaMCDAOQAZ4BBIEBKUtQUYYBQscDDxPSARA1oAEHAWmnAsMB2wFyywGLAxol+wImlwOOA80CtwN26A0WjwJVbQEJPAH+BRDeAfkHK/ABASEBCSAaHQemAzkaRiu2Ad8BdXeiAwEBGBUBBN4LEIABK4gB2AFLfwECAdoENq0CkQGMBsIBiQEtiwGgA1zyAUQ4uwS8AwhsvgPyAcEDF27vApsBHaICGhl3GSKxAR8MC6cBAgItmQYG9QIeywLvAeYBDArLAh8HASI4ELICDVmVBgsY/gHWARtcAsMBpALiAdsBA7QBpAJmIArpByn0AyAKBwHTARIHAX8D+AMBcRIBBbEDmwUBMacCHAciNp0BAQF0OgQLJDuSAh54kwFSP0eeAQQ4M5EBQgMEmwFXywFo0gFyWwMcapQBBugBPUW2AVgBKmy3AR6PAbMBGQxrUJECvQR+8gFoWDsYgQNwRSczBRXQAgtRswEW0ALMAREYAUEBIG6yATYCRE8OxgER8gMBvQEDRkwLc8MBTwHZAUOnAXiiBakDIbYBNNcCIUmuArIBSakBrgFHKs0EgwV/G3AD0wE6LgECtQJ4xQFwFbUCjQPkBS6vAQqEAUZF3QIM9wEhCoYCQhXsBCyZArQDugIziALWAdIBlQHwBdUErQE6qQaSA4EEIvYBHir9AQVLmgMCApsCKAwHuwgrENsBAjNYswEVmgIt7QJnN4wDEnta+wGfAcUBxgEtEFXQAQWdAUAeBcwBAQM7rAEJATJ0LENrdh73A6UBhAE+qwEeASxLZUMhDREuH0CGARbd7K0GlQo"}
   req.Header["X-Dfe-Phenotype"] = []string{"H4sIAAAAAAAAAB3OO3KjMAAA0KRNuWXukBkBQkAJ2MhgAZb5u2GCwQZbCH_EJ77QHmgvtDtbv-Z9_H63zXXU0NVPB1odlyGy7751Q3CitlPDvFd8lxhz3tpNmz7P92CFw73zdHU2Ie0Ad2kmR8lxhiErTFLt3RPGfJQHSDy7Clw10bg8kqf2owLokN4SecJTLoSwBnzQSd652_MOf2d1vKBNVedzg4ciPoLz2mQ8efGAgYeLou-l-PXn_7Sna1MfhHuySxt-4esulEDp8Sbq54CPPKjpANW-lkU2IZ0F92LBI-ukCKSptqeq1eXU96LD9nZfhKHdtjSWwJqUm_2r6pMHOxk01saVanmNopjX3YxQafC4iC6T55aRbC8nTI98AF_kItIQAJb5EQxnKTO7TZDWnr01HVPxelb9A2OWX6poidMWl16K54kcu_jhXw-JSBQkVcD_fPsLSZu6joIBAAA"}
   
   res, err := http.DefaultClient.Do(req)
   if err != nil {
      return err
   }
   defer res.Body.Close()
   if res.StatusCode != http.StatusOK {
      return errors.New(res.Status)
   }
   return nil
}
