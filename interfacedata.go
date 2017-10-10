package main

import (
	
		_"fmt"   
	)


type Phone interface{
    U_data(i_data int)  int
    G_data(str1 string,str2 string) int
}

type Changeinfo interface{
    Get_data(str1 string,str2 string) (rec int)
    Udate_data(i_data int) (recdata int)
}
