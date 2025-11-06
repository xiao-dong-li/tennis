# 🎮 Tetris in Go (Ebitengine)

A classic **Tetris clone** written in **Go**, powered by [Ebitengine](https://ebitengine.org/).  
This project demonstrates how to build a 2D puzzle game from scratch — including rendering, game loop, input, and logic management.

---

## 🚀 Features (Planned and Implemented)

* Standard Tetris playfield (10×10 grid)
* All seven tetromino shapes (I, J, L, O, S, T, Z)
* Piece rotation and horizontal movement
* Soft drop and hard drop
* Line clearing and scoring
* Next-piece preview
* Hold piece functionality (optional if implemented)
* Pause / resume and restart
* Simple sound effects and basic UI overlays (score, level, lines)

---

## 🕹️ Controls

* Left Arrow: Move piece left
* Right Arrow: Move piece right
* Down Arrow: Soft drop
* `Space` bar: Hard drop
* `C` (or `Shift`): Hold piece
* `X`: Rotate clockwise
* `Z`: Rotate counterclockwise
* `P` or `Esc`: Pause / Resume
* `R`: Restart game

---

## 📦 Dependencies

* Go 1.24+
* [Ebitengine](https://ebitengine.org)

---

## ⚙️ Build & Run

1. Clone the repository:

```bash
git clone <your-repo-url>
cd <repo-folder>
```

2. Fetch dependencies and run:

```bash
go mod tidy
go run .
```

---

## 📄 License

This project is released under the MIT License. See `LICENSE` for details.

---

## 🙏 Acknowledgements

* Ebitengine / Ebiten for the game framework
* Classic Tetris community and design principles

Enjoy building and playing!
