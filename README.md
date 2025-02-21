Выполненное тестовое задание.  
Использовалось fiber, reform, logrus, goose, godotenv.  
Для запуска необходимо в IDE перейти в консоли папку project и там прописать make up_build, после чего проект забилдится в Docker'e и применятся миграции для создания таблиц.  
Также в news-service/cmd/api/migrations можно из папки add_mock_info перенести файлы в папку migrations чтобы сразу были созданы экзмепляры в таблицах.  
Используется Authorization заголовок с вручную установленным токеном.  
Для проверки работоспособности использовался Postman.  
Успешный GET запрос с токеном и сам запрос:  
![изображение](https://github.com/user-attachments/assets/8705565b-4418-44b6-81ad-74f5696be15f)  
![изображение](https://github.com/user-attachments/assets/c3a3c828-c773-4c0b-82c5-645e883b023c)  
Успешный POST запрос и его результат:  
![изображение](https://github.com/user-attachments/assets/a13bb665-25f2-4587-8834-dea3fe3deff5)  
![изображение](https://github.com/user-attachments/assets/a7deceab-9c4a-4e85-8d7a-4a62517da604)  


