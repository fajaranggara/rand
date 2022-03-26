package main

import (
	"api-mysql/functions"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	
	router := httprouter.New()
	router.GET("/mahasiswa", functions.GetMahasiswa)
	router.POST("/mahasiswa/create", functions.PostMahasiswa)
	router.PUT("/mahasiswa/:id/update", functions.UpdateMahasiswa)
	router.DELETE("/mahasiswa/:id/delete", functions.DeleteMahasiswa)

	
	router.GET("/mata-kuliah", functions.GetMataKuliah)
	router.POST("/mata-kuliah/create", functions.PostMataKuliah)
	router.PUT("/mata-kuliah/:id/update", functions.UpdateMataKuliah)
	router.DELETE("/mata-kuliah/:id/delete", functions.DeleteMataKuliah)

	
	router.GET("/nilai", functions.GetNilai)
	router.POST("/nilai/create", functions.PostNilai)
	router.PUT("/nilai/:id/update", functions.UpdateNilai)
	router.DELETE("/nilai/:id/delete", functions.DeleteNilai)

	fmt.Println("server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}

