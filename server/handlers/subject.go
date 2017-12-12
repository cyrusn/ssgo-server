package handlers

// import "net/http"

// // GetSubjectHandler ...
// func (env *Env) GetSubjectHandler(w http.ResponseWriter, r *http.Request) {
// 	subjectCode := r.FormValue("code")
// 	subject, err := env.SubjectStore.Get(subjectCode)
// 	errCode := http.StatusBadRequest
// 	if err != nil {
// 		errPrint(w, err, errCode)
// 		return
// 	}

// 	jsonPrint(w, subject, errCode)
// 	return
// }
