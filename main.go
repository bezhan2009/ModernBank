package main

import (
	"OnlineBanking/db"
	"OnlineBanking/models"
	"OnlineBanking/pkg/core"
	"database/sql"
	"fmt"
	"github.com/fatih/color"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func main() {
	database, err := sql.Open("sqlite3", "OnlineBankingDB")
	if err != nil {
		log.Fatal("Can't open DB. Error is:", err)
	}
	defer database.Close()

	db.DBinit(database)
	fmt.Println(core.GetAllUsersInDB(database))

	Start(database)
}

func Start(database *sql.DB) {
	WelcomeWindow(database)
}

func WelcomeWindow(database *sql.DB) {
	cyan := color.New(color.FgCyan).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	WelcomeText := cyan("╭────────────────────────────────────────╮\n") +
		cyan("│") + green("                Вход                    ") + cyan("│\n") +
		cyan("├────────────────────────────────────────┤\n") +
		cyan("│") + green("          1. Авторизация                ") + cyan("│\n") +
		cyan("│") + green("          2. Список банкоматов          ") + cyan("│\n") +
		cyan("│") + green("          0. Выход                      ") + cyan("│\n") +
		cyan("╰────────────────────────────────────────╯")

	fmt.Println("Добро пожаловать!")
	for {
		fmt.Println(WelcomeText)
		fmt.Println(color.CyanString("         |" + "Введите команду" + "|"))
		fmt.Println(color.CyanString("         ─────────────────"))
		fmt.Print("                ")
		var cmd int
		fmt.Scan(&cmd)
		switch cmd {
		case 1:
			core.Authorization(database)
		case 2:
			models.PrintingListOfATMs(database)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Вы ввели некорректные данные. Попробуйте ещё раз!")
		}
	}
}
