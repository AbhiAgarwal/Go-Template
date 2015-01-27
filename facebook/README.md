Running webserver

`go run main.go`

Notes about Facebook

- `/login` to login to Facebook
- `/logout` to logout of Facebook
- `/oauth2callback` should be the callback
	- must be `http://localhost:3000/oauth2callback`, etc.
	- that will redirect to `/`