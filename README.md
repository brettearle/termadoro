
# Termadoro

**Termadoro** is a lightweight, terminal-based Pomodoro timer designed for simplicity and productivity. It allows you to structure focused work sessions and breaks using the classic Pomodoro technique. Directly from your command line.

---

## 🚀 Installation

You can install **Termadoro** (as the `tdoro` command) using a simple shell script. This works on most Linux and macOS systems with either `bash` or `zsh`.

### ✅ One-Line Install (bash/zsh)

Run this command in your terminal:

```bash
curl -sSL https://raw.githubusercontent.com/brettearle/termadoro/main/install.sh | bash
```

Or, if you're using `zsh`:

```zsh
curl -sSL https://raw.githubusercontent.com/brettearle/termadoro/main/install.sh | zsh
```

This will:
- Detect your OS and architecture
- Download the correct binary from the latest GitHub release
- Install it as `tdoro` in `/usr/local/bin`

### 🧪 Manual Install

If you’d prefer to review the script first:

```bash
curl -sSL https://raw.githubusercontent.com/brettearle/termadoro/main/install.sh -o install.sh
chmod +x install.sh
./install.sh   # or bash install.sh / zsh install.sh
```
You can install Termadoro using `go install`:

```bash
go install github.com/brettearle/termadoro@latest
```

Alternatively, clone and build it manually:

```bash
git clone https://github.com/brettearle/termadoro.git
cd termadoro
go build -o tdoro
```

---

## 🕒 Usage

Run Termadoro with your desired focus and break durations in minutes:

```bash
tdoro 25.00 5.00
```

- `25.00` — Focus session duration (in minutes)
- `5.00` — Break duration (in minutes)

The timer will run directly in your terminal, providing a minimal and distraction-free environment.

---

## 🤝 Contributing

We welcome contributions! To contribute:

1. **Fork** the repository
2. Create a **feature branch**
3. Submit a **pull request**

Please keep PRs focused and clearly documented. For any bugs or feature requests, feel free to [open an issue](https://github.com/brettearle/termadoro/issues).

---

## 📄 License

This project is licensed under the MIT License. See the [LICENSE.md](LICENSE.md) file for details.

---

## 🙋‍♂️ Support

If you encounter any problems or have suggestions for improvement, please raise an issue in the [GitHub Issues](https://github.com/brettearle/termadoro/issues) section.

---

## 🔧 Requirements

- Go 1.18 or later

---

## 📌 Features

- Terminal-native Pomodoro timer
- Lightweight and fast
- Customizable focus and break times
- No GUI — pure terminal experience

---

Stay focused. Stay productive.
