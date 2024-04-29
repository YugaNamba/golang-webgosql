package controllers

import (
	"net/http"
	"todo_app/app/models"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", int(Found))
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", int(Found))
	} else {
		user, err := sess.GetUserBySession()
		if err != nil {
			http.Redirect(w, r, "/login", int(Found))
		}
		todos, err := user.GetTodosByUser()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		user.Todos = todos
		generateHTML(w, user, "layout", "private_navbar", "index")
	}
}

func todoNew(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", int(Found))
	} else {
		generateHTML(w, sess, "layout", "private_navbar", "todo_new")
	}
}

func todoSave(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		sess, err := session(w, r)
		if err != nil {
			http.Redirect(w, r, "/login", int(Found))
		} else {
			user, err := sess.GetUserBySession()
			if err != nil {
				http.Redirect(w, r, "/login", int(Found))
			}
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			content := r.PostFormValue("content")
			if err := user.CreateTodo(content); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/todos", int(Found))
		}
	}
}

func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", int(Found))
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			http.Redirect(w, r, "/login", int(Found))
		}
		todo, err := models.GetTodo(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		generateHTML(w, todo, "layout", "private_navbar", "todo_edit")
	}	
}

func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", int(Found))
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			http.Redirect(w, r, "/login", int(Found))
		}
		if r.Method == "POST" {
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			content := r.PostFormValue("content")
			todo, err := models.GetTodo(id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			todo.Content = content
			if err := todo.UpdateTodo(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/todos", int(Found))
		}
	}
}

func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", int(Found))
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			http.Redirect(w, r, "/login", int(Found))
		}
		todo, err := models.GetTodo(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if err := todo.DeleteTodo(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/todos", int(Found))
	}
}
