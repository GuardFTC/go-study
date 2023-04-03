// Package test  @Author:冯铁城 [17615007230@163.com] 2023-04-03 10:16:19
package test

// 定义游戏
type game struct {
	gooses []goose
	ducks  []duck
}

// 定义加载游戏方法
func (g *game) loadPlayer(gooses []goose, ducks []duck) {
	g.gooses = gooses
	g.ducks = ducks
}

// 定义开始游戏方法
func (g *game) startGame() {

	//1.鹅释放技能
	for _, _goose := range g.gooses {
		_goose.releaseSkills()
	}

	//2.鸭释放技能
	for _, _duck := range g.ducks {
		_duck.kill()
	}
}

// Gaming 进行游戏
func Gaming() {

	//1.创建鹅
	_bigGoose := &bigGoose{"白板大鹅"}
	_sheriff := &sheriff{&bigGoose{"警长"}}
	_messengerJustice := &messengerJustice{&bigGoose{"正义使者"}}

	//2.创建鸭
	_bigDuck := &bigDuck{"鸭子"}
	_assassin := &assassin{&bigDuck{"刺客"}}
	_superPowers := &superPowers{&bigDuck{"超能力者"}}

	//3.创建游戏变量
	_game := new(game)

	//4.加载玩家
	_game.loadPlayer([]goose{_bigGoose, _sheriff, _messengerJustice}, []duck{_bigDuck, _assassin, _superPowers})

	//5.开始游戏
	_game.startGame()
}
