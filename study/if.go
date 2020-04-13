package main

import "fmt"

func main() {
  //반복문 - For
  //Go에서 유일한 반복문
  //다양한 사용법 숙지

  // var a int =20
  // b :=30
  //
  // //예제1
  // if a>=15{
  //   fmt.Println("15 이상")
  // }
  // if(b>=25){
  //   fmt.Println("25 이상")
  // }
  //
  // if c:=40; c>=30{
  //   fmt.Println("30 이상")
  // }

  var a int =20

  if a>=65{
    fmt.Println("65 이상 ")
  } else{
    fmt.Println("65 이하 ")
  }

  i :=100
  if i>200{
    fmt.Println("120 이상")
  } else if i>=100 && i<120{
    fmt.Println("120 이상 120 미만")
  }else if i<100 && i>=50{
    fmt.Println("120 이상 120 미만")
  }
}
