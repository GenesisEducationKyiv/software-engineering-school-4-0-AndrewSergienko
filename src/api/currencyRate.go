package api

import (
	"encoding/json"
	"fmt"
	"go_service/src/common"
	"net/http"
)

func GetCurrencyHandler(cr common.CurrencyReader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			_, err := fmt.Fprintf(w, "Method Not Allowed")
			if err != nil {
				return
			}
		}

		rate, err := cr.GetUSDCurrencyRate()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		response := map[string]interface{}{
			"rate": rate,
		}

		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			return
		}
	}
}
