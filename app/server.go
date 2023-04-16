package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/herizal95/hisabia_api/config"
	"github.com/herizal95/hisabia_api/helper"
	"github.com/herizal95/hisabia_api/routes"
	dataroutes "github.com/herizal95/hisabia_api/routes/dataRoutes"
	transaksiroutes "github.com/herizal95/hisabia_api/routes/transaksiRoutes"
	wilayahroutes "github.com/herizal95/hisabia_api/routes/wilayahRoutes"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

type AppConfig struct {
	AppName string
	AppEnv  string
	AppPort string
}

func (server *Server) Initialize(appConfig AppConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)

}

func (server *Server) Run(addr string) {

	router := mux.NewRouter()
	route := router.PathPrefix(helper.ApiV1).Subrouter()

	config.ConnectDatabase()
	routes.AuthenticationRoutes(route)
	routes.UserRoutes(route)
	wilayahroutes.DesaRoutes(route)
	dataroutes.DataRoutes(route)
	transaksiroutes.TransaksiRoutes(route)

	t := time.Now()

	formattedTime := t.Format("2006-01-02 15:04:05")
	fmt.Println("Server Starting on : ", formattedTime)

	fmt.Printf("Server running on port %s", addr)

	log.Fatal(http.ListenAndServe(addr, route))
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func Run() {

	var server = Server{}
	var appconfig = AppConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error on Loading .env file")
	}

	appconfig.AppName = GetEnv("APP_NAME", "appname")
	appconfig.AppEnv = GetEnv("APP_ENV", "appenv")
	appconfig.AppPort = GetEnv("APP_PORT", "9999")

	server.Initialize(appconfig)
	server.Run(":" + appconfig.AppPort)
}
