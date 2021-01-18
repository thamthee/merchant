package dataloader

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/thamthee/merchant/business/adapter"
	"github.com/thamthee/merchant/business/data/product"
	"github.com/thamthee/merchant/business/data/seller"
	"github.com/thamthee/merchant/business/graph/models"
)

//go:generate go run github.com/vektah/dataloaden SoftwareLoader string *github.com/thamthee/merchant/business/graph/models.Software
//go:generate go run github.com/vektah/dataloaden SellerLoader string *github.com/thamthee/merchant/business/graph/models.Seller

type ctxKey int

const Key = 2

type (
	Loader struct {
		SoftwareByID     *SoftwareLoader
		SellerBySoftware *SellerLoader
	}

	DataLoader struct {
		log *logrus.Logger
		p   product.Product
		s   seller.Seller
	}

	SoftwareKeysLoadFn func(keys []string) ([]*models.Software, []error)
)

func New(log *logrus.Logger, p product.Product, s seller.Seller) DataLoader {
	return DataLoader{
		log: log,
		p:   p,
		s:   s,
	}
}

func NewSwLoader(wait time.Duration, max int, fetch SoftwareKeysLoadFn) *SoftwareLoader {
	return &SoftwareLoader{
		fetch:    fetch,
		wait:     wait,
		maxBatch: max,
	}
}

func (dl DataLoader) LoadSoftwareByKeys(ctx context.Context) SoftwareKeysLoadFn {
	return func(keys []string) ([]*models.Software, []error) {
		sfs, err := dl.p.QuerySoftwareByIDs(ctx, keys)
		if err != nil {
			return nil, []error{err}
		}

		Sellers := make(map[string]seller.Info)
		for _, sf := range sfs {
			Sellers[sf.Seller] = seller.Info{}
		}

		SellerKeys := func(list map[string]seller.Info) []string {
			var ids []string
			for id, _ := range list {
				ids = append(ids, id)
			}
			return ids
		}(Sellers)

		sellers, err := dl.s.QueryByIDs(ctx, SellerKeys)
		if err != nil {
			return nil, []error{err}
		}

		for _, info := range sellers {
			Sellers[info.ID.Hex()] = info
		}

		res := make([]*models.Software, len(sfs))
		for i, sf := range sfs {
			software := adapter.SoftwareDBToGraph(ctx, sf, Sellers[sf.Seller])

			res[i] = &software
		}

		return res, nil
	}
}
