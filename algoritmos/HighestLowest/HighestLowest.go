package main

import (
	"fmt"
	"regexp"
	"strconv"
	"math"
  )
  
  func HighAndLow(in string) string {
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(in, -1)
	
	high := math.MinInt
	low := math.MaxInt
	
	for _, match := range matches {
	  num, err := strconv.Atoi(match)
	  
	  if err != nil {
		return "erro"
	  }
	  
	  if num > high {
		high = num
	  }
	  if num < low {
		low = num
	  }
	}
	
	return fmt.Sprintf("%d %d", high, low)
  }