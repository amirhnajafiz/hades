class Settings:
	# the game settings to display the frame with the style we choose
	def __init__(self):
		# class constructor
		self.screen_width = 1200
		self.screen_height = 800
		self.bg_color = (230, 230, 230)
		# ship settings
		self.ship_limit = 3
		# bullet settings
		self.bullet_width = 100
		self.bullet_height = 15
		self.bullet_color = (60, 60, 60)
		self.bullets_allowed = 3
		# alienss settings
		self.fleet_drop_down = 50 # vertical movement of alien
		self.speedup_scale = 1.1
		self.score_scale = 1.5
		#
		self.initialize_dynamic_settings()

	def initialize_dynamic_settings(self):
		# this method sets the game speed
		self.ship_speed = 10
		self.bullet_speed = 2
		self.alien_speed = 0.5 # horizontall movement
		self.fleet_direction = 1 # 1 is for going right and -1 is for going left
		self.alien_point = 50

	def increase_speed(self):
		# this method increases the game speed each time we pass a level
		self.ship_speed *= self.speedup_scale
		self.bullet_speed *= self.speedup_scale
		self.alien_speed *= self.speedup_scale # horizontall movement of alien
		self.alien_point = int(self.alien_point * self.score_scale)
			