package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// генерирует количество очков атаки в зависимости от выбранного типа персонажа и возвращает строковое сообщение о проведённой атаке.
func attack(charName, charClass string) string {
	if charClass == "warrior" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(3, 5))
	}

	if charClass == "mage" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(5, 10))
	}

	if charClass == "healer" {
		return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, 5+randint(-3, -1))
	}
	return "неизвестный класс персонажа"
}

// генерирует количество очков защиты в зависимости от выбранного типа персонажа и возвращает строковое сообщение о выполненном блокировании атаки.
// обратите внимание на "if else" и на "else"
func defence(char_name, char_class string) string {
	if char_class == "warrior" {
		return fmt.Sprintf("%s блокировал %d урона.", char_name, 10+randint(5, 10))
	}
	if char_class == "mage" {
		return fmt.Sprintf("%s блокировал %d урона.", char_name, 10+randint(-2, 2))
	}
	if char_class == "healer" {
		return fmt.Sprintf("%s блокировал %d урона.", char_name, 10+randint(2, 5))
	}
	return "неизвестный класс персонажа"
}

// в зависимости от выбранного типа персонажа возвращает сообщение о применении специального умения.
// обратите внимание на "if else" и на "else"
func special(charName, charClass string) string {
	if charClass == "warrior" {
		return fmt.Sprintf("%s применил специальное умение `Выносливость %d`", charName, 80+25)
	}
	if charClass == "mage" {
		return fmt.Sprintf("%s применил специальное умение `Атака %d`", charName, 5+40)
	}
	if charClass == "healer" {
		return fmt.Sprintf("%s применил специальное умение `Защита %d`", charName, 10+30)
	}
	return "неизвестный класс персонажа"
}

// запускает цикл тренировки навыков персонажа. В качестве параметров она получает введённое игроком имя персонажа и выбранный тип персонажа.
// здесь обратите внимание на имена параметров
func startTraining(charName, charClass string) string {
	if charClass == "warrior" {
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", charName)
	}

	if charClass == "mage" {
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", charName)
	}

	if charClass == "healer" {
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", charName)
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд: attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)

		if cmd == "attack" {
			fmt.Println(attack(charName, charClass))
		}

		if cmd == "defence" {
			fmt.Println(defence(charName, charClass))
		}

		if cmd == "special" {
			fmt.Println(special(charName, charClass))
		}
	}

	return "тренировка окончена"
}

// позволяет игроку выбрать тип игрового персонажа и возвращает выбранный вариант.
// обратите внимание на имя функции и имена переменных
func choiseCharClass() string {
	var approve_choice string
	var charClass string

	for approve_choice != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &charClass)
		if charClass == "warrior" {
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		} else if charClass == "mage" {
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		} else if charClass == "healer" {
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		}
		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &approve_choice)
		approve_choice = strings.ToLower(approve_choice)
	}
	return charClass
}

// обратите внимание на имена переменных
func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var charName string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &charName)

	fmt.Printf("Здравствуй, %s\n", charName)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	char_class := choiseCharClass()

	fmt.Println(startTraining(charName, char_class))
}

// вспомогательная функция, которая возвращает случайное целое число в заданном диапазоне.
func randint(min, max int) int {
	return rand.Intn(max-min) + min
}
