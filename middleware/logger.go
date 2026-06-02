package middleware

import (
	"fmt"
	"net/http"
	"time"
)

func LoggerMiddleware(next http.Handler) http.Handler {

	// var array [2]string = [2]string{"sajib", "saha"}

	// fmt.Println(array)

	// var array2 []int

	// array2 = append(array2, 1)

	// fmt.Println(array2)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		startTime := time.Now()

		// Execute the actual handler
		next.ServeHTTP(w, r)

		diff := time.Since(startTime)

		fmt.Printf(
			"Path=%s Method=%s Time=%s\n",
			r.URL.Path,
			r.Method,
			diff,
		)
	})
}
