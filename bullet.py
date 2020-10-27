import pygame
from pygame.sprite import Sprite

class Bullet(Sprite):
	# this is a single bullet class which holds the bullet data
	def __init__(self, ai_game):
		# class constructor
		super().__init__() # Spirit class constructor
		self.screen = ai_game.screen
		self.settings = ai_game.settings
		self.color = (200, 200, 200)

		self.rect = pygame.Rect(0, 0, self.settings.bullet_width,
			self.settings.bullet_height) # building a rectangle for bullet

		self.rect.midtop = ai_game.ship.rect.midtop # setting it on the top of the ship

		self.y = float(self.rect.y) # creating the bullet var

	def update(self):
		# this method will update the bullet 
		self.y -= self.settings.bullet_speed
		self.rect.y = self.y

	def draw_bullet(self):
		# a method for drawing the bullet on the screen
		pygame.draw.rect(self.screen, self.color, self.rect) # drawing a shape		