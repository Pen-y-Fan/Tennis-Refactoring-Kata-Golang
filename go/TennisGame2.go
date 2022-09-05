package tenniskata

type tennisGame2 struct {
	Player1Points int
	Player2Points int

	player1Result string
	player2Result string
	player1Name   string
	player2Name   string
}

func TennisGame2(player1Name string, player2Name string) TennisGame {
	return &tennisGame2{player1Name: player1Name, player2Name: player2Name}
}

func (game *tennisGame2) GetScore() string {

	game.calculateResults()

	if game.isScoreDraw() && game.isScoringGame() {
		return game.player1Result + "-All"
	}

	if game.isScoreDraw() {
		return "Deuce"
	}

	if game.isWin() {
		return "Win for " + game.getLeadingPlayerName()
	}

	if game.isAdvantageOrWin() {
		return "Advantage " + game.getLeadingPlayerName()
	}

	return game.getResults()
}

func (game *tennisGame2) updatePlayer1Points() {
	game.Player1Points++
}

func (game *tennisGame2) updatePlayer2Points() {
	game.Player2Points++
}

func (game *tennisGame2) WonPoint(player string) {
	if player == game.player1Name {
		game.updatePlayer1Points()
		return
	}
	game.updatePlayer2Points()
}

func (game *tennisGame2) isScoreDraw() bool {
	return game.Player1Points == game.Player2Points
}

func (game *tennisGame2) isScoringGame() bool {
	return game.Player1Points < 3
}

func (game *tennisGame2) calculatePlayer1Result() {
	game.player1Result = getResultFor(game.Player1Points)
}

func (game *tennisGame2) calculatePlayer2Result() {
	game.player2Result = getResultFor(game.Player2Points)
}

func (game *tennisGame2) getResults() string {

	return game.player1Result + "-" + game.player2Result
}
func (game *tennisGame2) calculateResults() {
	game.calculatePlayer1Result()
	game.calculatePlayer2Result()
}

func (game *tennisGame2) getLeadingPlayerName() string {

	if game.Player1Points > game.Player2Points {
		return game.player1Name
	}
	return game.player2Name
}

func (game *tennisGame2) isAdvantageOrWin() bool {
	return (game.Player1Points > game.Player2Points && game.Player2Points >= 3) || (game.Player2Points > game.Player1Points && game.Player1Points >= 3)
}

func (game *tennisGame2) isWin() bool {
	return (game.Player1Points >= 4 && game.Player2Points >= 0 && (game.Player1Points-game.Player2Points) >= 2) || (game.Player2Points >= 4 && game.Player1Points >= 0 && (game.Player2Points-game.Player1Points) >= 2)
}

func getResultFor(playerPoints int) string {

	if playerPoints == 0 {
		return "Love"
	}
	if playerPoints == 1 {
		return "Fifteen"
	}
	if playerPoints == 2 {
		return "Thirty"
	}
	return "Forty"
}
