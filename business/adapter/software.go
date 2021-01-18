package adapter

import (
	"context"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/thamthee/merchant/business/data/product"
	"github.com/thamthee/merchant/business/data/seller"
	"github.com/thamthee/merchant/business/graph/models"
)

func NewSoftwareToDB(ctx context.Context, nf models.NewSoftware) product.NewSoftware {
	return product.NewSoftware{
		Title:       nf.Title,
		Price:       nf.Price,
		Currency:    nf.Currency,
		Description: nf.Description,
		SKU:         nf.Sku,
		Stock:       nf.Stock,
		License:     nf.License,
		Code:        nf.Code,
	}
}

func SoftwareDBToGraph(ctx context.Context, sw product.Software, sl seller.Info) models.Software {
	price, err := strconv.ParseFloat(sw.Price.String(), 64)
	if err != nil {
		logrus.Errorln("software db: parse decimal to float:", err)
	}

	return models.Software{
		ID:          sw.ID.Hex(),
		Title:       sw.Title,
		Price:       price,
		Currency:    sw.Currency,
		Description: sw.Description,
		Sku:         sw.SKU,
		Stock:       sw.Stock,
		Owner: &models.Seller{
			ID:          sl.ID.Hex(),
			Name:        sl.Title,
			Slug:        sl.Slug,
			Description: sl.Description,
		},
		CreateAt: sw.CreatedAt,
		License:  sw.License,
		Code:     sw.Code,
	}
}

func ProductDBToSoftwareGraph(ctx context.Context, info product.Info, sl seller.Info) models.Software {
	price, err := strconv.ParseFloat(info.Price.String(), 64)
	if err != nil {
		logrus.Errorln("product db: parse decimal to float:", err)
	}

	return models.Software{
		ID:          info.ID.Hex(),
		Title:       info.Title,
		Price:       price,
		Currency:    info.Currency,
		Description: info.Description,
		Sku:         info.SKU,
		Stock:       info.Stock,
		Owner: &models.Seller{
			ID:          sl.ID.Hex(),
			Name:        sl.Title,
			Slug:        sl.Slug,
			Description: sl.Description,
		},
		CreateAt: info.CreatedAt,
		License:  info.License,
		Code:     info.Code,
	}
}
