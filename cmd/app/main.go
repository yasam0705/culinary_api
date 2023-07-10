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
### Уровень 3 (Дедлайн 4 дня)
Все что в уровне 2 + дополнительные функции:
- Реализовать возможность добавления оценки. Пользователи могут оценивать рецепты по шкале от 1 до 5. Фильтрация и сортировка списка по рейтингу.
- Добавить возможность загружать фото для рецепта и шагов приготовления, реализовать загрузку и отдачу изображаний.
- Покрытие кода тестами
*/
