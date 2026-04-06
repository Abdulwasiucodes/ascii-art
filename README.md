# 🖼️ ASCII Art Generator (with Color Support)

A Go-based command-line application that converts text into ASCII art using banner files, with optional color support for full text or specific substrings.

---

## 📌 Overview

This project is an implementation of the classic **ASCII Art generator**, extended with **color handling capabilities**.

It reads a banner file and maps each character of the input string into a styled ASCII representation. Additionally, users can apply colors to:

* The entire output
* Specific substrings within the text

---

## 🚀 Features

### 🎯 Core Features (ASCII Art)

* Convert text into ASCII art
* Support printable ASCII characters (32–126)
* Handle multi-line input using `\n`
* Use banner file (`standard.txt`) for rendering

### 🎨 Color Features

* Apply color using `--color=<color>` flag
* Color specific substrings only
* Color multiple occurrences of a substring
* Fallback to normal output if color is invalid

### ⚙️ General

* Efficient string building using `strings.Builder`
* Clean modular structure
* Unit tests included

---

## 🛠️ Technologies Used

* **Language:** Go
* **Packages:** Standard Go libraries only

  * `fmt`
  * `os`
  * `strings`

---

## 📂 Project Structure

```id="p4u3mj"
ascii-art/
├── main.go         # CLI argument parsing
├── render.go       # ASCII rendering logic
├── color.go        # Color handling (ANSI codes)
├── standard.txt    # Banner file
├── render_test.go  # Unit tests
└── README.md
```

---

## ⚙️ Installation

```bash id="qv2o4r"
git clone https://acad.learn2earn.ng/git/wabdulra/ascii-art-color
cd ascii-art-color
```

Ensure Go is installed:

```bash id="q0xj3p"
go version
```

---

## 🧪 Usage

### ✅ Basic ASCII Art

```bash id="3h0gq1"
go run . "Hello"
```

---

### 🎨 Full Color Output

```bash id="e8r2d1"
go run . --color=red "" "Hello"
```

---

### 🎯 Color Specific Substring

```bash id="8ydp1c"
go run . --color=blue kit "a king kitten have kit"
```

👉 Colors:

* "kit" inside **kitten**
* "kit" at the end

---

### ↩️ Multi-line Input

```bash id="y4tq7w"
go run . "Hello\nWorld"
```

---

## 🎨 Supported Colors

* black
* red
* green
* yellow
* blue
* magenta
* cyan
* white
* orange

---

## ⚠️ Usage Rules (IMPORTANT)

The program must follow this exact format:

```bash id="f2j7lm"
go run . [OPTION] [STRING]

EX:
go run . --color=<color> <substring to be colored> "something"
```

### ❌ Invalid Example

```bash id="0l0o3f"
go run . --color red "Hello"
```

### Output:

```id="g1z9c0"
Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"
```

---

## 🧠 How It Works

1. Reads the banner file (`standard.txt`)
2. Splits input text using `\n`
3. Maps each character to ASCII art (8 rows per character)
4. Detects substring indices (if provided)
5. Applies ANSI color codes during rendering
6. Outputs formatted ASCII art to the terminal

---

## 🧪 Testing

Run tests with:

```bash id="q9j7pa"
go test ./...
```

---

## 📌 Example

```bash id="x9w2kp"
go run . --color=green "Go"
```

Outputs colored ASCII art in the terminal.

---

## 📈 Future Improvements

* 🌈 Gradient coloring
* 🎨 Multi-color (rainbow mode)
* 🔍 Regex-based substring coloring
* 🖼️ Multiple banner styles (shadow, thinkertoy, etc.)

---

## 👨‍💻 Author

**Abdulraufu Wasiu Olamilekan**
(*Abdulwasiucodes*)

---

## 📜 License

This project is for educational purposes as part of the Learn2Earn program.

---

If you want next:
👉 I can turn this into a **perfect GitHub repo (badges, preview images, stars-ready)**
👉 Or help you write a 
