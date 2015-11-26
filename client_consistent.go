package main

import (
	
	"fmt"
    "os"
	"hash/crc32"
    "net/http"
	"encoding/json"
	"io/ioutil"
)





type Response struct{
       
       Key string `json:"key"`
       Value string `json:"value"`
}



func Circle_implementation(keyy string) string {                      //uint32
 
		var server_mapping string
		hash := crc32.ChecksumIEEE
       
        first_server:=hash([]byte("3000"))                 // hashed the server localhost:3000
        second_server:=hash([]byte("3001"))               // hased the server localhost:3001
        third_server :=hash([]byte("3002"))               // hased the server localhost:3002
        
        
         
        hashing:=hash([]byte(keyy))
        
        var a bool = hashing>second_server
         
         var b bool = hashing<third_server
         //fmt.Println(b)
         var c bool = hashing>third_server
         //fmt.Println(c)
         var d bool = hashing<first_server
         //fmt.Println(d)
         var e bool = hashing<second_server
        // fmt.Println(e)
         var f bool = hashing>first_server
        // fmt.Println(f)

         if a && b {
            //fmt.Println("in If a & b")
         	server_mapping="3002"
         }else if c && d{
         	server_mapping="3000"
         } else if e || f{
         	server_mapping="3001"
         } else{
         	//fmt.Println("Do Nothing")
         }

        
        return  server_mapping 

}


func f(method string,keys string, values string) {
	// fix the input for implementing consistent hashing
	reply:=Response{}
	reply1:=Response{}
    url_entry :=Circle_implementation(keys)
    fmt.Println("\n\n")
    fmt.Println("The key "+keys+" is getting mapped to server "+url_entry)
    //fmt.Println("Server-------->mapping")
    //fmt.Println(url_entry)


                                                   
      if(method=="PUT"){
     	
        
          //client := &http.Client{}
         url:= "http://localhost:"+url_entry+"/keys/"+keys+"/"+values
         fmt.Println("\nPUT URL-"+url)
	     request, err := http.NewRequest("PUT", url, nil)
	     if err != nil {
                       panic(err)
                         }
           
                   request.Header.Set("Content-Type", "application/json")
                   
                   Client1:= &http.Client{}
                   resp1, _ := Client1.Do(request)
                   //fmt.Println("StatusCode")
                   //fmt.Println(resp1.StatusCode)
                   if err != nil {
                                      fmt.Printf("%s", err)
                                      os.Exit(1)
                                    } else {
                                      defer resp1.Body.Close()
                                      contents, err := ioutil.ReadAll(resp1.Body)
                                      if err != nil {
                                          fmt.Printf("%s", err)
                                          os.Exit(1)
                                      }
                                      //fmt.Printf("%s\n", string(contents))
                                      //var f interface{}
                                      err=json.Unmarshal(contents, &reply1)
                                      //fmt.Println(reply.Value)
                                     }




        

	   } else if(method=="GET") {

	   	puturl :=fmt.Sprint("http://localhost:"+url_entry+"/keys/"+keys)
        read1, err := http.Get(puturl)
        if err != nil {
                                   fmt.Printf("%s", err)
                                   os.Exit(1)
                                   } else {
                                   defer read1.Body.Close()
                                   content, err := ioutil.ReadAll(read1.Body)
                                   //fmt.Println(content)

                                   if err != nil {
                                    fmt.Printf("%s", err)
                                    os.Exit(1)
                                                  }

                                    err=json.Unmarshal(content, &reply)
                                    fmt.Println("Value of the key is "+reply.Value)
                        }


                       //fmt.Println("here is GET function")

     } else {

     	//fmt.Println("Print nothing")
     }

}


func main(){

   length := len(os.Args)
   var first string
   var second string 
   var third string
   

           if length==4 {                              //go run client_consistent.go PUT 1 "a"
           	   first = os.Args[1];
			   second = os.Args[2];
			   third = os.Args[3];
			  

           } else if length==3 {                     //go run client_consistent.go GET 1 

                  first = os.Args[1];
			      second = os.Args[2];
			       third = "";

           }
   

        f(first,second,third)

}