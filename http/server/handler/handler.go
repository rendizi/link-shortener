package handler

import (
	"encoding/json"
	"fmt"
	"module_name/pkg/shortener"
	"net/http"
	"strings"
)

func IsValid(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/"):]

	segments := strings.Split(path, "/")

	if len(segments) > 0 && segments[0] != "" {
		short, err := shortener.Get(segments[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Redirect(w, r, short, http.StatusFound)
		return
	}
	html := `
		<!DOCTYPE html>
<html>
<head>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>URL Shortener</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            display: flex;
            justify-content: center;
            align-items: center;
            height: 100vh;
            margin: 0;
            background-color: #f5f5f5;
        }

        .url-box {
            background-color: #fff;
            border-radius: 5px;
            box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
            padding: 20px;
            width: 100%;
            max-width: 600px;
        }

        .url-box input {
            width: 100%;
            padding: 10px;
            margin-bottom: 20px;
            border: 1px solid #ddd;
            border-radius: 5px;
            outline: none;
            font-size: 16px;
            width: 100%;
        }

        .url-box button {
            background-color: #4CAF50;
            color: white;
            padding: 10px 20px;
            border: none;
            border-radius: 5px;
            cursor: pointer;
            font-size: 16px;
            width: 100%;
        }

        .url-box button:hover {
            background-color: #45a049;
        }
    </style>
</head>
<body>
    <div class="url-box">
        <input type="url" id="url" name="url" placeholder="Enter URL">
        <button id="shorten-btn">Shorten</button>
    </div>
    <script>
    document.getElementById('shorten-btn').addEventListener('click', function() {
        var inputUrl = document.getElementById('url').value; 
        var data = { url: inputUrl }; 

        fetch('/insert', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(data) 
        })
        .then(response => {
            if (response.ok) {
                return response.text(); 
            } else {
                throw new Error('Failed to insert URL');
            }
        })
        .then(data => {
            alert(data);
        })
        .catch(error => {
            console.error('Network error:', error);
        });
    });
</script>
</body>
</html>
`
	fmt.Fprint(w, html)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Request body is empty", http.StatusBadRequest)
		return
	}

	var p struct {
		Url string `json:"url"`
	}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	short, err := shortener.Insert(p.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "The link has been successfully shortened! To access: 127.0.0.1:3333/%s", short)
}
