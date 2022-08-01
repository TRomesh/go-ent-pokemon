package main

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/viper"
	"github.com/tromesh/go-ent-pokemon/ent"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()

	ctx := context.Background()
	app := fiber.New()

	url := fmt.Sprintf("%s:%s@(%s:%s)/%s?parseTime=True", viper.Get("DB_USER"), viper.Get("DB_PASSWORD"), viper.Get("DB_HOST"), viper.Get("DB_PORT"), viper.Get("DB_NAME"))
	client, err := ent.Open(dialect.MySQL, url)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/all", func(c *fiber.Ctx) error {
		pokemons, err := client.Pokemon.Query().All(ctx)
		if err != nil {
			return fmt.Errorf("failed querying pokemons: %w", err)
		}
		log.Println("returned pokemons:", pokemons)

		return c.Status(200).JSON(pokemons)
	})

	app.Post("/create", func(c *fiber.Ctx) error {
		payload := struct {
			Name        string  `json:"name"`
			Description string  `json:"description"`
			Weight      float64 `json:"weight"`
			Height      float64 `json:"height"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		pokemon, err := client.Pokemon.
			Create().
			SetName(payload.Name).
			SetDescription(payload.Description).
			SetWeight(payload.Weight).
			SetHeight(payload.Height).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed creating pokemon: %w", err)
		}
		log.Println("pokemon created: ", pokemon)
		return c.Status(200).JSON(pokemon)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {

		payload := struct {
			Id int `json:"id"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		err := client.Pokemon.DeleteOneID(payload.Id).Exec(ctx)
		if err != nil {
			return fmt.Errorf("failed deleting pokemon: %w", err)
		}
		return c.Status(200).SendString("Success")
	})

	app.Put("/update", func(c *fiber.Ctx) error {
		payload := struct {
			Id          int     `json:"id"`
			Name        string  `json:"name"`
			Description string  `json:"description"`
			Weight      float64 `json:"weight"`
			Height      float64 `json:"height"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		pokemon, err := client.Pokemon.
			UpdateOneID(payload.Id).
			SetName(payload.Name).
			SetDescription(payload.Description).
			SetWeight(payload.Weight).
			SetHeight(payload.Height).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed updating pokemon: %w", err)
		}
		return c.Status(200).JSON(pokemon)
	})

	app.Post("/create/battle", func(c *fiber.Ctx) error {
		payload := struct {
			Result    string `json:"result"`
			Contender int    `json:"contender"`
			Oponent   int    `json:"oponent"`
		}{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		battle, err := client.Battle.
			Create().
			SetResult(payload.Result).
			SetContenderID(payload.Contender).
			SetOponentID(payload.Oponent).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("failed creating battle: %w", err)
		}
		log.Println("battle created: ", battle)
		return c.Status(200).JSON(battle)
	})

	app.Get("/all/battle", func(c *fiber.Ctx) error {
		battles, err := client.Battle.Query().WithContender().WithOponent().All(ctx)
		if err != nil {
			return fmt.Errorf("failed querying battles: %w", err)
		}
		log.Println("returned battles:", battles)

		return c.Status(200).JSON(battles)
	})

	log.Fatal(app.Listen(fmt.Sprintf(":%s", viper.Get("APP_PORT"))))
}
