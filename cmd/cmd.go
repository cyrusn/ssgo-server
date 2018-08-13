package cmd

import (
	"fmt"
	"log"
	"os"

	auth "github.com/cyrusn/goJWTAuthHelper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	CONFIG_PATH            = "./config.yaml"
	TEACHER_JSON_PATH      = "./data/teacher.json"
	STUDENT_JSON_PATH      = "./data/student.json"
	SUBJECT_JSON_PATH      = "./data/subject.json"
	DEFAULT_DSN            = "root@/ssgoTestDB"
	CONTEXT_KEY_NAME       = "authClaim"
	JWT_KEY_NAME           = "jwt"
	ROLE_KEY_NAME          = "Role"
	PRIVATE_KEY            = "skill-vein-planet-neigh-envoi"
	DEFAULT_PORT           = ":5000"
	STATIC_FOLDER_LOCATION = "./public"
	DEFAULT_LIFE_TIME      = 30
)

var (
	cfgFile              string
	port                 string
	staticFolderLocation string
	teacherJSONPath      string
	studentJSONPath      string
	subjectJSONPath      string
	DSN                  string
	isOverwrite          bool
	privateKey           string
	lifeTime             int64
	secret               auth.Secret
)

func initConfig() {
	viper.SetConfigFile(cfgFile)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}

	fmt.Println("Using config file:", viper.ConfigFileUsed())
}

func initSecret() {
	secret = auth.New(
		CONTEXT_KEY_NAME, JWT_KEY_NAME, ROLE_KEY_NAME, []byte(privateKey),
	)
}

func initVariables() {
	privateKey = viper.GetString("key")
	DSN = viper.GetString("dsn")
	isOverwrite = viper.GetBool("overwrite")
	teacherJSONPath = viper.GetString("teacher")
	studentJSONPath = viper.GetString("student")
	subjectJSONPath = viper.GetString("subject")
	port = viper.GetString("port")
	staticFolderLocation = viper.GetString("static")
	lifeTime = viper.GetInt64("time")
}

// Execute excute all commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cobra.OnInitialize(initConfig, initVariables, initSecret)
}
