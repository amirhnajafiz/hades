import pygame
from pygame.sprite import Sprite 

class Ship(Sprite):
	# this class is for the player ship in the game
	def __init__(self, ai_game):
		# class constructor to open the image and set the frames
		super().__init__()
		self.screen = ai_game.screen
		self.settings = ai_game.settings
		self.screen_rect = ai_game.screen.get_rect()
		#
		self.image = pygame.image.load('images/ship.png') # how to open an image
		self.image = pygame.transform.scale(self.image, (15, 35)) # this is how we resize in pygame
		#
		self.rect = self.image.get_rect() # this is for converting the image to rectangle
		#
		self.rect.midbottom = self.screen_rect.midbottom # this is for placing the image at the center bottom
		#
		self.x = float(self.rect.x) # since rect.x stores int values we made a x var in ship so we can handel the float numbers input
		#
		self.moving_right = False # movement status flags
		self.moving_left = False

	def blitme(self):
		# this method will draw the image into the main frame of the game
		self.screen.blit(self.image, self.rect)	# drawing image

	def update(self):
		# this method will make the changes based on the flags situations
		if self.moving_right and self.rect.right < self.screen_rect.right: # the second part will check for the right sides of two rects 
			self.x += self.settings.ship_speed	
		if self.moving_left and self.rect.left > self.screen_rect.left: # the second part will check for left sides of two rects
			self.x -= self.settings.ship_speed
		self.rect.x = self.x # we update our ship vars to the rect vars		

	def center_ship(self):
		# this method will replace the use ship at the center
		self.rect.midbottom = self.screen_rect.midbottom
		self.x = float(self.rect.x)	# notice that when the rect midbottom changes the rect x also changes