package health

import "net/http"

// Status godoc
//
//	@summary        Status health
//	@description    Status health
//	@tags           health
//	@success        200
//	@router         /../status [get]
func Status(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I'm fine"))
}
