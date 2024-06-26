package controllers

import (
	"log"
	"net/http"
	"todo_app/app/models"
)

func signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
		} else {
			http.Redirect(w, r, "/todos", int(Found))
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Cannot read body", http.StatusInternalServerError)
		}
		user := models.User{
			Name: r.PostFormValue("name"),
			Email: r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
			http.Error(w, "Cannot create user", http.StatusInternalServerError)
		}
		http.Redirect(w, r, "/", int(Found))
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "login")
		} else {
			http.Redirect(w, r, "/todos", int(Found))
		}
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Cannot read body", http.StatusInternalServerError)
		}
	}
}

func authenticate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		user, err := models.GetUserByEmail(r.PostFormValue("email"))
		if err != nil {
			log.Println(err)
			http.Redirect(w, r, "/login", int(Unauthorized))
		}
		if user.Password == models.Encrypt(r.PostFormValue("password")) {
			session, err := user.CreateSession()
			if err != nil {
				log.Println(err)
				http.Error(w, "Cannot create session", http.StatusInternalServerError)
			}
			cookie := http.Cookie{
				Name: "_cookie",
				Value: session.UUID,
				HttpOnly: true,
			}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/", int(Found))
		} else {
			http.Redirect(w, r, "/login", int(Unauthorized))
		}
	}
}

func logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}

	if err != http.ErrNoCookie {
		session := models.Session{UUID: cookie.Value}
		session.DeleteSessionByUUID()
		http.Redirect(w, r, "/login", int(Found))
	}
}

