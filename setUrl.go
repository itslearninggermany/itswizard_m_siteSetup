package itswizard_m_siteSetup

import (
	"github.com/itslearninggermany/itswizard_m_jwt"
	"net/url"
)

func (p *SiteSetup) SetTargetPath(target string) *SiteSetup {
	p.u.Path = target
	return p
}

func (p *SiteSetup) AddOpaque(Opaque string) *SiteSetup {
	p.u.Opaque = Opaque
	return p
}

func (p *SiteSetup) AddUserToUrl(username, password string) *SiteSetup {
	if password == "" {
		p.u.User = url.User(username)
	} else {
		p.u.User = url.UserPassword(username, password)
	}
	return p
}

func (p *SiteSetup) AddRawPath(rawPath string) *SiteSetup {
	p.u.RawPath = rawPath
	return p
}

func (p *SiteSetup) SetForceQuery(forceQuery bool) *SiteSetup {
	p.u.ForceQuery = forceQuery
	return p
}

func (p *SiteSetup) AddRawQuery(rawQuery string) *SiteSetup {
	p.u.RawQuery = rawQuery
	return p
}

func (p *SiteSetup) AddFragment(fragment string) *SiteSetup {
	p.u.Fragment = fragment
	return p
}

func (p *SiteSetup) AddValue(value, key string) *SiteSetup {
	p.q.Add(value, key)
	return p
}

func (p *SiteSetup) GetURL() (url string, err error) {
	_, err = p.MakeSite()
	if err != nil {
		return "", err
	}
	return p.u.String(), nil
}

func (p *SiteSetup) GetJWTToken() (string, error) {
	return itswizard_m_jwt.ReAuthentificate(p.r, p.dbWebserver, p.dbClient)
}

func (p *SiteSetup) GetQuery() (query string, err error) {
	_, err = p.MakeSite()
	if err != nil {
		return "", err
	}
	return p.q.Encode(), nil
}
