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


func GetMahasiswa(rw http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()
	
	mahasiswa, err := query.GetAllMahasiswa(ctx)

	if err != nil {
		fmt.Println(err)
	}

	utils.ResponseJSON(rw, mahasiswa, http.StatusOK)
}

func PostMahasiswa(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//pengecekan inputannya json atau bukan
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan Content-Type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	var mahasiswa models.Mahasiswa

	if err := json.NewDecoder(r.Body).Decode(&mahasiswa); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	if err := query.InsertMahasiswa(ctx, mahasiswa); err != nil {
		utils.ResponseJSON(rw, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"Status": "Successfully",
	}

	utils.ResponseJSON(rw, res, http.StatusOK)
}


func UpdateMahasiswa(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//pengecekan inputannya json atau bukan
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "Gunakan Content-Type application/json", http.StatusBadRequest)
		return
	}

	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()

	var mahasiswa models.Mahasiswa

	if err := json.NewDecoder(r.Body).Decode(&mahasiswa); err != nil {
		utils.ResponseJSON(rw, err, http.StatusBadRequest)
		return
	}

	var idMahasiswa = ps.ByName("id")
	if err := query.UpdateMahasiswa(ctx, mahasiswa, idMahasiswa); err != nil {
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


func DeleteMahasiswa(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx, cancle := context.WithCancel(context.Background())
	defer cancle()


	var idMahasiswa = ps.ByName("id")

	if err := query.DeleteMahasiswa(ctx, idMahasiswa); err != nil {
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
