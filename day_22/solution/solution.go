package solution

import (
	"fmt"
	"time"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func mix(secret int, next_secret int) int {
	return secret ^ next_secret
}

func prune(secret int) int {
	return secret % 16777216
}

func first_step(secret int) int {
	result := secret * 64
	result = mix(secret, result)
	result = prune(result)

	return result
}

func seconde_step(secret int) int {
	result := secret / 32

	result = mix(secret, result)
	result = prune(result)
	return result
}

func third_step(secret int) int {
	result := secret * 2048

	result = mix(secret, result)
	result = prune(result)
	return result
}

func makePathKey(a, b, c, d int) string {
	return fmt.Sprintf("%d,%d,%d,%d", a, b, c, d)
}

func Part_One(initial_secret []int) int {
	defer Track(time.Now(), "Part 1")

	total := 0
	for _, secret := range initial_secret {
		res := secret
		for i := 0; i < 2000; i++ {
			res = first_step(res)
			res = seconde_step(res)
			res = third_step(res)
		}
		// fmt.Printf("res: %v\n", res)
		total += res
	}
	return total
}

func Part_Two(initial_secret []int) int {
	defer Track(time.Now(), "Part 2")

	prices_for_sequences := make(map[string]int)
	for _, secret := range initial_secret {

		res := secret
		list_price := []int{}
		last_price := secret % 10
		visited := make(map[string]bool)
		for i := 0; i < 2000; i++ {
			res = first_step(res)
			res = seconde_step(res)
			res = third_step(res)

			current_price := res % 10

			list_price = append(list_price, current_price-last_price)

			if i >= 3 {
				key := makePathKey(list_price[i-3], list_price[i-2], list_price[i-1], list_price[i])
				if visited[key] {
					continue
				}
				if key == "1,0,-1,1" {
					fmt.Printf("%v\n", current_price)
					fmt.Printf("%v\n", secret)
				}

				visited[key] = true
				prices_for_sequences[key] += current_price
			}

			last_price = current_price
		}

		// fmt.Printf("list_price: %v\n", list_price)
		// fmt.Printf("list_price: %v\n", max_value)
		// prices = append(prices)
		// fmt.Printf("res: %v\n", res)
	}

	max_value := 0
	for _, prices := range prices_for_sequences {
		if prices > max_value {
			max_value = prices
			// fmt.Printf("prices: %v\n", key)
			fmt.Printf("prices: %v\n", prices)
		}
	}
	fmt.Printf("max_value: %v\n", prices_for_sequences["1,0,-1,1"])
	return max_value
}
