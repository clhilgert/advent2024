# 🎄 🎁 Advent of Code 2024 🎁 🎄

Advent of Code is an annual event that runs during the first 25 days of December. Each day, a new programming puzzle is released on the [Advent of Code website](https://adventofcode.com/). The puzzles typically involve problems related to algorithms, data structures, and computational thinking.

This repository contains my solutions to the Advent of Code 2024 challenges, written in **Go**. I'll be updating this repo daily with new solutions as the puzzles are released.

# AOC cli Tool
### Additionally, this repo includes a CLI tool that can be used to easily fetch each day's input.
#### Flags:
- -n (int, required): Day number (1-25)
- -y (int, optional): Year (default 2024)
- -o (string, optional): Output filename (default 'input.txt')
- --dir (string, optional): Output directory (default: day01 - day25, based on -n flag)

#### Examples:
* Day 1 default
```sh
aoc -n 1
```
>day01/input.txt
* Day 1 of 2023
```sh
aoc -n 1 -y 2023 -o file.txt --dir .
```
>file.txt

### Installation
Clone this repository
```sh
git clone git@github.com:clhilgert/advent2024.git
cd advent2024
```
Build binary
```sh
make build
```
Create ~/bin/ if it doesn't already exist and add to PATH
```sh
mkdir -p ~/bin
mv target/aoc ~/bin/
echo 'export PATH="$HOME/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```