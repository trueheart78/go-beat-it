# Go! Beat It!

A command-line client that reaches out to [HowLongToBeat.com](https://howlongtobeat.com) and returns related data.


## Usage

It returns the first page of results (18 max), so being specific helps.

```sh
hltb -g "final fantasy 6"

1. Final Fantasy VI (Final Fantasy III NA)
game.php?id=3519
Main    : 35.5
Extra   : 40.5
Complete: 65
Average : 41
```

```sh
hltb -g "horizon zero dawn"

1. Horizon Zero Dawn
game.php?id=26784
Main    : 22
Extra   : 43.5
Complete: 59.5
Average : 48.5

2. Horizon Zero Dawn - Complete Edition
game.php?id=53247
Main    : 37.5
Extra   : 53.5
Complete: 82.5
Average : 65

3. Horizon Zero Dawn: The Frozen Wilds
game.php?id=46430
Main    : 8.5
Extra   : 13
Complete: 19.5
Average : 14.5
```
