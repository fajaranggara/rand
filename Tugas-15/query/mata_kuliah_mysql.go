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
	tabelMataKuliah          = "mata_kuliah"
	//layoutDateTime = "2006-01-02 15:04:05"
)

//GetAllNilai
func GetAllMataKuliah(ctx context.Context) ([]models.MataKuliah, error){
	var listMataKuliah []models.MataKuliah

	db, err := config.MySQL()

	//config db check
	if err != nil {
		log.Fatal("Cannot connect to mysql", err)
	}

	//query
	queryText := fmt.Sprintf("SELECT * FROM %v Order By created_at DESC", tabelMataKuliah)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next(){
		var mataKuliah models.MataKuliah
		//var createdAt, updatedAt string
		if err = rowQuery.Scan(&mataKuliah.ID,
		&mataKuliah.Nama, 
		&mataKuliah.CreatedAt, 
		&mataKuliah.UpdatedAt); err != nil {
			return nil, err
		}

		listMataKuliah = append(listMataKuliah, mataKuliah)
	}
	return listMataKuliah, nil
}


func InsertMataKuliah(ctx context.Context, mataKuliah models.MataKuliah) error{
	db, err := config.MySQL()

	//config db check
	if err != nil {
		log.Fatal("Cannot connect to mysql", err)
	}

	//query
	queryText := fmt.Sprintf("INSERT INTO %v (nama, created_at, updated_at) VALUES ('%v', NOW(), NOW())",
	tabelMataKuliah, mataKuliah.Nama)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil

}


func UpdateMataKuliah(ctx context.Context, mataKuliah models.MataKuliah, idMataKuliah string) error{
	db, err := config.MySQL()

	//config db check
	if err != nil {
		log.Fatal("Cannot connect to mysql", err)
	}

	queryText := fmt.Sprintf("UPDATE %v SET nama='%v', updated_at=NOW() WHERE ID=%v",
	tabelMataKuliah, mataKuliah.Nama, idMataKuliah)
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


func DeleteMataKuliah(ctx context.Context, idMataKuliah string) error{
	db, err := config.MySQL()

	//config db check
	if err != nil {
		log.Fatal("Cannot connect to mysql", err)
	}

	//query
	queryText := fmt.Sprintf("DELETE FROM %v WHERE ID=%v", tabelMataKuliah, idMataKuliah)
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