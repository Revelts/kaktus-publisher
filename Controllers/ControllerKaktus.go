package Controllers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/rabbitmq/amqp091-go"
	"kaktus-services/Connections"
	"kaktus-services/Controllers/Dto/request"
	"kaktus-services/Controllers/Dto/response"
	"kaktus-services/Library"
	"log"
	"net/http"
)

func ViewAllThreads(w http.ResponseWriter, r *http.Request) {
	RabbitChannel, err := Connections.RabbitConnection.Channel()
	if err != nil {
		log.Panicf("[ERROR] : %s", err)
		return
	}

	q, err := RabbitChannel.QueueDeclare("",
		false,
		true,
		true,
		false,
		nil)

	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	correlationId := uuid.New()
	err = RabbitChannel.Publish("kaktus", "view_all_thread",
		false,
		false,
		amqp091.Publishing{
			ContentType:   "text/plaint",
			CorrelationId: correlationId.String(),
			ReplyTo:       q.Name,
		})

	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := RabbitChannel.Consume(q.Name, "",
		true,
		false,
		false,
		false,
		nil,
	)

	var resp []response.ViewAllThread
	for d := range data {
		if d.CorrelationId == correlationId.String() {
			err = json.Unmarshal(d.Body, &resp)
			if err != nil {
				Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
				return
			}
			break
		}
	}
	err = RabbitChannel.Close()
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	Library.HttpResponseSuccess(w, r, resp)
}

func ViewThreadDetails(w http.ResponseWriter, r *http.Request) {
	RabbitChannel, err := Connections.RabbitConnection.Channel()
	if err != nil {
		log.Panicf("[ERROR] : %s", err)
		return
	}

	q, err := RabbitChannel.QueueDeclare("",
		false,
		true,
		true,
		false,
		nil)

	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	tId := chi.URLParam(r, "id")
	correlationId := uuid.New()
	dataByte, _ := json.Marshal(tId)

	err = RabbitChannel.Publish("kaktus", "view_thread_details",
		false,
		false,
		amqp091.Publishing{
			ContentType:   "text/plaint",
			CorrelationId: correlationId.String(),
			ReplyTo:       q.Name,
			Body:          dataByte,
		})

	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := RabbitChannel.Consume(q.Name, "",
		true,
		false,
		false,
		false,
		nil,
	)

	var resp response.ViewThreadDetail
	for d := range data {
		if d.CorrelationId == correlationId.String() {
			err = json.Unmarshal(d.Body, &resp)
			if err != nil {
				Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
				return
			}
			break
		}
	}
	err = RabbitChannel.Close()
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	Library.HttpResponseSuccess(w, r, resp)
}

func CommentThread(w http.ResponseWriter, r *http.Request) {
	var requestBody request.CommentThread
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	RabbitChannel, err := Connections.RabbitConnection.Channel()
	if err != nil {
		log.Panicf("[ERROR] : %s", err)
		return
	}

	q, err := RabbitChannel.QueueDeclare("",
		false,
		true,
		true,
		false,
		nil)

	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	correlationId := uuid.New()
	dataByte, _ := json.Marshal(requestBody)
	err = RabbitChannel.Publish("kaktus", "comment_thread",
		false,
		false,
		amqp091.Publishing{
			ContentType:   "text/plaint",
			CorrelationId: correlationId.String(),
			ReplyTo:       q.Name,
			Body:          dataByte,
		})

	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	Library.HttpResponseSuccess(w, r, http.StatusOK)
}

func ReplyComment(w http.ResponseWriter, r *http.Request) {
	var requestBody request.ReplyComment
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	RabbitChannel, err := Connections.RabbitConnection.Channel()
	if err != nil {
		log.Panicf("[ERROR] : %s", err)
		return
	}

	q, err := RabbitChannel.QueueDeclare("",
		false,
		true,
		true,
		false,
		nil)

	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	correlationId := uuid.New()
	dataByte, _ := json.Marshal(requestBody)
	err = RabbitChannel.Publish("kaktus", "reply_comment",
		false,
		false,
		amqp091.Publishing{
			ContentType:   "text/plaint",
			CorrelationId: correlationId.String(),
			ReplyTo:       q.Name,
			Body:          dataByte,
		})

	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	Library.HttpResponseSuccess(w, r, http.StatusOK)
}

func CreateThread(w http.ResponseWriter, r *http.Request) {
	var requestBody request.CreateThread
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	RabbitChannel, err := Connections.RabbitConnection.Channel()
	if err != nil {
		log.Panicf("[ERROR] : %s", err)
		return
	}

	q, err := RabbitChannel.QueueDeclare("",
		false,
		true,
		true,
		false,
		nil)

	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	correlationId := uuid.New()
	dataByte, _ := json.Marshal(requestBody)
	err = RabbitChannel.Publish("kaktus", "create_thread",
		false,
		false,
		amqp091.Publishing{
			ContentType:   "text/plaint",
			CorrelationId: correlationId.String(),
			ReplyTo:       q.Name,
			Body:          dataByte,
		})

	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	Library.HttpResponseSuccess(w, r, http.StatusOK)
}

func LikeThread(w http.ResponseWriter, r *http.Request) {
	var requestBody request.LikeThread
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusBadRequest)
		return
	}

	RabbitChannel, err := Connections.RabbitConnection.Channel()
	if err != nil {
		log.Panicf("[ERROR] : %s", err)
		return
	}

	q, err := RabbitChannel.QueueDeclare("",
		false,
		true,
		true,
		false,
		nil)

	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}
	correlationId := uuid.New()
	dataByte, _ := json.Marshal(requestBody)
	err = RabbitChannel.Publish("kaktus", "like_thread",
		false,
		false,
		amqp091.Publishing{
			ContentType:   "text/plaint",
			CorrelationId: correlationId.String(),
			ReplyTo:       q.Name,
			Body:          dataByte,
		})

	if err != nil {
		Library.HttpResponseError(w, r, err.Error(), http.StatusInternalServerError)
		return
	}

	Library.HttpResponseSuccess(w, r, http.StatusOK)
}
