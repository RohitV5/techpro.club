package projects

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"techpro.club/sources/common"
	"techpro.club/sources/pages"
	"techpro.club/sources/users"
)

type FinalProjectPreviewOutStruct struct{
	ProjectPreview common.FetchProjectStruct `json:"projectsList"`
	UserNameImage common.UsernameImageStruct `json:"userNameImage"`
}

func ProjectPreview(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/projects/view" {
        pages.ErrorHandler(w, r, http.StatusNotFound)
        return
    }

	// Session check
	sessionOk, userID := users.ValidateDbSession(w, r)
	if(!sessionOk){
		
		// Delete cookies
		users.DeleteSessionCookie(w, r)
		users.DeleteUserCookie(w, r)

		http.Redirect(w, r, "/projects", http.StatusSeeOther)
	}

	var finalOutStruct FinalProjectPreviewOutStruct
	var userNameImage common.UsernameImageStruct

	// Fetch user name and image from saved browser cookies
	status, msg, userName, image := pages.FetchUsernameImage(w, r)

	if(!status){
		log.Println(msg)
	} else {
		userNameImage  = common.UsernameImageStruct{userName,image}
	}

	projectID := r.URL.Query().Get("projectid")
	_, _, result := pages.FetchProjectDetails(projectID, userID)

	finalOutStruct = FinalProjectPreviewOutStruct{result, userNameImage}
	

	tmpl, err := template.New("").ParseFiles("templates/app/projects/projectpreview.gohtml", "templates/app/projects/common/base_new.gohtml")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		tmpl.ExecuteTemplate(w, "projectbase", finalOutStruct) 
	}
}