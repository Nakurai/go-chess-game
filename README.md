# A Simpe Chess Engine

This is a project to practice Go. It is a very naive implementaiton of a chess logic. I think it is also a very good playground.

![Chess Game Screenshot](/docs/chess-game.png)

# How to Run

`go mod tidy`

`go run .`

# More Details

This project uses the [Ebitengine](https://ebitengine.org/) library. It was a lot of fun.

The "AI" at the moment is only picking moves at random. You can take control at anytime during the game by hitting the `Esc` key and selecting "Player vs Player". This will allow you to play as both players. Go back and select "Player vs AI" if you want to revert.

# Roadmap

When I have the time, here are the things I would like to improve:

- [ ] Organize the code better by using packages.

- [ ] Allow the user to pick their side when playing vs the AI.

- [ ] Make the AI incrementally smarter. I have a lot of ideas about this. For a starter, I would like to build a heuristic function able to assess the intrinsic "value" of a board in any specific state. That would help determine the next move. I know stockfish has one, so it should be feasible.

- [ ] Make the UI part of the game better. For now, it does not even acknowledge a checkmate/draw state!

- [ ] Make the art look better. I would like to learn how to use 2D sprites with Ebitengine and make the game as a whole better.

- [ ] Optimize the chess logic. For now it is all very naïve, and given the simple "AI" it does not show its limitations. However, there are a lot of possible optimizations. As the logic gets more complicated, the idea would be to keep an eye for emerging bottlenecks and go from there.