/*
 * Copyright (C) 2025 Ian M. Fink.
 * All rights reserved.
 *
 * File:	main.go
 *
 * Purpose:	Experiment Ticker
 *
 */

package main

import (
	"fmt"
	"time"
)

/*
 * Globals
 */

var theTicks	int

/******************************************************************************/

func tickMeOff() {
	fmt.Printf("I've been ticked off! theTicks = %d\n", theTicks)
	theTicks++
} /* tickMeOff */

/******************************************************************************/

func main() {
	var ticker		*time.Ticker

	theTicks = 0
	fmt.Println("Get ready to be ticked off!!!")

	ticker = time.NewTicker(2 * time.Second)

	for {
		<-ticker.C		// Wait for the next tick
		tickMeOff()
	}

	// This will never be executed:
	ticker.Stop()
} /* main */

/******************************************************************************/

/*
 * End of file:	main.go
 */

