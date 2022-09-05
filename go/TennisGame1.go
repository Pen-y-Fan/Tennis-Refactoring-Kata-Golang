package tenniskata

type tennisGame1 struct {
	scorePlayer1 int
	scorePlayer2 int
	NamePlayer1  string
	NamePlayer2  string
}

func TennisGame1(player1Name string, player2Name string) TennisGame {
	return &tennisGame1{
		NamePlayer1: player1Name,
		NamePlayer2: player2Name}
}

func (game *tennisGame1) WonPoint(playerName string) {
	if playerName == game.NamePlayer1 {
		game.scorePlayer1++
	} else {
		game.scorePlayer2++
	}
}

func (game *tennisGame1) GetScore() string {
	if game.isScoreDraw() {
		return game.getScoreDraw()
	}

	if game.isAdvantageOrWin() {
		return game.getAdvantageOrWin()
	}

	return game.getGameScorePlayer1() + "-" + game.getGameScorePlayer2()

}

func (game *tennisGame1) isScoreDraw() bool {
	return game.scorePlayer1 == game.scorePlayer2
}

func (game *tennisGame1) getScoreDraw() string {
	if game.scorePlayer1 < 3 {
		return game.getGameScorePlayer1() + "-All"
	}
	return "Deuce"
}

func (game *tennisGame1) isAdvantageOrWin() bool {
	return game.scorePlayer1 >= 4 || game.scorePlayer2 >= 4
}

func (game *tennisGame1) getAdvantageOrWin() string {
	if getScoreDifference(game) == 1 {
		return "Advantage " + getLeadingPlayerName(game)
	}

	return "Win for " + getLeadingPlayerName(game)
}

func getScoreDifference(game *tennisGame1) int {
	return abs(game.scorePlayer1 - game.scorePlayer2)
}

func (game *tennisGame1) getGameScorePlayer1() string {
	return getGameScore(game.scorePlayer1)
}

func (game *tennisGame1) getGameScorePlayer2() string {
	return getGameScore(game.scorePlayer2)
}

func getGameScore(playerScore int) string {
	switch playerScore {
	case 0:
		return "Love"
	case 1:
		return "Fifteen"
	case 2:
		return "Thirty"
	}

	return "Forty"
}

func getLeadingPlayerName(game *tennisGame1) string {
	if game.scorePlayer1 > game.scorePlayer2 {
		return game.NamePlayer1
	}
	return game.NamePlayer2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
