package main

import (
  "fmt"
  "time"
  "math"
)

func main() {

  fmt.Println("heyo buddy!")

      var n float64
      n = 2.0

    output:=make(chan float64)

    i:= 2.0

    for i > 0 {
      go func() {output <- dozenIt(i)}()

      i--
    }

    var hello string
    hello = "hello"
    //
    fmt.Println(hello)

    greeting(y(), m(), d())

    validateValidaiton(math.Mod(n,dozenIt(n)))
    //fmt.Println(math.Pow(n,i))
    // var power float64
    //
    // power = math.Pow(n, i)

    //fmt.Println(power)

    // for b:=power; b<=n; b++{
    //   fmt.Println("i have to tell u something!")
    //   good()

    //}
    //
    p,q := learnMemory()
    fmt.Println("address of p:",p,"address of q: ", q)

    fmt.Println("value of p:",*p,"value of q:", *q)


}

func learnMemory()(p,q *int){
  p = new(int)
  s:=make([]int, 20)
  s[3] = 7
  r:=-2
  return &s[3], &r
}

func validateValidaiton(i float64){
  for j := i; j<=10; j++{
    fmt.Println("Yeah!!!")

  }
  fmt.Println("yeah good move")
}

func ohYes(right bool){
  fmt.Println("she did that!")
}

func add(i, n float64)float64 {
  return i + n
}

func dozenIt(a float64) float64{
  return (a)*12
}

func good(){
  fmt.Println("you ROCK!")
}

func whatTimeIsItUTC() string{

  t := (time.Now().UTC())//.UnixNano())

  var outtie string

  outtie = t.String()
  return outtie
}

func y() int{
  t := (time.Now().UTC())
  var yearObj int

  yearObj = t.Year()

  return yearObj
}

func m() string{
  t := (time.Now().UTC())
  var monthObj time.Month
  var monthString string

  monthObj = t.Month()
  monthString = monthObj.String()

  return monthString
}

func d() int{
  t := (time.Now().UTC())
  var dayObj int

  dayObj = t.Day()

  return dayObj
}

func greeting( year int, month string, day int) {
  fmt.Println("Today is",month,",",day,",",year," and I believe in you!")

  for i:=day; i<=0;i--{
    fmt.Println("you can do it!")
  }
}
