package tenniskata

type tennisGame3 struct {
	player1Score int
	player2Score int
	player1Name  string
	player2Name  string
}

func TennisGame3(player1Name string, player2Name string) TennisGame {
	return &tennisGame3{player1Name: player1Name, player2Name: player2Name}
}

func (game *tennisGame3) GetScore() string {
	if game.isScoring() {
		if game.isDraw() {
			return game.getPlayer1Result() + "-All"
		}
		return game.getPlayer1Result() + "-" + game.getPlayer2Result()
	}

	if game.isDraw() {
		return "Deuce"
	}

	if game.isAdvantage() {
		return "Advantage " + game.getLeadingPlayerName()
	}

	return "Win for " + game.getLeadingPlayerName()
}

func (game *tennisGame3) WonPoint(playerName string) {
	if playerName == game.player1Name {
		game.player1Score++
		return
	}
	game.player2Score++
}

func (game *tennisGame3) getLeadingPlayerName() string {
	if game.player1Score > game.player2Score {
		return game.player1Name
	}
	return game.player2Name
}

func (game *tennisGame3) getPlayer1Result() string {
	return results()[game.player1Score]
}

func (game *tennisGame3) getPlayer2Result() string {
	return results()[game.player2Score]
}

func (game *tennisGame3) isScoring() bool {
	return game.player1Score < 4 && game.player2Score < 4 && !(game.player1Score+game.player2Score == 6)
}

func (game *tennisGame3) isDraw() bool {
	return game.player1Score == game.player2Score
}

func (game *tennisGame3) isAdvantage() bool {
	return (game.player1Score-game.player2Score)*(game.player1Score-game.player2Score) == 1
}

func results() []string {
	return []string{"Love", "Fifteen", "Thirty", "Forty"}
}
