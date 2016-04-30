package app

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/favclip/ucon"
)

func init() {
	ucon.Orthodox()

	ucon.HandleFunc("POST", "/command", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		data := map[string]string{
			"response_type": "in_channel",
		}
		if t := r.PostFormValue("text"); strings.Contains(t, "オープンソース") {
			data["text"] = "焼きそば！"
		} else {
			data["text"] = "オープンソース！"
		}

		var buf bytes.Buffer
		json.NewEncoder(&buf).Encode(data)
		w.Write(buf.Bytes())
	})

	ucon.DefaultMux.Prepare()
	http.Handle("/", ucon.DefaultMux)
}
