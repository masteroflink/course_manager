package controllers

import "main/api/middlewares"

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login, false)).Methods("POST")

	s.Router.HandleFunc("/user", middlewares.SetMiddlewareJSON(s.CreateUser, true)).Methods("POST")
	s.Router.HandleFunc("/user", middlewares.SetMiddlewareJSON(s.GetUsers, true)).Methods("GET")
	s.Router.HandleFunc("/user/{id}", middlewares.SetMiddlewareJSON(s.GetUser, true)).Methods("GET")
	s.Router.HandleFunc("/user/{id}", middlewares.SetMiddlewareJSON(s.UpdateUser, true)).Methods("PUT")
	s.Router.HandleFunc("/user/{id}", middlewares.SetMiddlewareJSON(s.DeleteUser, true)).Methods("DELETE")

	s.Router.HandleFunc("/student", middlewares.SetMiddlewareJSON(s.CreateStudent, true)).Methods("POST")
	s.Router.HandleFunc("/student", middlewares.SetMiddlewareJSON(s.GetStudents, true)).Methods("GET")
	s.Router.HandleFunc("/student/{id}", middlewares.SetMiddlewareJSON(s.GetStudent, true)).Methods("GET")
	s.Router.HandleFunc("/student/{id}", middlewares.SetMiddlewareJSON(s.UpdateStudent, true)).Methods("PUT")
	s.Router.HandleFunc("/student/{id}", middlewares.SetMiddlewareJSON(s.DeleteStudent, true)).Methods("DELETE")

	s.Router.HandleFunc("/professor", middlewares.SetMiddlewareJSON(s.CreateProfessor, true)).Methods("POST")
	s.Router.HandleFunc("/professor", middlewares.SetMiddlewareJSON(s.GetProfessors, true)).Methods("GET")
	s.Router.HandleFunc("/professor/{id}", middlewares.SetMiddlewareJSON(s.GetProfessor, true)).Methods("GET")
	s.Router.HandleFunc("/professor/{id}", middlewares.SetMiddlewareJSON(s.UpdateProfessor, true)).Methods("PUT")
	s.Router.HandleFunc("/professor/{id}", middlewares.SetMiddlewareJSON(s.DeleteProfessor, true)).Methods("DELETE")

	s.Router.HandleFunc("/course", middlewares.SetMiddlewareJSON(s.CreateCourse, true)).Methods("POST")
	s.Router.HandleFunc("/course", middlewares.SetMiddlewareJSON(s.GetCourses, true)).Methods("GET")
	s.Router.HandleFunc("/course/{id}", middlewares.SetMiddlewareJSON(s.GetCourse, true)).Methods("GET")
	s.Router.HandleFunc("/course/{id}", middlewares.SetMiddlewareJSON(s.UpdateCourse, true)).Methods("PUT")
	s.Router.HandleFunc("/course/{id}", middlewares.SetMiddlewareJSON(s.DeleteCourse, true)).Methods("DELETE")
	s.Router.HandleFunc("/course/{cid}/student/{sid}/enroll", middlewares.SetMiddlewareJSON(s.EnrollStudent, true)).Methods("PUT")
	s.Router.HandleFunc("/course/{cid}/student/{sid}/remove", middlewares.SetMiddlewareJSON(s.RemoveStudent, true)).Methods("PUT")
	s.Router.HandleFunc("/course/{cid}/professor/{pid}/assign", middlewares.SetMiddlewareJSON(s.AssignProfessor, true)).Methods("PUT")
	s.Router.HandleFunc("/course/{cid}/professor/{pid}/remove", middlewares.SetMiddlewareJSON(s.RemoveProfessor, true)).Methods("PUT")
}
