package handlers

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

const uploadPath = "./upload"
const indexFile = "index.html"

// если нет директории для загрузок то создать
func makeUploadDir() error {
	err := os.Mkdir(uploadPath, 0755)
	if err != nil && !errors.Is(err, os.ErrExist) {
		return err
	}
	return nil
}

// основной handler
func MainHandler(w http.ResponseWriter, r *http.Request) {
	indexFile, err := os.ReadFile(indexFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			errMesssage := "file" + string(indexFile) + " not found"
			log.Println(errMesssage)
			http.Error(w, errMesssage, http.StatusBadRequest)
			return
		} else {
			errMessage := fmt.Sprintf("internal server error on read %s: %v",
				indexFile, err)
			log.Println(errMessage)
			http.Error(w, errMessage, http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "text-plain")
	w.WriteHeader(http.StatusOK)
	w.Write(indexFile)
}

// handler для загрузки
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(100)
	mForm := r.MultipartForm
	for k := range mForm.File {
		file, fileHeader, err := r.FormFile(k)
		if err != nil {
			errMessage := fmt.Sprintf("invoke FormFile error: %v\n", err)
			log.Print(errMessage)
			http.Error(w, errMessage, http.StatusInternalServerError)
			return
		}
		defer file.Close()
		// -----------------------------------------------------------------
		// check & create upload dir
		err = makeUploadDir()
		if err != nil {
			errMessage := "couldn`t create upload dir"
			log.Println(errMessage, err)
			http.Error(w, errMessage, http.StatusInternalServerError)
			return
		}
		// -----------------------------------------------------------------
		// файл для сохранения результата
		extension := filepath.Ext(fileHeader.Filename)
		localFileName := uploadPath + "/" + time.Now().UTC().String() + extension
		out, err := os.Create(localFileName)
		if err != nil {
			errMessage := fmt.Sprintf("failed to open file %s for writing", localFileName)
			log.Println(errMessage)
			http.Error(w, errMessage, http.StatusInternalServerError)
			return
		}
		defer out.Close()
		// -----------------------------------------------------------------
		// читаем и конвертируем
		err = convertToFile(file, out)
		if err != nil {
			errMessage := fmt.Sprintf("error read file %s:\n", err)
			fmt.Print(errMessage)
			http.Error(w, errMessage, http.StatusInternalServerError)
		}
		// -----------------------------------------------------------------
		// все ок
		log.Printf("file %s uploaded ok\n", localFileName)
		log.Println("file uploaded & converted")

		w.Header().Set("Content-Type", "text-plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("file uploaded & converted"))
	}
}

// читаем через scanner и конвертируем
func convertToFile(r io.Reader, w io.Writer) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := scanner.Text()
		str, err := service.Convert(s)
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "%s", str)
	}
	return nil
}
