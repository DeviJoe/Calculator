package controller

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

func RestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		b, err := io.ReadAll(r.Body)
		if err != nil {
			slog.Error("Cannot read body from request", "err", err.Error())
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var input []map[string]interface{}
		err = json.Unmarshal(b, &input)
		if err != nil {
			slog.Error("Cannot unmarshal body", "err", err.Error())
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		res, err := CalculateSentence(r.Context(), input...)

		if err != nil {
			slog.Error("Cannot calculate sentence", "err", err.Error())
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		type Item struct {
			Var   string `json:"var"`
			Value int64  `json:"value"`
		}
		type Result struct {
			Items []Item `json:"items"`
		}
		result := &Result{
			make([]Item, 0),
		}

		for key, item := range res {
			result.Items = append(result.Items, Item{
				key,
				item,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(result)
		if err != nil {
			slog.Error("Cannot parse result to json", "err", err.Error())
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
		}
		slog.Info("Successfully calculated sentence", "result", res, "ip", r.RemoteAddr)
		return
	}
}
