package main

import (
	"fmt"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// Note: my purpose for this is to make collision free strings via hashing

func main() {
	byte_data_to_hash := []byte( `Example string to hash` )
	
	fmt.Print( "MD5: " )
	md5_instance := md5.New()
	md5_instance.Write( byte_data_to_hash )
	fmt.Println( hex.EncodeToString( md5_instance.Sum( nil ) ) )
	
	fmt.Print( "SHA1: " )
	sha1_instance := sha1.New()
	sha1_instance.Write( byte_data_to_hash )
	fmt.Println( hex.EncodeToString( sha1_instance.Sum( nil ) ) )
	
	fmt.Print( "SHA256: " )
	sha256_instance := sha256.New()
	sha256_instance.Write( byte_data_to_hash )
	fmt.Println( hex.EncodeToString( sha256_instance.Sum( nil ) ) )
}