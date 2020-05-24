# Gophercises

A collection of projects written in Go

## Project 1: QUiz Game
A CLI program where a csv file is read which contains a list of questions and answers. It will then quiz the user and keep track of how many correct / incorrect answers. 

Regardless if an answer is correct or incorrect the next question shall be immediately asked. 

The CSV file should default to problems.cvs but the user should be able to customize the filename via a flag.

The CSV file will be in a format like below where the first column is the question and the second column in the same row is the answer to that question

```
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

At the end of the quiz game out the total number of questions correct and how many questions there were in total. 