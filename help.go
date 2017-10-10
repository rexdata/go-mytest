package main

import (

    "fmt"   
)

type NokiaPhone struct {
    // Name string
    // id string
}



func main() {
	fmt.Println("Hello,world")
	
    fmt.Println("Sorry,data")
    var rec int
    var rec1 int
    rec = getdata();
    fmt.Printf("data count:%d\n",rec)

    // NokiaPhone := NokiaPhone{"张三","22"}
    
    var phone11 Phone
    // var phone1 Phone
    
    phone11 = new(NokiaPhone)
    
    // phone1 = new(NokiaPhone)

    rec1=phone11.U_data(8)
  
    rec1=rec1+phone11.G_data("123","543")

    fmt.Printf("data:%d\n",rec1)

}

func (p NokiaPhone) G_data(str1 string,str2 string) int {
    //var aa string=p.id
    fmt.Printf("G_data:%v or  %v \n",str1,str2)
	return 1
}

func (p NokiaPhone) U_data(i_data int) int {
	return 5+i_data
}


 

