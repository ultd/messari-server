package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/ultd/messari-server/messari"
)

// GetAllAssetsHandler func is a gin route controller for handling assets
func GetAllAssetsHandler(apiKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page := ctx.Query("page")
		m := messari.New(apiKey)
		var opts *messari.GetAllAssetsOptions
		if page != "" {
			pg, err := strconv.Atoi(page)
			if err != nil {
				ctx.JSON(400, gin.H{"message": "Invalid page specified in query."})
				return
			}
			opts = &messari.GetAllAssetsOptions{
				Page: intPtr(pg),
			}
		}

		resp, err := m.GetAllAssets(opts)
		if err != nil {
			logrus.Error(err)
			ctx.JSON(400, gin.H{"message": "An error occured getting all asset metrics."})
			return
		}
		ctx.JSON(200, resp.Data)
	}
}

// GetAssetHandler func returns a hanlder for getting metadata of an asset using a symbol or a slug
func GetAssetHandler(apiKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		symbolOrSlug := ctx.Param("symbolOrSlug")
		if symbolOrSlug == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "No slug or symbol provided in URL."})
			return
		}
		m := messari.New(apiKey)
		resp, err := m.GetAsset(symbolOrSlug, &messari.GetAssetOptions{
			// have to put all fields in single string comma seperated as API dictates
			// Fields: []string{"symbol,name"},
			Fields: nil,
		})
		if err != nil {
			logrus.Error(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("An error occured getting %s's metrics.", symbolOrSlug)})
			return
		}
		ctx.JSON(200, resp.Data)
	}
}

// GetAssetMetricsHandler func returns a hanlder for getting metrics of an asset using a symbol or a slug
func GetAssetMetricsHandler(apiKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		symbolOrSlug := ctx.Param("symbolOrSlug")
		if symbolOrSlug == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "No slug or symbol provided in URL."})
			return
		}
		m := messari.New(apiKey)
		resp, err := m.GetAssetMetrics(symbolOrSlug, &messari.GetAssetMetricsOptions{
			// have to put all fields in single string comma seperated as API dictates
			// Fields: []string{"symbol,name"},
			Fields: nil,
		})
		if err != nil {
			logrus.Error(err)
			ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("An error occured getting %s's metrics.", symbolOrSlug)})
			return
		}
		ctx.JSON(200, resp.Data)
	}
}

// AssetAggregateMetrics struct is the response json of GetAssetMetricsAggregateHandler
type assetAggregateMetrics struct {
	Tags                 []string `json:"tags,omitempty"`
	Sector               []string `json:"sector,omitempty"`
	Volume               float64  `json:"volume,omitempty"`
	TwentyFourHourChange float64  `json:"24HourChange,omitempty"`
	MarketCap            float64  `json:"marketcap,omitempty"`
}

// GetAssetMetricsAggregateHandler func returns a hanlder for getting metrics of an asset using a symbol or a slug
func GetAssetMetricsAggregateHandler(apiKey string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		symbolOrSlug := ctx.Param("symbolOrSlug")

		tags := ctx.Query("tags")
		sector := ctx.Query("sector")

		m := messari.New(apiKey)

		page := 1

		var assetMetricsAggregate []messari.Asset = make([]messari.Asset, 0, 500)
		var resp *messari.GetAllAssetsResp
		var err error
		for {
			if resp != nil && resp.Data == nil {
				break
			}
			resp, err = m.GetAllAssets(&messari.GetAllAssetsOptions{
				Page:             &page,
				Limit:            intPtr(500),
				Fields:           []string{"id,name,symbol,slug,metrics,profile/general/overview/tags,profile/general/overview/sector"},
				WithMertricsOnly: boolPtr(true),
				WithProfilesOnly: boolPtr(true),
			})
			if err != nil {
				logrus.Error(err)
				ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("An error occured getting %s's metrics.", symbolOrSlug)})
				return
			}
			for _, asset := range resp.Data {
				// if getting all asset metrics (not filtered by tags or sector), only return assets with
				// greater than 20,000,000 market cap.
				if tags == "" && sector == "" && asset.Metrics.Marketcap.CurrentMarketcapUsd < 20000000 {
					break
				}

				// if only tags specified in query
				if tags != "" && sector == "" {
					if asset.Profile.General.Overview.Tags != nil && *asset.Profile.General.Overview.Tags == tags {
						assetMetricsAggregate = append(assetMetricsAggregate, asset)
					}

					// if only sector specified in query
				} else if tags == "" && sector != "" {
					if asset.Profile.General.Overview.Sector != nil && *asset.Profile.General.Overview.Sector == sector {
						assetMetricsAggregate = append(assetMetricsAggregate, asset)
					}

					// if tags & sector specified in query
				} else if tags != "" && sector != "" {
					if asset.Profile.General.Overview.Sector != nil && *asset.Profile.General.Overview.Sector == sector &&
						asset.Profile.General.Overview.Tags != nil && *asset.Profile.General.Overview.Tags == tags {
						assetMetricsAggregate = append(assetMetricsAggregate, asset)
					}

					// if tags & sector NOT specified in query, append as normal
				} else {
					assetMetricsAggregate = append(assetMetricsAggregate, asset)
				}
			}
			page++
		}

		allTags := []string{}
		allSectors := []string{}
		volume := 0.0
		marketCap := 0.0
		twentyFourHourChangeAgg := 0.0

		for _, asset := range assetMetricsAggregate {
			volume += asset.Metrics.MarketData.VolumeLast24Hours
			marketCap += asset.Metrics.Marketcap.CurrentMarketcapUsd
			twentyFourHourChangeAgg += asset.Metrics.MarketData.PercentChangeUsdLast24Hours
			if asset.Profile.General.Overview.Tags != nil &&
				*asset.Profile.General.Overview.Tags != "" &&
				!includesString(allTags, *asset.Profile.General.Overview.Tags) {
				allTags = append(allTags, *asset.Profile.General.Overview.Tags)
			}
			if asset.Profile.General.Overview.Sector != nil &&
				*asset.Profile.General.Overview.Sector != "" &&
				!includesString(allSectors, *asset.Profile.General.Overview.Sector) {
				allSectors = append(allSectors, *asset.Profile.General.Overview.Sector)
			}
		}

		agg := &assetAggregateMetrics{
			Tags:                 allTags,
			Sector:               allSectors,
			Volume:               normalizeFloat(volume),
			MarketCap:            normalizeFloat(marketCap),
			TwentyFourHourChange: normalizeFloat(twentyFourHourChangeAgg / float64(len(assetMetricsAggregate))),
		}

		ctx.JSON(200, agg)
	}
}

func intPtr(v int) *int {
	return &v
}

func strPtr(v string) *string {
	return &v
}

func boolPtr(v bool) *bool {
	return &v
}

func normalizeFloat(v float64) float64 {
	return math.Round(v*100) / 100
}

func includesString(v []string, s string) bool {
	for _, val := range v {
		if val == s {
			return true
		}
	}
	return false
}
