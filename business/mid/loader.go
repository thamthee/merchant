package mid

import (
	"context"
	"net/http"
	"time"

	"github.com/thamthee/merchant/business/dataloader"
)

func Loader(dl dataloader.DataLoader) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			loader := dataloader.Loader{}

			wait := 5 * time.Second

			loader.SoftwareByID = dataloader.NewSwLoader(
				wait,
				2,
				dl.LoadSoftwareByKeys(r.Context()))

			c := context.WithValue(r.Context(), dataloader.Key, &loader)

			next.ServeHTTP(w, r.WithContext(c))
		})
	}
}
