package main

/*
func Start(mem *model.Cashe) {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/badsearch", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/badsearch.html")
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		data := r.FormValue("order")
		if data == "" {
			http.Redirect(w, r, "/badsearch", http.StatusSeeOther)
		}
		mem.RLock()
		if Order, ok := mem.Memory[data]; ok {
			//http.Redirect(w, r, "/data", http.StatusSeeOther)
			t, err := template.ParseFiles("static/search.html")
			if err != nil {
				http.Error(w, "Не удалось:"+err.Error(), http.StatusInternalServerError)
			}
			t.Execute(w, Order)
		} else {
			http.Redirect(w, r, "/badsearch", http.StatusSeeOther)
		}

		mem.RUnlock()
	})

}
*/
