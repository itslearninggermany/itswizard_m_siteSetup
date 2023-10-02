package itswizard_m_siteSetup


import (
	"github.com/jinzhu/gorm"
	"net/http"
	"net/url"
)

type SiteSetup struct {
	u           *url.URL
	q           url.Values
	site        Site
	dbWebserver *gorm.DB
	dbClient    *gorm.DB
	r           *http.Request
	newAuth     bool
	username    string
}

/*
Wichtige Information
if username == "" ==> Dann ist es keine Erneuerung
*/
func InitialSite(scheme, host string, username string, dbWebserver, dbClient *gorm.DB, r *http.Request) *SiteSetup {
	s := new(SiteSetup)
	s.u = new(url.URL)
	s.q = url.Values{}
	s.r = r

	s.u.Scheme = scheme
	s.u.Host = host
	if username != "" {
		s.newAuth = true
		s.username = username
	} else {
		s.newAuth = false
	}
	s.dbClient = dbClient
	s.dbWebserver = dbWebserver
	return s
}

