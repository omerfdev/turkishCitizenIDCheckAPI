package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"unicode"
)

func validateTCK(tck string) error {
	if tck == "" {
		return fmt.Errorf("TCK cannot be empty.")
	} else if len(tck) != 11 {
		return fmt.Errorf("TCK must consist of 11 digits.")
	} else if tck[0] == '0' {
		return fmt.Errorf("The first digit of TCK cannot be zero.")
	}

	var total int
	for i := 0; i < 10; i++ {
		total += int(tck[i] - '0')
	}

	if total%10 != int(tck[10]-'0') {
		return fmt.Errorf("Invalid TCK.")
	}

	onlyDigits := true
	for _, c := range tck {
		if !unicode.IsDigit(c) {
			onlyDigits = false
			break
		}
	}

	digit1 := int(tck[0] - '0')
	digit2 := int(tck[1] - '0')
	digit3 := int(tck[2] - '0')
	digit4 := int(tck[3] - '0')
	digit5 := int(tck[4] - '0')
	digit6 := int(tck[5] - '0')
	digit7 := int(tck[6] - '0')
	digit8 := int(tck[7] - '0')
	digit9 := int(tck[8] - '0')
	digit10 := int(tck[9] - '0')
	digit11 := int(tck[10] - '0')

	oddSum := digit1 + digit3 + digit5 + digit7 + digit9
	evenSum := digit2 + digit4 + digit6 + digit8
	result := (oddSum*7 - evenSum) % 10

	if result != digit10 {
		return fmt.Errorf("Invalid TCK.")
	}

	if digit11 != (total % 10) {
		return fmt.Errorf("Invalid TCK.")
	}

	if !onlyDigits {
		return fmt.Errorf("Invalid TCK.")
	}

	return nil
}

func validateHandler(c *gin.Context) {
	tck := c.Query("tck")

	err := validateTCK(tck)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "TCK is valid."})
	}
}

func main() {
	r := gin.Default()

	// Endpoint to validate TCK
	r.GET("/validateTCK", validateHandler)

	// Run the server
	r.Run(":8080")
}
