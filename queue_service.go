package matching

type Player struct {
  token string
}

type PlayerQueue struct {
  players []Player
  playersNeeded int
  front int
}

func (pq *PlayerQueue) addPlayer(p Player) {
  // check if full
  if len(pq.players) >= pq.playersNeeded {
    pq.initGame()
  }

  // check if just filled
  pq.players[pq.front] = p
  pq.front++
  if len(pq.players) >= pq.playersNeeded {
    pq.initGame()
  }
}

func (pq *PlayerQueue) initGame() []Player {
  res := make([]Player, len(pq.players))
  copy(res, pq.players)
  pq.front = 0
  return res
}

func hello() string {
  s := "Hello, World"
  return s
}
