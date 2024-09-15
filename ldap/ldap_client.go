package ldap

import (
	"crypto/tls"
	"github.com/go-ldap/ldap/v3"
	"log"
)

func ConnectToLDAP(ldapURL, username, password string) (*ldap.Conn, error) {
	l, err := ldap.DialURL(ldapURL, ldap.DialWithTLSConfig(&tls.Config{InsecureSkipVerify: true}))
	if err != nil {
		return nil, err
	}

	err = l.Bind(username, password)
	if err != nil {
		return nil, err
	}

	log.Println("LDAP Bind Successful")
	return l, nil
}

func SearchLDAP(conn *ldap.Conn, baseDN, filter string) (*ldap.SearchResult, error) {
	searchRequest := ldap.NewSearchRequest(
		baseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,
		[]string{"dn", "cn"}, // Attributes to return
		nil,
	)

	sr, err := conn.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	for _, entry := range sr.Entries {
		log.Printf("DN: %s, CN: %s\n", entry.DN, entry.GetAttributeValue("cn"))
	}

	return sr, nil
}
