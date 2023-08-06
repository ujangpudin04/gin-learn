package upload

import (
	"encoding/json"
	"fmt"
	"gin-learn/service/result"
	"gin-learn/storage"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UploadService struct{
	storage *storage.Storage
}



func NewUpload(storage *storage.Storage) (*UploadService, error) {
	return &UploadService{storage: storage}, nil
}

// func (u *UploadService) GetUpload(c *gin.Context) {
// 	uploads := u.storage.GetDataUpload()
// 	c.JSON(http.StatusOK,uploads)
// }

func (u *UploadService) PostUpload(c *gin.Context){

	c.Writer.Header().Set("Content-Type", "application/json")
	// Get the file from the form data
	file, handler,err := c.Request.FormFile("filename")
	log.Println("file----",file)
	log.Println("handler----",handler.Filename)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()
	
	// setup file type filtering
	buff := make([]byte, 512)
	_, err = file.Read(buff)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	filetype := http.DetectContentType(buff)
		fmt.Println(filetype)
		if filetype != "image/jpeg" && filetype != "image/png" && filetype != "image/jpg" && filetype != "video/mp4" {
			c.Writer.WriteHeader(http.StatusBadRequest)
			response := result.ErrorResult{Code: http.StatusBadRequest, Message: "The provided file format is not allowed. Please upload a JPEG or PNG image"}
			json.NewEncoder(c.Writer).Encode(response)
			return
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			c.Writer.WriteHeader(http.StatusInternalServerError)
			response := result.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
			json.NewEncoder(c.Writer).Encode(response)
			return
		}

		// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		// fmt.Printf("File Size: %+v\n", handler.Size)
		// fmt.Printf("MIME Header: %+v\n", handler.Header)
		const MAX_UPLOAD_SIZE = 100 << 20 // 1MB
		// Parse our multipart form, 10 << 20 specifies a maximum
		// upload of 1 MB files.
		c.Request.ParseMultipartForm(MAX_UPLOAD_SIZE)
		if c.Request.ContentLength > MAX_UPLOAD_SIZE {
			c.Writer.WriteHeader(http.StatusBadRequest)
			response := Result{Code: http.StatusBadRequest, Message: "Max size in 1mb"}
			json.NewEncoder(c.Writer).Encode(response)
			return
		}

		// Create a temporary file within our temp-images directory that follows
		// a particular naming pattern
		// image/png
		// video/mp4
		fileTypeSplit := strings.Split(filetype, "/")
		tempFile, err :=os.CreateTemp("uploads", fileTypeSplit[0]+"-*."+fileTypeSplit[1])
		if err != nil {
			fmt.Println(err)
			fmt.Println("path upload error")
			json.NewEncoder(c.Writer).Encode(err)
			return
		}
		defer tempFile.Close()

		// read all of the contents of our uploaded file into a
		// byte array
		fileBytes, err := io.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}

		// write this byte array to our temporary file
		tempFile.Write(fileBytes)

		data := tempFile.Name()
		fmt.Println("ini data",data)    // uploads/image-89312783912.jpg
		filename := data[8:] // split uploads/



	// Define the path where the file will be saved
	// filePath := filepath.Join("uploads", file.Filename)
	log.Println("filename",filename)


	// Save the file to the defined path
	if err := c.SaveUploadedFile(handler,data); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	fileUpload := ImageDataRequest{
		Filename: data,
	}

	log.Println(fileUpload)

	validation := validator.New()
	errStruct:=validation.Struct(fileUpload)
	if errStruct!=nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file data"})
  		return
	}
	
	
	c.JSON(http.StatusOK, result.SuccessResult{
		Code: http.StatusOK,
		Data: fileUpload,
	})
	// u.storage.UploadImage()
	
	// Return a success message and the file metadata
}

type Result struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}