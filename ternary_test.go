package ternary

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func ExampleGiveOnSuccess() {
	fmt.Println(
		GiveOnSuccess(strconv.ParseFloat("-9.0", 64)).
			Else(math.NaN()),
	)
	// Output: -9
}

func ExampleGiveOnSuccess_second() {
	fmt.Println(
		GiveOnSuccess(strconv.ParseFloat("-9.0.0", 64)).
			Else(math.NaN()),
	)
	// Output: NaN
}

func ExampleGiveOnSuccess_third() {
	fmt.Println(
		GiveOnSuccess(strconv.ParseFloat("-9.0.0", 64)).Else(
			GiveOnSuccess(strconv.ParseFloat("-8.0.0", 64)).Else(
				GiveOnSuccess(strconv.ParseFloat("-7.0", 64)).Else(
					math.NaN(),
				),
			),
		),
	)
	// Output: -7
}

func ExampleGiveOnOK() {
	x := struct {
		Y string `tagged:"yes"`
	}{}
	fmt.Println(
		GiveOnOK(
			GiveOnOK(
				reflect.TypeOf(x).FieldByName("Y"),
			).Else(reflect.StructField{}).Tag.Lookup("tagged"),
		).Else("no tagged"),
	)
	// Output: yes
}
func ExampleGiveOnOK_second() {
	x := struct {
		Y string `tagged:"yes"`
	}{}
	fmt.Println(
		GiveOnOK(
			GiveOnOK(
				reflect.TypeOf(x).FieldByName("N"),
			).Else(reflect.StructField{}).Tag.Lookup("tagged"),
		).Else("no tagged"),
	)
	// Output: no tagged
}
func ExampleGiveOnOK_third() {
	x := struct {
		Y string `bagged:"yes"`
	}{}
	fmt.Println(
		GiveOnOK(
			GiveOnOK(
				reflect.TypeOf(x).FieldByName("Y"),
			).Else(reflect.StructField{}).Tag.Lookup("tagged"),
		).Else("no tagged"),
	)
	// Output: no tagged
}
