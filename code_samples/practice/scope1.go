package main

var a = 43 // package scope
var b int  // package scope

func main()  {
    //var c = 44 // function scope, not package scope
    //d := 45    // function scope
    foo()
}