package handlers

import (
	"fmt"
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

		ctx.JSON(200, assetMetricsAggregate)
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
