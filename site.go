package itswizard_m_siteSetup

import (
	"github.com/itslearninggermany/itswizard_m_objects"
	"html/template"
)

type Site struct {
	Navigation  template.HTML
	SessionUser itswizard_m_objects.SessionUser
	SiteName    string
	Special     interface{}
	URLQuery    string
}
