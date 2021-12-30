package main

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
)

func main() {
	fmt.Println("LDAP")
	fmt.Println()

	url := "ldap://localhost:1389"
	fmt.Printf("Connecting to %s...\n", url)
	conn, err := ldap.DialURL(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer conn.Close()

	bindDN := "cn=admin,dc=com"
	fmt.Printf("Binding to DN: %s...\n", bindDN)
	conn.Bind(bindDN, "MmsIsAw3s0me")
	search(conn, bindDN, "(objectClass=*)") // 1 result

	baseDN := "dc=babypearfoo,dc=com"
	search(conn, baseDN, "(objectClass=*)") // 63 results
	search(conn, baseDN, "(uid=*comma*)") // 4 results
	search(conn, baseDN, "(uid=comma*)") // 3 results
	search(conn, baseDN, "(uid=commas,*)") // 1 result
	search(conn, baseDN, "(uid=commas\\,*)") // Error
	search(conn, baseDN, "(uid=commas\\2C*)") // 1 result
	search(conn, baseDN, "(uid=commas\\2c*)") // 1 result
	search(conn, baseDN, "(uid=commas,and,asterisk*)") // 1 result, though it would also match any uid starting with: commas,and,asterisk
	search(conn, baseDN, "(uid=commas,and,asterisk\\2A)") // 1 result, this time specifically for uid ending with: *
	search(conn, baseDN, "(uid=commas\\2C*\\2A)") // 1 result
	search(conn, baseDN, "(uid=commas\\2Cand\\2Casterisk\\2A)") // 1 result
}

func search(conn *ldap.Conn, baseDN string, filter string) {
	fmt.Println("------------------------------------------------")
	fmt.Printf("searching baseDN '%s' using filter '%s'...\n", baseDN, filter)
	searchReq := ldap.NewSearchRequest(baseDN, ldap.ScopeWholeSubtree, 0, 0, 0, false, filter, []string{}, nil)
	result, err := conn.Search(searchReq)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println("Got", len(result.Entries), "search results")
	for i := 0; i < len(result.Entries); i++ {
		fmt.Println()
		result.Entries[i].Print()
	}
	fmt.Println()
}

