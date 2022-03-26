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


func GetMataKuliah(rw http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()
	
	mataKuliah, err := query.GetAllMataKuliah(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(rw, mataKuliah, http.StatusOK)
}

func PostMataKuliah(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//pengecekan inputannya json atau bukan
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan Content-Type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	var mataKuliah models.MataKuliah

	if err := json.NewDecoder(r.Body).Decode(&mataKuliah); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	if err := query.InsertMataKuliah(ctx, mataKuliah); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"Status": "Successfully",
	}

	utils.ResponseJSON(rw, res, http.StatusOK)
}


func UpdateMataKuliah(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//pengecekan inputannya json atau bukan
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan Content-Type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	var mataKuliah models.MataKuliah

	if err := json.NewDecoder(r.Body).Decode(&mataKuliah); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	var idMataKuliah = ps.ByName("id")

	if err := query.UpdateMataKuliah(ctx, mataKuliah, idMataKuliah); err != nil {
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


func DeleteMataKuliah(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()


	var idMataKuliah = ps.ByName("id")

	if err := query.DeleteMataKuliah(ctx, idMataKuliah); err != nil {
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
