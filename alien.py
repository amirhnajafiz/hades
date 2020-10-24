import pygame
from pygame.sprite import Sprite 

class Alien(Sprite):
	# this is the alien class to store the aliens information on the game
	def __init__(self, ai_game):
		# class constructor
		super().__init__()
		self.screen = ai_game.screen
		self.settings = ai_game.settings

		self.image = pygame.image.load('images/alien.png') # alien image

		self.rect = self.image.get_rect()

		self.rect.x = self.rect.width
		self.rect.y = self.rect.height

		self.x = float(self.rect.x)

	def update(self):
		# this method updates the place of the alien
		self.x += (self.settings.alien_speed * self.settings.fleet_direction)
		self.rect.x = self.x

	def check_edges(self):
		# this method checks if this alien overlaps any side walls
		screen_rect = self.screen.get_rect()
		if self.rect.right >= screen_rect.right or self.rect.left <= screen_rect.left:
			return True		