import (
    "cmp"
    "slices"
)

func carFleet(target int, position []int, speed []int) int {
    stacks := []car{}
    for i, pos := range position {
        c := car{
            position: float64(pos),
            speed: float64(speed[i]),
        }
        stacks = append(stacks, c)
    }
    //Now we sort the stacks of cars
    slices.SortFunc(stacks, func(a, b car) int {
        return cmp.Compare(a.position, b.position)
    })

    total := 0
    for i:=len(stacks) - 1; i >=0; {
        total ++
        j:= i-1
        for j >= 0 && timeToTarget(stacks[i], target) >= timeToTarget(stacks[j], target) {
            j--
        }
        i = j
    }

    return total
}

type car struct {
    position float64
    speed float64
}

func timeToTarget(c car, t int) float64 {
    return (float64(t) - c.position)/c.speed
}

//Earlier solution didnt convert to float64 and just assumed discrete math over the integers.
