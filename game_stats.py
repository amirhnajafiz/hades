class GameStats:
	# this class holdes the game stats like the number of ships
	def __init__(self, ai_game):
		# class constructor
		self.settings = ai_game.settings
		self.reset_stats()
		self.game_active = False
		self.high_score = 0
		self.level = 1

	def reset_stats(self):
		# this method resets the state mode
		self.ships_left = self.settings.ship_limit
		self.score = 0
		self.level = 1
