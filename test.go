package main

import (
	"testing"
	"time"
)

type testpairFib struct {
	values int
	fib int
}


var testfib = []testpairFib{
	{ 0, 0 },
	{ 1, 1 },
	{ 10,55 },
}

func TestFib(t *testing.T) {
	for _, pair := range testfib {
		v := fib(pair.values)
		if v != pair.fib {
			t.Error(
				"For", pair.values,
				"expected", pair.fib,
				"got", v,
			)
		}
	}
}


type testpairRecPeri struct {
	values Rectangle
	perimeter float64
}

var testRecPeri = []testpairRecPeri{
	{ Rectangle{0,0,1,1},4},
	{ Rectangle{0,0,0,0},0 },
	{ Rectangle{1,1,100,10},216 },
}

func TestRecPerimeter(t *testing.T) {
	for _, pair := range testRecPeri {
		v := pair.values.perimeter()
		if v != pair.perimeter {
			t.Error(
				"For", pair.values,
				"expected", pair.perimeter,
				"got", v,
			)
		}
	}
}

type testpairCirPeri struct {
	values Circle
	perimeter float64
}

var testCirPeri = []testpairCirPeri{
	{ Circle{0,0,4},25.132741228718345},
	{ Circle{1,1,5},31.41592653589793 },
	{ Circle{10,10,10},62.83185307179586 },
}

func TestCirPerimeter(t *testing.T) {
	for _, pair := range testCirPeri {
		v := pair.values.perimeter()
		if v != pair.perimeter {
			t.Error(
				"For", pair.values,
				"expected", pair.perimeter,
				"got", v,
			)
		}
	}
}

type testpairSleep struct {
	duration time.Duration
	time time.Duration
}

var testSleep = []testpairSleep{
	{ time.Duration(10),time.Second},
	//{ time.Duration(256),time.Millisecond },
	//{ time.Duration(0),time.Hour },
}

func TestSleep(t *testing.T) {                              
	for _, pair := range testSleep {                        
		v := pair.duration*pair.time                        
		s1 := time.Now();                                   
		s2 := s1.Add(v);                                    
		Sleep(v)                                            
		if time.Now().Unix() != s2.Unix() {                 
			t.Error(                                        
				"For", pair.duration,                       
				"expected", v,                              
				"got",time.Now(),"s",s1,                    
			)                                               
		}                                                   
	}                                                       
}                                                           
