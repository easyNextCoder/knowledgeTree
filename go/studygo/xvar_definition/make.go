package xvar_definition

import "fmt"

//make 只是用来定义chan slice map 并且生成的不是指针
func make0() {
	var arr0 = make([]interface{}, 0)
	arr0 = append(arr0, new(interface{}))
	v := new([]int)
	*v = append(*v, 1)
	fmt.Println("make0 arr0:", arr0, *v)
}

func make1() {
	var arr1 = make([]interface{}, 1)
	arr1 = append(arr1, new(interface{}))
	fmt.Println("make1 arr1:", arr1)

	okey_solution := (solution)
	var bakPlay1, bakPlay2 *Play
	var bakOkey1, bakOkey2 int
	if best_score_1.index >= 0 {
		bakPlay1 = &okey_solution.plays[best_score_1.index]
		okey_solution.plays[best_score_1.index] = best_score_1.play
		bakOkey1 = okey_solution.okey
		okey_solution.okey -= 1
	}
	if best_score_2.index >= 0 {
		bakPlay2 = &okey_solution.plays[best_score_2.index]
		okey_solution.plays[best_score_2.index] = best_score_2.play
		bakOkey2 = okey_solution.okey
		okey_solution.okey -= 1
	}

	if self.TryAddBest(okey_solution) {
		okey_solution := copySolution(okey_solution)
		self.AddBest(okey_solution)

		if best_score_2.index >= 0 {
			solution.plays[best_score_1.index] = *bakPlay1
			solution.okey = bakOkey1
		}
		if best_score_1.index >= 0 {
			solution.plays[best_score_2.index] = *bakPlay2
			solution.okey = bakOkey2
		}

	} else {
		if best_score_2.index >= 0 {
			solution.plays[best_score_1.index] = *bakPlay1
			solution.okey = bakOkey1
		}
		if best_score_1.index >= 0 {
			solution.plays[best_score_2.index] = *bakPlay2
			solution.okey = bakOkey2
		}
	}
}
