# Math Skills — Statistics CLI

**Simple command-line program** to compute basic statistics from a file containing one numeric value per line.

---

## Overview

This project reads a text file where each line is a numeric value (the data represents a **population**) and prints the following statistics as **rounded integers**:

* Average (mean)
* Median
* Variance (population)
* Standard Deviation (population)

The program must accept the input file path as a command-line argument and print results using the exact format below so an auditor program can compare outputs easily.

---

## Objectives

* Practice file I/O and argument parsing
* Implement basic statistical calculations
* Handle common edge cases reliably
* Produce deterministic, reproducible output

---

## Input format

A plain text file where each line contains a single integer number, for example:

```
189
113
121
114
145
110
...
```

The file represents a statistical population: there are no weights or grouping—each line is one observation.

---

## Output format (required)

Print the four statistics, one per line, with labels exactly as shown (values are rounded integers):

```
Average: <integer>
Median: <integer>
Variance: <integer>
Standard Deviation: <integer>
```

Example:

```
Average: 35
Median: 4
Variance: 5
Standard Deviation: 65
```

---

## Mathematical definitions & rounding

This project uses **population** formulas (denominator = *n*):

* **Average (mean)**: ( \mu = \frac{1}{n} \sum_{i=1}^{n} x_i )
* **Median**: middle value after sorting. If *n* is even, median is ( (x_{n/2}+x_{n/2+1})/2 ) (1-based indices) then rounded.
* **Variance (population)**: ( \sigma^2 = \frac{1}{n} \sum_{i=1}^{n} (x_i - \mu)^2 )
* **Standard deviation (population)**: ( \sigma = \sqrt{\sigma^2} )

**Rounding policy:** round to the nearest integer. For values exactly at .5, use the language/runtime default "round half away from zero" or standard `math.Round()` (for Go use `math.Round`).

Internally use `float64` for intermediate calculations to avoid overflow and preserve precision.

---

## Behavior & Error handling

* If the input file does not exist or cannot be opened: print a meaningful error to `stderr` and exit with a non-zero code.
* If the file is empty: print an error to `stderr` and exit non-zero.
* If any line cannot be parsed as a number (non-numeric content): print an error indicating the faulty line (line number) and exit non-zero.
* For a single value: `Variance` and `Standard Deviation` should be `0`.

This strict behavior helps the auditor detect malformed input quickly.

---

## Implementation notes (recommendations)

* Read the entire file into memory (acceptable for typical auditor tests). If you expect extremely large files, implement a streaming approach for computing mean and variance online (Welford's algorithm).
* For median, you will need to sort the data set. Use built-in sort utilities.
* Use `float64` for sums and variance calculation; convert final results to integers with rounding before printing.

**Suggested approach in Go:**

1. Read file into `[]string` lines.
2. Parse lines to `[]float64` (or `[]int` then convert) and validate.
3. Compute `sum` and `n` to get `mean`.
4. Sort the numeric slice and compute `median`.
5. Compute variance by summing squared differences from `mean` and dividing by `n`.
6. Compute standard deviation as `math.Sqrt(variance)`.
7. Print values using `math.Round` then convert to `int`.

---

## Usage examples (Go)

**Run without building:**

```bash
$ go run main.go data.txt
```

**Build and run:**

```bash
$ go build -o math-skills
$ ./math-skills data.txt
```

**Expected stdout** (example):

```
Average: 35
Median: 4
Variance: 5
Standard Deviation: 65
```

Errors are printed to `stderr`.

---

## Testing notes for the auditor

* The auditor will provide randomly generated datasets and a reference program. Ensure your program prints **exactly** the labeled lines and uses integer values after rounding—no extra whitespace lines or debug prints.
* Make sure the program exits non-zero for invalid inputs (missing file, malformed line, empty file).

---

## Example input & computed steps

Input:

```
10
20
30
40
```

* Average: (10+20+30+40)/4 = 25 → `25`
* Median: (20+30)/2 = 25 → `25`
* Variance (population): average of squared deviations = ((225+25+25+225)/4) = 125 → `125`
* Std dev: sqrt(125) ≈ 11.18034 → rounded: `11`

Output:

```
Average: 25
Median: 25
Variance: 125
Standard Deviation: 11
```

---

## Contributing

Pull requests are welcome. Keep outputs deterministic and maintain the exact output format.

---
