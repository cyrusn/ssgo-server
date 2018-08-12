package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/cyrusn/ssgo-server/model/auth"
	"github.com/cyrusn/ssgo-server/model/student"
	"github.com/cyrusn/ssgo-server/model/subject"
	"github.com/cyrusn/ssgo-server/route"

	helper "github.com/cyrusn/goHTTPHelper"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	_ "github.com/mattn/go-sqlite3"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start Subject Selection System Backend Server",
	Run: func(cmd *cobra.Command, args []string) {
		auth.UpdateLifeTime(lifeTime)
		checkPathExist(dbPath, staticFolderLocation)
		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			log.Fatal(err)
		}

		env := route.Env{
			&auth.DB{db, &secret},
			&student.DB{db},
			&subject.DB{db},
		}

		Serve(&env)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.PersistentFlags().StringVarP(
		&port,
		"port",
		"p",
		DEFAULT_PORT,
		"port value",
	)
	serveCmd.PersistentFlags().StringVarP(
		&staticFolderLocation,
		"static",
		"s",
		STATIC_FOLDER_LOCATION,
		"location of static folder for serving",
	)
	serveCmd.PersistentFlags().Int64VarP(
		&lifeTime,
		"time",
		"t",
		DEFAULT_LIFE_TIME,
		"update the life time (minutes) of jwt",
	)

	viper.BindPFlags(serveCmd.PersistentFlags())
}

// Serve serve the routers
func Serve(env *route.Env) {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.

	routes := env.Routes()
	for _, ro := range routes {
		handler := http.HandlerFunc(ro.Handler)

		// pass Access to handler first
		if len(ro.Scopes) != 0 {
			handler = secret.Access(ro.Scopes, handler).(http.HandlerFunc)
		}

		// then pass Authenticate at last
		if ro.Auth {
			handler = secret.Authenticate(handler).(http.HandlerFunc)
		}

		r.
			PathPrefix("/api/").
			Methods(ro.Methods...).
			Path(ro.Path).
			HandlerFunc(handler)
	}

	srv := &http.Server{
		Handler: helper.Logger(r),
		Addr:    "localhost" + port,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Bind to a port and pass our router in
	fmt.Println("Available on http://localhost" + port)
	log.Fatal(srv.ListenAndServe())
}
