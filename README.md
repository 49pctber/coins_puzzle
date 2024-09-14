# Coins Puzzle

Read the following [following tweet](https://x.com/littmath/status/1834273354628424080):

> Flip 100 coins, labeled 1 through 100. Alice checks the coins in order (1, 2, 3, …) while Bob checks the odd-labeled coins, then the even-labeled ones (so 1, 3, 5, …, 99, 2, 4, 6, …). Who is more likely to see two heads *first*?

What does your gut tell you? Does either player has an advantage? If so, which one?

Spoilers abound below, so be sure to think about it before you read on. But before you do, I really appreciated [this tweet](https://x.com/OnTheTurnipTruk/status/1834274223885283833):

> This better be equally likely or I’m gonna lobby for probability studies to be banned

## Simulation #1

I simulated this game using `runGame()` one million times, and here are the results:

```
2024/09/12 15:20:43 Alice wins: 311479
2024/09/12 15:20:43 Bob wins  : 416711
2024/09/12 15:20:43 Ties      : 271810
```

### Analysis #1

Ignoring ties, Bob has about a 57% chance of winning.

The fact that Bob has a clear advantage is simple. Say Alice doesn't win after looking at the first two coins. We know for a fact that she will *not* win after the third coin; if the third coin were a head, Bob would have already won. Yet Bob's third coin could be a winner.

Using a similar argument, we can see that Alice is at a disadvantage on every odd coin. Bob doesn't have any such disadvantage, as each coin he looks at is "fresh". (This isn't true once Bob wraps around, but that chance of that even happening is miniscule.)

This might be counterintuitive, but once you think through it, it makes sense.

## Simulation #2

I initially thought this question was about who would find two heads *in a row*. I tweaked my code into `runGame2()` to simulate the new game.

I was expecting that neither of the two strategies would have an advantage, but take a look at the results:

```
2024/09/12 15:22:38 Alice wins: 340360
2024/09/12 15:22:38 Bob wins  : 451580
2024/09/12 15:22:38 Ties      : 208060
```

### Analysis #2

Ignoring ties, Bob will win about 57% of the time. **That is the same advantage as before.** (There are simply fewer ties.)

This fact was surprising to me. I suppose a similar argument could be made: for Alice to see two coins in a row, she would (practically) have to use a coin that Bob already looked at. Those coins will be biased toward not being a head.

But why are there substantially fewer ties? I think this has to do with the fact that when you see a tail, you have to restart your counter. If the other player just saw a head, you can't "catch up", leaving less opportunity for ties.

I don't know about you, but that is even more surprising than the original question.

## Simulation #3

This version of the game is inspired by [this tweet](https://x.com/49pctber/status/1834700741245858239).

> Flip 100 coins, marked 1-100. Each second, Alice and Bob simultaneously check one coin. Alice goes in order (1, 2, 3, …); Bob checks the odd coins, then the even (so 1, 3, 5, …, 99, 2, 4, 6, …). Who is more likely to see two consecutive heads OR two consecutive tails *first*?

```
2024/09/13 23:34:22 Alice wins: 374719
2024/09/13 23:34:22 Bob wins  : 250723
2024/09/13 23:34:22 Ties      : 374558
```

### Analysis #3

Now Alice has a 60% probability of winning ignoring ties.

This is the easiest to analyze. The game is completely determined by four coins 1, 2, 3, and 5. There are only sixteen possible permutations of those four coins, so we can simply enumerate them and note the outcome of each game:

- `HHH?H` (tie)
- `HHH?T` (tie)
- `HHT?H` (Alice win)
- `HHT?T` (Alice win)
- `HTH?H` (Bob win)
- `HTH?T` (Bob win)
- `HTT?H` (Alice win)
- `HTT?T` (tie)
- `THH?H` (tie)
- `THH?T` (Alice win)
- `THT?H` (Bob win)
- `THT?T` (Bob win)
- `TTH?H` (Alice win)
- `TTH?T` (Alice win)
- `TTT?H` (tie)
- `TTT?T` (tie)

Since each sequence of coins is equally likely to occur, we can count the ties and wins for each player and divide by 16 to get the probability of that happening.

There are 6 sequences that lead to an Alice win, 6 that lead to a tie, and 4 that lead to a Bob win. Those correspond to probabilities of 0.375, 0.375, and 0.25, respectively, which closely matches the simulated results.

Intuitively, this advantage can be explained by the fact that if neither Alice nor Bob won on their second coin, Alice can only win or tie (but not lose) on her third. This is because coins 1 and 3 are necessarily different because Bob didn't win on his second coin. (Equivalently, coins 2 and 3 must be the same.) The best Bob can do in that situation is tie when coins 3 and 5 are the same.

## Related Puzzle

Daniel Litt quoted [this tweet](https://x.com/GilKalai/status/1830970396352618561) as being a "closely related...beautiful puzzle". The full blog post can be found [here](https://gilkalai.wordpress.com/2024/09/03/test-your-intuition-56-fifteen-boxes-puzzle/).
