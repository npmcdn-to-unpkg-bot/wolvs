package commonWeb

import "net/http"

func WriteMessageOrError(writer http.ResponseWriter, message string, err error) {
	if err != nil {
		http.Error(writer, message, 400)
	} else {
		writer.Write([]byte(message))
	}
}

func WriteAcceptedMessageOrError(writer http.ResponseWriter, message string, err error) {
	if err != nil {
		http.Error(writer, message, 400)
	} else {
		writer.WriteHeader(200)
		writer.Write([]byte(message))
	}
}
