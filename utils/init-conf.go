package utils

import (
	"log"
	"streaming/utils/confs"
)

// InitConf inicializa la configuración y devuelve un error si falla
func InitConf() error {
	errEnv := confs.LoadEnv()
	errDataBase := confs.LoadConfDatabase()
	if errEnv != nil {
		log.Println("No se pudo cargar el archivo .env:", errEnv)
		return errEnv
	}

	if errDataBase != nil {
				log.Fatal("Failed to migrate models:", errDataBase)
		return errDataBase
	}

	log.Println("Configuración inicializada correctamente")
	return nil
}