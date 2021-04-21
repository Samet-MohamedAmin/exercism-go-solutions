package tournament

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

type team struct {
	name string
	win  int
	draw int
	loss int
}

var teams map[string]*team

func (t team) getMatchesPlayed() int {
	return t.win + t.draw + t.loss
}

func (t team) getPoints() int {
	return t.win*3 + t.draw
}

// updateTeamScore update teams score for a given game
func updateTeamScore(n1, n2, state string) {
	if teams[n1] == nil {
		teams[n1] = &team{name: n1}
	}
	if teams[n2] == nil {
		teams[n2] = &team{name: n2}
	}
	switch state {
	case "win":
		teams[n1].win++
		teams[n2].loss++
	case "draw":
		teams[n1].draw++
		teams[n2].draw++
	case "loss":
		teams[n1].loss++
		teams[n2].win++
	}
}

// getSortedTeams returns sorted array of teams
func getSortedTeams() []team {
	ta := []team{}
	for _, t := range teams {
		ta = append(ta, *t)
	}

	sort.SliceStable(ta, func(i, j int) bool {
		p1, p2 := ta[i].getPoints(), ta[j].getPoints()
		if p1 == p2 {
			return ta[i].name[0] < ta[j].name[0]
		}
		return p1 > p2
	})

	return ta
}

// extractLine extracts team1 name, team2 name and state from a given name
func extractLine(l string) (string, string, string, error) {
	game := strings.Split(l, ";")
	if l == "" || l[0] == '#' {
		return "", "", "", nil
	}
	if len(game) != 3 {
		return "", "", "", errors.New("bad format. should be .*;.*;.*")
	}
	t1, t2, state := game[0], game[1], game[2]
	if t1 == t2 {
		return "", "", "", errors.New("teams have the same name")
	}
	if state != "win" && state != "draw" && state != "loss" {
		return "", "", "", errors.New("bad state. should be \"win\" or \"draw\" or \"loss\"")
	}
	return t1, t2, state, nil
}

// extractData extract trournament data from given reader
// and update teams scores
func extractData(reader io.Reader) error {
	buf, _ := ioutil.ReadAll(reader)
	lines := strings.Split(string(buf), "\n")
	for _, l := range lines {
		game := strings.Split(l, ";")
		if l == "" || l[0] == '#' {
			continue
		}
		if len(game) != 3 {
			return errors.New("bad format. should be .*;.*;.*")
		}
		t1, t2, state := game[0], game[1], game[2]
		if t1 == t2 {
			return errors.New("teams have the same name")
		}
		if state != "win" && state != "draw" && state != "loss" {
			return errors.New("bad state. should be \"win\" or \"draw\" or \"loss\"")
		}
		updateTeamScore(t1, t2, state)
	}
	return nil
}

// writeResult writes teams result in given writer buffer
func writeResult(writer io.Writer) {
	header := fmt.Sprintf("%-30s | MP |  W |  D |  L |  P\n", "Team")
	writer.Write([]byte(header))
	for _, t := range getSortedTeams() {
		line := fmt.Sprintf(
			"%-30s | %2d | %2d | %2d | %2d | %2d\n",
			t.name,
			t.getMatchesPlayed(),
			t.win,
			t.draw,
			t.loss,
			t.getPoints(),
		)
		writer.Write([]byte(line))
	}
}

// Tally extracts played trournament matches data from reader input
// and writes tournament results in given writer
func Tally(reader io.Reader, writer io.Writer) error {

	teams = map[string]*team{}

	err := extractData(reader)
	if err != nil {
		return err
	}
	writeResult(writer)

	return nil
}
