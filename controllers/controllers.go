package controllers

import 

http.HandleFunc("/",MainPageHandler)
http.HandleFunc("/entry",loginPageHandler)
http.HandleFunc("/registerForm", registerPageHandler )	
http.HandleFunc("/login",  middlewares.VerifyLogin(loginHandler) )
http.HandleFunc("/register", middlewares.VerifyUser(registerHandler) )
http.HandleFunc("/secretData", middlewares.VerifyJWT(secretHandler))

log.Fatal(http.ListenAndServe(":8000",nil))