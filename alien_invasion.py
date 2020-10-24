import sys # the system library for controlling the screen 
import pygame # this library is for creating 2D games
from time import sleep

from settings import Settings # using the 'Setting' class we build
from ship import Ship
from bullet import Bullet 
from alien import Alien 
from game_stats import GameStats 
from button import Button 
from scoreboard import ScoreBoard 
from util import save, load

class AlienInvasion:
	# the game class 
	def __init__(self):
		# class constructor
		pygame.init()

		self.settings = Settings() # setting instance

		# manually set the sizes of the frame
		self.screen = pygame.display.set_mode(
			(self.settings.screen_width, self.settings.screen_height)) # only takes one argument

		self.bg_image = pygame.image.load('images/background.png') # this is for setting the background image
		self.bg_image = pygame.transform.scale(self.bg_image, (self.settings.screen_width, self.settings.screen_height))
		self.bg_rect = self.bg_image.get_rect()
		self.bg_rect.midbottom = self.screen.get_rect().midbottom

		#self.screen = pygame.display.set_mode((0, 0), pygame.FULLSCREEN) # use fullscreen in pygame
		#self.settings.screen_width = self.screen.get_rect().width 
		#self.settings.screen_height = self.screen.get_rect().height 

		pygame.display.set_caption("Alien Invasion")

		self.stats = GameStats(self)
		self.stats.high_score = load()
		self.sb = ScoreBoard(self)

		self.ship = Ship(self)
		self.bullets = pygame.sprite.Group() # creating a group sprite which is usefull in updating
		self.aliens = pygame.sprite.Group() # creating a group sprite for aliens same as bullets

		self._create_fleet()

		self.play_button = Button(self, "Play") # the game button

	def _create_fleet(self):
		# this method places the aliens in rows and colls
		alien = Alien(self)

		alien_width, alien_height = alien.rect.size # getting the bounds
		available_space_x = self.settings.screen_width - (2 * alien_width) # calculations to find the maximum number of aliens in a line
		number_alien_x = available_space_x // (3 * alien_width) # the '//' is floor division

		ship_height = self.ship.rect.height
		available_space_y = (self.settings.screen_height - 10 * ship_height - (4 * alien_height)) # calculations to find the maximum number of lines in screen
		number_alien_y = available_space_y // (2 * alien_height)

		for row_number in range(number_alien_y):
			for alien_number in range(number_alien_x): # creating a row of aliens
				self._create_alien(alien_number, row_number)

	def _check_fleet_edges(self):
		# this method checks the overlapping of alien and walls
		for alien in self.aliens.sprites():
			if alien.check_edges():
				self._change_fleet_direction()
				break

	def _change_fleet_direction(self):
		# this method make the changes when aliens hit the walls
		for alien in self.aliens.sprites():
			alien.rect.y += self.settings.fleet_drop_down
		self.settings.fleet_direction *= -1		

	def _create_alien(self, alien_number, row_number):
		# this method creates a single alien 
		alien = Alien(self)
		alien_width, alien_height = alien.rect.size
		alien.x = alien_width + 3 * alien_width * alien_number
		alien.rect.x = alien.x 
		alien.rect.y = 3 * alien.rect.height + 2 * alien.rect.height * row_number
		self.aliens.add(alien)			

	def run_game(self):
		# a method for starting the game process
		# the game main loop
		while True:
			self._check_events()

			if self.stats.game_active: # check if the game is still on
				self.ship.update()
				self._update_bullets()
				self._update_aliens()	

			self._update_screen()

	def _check_events(self):
		# a method for getting the keyboard and mouse events
		for event in pygame.event.get():
			if event.type == pygame.QUIT:
				save(self.stats.high_score)
				sys.exit()	
			elif event.type == pygame.KEYDOWN:
				self._check_keydown_events(event)	
			elif event.type == pygame.KEYUP:
				self._check_keyup_events(event)
			elif event.type == pygame.MOUSEBUTTONDOWN:
				mouse_pos = pygame.mouse.get_pos()
				self._check_play_button(mouse_pos)	
				
	def _check_keydown_events(self, event):
		# this method handels the pressing keys
		if event.key == pygame.K_RIGHT:
			self.ship.moving_right = True
		if event.key == pygame.K_LEFT:
			self.ship.moving_left = True
		if event.key == pygame.K_q:
			save(self.stats.high_score)
			sys.exit()	
		if event.key == pygame.K_SPACE:
			self._fire_bullet()
		if event.key == pygame.K_p:
			self._do_play_click()
				
	def _check_keyup_events(self, event):	
		# this method handels the released keys								
		if event.key == pygame.K_RIGHT:
			self.ship.moving_right = False
		if event.key == pygame.K_LEFT:
			self.ship.moving_left = False

	def _fire_bullet(self):
		# this method creates and adds new bullet to the bullets list
		if len(self.bullets) < self.settings.bullets_allowed: # a limit for bullets added
			new_bullet = Bullet(self)
			self.bullets.add(new_bullet)

	def _check_play_button(self, mouse_pos):
		# this method checks the mouse position and the button position
		if self.play_button.rect.collidepoint(mouse_pos) and not self.stats.game_active:
			self._do_play_click()

	def _do_play_click(self):
		# this method gets the player in the game
		self.stats.reset_stats()
		self.settings.initialize_dynamic_settings()
		self.aliens.empty()
		self.bullets.empty()
		self._create_fleet()
		self.sb.prep_score()
		self.sb.prep_level()
		self.sb.prep_ships()
		self.ship.center_ship()
		self.play_button.button_color = (0, 100, 0) # mouse hovering
		self.play_button._preg_msg('Play')
		self._update_screen()
		sleep(0.5) # a short break
		self.stats.game_active = True	
		pygame.mouse.set_visible(False) # not showing the mouse		

	def _update_bullets(self):
		# this method will update the bullets
		self.bullets.update()
		for bullet in self.bullets.copy(): # we use the copy method to create an other list for iteration
			if bullet.rect.bottom < 1:
				self.bullets.remove(bullet)

		self._check_alien_bullet_collision()		

	def _check_alien_bullet_collision(self):
		'''
		The method bellow will get two groups and will check for any overlapping between
		the elements in each group. If so it will remove them from their lists.
		'''		
		collisions = pygame.sprite.groupcollide(self.bullets, self.aliens, True, True)

		if collisions:
			for aliens in collisions.values():
				self.stats.score += self.settings.alien_point * len(aliens)
			self.sb.prep_score()	
			self.sb.check_high_score()	

		if not self.aliens: # this is for recreating aliens when all are dead
			self.bullets.empty()	
			self._create_fleet()
			self.settings.increase_speed()

			self.stats.level += 1
			self.sb.prep_level()

	def _check_aliens_bottom(self):
		# this method checks if the aliens have reached the bottom of the screen
		screen_rect = self.screen.get_rect()
		for alien in self.aliens.sprites().copy():
			if alien.rect.bottom >= screen_rect.bottom:
				self._ship_hit()
				break

	def _ship_hit(self):
		# this method is for restarting a game and make the changes
		if self.stats.ships_left > 0:
			self.stats.ships_left -= 1
			self.sb.prep_ships()

			self.aliens.empty()
			self.bullets.empty()

			self._create_fleet()
			self.ship.center_ship()		

			sleep(1) # a puse for the user	
		else:
			self.stats.game_active = False
			pygame.mouse.set_visible(True)	
			self.play_button.button_color = (0, 255, 0)
			self.play_button._preg_msg('Play')

	def _update_aliens(self):
		# this method updates the aliens movements
		self._check_fleet_edges()	
		self.aliens.update()

		if pygame.sprite.spritecollideany(self.ship, self.aliens):
			self._ship_hit()

		self._check_aliens_bottom()			

	def _update_screen(self):
		# this method creates and shows the game screen to the user	
		self.screen.blit(self.bg_image, self.bg_rect) # drawing the background image
		self.ship.blitme()
		for bullet in self.bullets.sprites(): # group object only has update method which we overrode it
			bullet.draw_bullet()
		self.aliens.draw(self.screen) # drawing the aliens on the screen
		self.sb.show_score()
		if not self.stats.game_active:
			self.play_button.draw_button()

		pygame.display.flip() # making the frame visible	
		
			
if __name__ == '__main__': # creating an instance and start
	ai = AlienInvasion()
	ai.run_game()				