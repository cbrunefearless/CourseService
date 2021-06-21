package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
    "CourseService/app/v1/models"
	"github.com/gorilla/mux"
)

var Courses = []models.Course{
	models.Course{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	models.Course{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - Health Check"))
}

func ReturnAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllCourses")
	json.NewEncoder(w).Encode(Courses)
}

func ReturnSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnSingleCourse")
	vars := mux.Vars(r)
	key := vars["id"]

	fmt.Fprintf(w, "Key: "+key)
}

func CreateNewCourse(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	fmt.Println("Endpoint Hit: createCourse")
	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	fmt.Println("Endpoint Hit: updateCourse")
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	fmt.Println("Endpoint Hit: deleteCourse")
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]

	// we then need to loop through all our articles
	for index, course := range Courses {
		// if our id path parameter matches one of our
		// articles
		if course.Id == id {
			// updates our Articles array to remove the
			// article
			Courses = append(Courses[:index], Courses[index+1:]...)
		}
	}

}
