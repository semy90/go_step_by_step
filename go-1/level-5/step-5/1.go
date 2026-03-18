package main

/*Программист Арсений заинтересовался футболом, но просто играть ему мало. Он решил проанализировать результаты своей любимой команды и помочь улучшить их.

Программа для анализа должна:

Хранить данные о каждом футболисте: Имя, Голы, Промахи и Помощь.
Рассчитывать рейтинг для каждого игрока по формуле: (Голы + Помощь / 2) / Промахи (если количество промахов равно нулю, то Голы + Помощь / 2).
Сортировка по:
Убыванию количества голов (функция goalsSort(players []Player) []Player)
Убыванию рейтинга (функция ratingSort(players []Player) []Player)
Убыванию отношения голов к промахам (функция gmSort(players []Player) []Player) Если рейтинг одинаковый, то во всех функциях необходимо сортировать по имени (Name) в алфавитном порядке.
Также нужно реализовать такую структуру:
type Player struct {
    Name      string
    Goals     int
    Misses    int
    Assists   int
    Rating    float64
}
и вспомогательный метод calculateRating() для расчёта поля Rating на основе входных данных. Конструктор NewPlayer(name string, goals, misses, assists int) Player создаёт нового игрока и вычисляет его рейтинг с помощью calculateRating().*/
import (
	"sort"
)

type Player struct {
	Name    string
	Goals   int
	Misses  int
	Assists int
	Rating  float64
}

func (player *Player) calculateRating() float64 {
	player.Rating = float64(player.Goals) + float64(player.Assists)/2
	if player.Misses != 0 {
		player.Rating /= float64(player.Misses)
	}
	return player.Rating
}

func NewPlayer(name string, goals, misses, assists int) Player {
	player := Player{Name: name, Goals: goals, Misses: misses, Assists: assists}
	player.calculateRating()
	return player
}

func goalsSort(players []Player) []Player {

	sort.Slice(players, func(i, j int) bool {
		if players[i].Goals > players[j].Goals {
			return true
		}
		if players[i].Goals < players[j].Goals {
			return false
		}
		return players[i].Name < players[j].Name
	})
	return players
}

func ratingSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		if players[i].Rating == players[j].Rating {
			return players[i].Name < players[j].Name
		}
		return players[i].Rating > players[j].Rating
	})
	return players
}

func gmSort(players []Player) []Player {
	sort.Slice(players, func(i, j int) bool {
		ratioI := float64(players[i].Goals)
		if players[i].Misses != 0 {
			ratioI /= float64(players[i].Misses)
		}

		ratioJ := float64(players[j].Goals)
		if players[j].Misses != 0 {
			ratioJ /= float64(players[j].Misses)
		}

		if ratioI == ratioJ {
			return players[i].Name < players[j].Name
		}
		return ratioI > ratioJ
	})
	return players
}
