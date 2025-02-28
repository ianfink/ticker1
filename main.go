/*
 * Copyright (C) 2025 Ian M. Fink.
 * All rights reserved.
 *
 * File:	main.go
 *
 * Purpose:	Experimental Golang Ticker
 *
 * This program is free software: you can redistribute it and/or modify it
 * under the terms of the GNU General Public License as published by the Free
 * Software Foundation, either version 3 of the License, or (at your option)
 * any later version.
 *
 * This program is distributed in the hope that it will be useful, but
 * WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY
 * or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for
 * more details.
 * 
 * You should have received a copy of the GNU General Public License along
 * with this program. If not, please see: https://www.gnu.org/licenses
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

