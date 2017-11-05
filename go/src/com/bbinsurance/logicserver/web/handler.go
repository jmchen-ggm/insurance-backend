package web

import (
	"com/bbinsurance/log"
	"com/bbinsurance/logicserver/database"
	"fmt"
	"io"
	"net/http"
	"os"
)

func HandleCreateArticle(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		log.Error("Invalid Request Method: %s Url: %s", request.Method, request.URL)
		fmt.Fprintf(writer, "Invalid Requst, Please Use Http Get")
		return
	}
	request.ParseMultipartForm(32 << 20)
	file, fileHandler, err := request.FormFile("uploadfile")
	if err != nil {
		log.Error("Invalid File %s", err)
		fmt.Fprintf(writer, "Invalid Requst File")
		return
	}
	defer file.Close()
	title := request.FormValue("title")
	desc := request.FormValue("desc")
	url := request.FormValue("url")

	log.Info("CreateArticle: title=%s desc=%s url=%s file=%s", title, desc, url, fileHandler.Header)

	id, err := database.InsertArticle(title, desc, url, "")
	if err != nil {
		log.Error("Invalid File %s", err)
		fmt.Fprintf(writer, "Server Error")
		return
	}
	thumbUrl := fmt.Sprint(id) + ".png"
	database.UpdateArticleThumbUrl(id, thumbUrl)
	fis, err := os.OpenFile("./"+thumbUrl, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Error("Save File Err %s", err)
		fmt.Fprintf(writer, "Server Error")
		return
	}
	defer fis.Close()
	io.Copy(fis, file)
	fmt.Fprintf(writer, "Handle Success")
}
