# Welcome to Simone

This is a pattern memorization game, inspired by "Simon says", hence the name Simone.

<img width="416" height="768" alt="Screenshot 2025-10-28 102059" src="https://github.com/user-attachments/assets/f81d82f8-fa08-4073-87a6-d4f7a571ea46" />

# Play
You can play the game at [my homepage](https://paraschis.gr/filepage.php?key=simone).

# Changelog
v1.0
- Initial version

v1.1
- Change of peg graphics
- Introduction of error sound and life win sound
- On each round, a new sound/move is added to the existing sequence, instead of a whole new sequence.
- Removal of background
- Removal of color flash on error
- Fix of bug where multiple animations ended up messing up the peg size
- Every 4 levels a life is won

v1.2
- New sounds for peg hits
- A game over sound
- Animation tweaks

# About
Simone has been created using Defold and specifically the "mobile" project template. This means that the settings in ["game.project"](defold://open?path=/game.project) have been changed to be suitable for a mobile game:

- The screen size is set to 640x1136
- The projection is set to Fixed Fit
- Orientation is fixed vertically
- Android and iOS icons are set
- Mouse click/single touch is bound to action "touch"
- A simple script in a game object is set up to receive and react to input
- Accelerometer input is turned off (for better battery life)

# Credits
- Game over sound: Oiboo
- Peg sounds: wobbleboxx.com
- Everything else: spyros1973
