package api

import (
	"fmt"
	"strconv"
	"tugas/model"

	"tugas/db"
	"tugas/helpers"

	"net/http"

	"github.com/xuri/excelize/v2"

	"github.com/labstack/echo"
)

func MusliGGWP(c echo.Context) error {
	var kirim error

	//excelize dengan formfile untuk open file yang sudah dimasukan dalam postman
	db := db.DbManager() // db manager gorm
	book, _ := c.FormFile("EXCEL")
	MusliGanteng, err := book.Open() //multipart file header
	if err != nil {
		return err

	}

	//membaca sheet == excelize
	f, err := excelize.OpenReader(MusliGanteng)
	if err != nil {
		fmt.Println(err)
		return err

	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	var i int = 1
	// i := 1
	for {
		i++                                                          //infinite loop. pada saat membaca isi dari excel
		book, err := f.GetCellValue("Sheet1", fmt.Sprintf("B%d", i)) //membaca excel perkolom
		if err != nil {                                              //jika ada err maka akan mengembalikan nilai err
			// fmt.Println(err)
			return err
		}
		fmt.Println(book)

		author, err := f.GetCellValue("Sheet1", fmt.Sprintf("C%d", i)) // membaca excel perkolom
		if err != nil {                                                //jika ada err maka akan mengembalikan err
			fmt.Println(err)
			return err
		}
		if book == "" && author == "" { //jika nilai pada book atau author (kolom) koosong maka akan berhenti
			break
		}

		// fmt.Println(author)

		books := model.Books{ //setelah dibaca akan masuk ke dalam struct dan akan membuat db
			Books:   book,
			Authors: author,
		}
		db.Create(&books)

	}
	response := helpers.Response{ //jika berhasil maka akan membuat status err ok
		StatusCode: http.StatusOK,
		Message:    "ok",
		Data:       nil,
	}
	kirim = c.JSON(http.StatusOK, response)
	return kirim
}

func GetAllBooks(c echo.Context) error {
	page := c.QueryParam("page") //mengambil semua data yang ada di dalam db

	num, err := strconv.Atoi(page)
	if err != nil {
		fmt.Println(num)
	}

	limit := c.QueryParam("limit")
	xxx, err := strconv.Atoi(limit)
	if err != nil {
		fmt.Println(num)
	}

	db := db.DbManager()
	bookk := []model.Books{}

	var kirim error
	pagas := (num - 1) * xxx
	err = db.Limit(limit).Offset(pagas).Find(&bookk).Error
	response := helpers.Response{}

	if err != nil {
		response = helpers.Response{
			StatusCode: http.StatusNoContent,
			Message:    "failed",
			Data:       err.Error(),
			Halaman:    num,
		}
		kirim = c.JSON(http.StatusNoContent, response)
	} else {
		response = helpers.Response{
			StatusCode: http.StatusOK,
			Message:    "ok",
			Data:       bookk,
			Halaman:    num,
		}
		kirim = c.JSON(http.StatusOK, response)
	}

	return kirim
}
