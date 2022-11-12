package main

import {
	
}
func main(){
	a, err := os.Open("abcd.txt")
	if err != nil{
		log.fatal(err)
	}
	defer a.Close()

	b, err := ioutil.ReadAll(a)
	if err != nil {
		log.Fatal(err)
	}

	abc := string(b)
	newQR := "abcd-qr.png"

	err = qrcode.WriteFile(abc, qrcode.Medium, 512, newQR)

	if err != 

}