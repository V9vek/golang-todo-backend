package utils

import "net/http"

func GetParams(r *http.Request) map[string]string {
	if params, ok := r.Context().Value("params").(map[string]string); ok {
		return params
	}
	return nil
}
