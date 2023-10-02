package itswizard_m_siteSetup

import (
	"github.com/itslearninggermany/itswizard_m_jwt"
	"github.com/itslearninggermany/itswizard_m_objects"
)

func (p *SiteSetup) MakeSite() (site Site, err error) {
	if p.newAuth {
		authToken, user, err := itswizard_m_jwt.CreateToken(p.r, p.username, p.dbClient, p.dbWebserver)
		if err != nil {
			return site, err
		}
		p.q.Add("key", authToken)
		p.site.SessionUser = user
	} else {
		p.site.SessionUser, err = itswizard_m_jwt.GetUser(p.r, p.dbWebserver)
		if err != nil {
			return p.site, err
		}
		authtoken, err := itswizard_m_jwt.ReAuthentificate(p.r, p.dbWebserver, p.dbClient)
		if err != nil {
			return site, err
		}
		p.q.Add("key", authtoken)
	}

	p.u.RawQuery = p.q.Encode()
	p.site.URLQuery = "?" + p.u.RawQuery

	if p.site.SessionUser.Admin {
		p.site.Navigation = createAdminNavi(p.site.URLQuery)
	} else {
		p.site.Navigation = createClientNavi(p.site.URLQuery, p.site.SessionUser)
	}
	return p.site, err
}

func (p *SiteSetup) AddSiteName(sitename string) *SiteSetup {
	p.site.SiteName = sitename
	return p
}

func (p *SiteSetup) AddDataInSite(data interface{}) *SiteSetup {
	p.site.Special = data
	return p
}

func (p *SiteSetup) GetUser() (user itswizard_m_objects.SessionUser, err error) {
	if p.newAuth {
		authToken, user, err := itswizard_m_jwt.CreateToken(p.r, p.username, p.dbClient, p.dbWebserver)
		if err != nil {
			return user, err
		}
		p.q.Add("key", authToken)
		p.site.SessionUser = user
	} else {
		p.site.SessionUser, err = itswizard_m_jwt.GetUser(p.r, p.dbWebserver)
	}
	return p.site.SessionUser, err

}
