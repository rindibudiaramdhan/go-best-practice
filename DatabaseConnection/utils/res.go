package utils

import (
	"encoding/json"
	"net/http"
)

/**
* Kode di ini berguna untuk menampilkan data dengan bentuk JSON di browser. 
* Fungsi di atas nantinya akan kita panggil dari file main.go.
*/

func ResponseJSON(w http.ResponseWriter, p interface{}, status int) {
	ubahKeByte, err := json.Marshal(p)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		http.Error(w, "Error Mang", http.SatusBadRequest)
	}

}
w.Header().Set("Content-Type", "application/json")
w.writeHeader(status)
w.Write([]byte(ubahKeByte))