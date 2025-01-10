package main

import (
	//"bufio"
	"os"
)

func main(){
	// f,err := os.Open("example.txt")

	// if err != nil{
	// 	panic("Error Occured while reading")  // panic in Go is used for error handling , panic returns a string
	// }
	
	// filestatus,err := f.Stat() 

	// if err != nil {
	// 	panic("Issue in file status")
	// }

	// fmt.Printf("File name %s\n",filestatus.Name())
	// fmt.Printf("File Size %d\n",filestatus.Size())
	// //fmt.Printf("%s\n",filestatus.Sys()) // gives the system specific information like inode number etc , inode number is a unique number given to each file in a filesystem
	// fmt.Printf("Modified Time %s\n",filestatus.ModTime())
	// fmt.Printf("File or Folder %t\n",filestatus.IsDir()) // returns true if it is a folder and false if it is a file , %t can be used to print boolean values
	// fmt.Printf("Permission %s\n",filestatus.Mode())





	// f,err := os.Open("example.txt")

	// defer f.Close()

	// if err != nil {
	// 	panic(err)
	// }
	// filStatus,err := f.Stat()

	// if err != nil {
	// 	panic(err)
	// }

	// //read the file

	// buf := make([]byte,filStatus.Size())

	// len,err:= f.Read(buf) // returns length and error

	// if err != nil{
	// 	panic(err)
	// }
	// println(len)
	// println(buf)

	// for _,data := range(buf){
	// 	println("data 1 is ",string(data))  // converts bits into string 
	// }



	// easy way to read a file 


	// f,error := os.ReadFile("example.txt")

	// if error != nil {
	// 	panic(error)
	// }
	// println(string(f))


	// Read file

	// dir, err := os.Open("..") // it can read both files and folders
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close() // its import to close the file after reading it because it will consume memory

	// fileInfo,err:= dir.ReadDir(0)  // 0 is the number of files to read , 0 means all files, 


	// for _,file := range fileInfo{
	// 	println(file.Name(), file.IsDir())
	// }


	//write in a file

	// f, err := os.Create("navamita.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// data := []byte("hello mom") // file is made up of bites
	// f.Write(data)
	// f.WriteString("\nNavamita i will find you")





	// transfer text from one file to other 

	// fileRead, err := os.Open("example.txt")


	// if err != nil {
	// 	panic(err)
	// }

	// defer fileRead.Close()


	// fileWrite, err := os.Create("example2.txt")

	
	// if err != nil {
	// 	panic(err)
	// }
	
	// defer fileWrite.Close()


	// read := bufio.NewReader(fileRead)
	// write := bufio.NewWriter(fileWrite)


	// for {
	// 	b,err := read.ReadByte()

	// 	if err != nil {

	// 		if err.Error() != "EOF"{
	// 			panic(err)
	// 		}
	// 		break
			
	// 	}


	// 	e := write.WriteByte(b)

	// 	if e != nil{
	// 		panic(e)
	// 	}
	// }

	// write.Flush()

	// println("file writting is complete")


	// Copy file function is also there

	//remove file

	err := os.Remove("example2.txt")


	if nil != err {
		panic(err)
	}

	print("file deleted succesfully")


	



	


	
	

	

}

