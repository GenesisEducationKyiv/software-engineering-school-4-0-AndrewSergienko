package src

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetCurrencyHandler(cr CurrencyReader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			_, err := fmt.Fprintf(w, "Method Not Allowed")
			if err != nil {
				return
			}
		}

		rate, err := cr.getUSDCurrencyRate()
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
