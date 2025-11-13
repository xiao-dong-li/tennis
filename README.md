# 🎮 Tetris in Go (Ebitengine)

A classic **Tetris clone** written in **Go**, powered by [Ebitengine](https://ebitengine.org/).  
This project showcases how to build a 2D puzzle game from scratch — covering rendering, game loops, input handling, and
gameplay logic.

---

## 🚀 Features (Planned and Implemented)

### ✅ Implemented

* Standard **10×20** playfield
* All seven tetromino shapes (**I, J, L, O, S, T, Z**)
* Piece rotation and horizontal movement
* Soft drop & hard drop
* Line clearing and scoring
* Next-piece preview
* Pause / Resume / Restart
* Scene transitions (fade-in / fade-out)
* Simple UI overlays (Score, Level, Lines)

### 🧩 Planned

* Hold piece functionality
* Sound effects & background music

---

## 🕹️ Controls

| Action                  | Key          |
|-------------------------|--------------|
| Move Left               | ←            |
| Move Right              | →            |
| Soft Drop               | ↓            |
| Hard Drop               | `Space`      |
| Rotate Clockwise        | `X`          |
| Rotate Counterclockwise | `Z`          |
| Hold Piece              | `C`          |
| Pause / Resume          | `P` or `Esc` |
| Restart                 | `R`          |

---

## ⚙️ Build & Run

**Clone the repository**

```bash
git clone https://github.com/xiao-dong-li/tennis.git
cd tennis
```

**📦 Install dependencies**

```bash
make deps
```

**🛠️ Build (Linux + Windows)**

```bash
make all
```

**💻 Or run directly**

```bash
make run
```

**📖 Show Help**

```bash
make help
```

### 🪟 Windows (No Make Installed)

**▶️ Run the Game**

```bash
go mod tidy && go run .
```

---

## 📦 Dependencies

* Go 1.24+
* [Ebitengine](https://ebitengine.org)

---

## 📄 License

This project is released under the MIT License. See `LICENSE` for details.

---

## 🙏 Acknowledgements

* Ebitengine / Ebiten for the game framework
* Classic Tetris community and design principles

Enjoy building and playing!
