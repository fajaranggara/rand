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
	tableMahasiswa          = "mahasiswa"
	layoutDateTime = "2006-01-02 15:04:05"
)

//GetAllNilai
func GetAllMahasiswa(ctx context.Context) ([]models.Mahasiswa, error){
	var listMahasiswa []models.Mahasiswa

	db, err := config.MySQL()

	//config db check
	if err != nil {
		log.Fatal("Cannot connect to mysql", err)
	}

	//query
	queryText := fmt.Sprintf("SELECT * FROM %v Order By created_at DESC", tableMahasiswa)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next(){
		var mahasiswa models.Mahasiswa
		//var createdAt, updatedAt string
		if err = rowQuery.Scan(&mahasiswa.ID,
		&mahasiswa.Nama, 
		&mahasiswa.CreatedAt, 
		&mahasiswa.UpdatedAt); err != nil {
			return nil, err
		}

		listMahasiswa = append(listMahasiswa, mahasiswa)
	}
	return listMahasiswa, nil
}


func InsertMahasiswa(ctx context.Context, mahasiswa models.Mahasiswa) error{
	db, err := config.MySQL()

	//config db check
	if err != nil {
		log.Fatal("Cannot connect to mysql", err)
	}

	//query
	queryText := fmt.Sprintf("INSERT INTO %v (nama, created_at, updated_at) VALUES ('%v', NOW(), NOW())",
	tableMahasiswa, mahasiswa.Nama)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}

	return nil

}


func UpdateMahasiswa(ctx context.Context, mahasiswa models.Mahasiswa, idMahasiswa string) error{
	db, err := config.MySQL()

	//config db check
	if err != nil {
		log.Fatal("Cannot connect to mysql", err)
	}

	queryText := fmt.Sprintf("UPDATE %v SET nama='%v', updated_at=NOW() WHERE id=%v",
	tableMahasiswa, mahasiswa.Nama, idMahasiswa)
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


func DeleteMahasiswa(ctx context.Context, idMahasiswa string) error{
	db, err := config.MySQL()

	//config db check
	if err != nil {
		log.Fatal("Cannot connect to mysql", err)
	}

	//query
	queryText := fmt.Sprintf("DELETE FROM %v WHERE ID=%v", tableMahasiswa, idMahasiswa)
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