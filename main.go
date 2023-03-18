package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os" // Sustituye ioutil (deprecated)

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

type Config struct {
	WebPort          string `json:"webPort,omitempty"`
	WebHost          string `json:"webHost,omitempty"`
	BucketName       string `json:"bucketName,omitempty"`
	BucketFolder     string `json:"bucketFolder,omitempty"`
	ContentType      string `json:"contentType,omitempty"`
	CeredentialsFile string `json:"ceredentialsFile,omitempty"`
}

func GetConfig() Config {
	byteValue, _ := os.ReadFile("config.json")
	var config Config
	json.Unmarshal(byteValue, &config)
	return config
}

func main() {
	fmt.Printf("[INFO] WebHost: %s\n", GetConfig().WebHost)
	fmt.Printf("[INFO] WebPort: %s\n", GetConfig().WebPort)

	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.POST("/", func(c *gin.Context) {
		// Obtine el contenido del formulario
		file, errFromHtml := c.FormFile("formFile")
		if errFromHtml != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": errFromHtml.Error(),
			})
			return
		}
		// Guarda el archivo en el servidor local
		errUploading := c.SaveUploadedFile(file, "assets/uploads/"+file.Filename)
		if errUploading != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": errUploading.Error(),
			})
			return
		}

		ctx := context.Background()
		credsFile := GetConfig().CeredentialsFile

		// Crea una nueva instancia de storage.Client
		client, err := storage.NewClient(ctx, option.WithCredentialsFile(credsFile))
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": fmt.Sprintf("No se pudo crear el cliente de Cloud Storage: %v", err),
			})
			return
		}

		bucketName := GetConfig().BucketName                   // ej. "my-bucket"
		objectName := GetConfig().BucketFolder + file.Filename // ej. "tests/" + file.Filename
		contentType := GetConfig().ContentType                 // ej. "image/*" para todos los tipos de imagen

		// Lee el contenido del archivo que deseas subir
		fileContent, err := os.ReadFile("assets/uploads/" + file.Filename)
		if err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": fmt.Sprintf("No se pudo leer el archivo local: %v", err),
			})
			return
		}

		// Sube el archivo a Cloud Storage
		wc := client.Bucket(bucketName).Object(objectName).NewWriter(ctx)
		wc.ContentType = contentType
		if _, err := wc.Write(fileContent); err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": fmt.Sprintf("No se pudo escribir el archivo en Cloud Storage: %v", err),
			})
			return
		}
		if err := wc.Close(); err != nil {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": fmt.Sprintf("No se pudo cerrar el archivo en Cloud Storage: %v", err),
			})
			return
		}

		// Si todo sale bien, muestra el link del archivo subido
		c.HTML(http.StatusOK, "index.html", gin.H{
			"success": "Archivo subido correctamente",
			"gcsUrl":  "https://storage.googleapis.com/" + bucketName + "/" + objectName,
		})
	})

	// No especificar el GetConfig().WebHost, para que escuche en todas las interfaces. Si se especifica, solo escucha en esa interfaz.
	r.Run(":" + GetConfig().WebPort)
}
