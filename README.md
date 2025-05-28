 * Запуск проекта/Project start
git clone https://github.com/patayyo/go-car-service.git
cd go-car-service

 * Создайте .env файл/Create .env file
DB_NAME=your_db_name
DB_USER=your_db_username
DB_PASSWORD=your_db_password
JWT_SECRET=your_jwt_secret_key

 * Собрать и запустить контейнеры/Create and start dockers
docker-compose up --build

 * REST API будет доступен: http://localhost:8080
 * Swagger: http://localhost:8080/swagger/index.html для инициализации swagger - swag init --dir ./cmd,./internal/handler --output ./docs --parseDependency --parseInternal

Доступрные маршруты
Auth
 * POST /signup - регистрация
 * POST /login - авторизация
Vehicles
 * POST /vehicles - создать авто (for example json body{"make": "toyota", "mark": "mark2", "year": 2000}) 
 * GET /vehicles/:id - получить по id
 * GET /vehicles - получить все авто

Если потребуется пересобрать .pb.go файлы
 * protoc --go_out=. --go-grpc_out=. proto/vehicle.proto
