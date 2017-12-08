package main

type program struct {
	Name        string
	Weight      int
	SubPrograms []string
}

type runtime struct {
	programs     map[string]*program
	totalWeights map[string]int
}

func newRuntime(programs []*program) *runtime {
	ps := make(map[string]*program, len(programs))
	for _, p := range programs {
		ps[p.Name] = p
	}

	return &runtime{
		programs:     ps,
		totalWeights: make(map[string]int),
	}
}

func (r *runtime) program(name string) *program {
	return r.programs[name]
}

func (r *runtime) totalWeight(name string) int {
	if w, found := r.totalWeights[name]; found {
		return w
	}

	p := r.program(name)

	w := p.Weight
	for _, n := range p.SubPrograms {
		w += r.totalWeight(n)
	}

	r.totalWeights[name] = w

	return w
}
