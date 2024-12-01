package controllers

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"streaming/services"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type VideoController struct {
	VideoService services.VideoService
}

func (vc *VideoController) StreamVideo(ctx *fiber.Ctx) error {
	// Obtener el ID del video desde los parámetros de la URL
	videoID := ctx.Params("id")
	if videoID == "" {
		return ctx.Status(fiber.StatusBadRequest).SendString("Video ID is required")
	}

	// Obtener el video desde el servicio
	video, err := vc.VideoService.GetVideoByID(videoID)
	if err != nil {
		log.Println("Error fetching video:", err)
		return ctx.Status(fiber.StatusNotFound).SendString("Video not found")
	}

	filePath := video.VideoPath

	// Abrir el archivo del video
	file, err := os.Open(filePath)
	if err != nil {
		log.Println("Error opening video file:", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}
	defer file.Close()

	// Obtener información del archivo
	fileInfo, err := file.Stat()
	if err != nil {
		log.Println("Error getting file information:", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
	}

	// Obtener el tipo MIME del archivo
	mimeType := mime.TypeByExtension(filepath.Ext(filePath))
	fileSize := fileInfo.Size()

	// Obtener el encabezado Range de la solicitud
	rangeHeader := ctx.Get("Range")
	if rangeHeader != "" {
		var start, end int64
		ranges := strings.Split(rangeHeader, "=")
		if len(ranges) != 2 {
			log.Println("Invalid Range Header")
			return ctx.Status(http.StatusInternalServerError).SendString("Internal Server Error")
		}

		byteRange := ranges[1]
		byteRanges := strings.Split(byteRange, "-")

		// Obtener el inicio del rango
		start, err = strconv.ParseInt(byteRanges[0], 10, 64)
		if err != nil {
			log.Println("Error parsing start byte position:", err)
			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		// Calcular el fin del rango
		if len(byteRanges) > 1 && byteRanges[1] != "" {
			end, err = strconv.ParseInt(byteRanges[1], 10, 64)
			if err != nil {
				log.Println("Error parsing end byte position:", err)
				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
			}
		} else {
			end = fileSize - 1
		}

		// Configurar los encabezados de respuesta
		ctx.Set(fiber.HeaderContentRange, fmt.Sprintf("bytes %d-%d/%d", start, end, fileSize))
		ctx.Set(fiber.HeaderContentLength, strconv.FormatInt(end-start+1, 10))
		ctx.Set(fiber.HeaderContentType, mimeType)
		ctx.Set(fiber.HeaderAcceptRanges, "bytes")
		ctx.Status(fiber.StatusPartialContent)

		// Ajustar la posición de lectura
		_, seekErr := file.Seek(start, io.SeekStart)
		if seekErr != nil {
			log.Println("Error seeking to start position:", seekErr)
			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}

		// Copiar los bytes seleccionados al cliente
		_, copyErr := io.CopyN(ctx.Response().BodyWriter(), file, end-start+1)
		if copyErr != nil {
			log.Println("Error copying bytes to response:", copyErr)
			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
	} else {
		// Si no hay encabezado Range, enviar todo el archivo
		ctx.Set("Content-Length", strconv.FormatInt(fileSize, 10))
		ctx.Set("Content-Type", mimeType)
		_, copyErr := io.Copy(ctx.Response().BodyWriter(), file)
		if copyErr != nil {
			log.Println("Error copying entire file to response:", copyErr)
			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		}
	}

	return nil
}
