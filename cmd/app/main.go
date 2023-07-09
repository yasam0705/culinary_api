package main

import (
	"github/culinary_api/config"
	"github/culinary_api/internal/app"
	"os"
	"os/signal"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	app, err := app.NewApp(cfg)
	if err != nil {
		panic(err)
	}

	if err = app.Run(); err != nil {
		panic(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	if err = app.Stop(); err != nil {
		panic(err)
	}
}

/*
### Уровень 2 (Дедлайн 3 дня)
Все что в уровне 1 + дополнительные функции:
- Фильтр по списку ингредиентов.

- Добавить в шаги приготовления время этапа. Добавить фильтр и сортировку списка по общему времени приготовления.
	1) добавить в cooking steps колонку "cooking_time"
	2) добавить в recipe колонку "cooking_time" это будет общее время приготовления
	3) добавить фильтр в list метод (cooking_time)

- Реализовать аутентификацию и авторизацию пользователей. Только авторизованные пользователи могут создавать, редактировать и удалять рецепты.
	1) создать таблицу users (guid, login, password, created_at, updated_at)
	2) repo, usecase. регистрация, авторизация
	3) добавить middleware на проверку авторизации, если не авторизован то ограничивать действия
*/
