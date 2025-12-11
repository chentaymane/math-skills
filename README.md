📘 Math Skills – Go Project

A simple Go program that reads a list of numeric values from a text file, converts them into integers, and calculates essential statistical metrics:

Average (Mean)

Median

Variance

Standard Deviation

This project is designed for practice with Go fundamentals such as file handling, slices, error management, and mathematical computations.

📂 Project Structure
math-skills/
│
├── main.go
├── go.mod
└── operations/
    ├── convert.go
    ├── average.go
    ├── median.go
    ├── variance.go
    └── standard_deviation.go

⚙️ How It Works

The program expects one argument, which is the path to a file.

The file must contain numbers separated by spaces or newlines.

The program:

Reads the file

Splits the content into strings

Converts them into integers

Performs basic mathematical operations

It prints all results cleanly in the terminal.

📄 Example Input File

numbers.txt

10 20 30 40 50 60 70

▶️ Usage

Run the program with:

go run . numbers.txt

Example Output
Average: 40
Median: 40
Variance: 400
Standard Deviation: 20

🛠️ Features
✔️ Average

Sum of all elements divided by count.

✔️ Median

If odd count → middle element

If even count → average of the two middle elements
The program sorts the slice internally.

✔️ Variance

Measures spread of values compared to the mean.

✔️ Standard Deviation

Square root of the variance (uses math.Sqrt).

❗ Error Handling

The program safely handles:

🔸 Missing argument
error : 2argument

🔸 File not found
open numbers.txt: no such file or directory

🔸 Non-numeric value in the file
error: invalid number "abc"


All errors are cleanly returned and stop execution.

🧪 Example File With Invalid Data
10 20 x 30


Output:

error: invalid number "x"

📌 Requirements

Go 1.18+

A valid text file containing numbers

🧠 What You Learn From This Project

Working with files in Go (os.ReadFile)

Splitting and parsing string data

Managing errors properly

Slices and conversions

Implementing statistical algorithms

Understanding float vs int behavior

Using Go packages and modular code organization

🤝 Contributing

Feel free to fork the repo, submit pull requests, or open issues.

📜 License

MIT License — free to use and modify.
