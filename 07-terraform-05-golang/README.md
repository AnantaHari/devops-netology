# Домашнее задание к занятию "7.5. Основы golang"

С `golang` в рамках курса, мы будем работать не много, поэтому можно использовать любой IDE.
Но рекомендуем ознакомиться с [GoLand](https://www.jetbrains.com/ru-ru/go/).  

## Задача 1. Установите golang.
1. Воспользуйтесь инструкций с официального сайта: [https://golang.org/](https://golang.org/).
2. Так же для тестирования кода можно использовать песочницу: [https://play.golang.org/](https://play.golang.org/).
```
anantahari@Ubuntu-IdeaPad:~$ go version
go version go1.17.6 linux/amd64
```

## Задача 2. Знакомство с gotour.
У Golang есть обучающая интерактивная консоль [https://tour.golang.org/](https://tour.golang.org/).
Рекомендуется изучить максимальное количество примеров. В консоли уже написан необходимый код,
осталось только с ним ознакомиться и поэкспериментировать как написано в инструкции в левой части экрана.  
```
Ознакомился
```

## Задача 3. Написание кода.
Цель этого задания закрепить знания о базовом синтаксисе языка. Можно использовать редактор кода
на своем компьютере, либо использовать песочницу: [https://play.golang.org/](https://play.golang.org/).

1. Напишите программу для перевода метров в футы (1 фут = 0.3048 метр). Можно запросить исходные данные
у пользователя, а можно статически задать в коде.
    Для взаимодействия с пользователем можно использовать функцию `Scanf`:
    ```
    package main

    import "fmt"

    func main() {
        fmt.Print("Enter a number: ")
        var input float64
        fmt.Scanf("%f", &input)

        output := input * 2

        fmt.Println(output)    
    }
    ```

    Ответ
    ```
    package main

    import "fmt"

    func main() {
    	fmt.Print("Введите длину в футах: ")
    	var input float64
    	fmt.Scanf("%f", &input)

    	output := input * 0.3048

    	fmt.Printf("%v футов равно %v метров.\n", input, output)
    }
    ```

1. Напишите программу, которая найдет наименьший элемент в любом заданном списке, например:
    ```
    x := []int{48,96,86,68,57,82,63,70,37,34,83,27,19,97,9,17,}
    ```

    Ответ
    ```
    package main

    import "fmt"

    func main() {
    	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}

    	fmt.Printf("Самое маленькое число - %v\n", little_element(x))

    }

    func little_element(x []int) int {
    	s := x[0]
    	for _, xx := range x {
    		// fmt.Println(i, xx)
    		if s > xx {
    			s = xx
    		}
    	}
    	return s
    }

    ```
1. Напишите программу, которая выводит числа от 1 до 100, которые делятся на 3. То есть `(3, 6, 9, …)`.

В виде решения ссылку на код или сам код.
```
package main

import (
	"fmt"
)

func main() {
	fmt.Println(div_three())
}

func div_three() []int {
	var numbers []int
	for i := 1; i < 101; i++ {
		if i%3 == 0 {
			numbers = append(numbers, i)
		}
	}
	return numbers
}

```

## Задача 4. Протестировать код (не обязательно).

Создайте тесты для функций из предыдущего задания.
```
package main

import "testing"

func TestConvert_fm(t *testing.T) {

	res := convert_fm(3)

	if res != "3 футов равно 0.9145 метров." {
		t.Error("Expected 0.9145, got ", res)
	}
}

func TestLittle_Element(t *testing.T) {
	x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
	res := little_element(x)

	if res != 9 {
		t.Error("Expected 9, got ", res)
	}
}

func TestDiv_three(t *testing.T) {
	numbers := []int{3, 6, 9, 12, 15, 18, 21, 24, 27, 30, 33, 36, 39, 42, 45, 48, 51, 54, 57, 60, 63, 66, 69, 72, 75, 78, 81, 84, 87, 90, 93, 96, 99}
	var res []int
	res = div_three()

	if res[5] != numbers[5] {
		t.Error("Expected other, got ", res)
	}
}
```
