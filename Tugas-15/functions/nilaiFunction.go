package functions

import (
	"api-mysql/models"
	"api-mysql/query"
	"api-mysql/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)


func GetNilai(rw http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()
	
	nilai, err := query.GetAllNilai(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(rw, nilai, http.StatusOK)
}

func PostNilai(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//pengecekan inputannya json atau bukan
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan Content-Type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	var nilai models.Nilai

	if err := json.NewDecoder(r.Body).Decode(&nilai); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	if nilai.Skor > 100 {
		http.Error(rw, "Nilai tidak boleh lebih dari 100", http.StatusBadRequest)
		return
	}

	if err := query.InsertNilai(ctx, nilai); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"Status": "Successfully",
	}

	utils.ResponseJSON(rw, res, http.StatusOK)
}


func UpdateNilai(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//pengecekan inputannya json atau bukan
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan Content-Type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	var nilai models.Nilai

	if err := json.NewDecoder(r.Body).Decode(&nilai); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	if nilai.Skor > 100 {
		http.Error(rw, "Nilai tidak boleh lebih dari 100", http.StatusBadRequest)
		return
	}

	var idNilai = ps.ByName("id")

	if err := query.UpdateNilai(ctx, nilai, idNilai); err != nil {
		errUpdate := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(rw, errUpdate, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"Status": "Successfully",
	}

	utils.ResponseJSON(rw, res, http.StatusOK)
}


func DeleteNilai(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	var idNilai = ps.ByName("id")

	if err := query.DeleteNilai(ctx, idNilai); err != nil {
		errDelete := map[string]string{
			"error": fmt.Sprintf("%v", err),
		}
		utils.ResponseJSON(rw, errDelete, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"Status": "Successfully",
	}

	utils.ResponseJSON(rw, res, http.StatusOK)
}
