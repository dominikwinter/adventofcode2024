# Advent of Code 2024

This repository is a collection of meticulously crafted solutions for the [Advent of Code 2024](https://adventofcode.com/2024) challenges. Each day's solution is systematically organized into dedicated directories.

## Project Structure

- `days/`: Houses subdirectories for each day of the challenge.
  - Each day is divided into two parts: `a` and `b`.
- `lib/`: Contains reusable utility functions leveraged across multiple days.
- `run.sh`: Execute the solutions for a specified day and part.

A typical day's structure is as follows:
```
days/
  4/
    a/
      input.txt <- Note: Input files are not included in the repository
      main_test.go
      main.go
```

## Run the Solutions

To execute the solution for a specific day and part, utilize the `run.sh` script. The script requires two parameters: the day number and the part (either `a` or `b`).

```bash
./run.sh <DAY> <PART>
```

For instance, to execute the solution for day 4, part b:

```bash
./run.sh 4 b
```

## Development Workflow

The project settings are configured to automatically run unit tests for the current file upon saving. This feature is particularly advantageous when tackling the code challenges, ensuring immediate feedback and facilitating a robust development process.
