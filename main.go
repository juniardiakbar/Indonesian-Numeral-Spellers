package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func hundreds(q int) string {
	// This function will convert number per thee digit number start from unit
	num := [12]string{"", "satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh", "sebelas"}
	if q < 12 {
		return num[q]
	} else if q < 20 {
		return num[q%10] + " belas"
	} else if q < 100 {
		var temp = q / 10
		return num[temp] + " puluh " + num[q%10]
	} else {
		var temp = q / 100
		if temp == 1 {
			if q%100 == 0 {
				return "seratus"
			} else if q%100 < 12 {
				return "seratus " + num[q%100]
			} else if q%100 < 20 {
				return "seratus " + num[q%100%10] + " belas"
			} else {
				var temp2 = (q % 100) / 10
				return "seratus " + num[temp2] + " puluh " + num[q%10]
			}
		} else {
			if q%100 == 0 {
				return num[temp] + " ratus "
			} else if q%100 < 12 {
				return num[temp] + " ratus " + num[q%100]
			} else if q%100 < 20 {
				return num[temp] + " ratus " + num[q%100%10] + " belas"
			} else {
				var temp2 = (q % 100) / 10
				return num[temp] + " ratus " + num[temp2] + " puluh " + num[q%10]
			}
		}
	}
}

func convert(q int, l int) string {
	hund := [4]string{"", " ribu ", " juta ", " miliyar "}
	temp := q % 1000
	spell := ""
	i := 0
	for q > 0 {
		// Pengejaan dilakukan per tiga digit
		if i == 1 && temp == 1 {
			spell = "seribu " + spell
		} else {
			spell = hundreds(temp) + hund[i] + spell
		}
		q = q / 1000
		temp = q % 1000
		i = i + 1
		fmt.Println(spell)
	}
	return spell
}

func spell(c *gin.Context) {
	q := c.Query("number")
	number, err := strconv.Atoi(q)
	if err == nil {
		if number > 0 {
			var s = convert(number, len(q))
			c.JSON(200, gin.H{
				"text": s,
			})
		} else {
			c.JSON(200, gin.H{
				"text": "error",
			})
		}
	} else {
		c.JSON(200, gin.H{
			"text": "error",
		})
	}
}

func toInt(x string) int {
	if x == "satu" {
		return 1
	} else if x == "dua" {
		return 2
	} else if x == "tiga" {
		return 3
	} else if x == "empat" {
		return 4
	} else if x == "lima" {
		return 5
	} else if x == "enam" {
		return 6
	} else if x == "tujuh" {
		return 7
	} else if x == "delapan" {
		return 8
	} else if x == "sembilan" {
		return 9
	} else if x == "sepuluh" {
		return 10
	} else if x == "sebelas" {
		return 11
	} else if x == "seratus" {
		return 100
	} else if x == "seribu" {
		return 1000
	} else {
		return -1
	}
}

func conv(x string) int {
	if x == "puluh" {
		return 10
	} else if x == "ratus" {
		return 100
	} else if x == "ribu" {
		return 1000
	} else if x == "juta" {
		return 1000000
	} else if x == "miliyar" {
		return 1000000000
	} else {
		return -1
	}
}

func cekValid(x []string) int {
	// THis function will check vilidity of user input
	i := 0
	for i < len(x) {
		if i+1 < len(x) {
			if toInt(x[i]) != -1 && toInt(x[i+1]) != -1 {
				if x[i] != "seribu" && x[i] != "seratus" {
					return toInt(x[i+1])
				}
			}
		}
		i = i + 1
	}

	i = 0
	for i < len(x) {
		if i+1 < len(x) {
			if (x[i] == "puluh" || x[i] == "ratus" || x[i] == "belas") && (x[i+1] == "puluh" || x[i+1] == "ratus" || x[i+1] == "belas") {
				return 2
			}
		}
		i = i + 1
	}

	i = 0
	for i < len(x) {
		if i+1 < len(x) {
			if (x[i] == "ribu" || x[i] == "juta" || x[i] == "miliyar") && (x[i+1] == "ribu" || x[i+1] == "juta" || x[i+1] == "miliyar") {
				return 3
			}
		}
		i = i + 1
	}

	i = 0
	for i < len(x) {
		if i+1 < len(x) {
			if (x[i] == "ribu" || x[i] == "juta" || x[i] == "miliyar") && (x[i+1] == "puluh" || x[i+1] == "belas" || x[i+1] == "ratus") {
				return 4
			}
		}
		i = i + 1
	}

	m := 99
	j := 99
	r := 99
	i = 0
	for i < len(x) {
		if x[i] == "miliyar" && m == 99 {
			m = i
		} else if x[i] == "miliyar" && m != 99 {
			return 5
		} else if x[i] == "juta" && j == 99 {
			j = i
		} else if x[i] == "juta" && j != 99 {
			return 6
		} else if x[i] == "ribu" && r == 99 {
			r = i
		} else if x[i] == "ribu" && r != 99 {
			return 7
		}
		i = i + 1
	}
	if m != 99 {
		if (m <= j) && (j <= r) {
			return -1
		} else {
			return 8
		}
	} else if j != 99 {
		if j <= r {
			return -1
		} else {
			return 9
		}
	} else {
		return -1
	}
}

func read(c *gin.Context) {
	q := c.Query("text")
	s := strings.Split(q, " ")
	i := 0
	r := 0
	valid := cekValid(s)
	fmt.Println(valid)
	if valid == -1 {
		for i < len(s) {
			ribu := 0
			for s[i] != "ribu" && s[i] != "juta" && s[i] != "miliyar" {
				// Read start from left to right and this loop will read per three digits
				temp := 0
				temp = temp + toInt(s[i])
				i = i + 1
				if i < len(s) {
					if conv(s[i]) != -1 && s[i] != "ribu" && s[i] != "juta" && s[i] != "miliyar" {
						temp = temp * conv(s[i])
						i = i + 1
					} else if s[i] == "belas" {
						temp = temp + 10
						i = i + 1
					}
				}
				ribu = ribu + temp
				if i >= len(s) {
					break
				}
			}
			if i < len(s) {
				ribu = ribu * conv(s[i])
			}
			r = r + ribu
			fmt.Println(r)
			i = i + 1
		}
		c.JSON(200, gin.H{
			"number": r,
		})
	} else {
		c.JSON(200, gin.H{
			"number": "error",
		})
	}
}

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AddAllowHeaders("*")
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET"}
	r.Use(cors.New(config))
	r.GET("/spell", spell)
	r.POST("/read", read)
	r.Run()
}
