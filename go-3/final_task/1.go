package main

/*Реализуйте HTTP-сервис на Go, который поддерживает работу с пользователями.

Требования
Структуры данных:

Определите структуру User:
type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Age  int    `json:"age"`
}
Определите структуру Store для хранения пользователей в памяти.
Хранилище пользователей:

В структуре Store реализуйте методы:
CreateUser(name string, age int) User — создаёт пользователя с уникальным id (автоинкремент, начиная с 1).
GetUser(id int) (User, bool) — возвращает пользователя по id и флаг, найден ли пользователь.
Реализуйте функцию-конструктор NewStore() *Store, которая возвращает новое хранилище.
Важно: Хранилище должно быть потокобезопасным (использовать sync.Mutex или sync.RWMutex).
HTTP-обработчики:

POST /users — создать пользователя. В теле запроса приходит JSON с полями name (string) и age (int). В ответе — созданный пользователь с уникальным id. Код ответа: 200 OK. При ошибке декодирования JSON — 400 Bad Request.
GET /users/{id} — получить пользователя по id. В ответе — JSON с пользователем и код 200 OK, или код 404 Not Found, если пользователь не найден. При некорректном id — 400 Bad Request.
Middleware для логирования:

Реализуйте функцию middleware с сигнатурой:
func loggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler
Middleware должен логировать каждый HTTP-запрос после его обработки.
Для перехвата кода статуса создайте wrapper над http.ResponseWriter:
type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}
Инициализируйте statusCode значением 200 (статус по умолчанию).
Формат логирования с помощью slog:
logger.Info("http request",
    slog.String("method", r.Method),
    slog.String("path", r.URL.Path),
    slog.Int("status", statusCode),
    slog.Duration("duration", elapsed))
Сервер:

В функции main создайте logger: slog.New(slog.NewTextHandler(os.Stdout, nil))
Создайте хранилище: store := NewStore()
Настройте маршруты с использованием http.ServeMux
Оберните mux в middleware: handler := loggingMiddleware(logger, mux)
Запустите сервер на порту :8080: http.ListenAndServe(":8080", handler)
Примеры запросов и ответов
Создание пользователя (POST /users)

Запрос:

{"name": "Ivan", "age": 25}
Ответ (200 OK):

{"id": 1, "name": "Ivan", "age": 25}
Лог:

time=2024-01-15T10:30:45.123Z level=INFO msg="http request" method=POST path=/users status=200 duration=1.234ms
Получение пользователя (GET /users/1)

Ответ (200 OK):

{"id": 1, "name": "Ivan", "age": 25}
Лог:

time=2024-01-15T10:30:46.456Z level=INFO msg="http request" method=GET path=/users/1 status=200 duration=0.123ms
Пользователь не найден (GET /users/999)

Ответ (404 Not Found):

not found
Лог:

time=2024-01-15T10:30:47.789Z level=INFO msg="http request" method=GET path=/users/999 status=404 duration=0.089ms
Некорректный JSON (POST /users)

Запрос:

{invalid json}
Ответ (400 Bad Request):

bad request
Лог:

time=2024-01-15T10:30:48.012Z level=INFO msg="http request" method=POST path=/users status=400 duration=0.045ms*/
import (
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	store *Store
)

type User struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Store struct {
	k  int
	mx sync.Mutex
	mp map[int]User
}

func (s *Store) CreateUser(name string, age int) User {
	s.mx.Lock()
	s.k++
	u := User{ID: s.k, Name: name, Age: age}
	s.mp[s.k] = u
	s.mx.Unlock()
	return u
}
func (s *Store) GetUser(id int) (User, bool) {
	s.mx.Lock()
	u, ok := s.mp[id]
	s.mx.Unlock()
	if !ok {
		return User{}, false
	}
	return u, true
}
func NewStore() *Store {
	return &Store{k: 0, mx: sync.Mutex{}, mp: map[int]User{}}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	u, ok := store.GetUser(id)
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}
	userJson, err := json.Marshal(u)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(userJson))
	return
}
func PostHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	user = store.CreateUser(user.Name, user.Age)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, user.ID)
	return
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(rw, r)
		end := time.Since(start)
		logger.Info("http request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("status", rw.statusCode),
			slog.Duration("duration", end))
	})
}

func main() {
	store = NewStore()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mux := http.NewServeMux()
	mux.HandleFunc("GET /users/{id}", GetHandler)
	mux.HandleFunc("POST /users", PostHandler)

	handler := loggingMiddleware(logger, mux)
	http.ListenAndServe(":8080", handler)
}
