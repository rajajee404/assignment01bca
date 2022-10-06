package main

import (
	//	"bufio"
	"crypto/sha256"
	"fmt"

	//	"os"
	"strconv"
	"strings"
)

type Block struct {
	nonce       int
	transaction string
	phash       string
	chash       string
}

func Newblock(n int, t string) *Block {
	s := new(Block)
	s.nonce = n
	s.transaction = t
	return s
}

type BlockList struct {
	list []*Block
}

func (ls *BlockList) AddBlock(n int, t string) *Block {
	st := Newblock(n, t)

	if VerifyChain(ls) {
		ls.list = append(ls.list, st)
		CalHash(ls)
		fmt.Println("Block Added")
		return st
	} else {
		return nil
	}
}
func ListBlocks(stud *BlockList) {
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

	for i := 0; i < len(stud.list); i++ {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Println("nonce    ", stud.list[i].nonce)
		fmt.Println("Transaction   ", stud.list[i].transaction)
		fmt.Println("previous hash   ", stud.list[i].phash)
		fmt.Println("current hash   ", stud.list[i].chash)

	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")

}
func (s *Block) GetString() string {

	var r = ""
	r += strconv.Itoa(s.nonce)
	r += s.transaction + s.phash

	return r
}

func CalHash(stud *BlockList) {

	for i := 0; i < len(stud.list); i++ {
		sum := sha256.Sum256([]byte(stud.list[i].GetString()))
		stud.list[i].chash = fmt.Sprintf("%x", sum)
		if i < len(stud.list)-1 {
			stud.list[i+1].phash = fmt.Sprintf("%x", sum)

		}
	}

}
func VerifyChain(stud *BlockList) bool {
	var st = ""
	for i := 0; i < len(stud.list); i++ {
		sum := sha256.Sum256([]byte(stud.list[i].GetString()))
		st = fmt.Sprintf("%x", sum)

		if st != stud.list[i].chash {
			fmt.Printf("Block is tempered, Block no %d\n", i)
			return false

		}

	}
	fmt.Println("Blocks are ok! ")
	return true

}
func ChangeBlock(stud *BlockList, n int, t string) { // n = nonce of the block, t= new transaction or change to the block

	for i := 0; i < len(stud.list); i++ {
		if n == stud.list[i].nonce {

			stud.list[i].transaction = t
			fmt.Println("change done ")
			return
		}
	}

	fmt.Println("block not found!")

}

// func main() {
// //	reader := bufio.NewReader(os.Stdin)
// 	blockchain := new(blocklist)
// 	var non = 200
// 	blockchain.addblock(non, "hello")
// 	non++
// 	blockchain.addblock(non, "hhi")
// 	non++
// 	blockchain.addblock(non, "asdwr")
// 	non++
// 	blockchain.addblock(non, "rgherwv")
// 	non++
// 	blockchain.addblock(non, "xyzc")
// 	non++

// 	listblocks(blockchain)

// 	fmt.Println("changing a block with nonce 202")

// 	changeblock(blockchain,202, "alice to me 9999" )

// 	listblocks(blockchain)

// 	fmt.Println("Adding the new block ")

// 	blockchain.addblock(non, "new transaction")
// 	non++

// 	fmt.Println("verifying the chain ")
// 	verifychain(blockchain)

// 	fmt.Println("Recalculating the hashes")

// 	calhash(blockchain)

// 	fmt.Println("verifying the chain ")
// 	verifychain(blockchain)

// 	fmt.Println("")
// 	fmt.Println("Adding the new block ")

// 	blockchain.addblock(non, "new transaction")
// 	non++

// 	/*
// 	   	b := 1

// 	   	for b ==1 {
// 	   	    fmt.Println("")
// 	   	    fmt.Println("")
// 	   		fmt.Println("1) Do you want to add new transaction? ")
// 	   		fmt.Println("2) Do you want to list blocks? ")
// 	   		fmt.Println("3) Do you want to change a block? ")
// 	   		fmt.Println("4) Do you want to verify chain? ")
// 	   		fmt.Println("5) Do you want to calculate new hash? ")
// 	   		fmt.Println("6) Exit ")
// 	   		fmt.Println("Enter option number  ")

// 	   		var op int
// 	   		fmt.Scanln(&op)
// 	   		fmt.Scanln(&op)

// 	   		if op==1{
// 				fmt.Println("Enter the transaction ")
// 				text, _ := reader.ReadString('\n')
// 		   		// convert CRLF to LF
// 		       		text = strings.Replace(text, "\n", "", -1)
// 				blockchain.addblock(non, text)
// 				non++

// 	   		} else if op ==2{

// 	   		    listblocks(blockchain)

// 	   		} else if op==3{
// 	   		    var n int
// 	   		    fmt.Scanf("Enter the block nonce %d",&n)
// 	   		    fmt.Println("Enter the transaction ")
// 				text, _ := reader.ReadString('\n')
// 		   		// convert CRLF to LF
// 		       		text = strings.Replace(text, "\n", "", -1)
// 				changeblock(blockchain,n,text)

// 	   		} else if op==4 {

// 	   		    verifychain(blockchain)
// 	   		} else if op==5 {

// 	   		    calhash(blockchain)
// 	   		} else {

// 	   		    b=0

// 	   		}

// 	   	}
// */

// }
