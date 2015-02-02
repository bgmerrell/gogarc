# gogarc

Gogarc is a turn-based fantasy game (in very early development) for group chat environments.  It includes an IRC bot client.

Gogarc draws inspiration from the [Talisman: Digital Edition](http://www.talisman-game.com) digital board game.

## Example

![00:32 -!- Gogarc [gogarc@li51-104.members.linode.com] has joined #gogarc
00:32 < bean> .join
00:33 < Gogarc> bean joined the game
00:33 < bgmerrell> .join
00:33 < Gogarc> bgmerrell joined the game
00:33 < bgmerrell> .start
00:33 < Gogarc> Game has begun.  Player order: bgmerrell, bean.
00:33 < bean> .travel
00:33 < Gogarc> bean: It's not your turn.
00:33 < bgmerrell> .travel
00:33 < Gogarc> bgmerrell: You encounter an Orc (Health: 1, Vigor: 6, Wits: 2) 
                that challenges your vigor.
00:33 < Gogarc> bgmerrell: Your enemy's attack score is 8 (2 +6)
00:33 < bgmerrell> .attack
00:33 < Gogarc> bgmerrell: Your attack score is 10 (6 +4). You win!
00:33 < Gogarc> It is bean's turn
00:33 < bean> .travel
00:33 < Gogarc> bean: You encounter a Kobold (Health: 1, Vigor: 1, Wits: 1) 
                that challenges your vigor.
00:33 < Gogarc> bean: Your enemy's attack score is 6 (5 +1)
00:33 < bean> .attack
00:33 < Gogarc> bean: Your attack score is 6 (2 +4). It's a draw.
00:33 < Gogarc> It is bgmerrell's turn
00:33 < bean> .stats
00:33 < Gogarc> bean: Health: 4, Luck: 4, Vigor: 4, Wits: 4
00:33 < bgmerrell> .stats
00:33 < Gogarc> bgmerrell: Health: 4, Luck: 4, Vigor: 4, Wits: 4](https://raw.githubusercontent.com/bgmerrell/gogarc/master/images/irc-example.png "Example IRC Screenshot")

## Getting and Running

`git clone https://github.com/bgmerrell/gogarc.git`

`cd gogarc && go run gogarc-irc.go`

Note: For now, Gogarc must be run from the root "gogarc" source directory so the game contents can be found.

## Commands

**.join**: Join a game

**.start**: Start a game

**.stats [nick ...]**: See a players stats.  If no nick is specified, the stats are shown for the player who issued the command.

**.travel**: Travel to a new destination with new occurences (currently only enemy encounters)

**.attack**: Attack an enemy
