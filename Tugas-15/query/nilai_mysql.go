package query

import (
	"api-mysql/config"
	"api-mysql/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

const (
	tabelNilai          = "nilai"
)

//GetAllNilai
func GetAllNilai(ctx context.Context) ([]models.Nilai, error){
	var listNilai []models.Nilai

	db, err := config.MySQL()

	//config db check
	if err != nil {
		log.Fatal("Cannot connect to mysql", err)
	}

	//query
	queryText := fmt.Sprintf("SELECT * FROM %v Order By created_at DESC", tabelNilai)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next(){
		var nilai models.Nilai
		//var createdAt, updatedAt string
		if err = rowQuery.Scan(&nilai.ID,
		&nilai.Indeks,
		&nilai.Skor, 
		&nilai.CreatedAt, 
		&nilai.UpdatedAt,
		&nilai.MahasiswaId,
		&nilai.MataKuliahId); err != nil {
			return nil, err
		}

		listNilai = append(listNilai, nilai)
	}
	return listNilai, nil
}


func InsertNilai(ctx context.Context, nilai models.Nilai) error{
	db, err := config.MySQL()

	//config db check
	if err != nil {
		log.Fatal("Cannot connect to mysql", err)
	}

	//get index by skor
	indeks := indeksNilai(nilai.Skor)
	//query
	queryText := fmt.Sprintf("INSERT INTO %v (skor, indeks, created_at, updated_at, mahasiswa_id, mata_kuliah_id) VALUES ('%v', '%v', NOW(), NOW(), %v, %v)",
	tabelNilai, nilai.Skor, indeks, nilai.MahasiswaId, nilai.MataKuliahId)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil

}

// menambahkan index nilai
func indeksNilai(nilai uint) string{
	var index string
	switch{
	case nilai >= 80: index = "A"
	case nilai >= 70: index = "B"
	case nilai >= 60: index = "C"
	case nilai >= 50: index = "D"
	default: index = "E"
	}
	return index
}

func UpdateNilai(ctx context.Context, nilai models.Nilai, idNilai string) error{
	db, err := config.MySQL()

	//config db check
	if err != nil {
		log.Fatal("Cannot connect to mysql", err)
	}

	//get index by skor
	indeks := indeksNilai(nilai.Skor)
	queryText := fmt.Sprintf("UPDATE %v SET indeks='%v', skor='%v', mahasiswa_id=%v, mata_kuliah_id=%v, updated_at=NOW() WHERE ID=%v",
	tabelNilai, indeks, nilai.Skor, nilai.MahasiswaId, nilai.MataKuliahId, idNilai)
	s, err := db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	check, err := s.RowsAffected()

	if check == 0 {
		return errors.New("id tidak ditemukan, tidak ada yang diupdate")
	}
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}


func DeleteNilai(ctx context.Context, idNilai string) error{
	db, err := config.MySQL()

	//config db check
	if err != nil {
		log.Fatal("Cannot connect to mysql", err)
	}

	//query
	queryText := fmt.Sprintf("DELETE FROM %v WHERE ID=%v", tabelNilai, idNilai)
	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}
	
	check, err := s.RowsAffected()

	if check == 0 {
		return errors.New("id tidak ditemukan")
	}
	if err != nil {
		fmt.Println(err.Error())
	}

	return nil

}