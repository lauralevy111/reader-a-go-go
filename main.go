package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main() {

  reader := bufio.NewReader(os.Stdin)
  fmt.Println("hey girl!")
  fmt.Println("-------------------------")

  for {
    fmt.Print("->")
    text, _ := reader.ReadString('\n')
    //convert CRLF to LF
    text = strings.Replace(text, "\n","",-1)

    if strings.Contains("hi",text) == true {
      fmt.Println("Hello, yourself")
  //    break;
    }
    if strings.Compare("bye",text) ==0 {
      fmt.Println("goodbye")
      break;
    }
    if strings.Compare("help",text)>=0 {
      fmt.Println("All you need is within you now!")
      //fmt.Println(strings.Compare("help",text))
    }

    if strings.Compare("hey",text)>=0 {
      fmt.Println("All you need is within you now!")
      //fmt.Println(strings.Compare("help",text))
    }

    if strings.Compare("hello",text)>=0 {
      fmt.Println("All you need is within you now!")
      //fmt.Println(strings.Compare("help",text))
    }

  }

}
