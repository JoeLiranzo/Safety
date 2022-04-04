package controllers

import "deepthinking.do/safetygo/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Apps routes
	s.Router.HandleFunc("/apps", middlewares.SetMiddlewareJSON(s.CreateApp)).Methods("POST")
	s.Router.HandleFunc("/apps", middlewares.SetMiddlewareJSON(s.GetApps)).Methods("GET")
	s.Router.HandleFunc("/apps/{id}", middlewares.SetMiddlewareJSON(s.GetApp)).Methods("GET")
	s.Router.HandleFunc("/apps/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateApp))).Methods("PUT")
	s.Router.HandleFunc("/apps/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteApp)).Methods("DELETE")

	//Roles routes
	s.Router.HandleFunc("/roles", middlewares.SetMiddlewareJSON(s.CreateRole)).Methods("POST")
	s.Router.HandleFunc("/roles", middlewares.SetMiddlewareJSON(s.GetRoles)).Methods("GET")
	s.Router.HandleFunc("/roles/{id}", middlewares.SetMiddlewareJSON(s.GetRole)).Methods("GET")
	s.Router.HandleFunc("/roles/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateRole))).Methods("PUT")
	s.Router.HandleFunc("/roles/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteRole)).Methods("DELETE")

	//AppsRoles routes
	s.Router.HandleFunc("/appsroles", middlewares.SetMiddlewareJSON(s.CreateAppRole)).Methods("POST")
	s.Router.HandleFunc("/appsroles", middlewares.SetMiddlewareJSON(s.GetAppsRoles)).Methods("GET")
	s.Router.HandleFunc("/appsroles/{id}", middlewares.SetMiddlewareJSON(s.GetAppRole)).Methods("GET")
	s.Router.HandleFunc("/appsroles/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateAppRole))).Methods("PUT")
	s.Router.HandleFunc("/appsroles/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteAppRole)).Methods("DELETE")

	//UsersAppsRoles routes
	s.Router.HandleFunc("/usersappsroles", middlewares.SetMiddlewareJSON(s.CreateUserAppRole)).Methods("POST")
	s.Router.HandleFunc("/usersappsroles", middlewares.SetMiddlewareJSON(s.GetUsersAppsRoles)).Methods("GET")
	s.Router.HandleFunc("/usersappsroles/{id}", middlewares.SetMiddlewareJSON(s.GetUserAppRole)).Methods("GET")
	s.Router.HandleFunc("/usersappsroles/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUserAppRole))).Methods("PUT")
	s.Router.HandleFunc("/usersappsroles/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUserAppRole)).Methods("DELETE")
}
