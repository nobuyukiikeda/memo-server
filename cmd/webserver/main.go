package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"google.golang.org/api/iterator"
	"google.golang.org/api/option"

	firebase "firebase.google.com/go"
)

const saveFile = "memo.text"

func main() {
	ctx := context.Background()
	sa := option.WithCredentialsFile("application_credentials.json")
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		fmt.Printf("error initializing app: %v", err)
		return
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		fmt.Printf("error initializing app: %v", err)
		return
	}
	iter := client.Collection("books").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if  err != nil {
			fmt.Printf("error get documents: %v", err)
		}
		fmt.Println(doc.Data()["shoppingMemo"])
	}

	defer client.Close()

	// print("server listening --- http://localhost:8888/\n")
	// fs := http.FileServer(http.Dir("web"))
	// http.Handle("/", fs)
	// http.ListenAndServe(":8888", nil)
}

func writeHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if (len(r.Form["text"]) == 0) {
		w.Write([]byte("フォームから投稿してください。"))
		return
	}
	text := r.Form["text"][0]
	ioutil.WriteFile(saveFile, []byte(text), 0644)
	fmt.Println("save: " + text)
	http.Redirect(w, r, "/", 301)
}



