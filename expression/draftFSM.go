package expression


// Finite State Machine realization
type State uint32

const (
	// Structure is empty
	Empty State = iota
	// Structure filled with integers and the operator
	Full
)

const (
	// Command to evaluate full structure
	CmdEvaluate = "Evaluate"
	// Command that structure has few arguments to evaluate
	CmdFewArgs = "Few arguments"
)

type Turnstile struct {
	State State
}

type CmdStateTupple struct {
	Cmd   string
	State State
}

type TransitionFunc func(state *State)

func (p *Turnstile) ExecuteCmd(cmd string) {
	tupple := CmdStateTupple{strings.TrimSpace(cmd), p.State}
	if f := StateTransitionTable[tupple]; f == nil {
		fmt.Println("unknown command, try again please")
	} else {
		f(&p.State)
	}
}

var StateTransitionTable = map[CmdStateTupple]TransitionFunc{
	{CmdFewArgs, Empty}: func(state *State) {
		fmt.Println("Sorry, need more args")
		*state = Empty
	},
	{CmdEvaluate, Full}: func(state *State) {
		fmt.Println("Ready to evaluate")
	},
}

func prompt(s State) {
	m := map[State]string{
		Empty: "Empty",
		Full:  "Full",
	}
	fmt.Printf("current state is [%s], please input command [Evaluate|Few arguments]\n", m[s])
}


process := &Turnstile{State: Empty}
	prompt(process.State)

	reader := bufio.NewReader(os.Stdin)
	for {
		cmd, err := reader.ReadString('n')
		if err != nil {
			log.Fatalln(err)
		}
		process.ExecuteCmd(cmd)
	}