package cmd

import (
	"log"

	"github.com/cyrusn/goJWTAuthHelper"
	"github.com/spf13/cobra"
)

const (
	TEACHER_JSON_PATH = "./data/teacher.json"
	STUDENT_JSON_PATH = "./data/student.json"
	SUBJECT_JSON_PATH = "./data/subject.json"
	DB_PATH           = "./test/test.db"
	CONTEXT_KEY_NAME  = "authClaim"
	JWT_KEY_NAME      = "jwt"
	ROLE_KEY_NAME     = "Role"
	PRIVATE_KEY       = "skill-vein-planet-neigh-envoi"
)

var (
	secret          auth.Secret
	teacherJSONPath string
	studentJSONPath string
	subjectJSONPath string
	dbPath          string
	isOverwrite     bool
	privateKey      string

	rootCmd = &cobra.Command{
		Use:   "ssgo",
		Short: "Welcome to Subject Selection System Backend Server",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cobra.OnInitialize(func() {
		secret = auth.New(
			CONTEXT_KEY_NAME, JWT_KEY_NAME, ROLE_KEY_NAME, []byte(privateKey),
		)
	})

	cmds := []*cobra.Command{
		versionCmd,
		createCmd,
		importCmd,
	}

	for _, cmd := range cmds {
		rootCmd.AddCommand(cmd)
	}

	importSubCmds := []*cobra.Command{
		teacherCmd,
		studentCmd,
		subjectCmd,
	}

	for _, cmd := range importSubCmds {
		importCmd.AddCommand(cmd)
	}

	rootCmd.PersistentFlags().StringVarP(
		&privateKey,
		"key",
		"k",
		PRIVATE_KEY,
		"change the private key for authentication on jwt",
	)

	rootCmd.PersistentFlags().StringVarP(
		&dbPath,
		"location",
		"l",
		DB_PATH,
		"location of sqlite3 database file",
	)

	createCmd.PersistentFlags().BoolVarP(
		&isOverwrite,
		"overwrite",
		"o",
		false,
		"overwrite database if database location exist",
	)

	teacherCmd.PersistentFlags().StringVarP(
		&teacherJSONPath,
		"teacher",
		"t",
		TEACHER_JSON_PATH,
		"path of teacher.json file\nplease check README.md for the schema",
	)

	studentCmd.PersistentFlags().StringVarP(
		&studentJSONPath,
		"student",
		"u",
		STUDENT_JSON_PATH,
		"path of student.json file\nplease check README.md for the schema",
	)

	subjectCmd.PersistentFlags().StringVarP(
		&subjectJSONPath,
		"subject",
		"s",
		SUBJECT_JSON_PATH,
		"path of subject.json file\nplease check README.md for the schema",
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
