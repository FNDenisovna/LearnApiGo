Учебный проект сервиса интеграции.

Приложение параллельного веб-сервиса, 
получает клиентские Rest/Grpc-запросы в необходимом формате JSON, 
обрабатывает по необходимой бизнес-логике, 
выполняет манипуляциии с БД, 
возвращает ответ в необходимом формате JSON.

##############################################################################################

#to generate swagger info and correct detected models
#run
PS D:\LearnApiGo> swag init -g internal/apis/endpoint.go

# GET "/albums" retrun all albums in db
curl http://localhost:8080/albums -X GET -H @{ "content-type" = "application/json"} 
curl.exe http://localhost:8080/albums -X GET -H "Content-Type: application/json"
#alternate command:
##Invoke-RestMethod -Uri http://localhost:18332/ -Credential bitcoinipvision -body $thisCanBeAPowerShellObject  

curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"


# POST "/albums" add new row to albums
curl.exe -i -X POST -H "Content-Type: application/json" -d "@body.json" http://localhost:8080/api/v1/albums

curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

#SWAGGER + пример подключения бд, авторизация
#https://github.com/MartinHeinz/go-project-blueprint/blob/rest-api/cmd/blueprint/main.go
