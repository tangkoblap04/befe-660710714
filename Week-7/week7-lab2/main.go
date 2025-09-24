package main

import (
	"fmt"
	"os"
)

func getEnv(key, defaultValue string) string{
	if value := os.Getenv(key) ; value != ""{
		return value
	}
	return defaultValue
}
func main(){
	host := getEnv("DB_HOST", "ILoveU")
	name := getEnv("DB_NAME", "ILoveU")
	user := getEnv("DB_USER", "ILoveU")
	password := getEnv("DB_PASSWORD", "ILoveU")
	port := getEnv("DB_PORT", "ILoveU")

	conSt := fmt.Sprintf("host = %s\nport = %s\nuser = %s\npassword = %s\nDB name = %s",host,port,user,password,name)

	fmt.Println(conSt)

}