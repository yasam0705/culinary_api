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
### Уровень 1 (Дедлайн 2 дня)
Разработать rest api для сервиса рецептов. Сервис должен поддерживать следующие базовые функции:
- Создание нового рецепта (название, описание, ингредиенты, шаги приготовления)
- Получение списка всех рецептов
- Получение информации о конкретном рецепте по ID
- Редактирование рецепта по ID
- Удаление рецепта по ID
*/
