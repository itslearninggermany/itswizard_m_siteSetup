package itswizard_m_siteSetup

import "html/template"

func createAdminNavi(query string) template.HTML {

	erg := `<div class="hero-head">
    <a role="button" class="navbar-burger" data-target="navMenu" aria-label="menu" aria-expanded="false">
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
    </a>


    <div class="navbar-menu" id="navMenu">
        <!-- Menü -->
        <div class="navbar-end">
            <a class="navbar-item" href="/client/startpage` + query + `">
                Start
            </a>

			<div class="navbar-item has-dropdown is-hoverable">
                <a class="navbar-link">
                    Manage Clients
                </a>

                <div class="navbar-dropdown">
                    <a class="navbar-item" href="/client/admin/newcustomer` + query + `">
                        New Customer
                    </a>
                    <a class="navbar-item" href="/client/admin/myCustomer` + query + `">
                        My Customer
                    </a>
                    <a class="navbar-item" href="/client/admin/MatchCustomer` + query + `">
                        Match Customer
                    </a>
 					<a class="navbar-item" href="/client/admin/integrationSupporter` + query + `">
                        New Integration Supporter
                    </a>
                    <hr class="navbar-divider">
                     <a class="navbar-item" href="https://bw.itswizard.de/client/admin/showInstitution` + query + `">
                        Client-Administration (Superadmin)
                    </a>
                     <a class="navbar-item" href="/client/admin/newitswizardadmin` + query + `">
                        New Itswizard Administrator (Superadmin)
                     </a>
                     <a class="navbar-item" href="/client/admin/itswizardusers` + query + `">
                        Show all Itswizard Users (Superadmin)
                     </a>
                </div>
            </div>

			 <div class="navbar-item has-dropdown is-hoverable">
                <a class="navbar-link">
            		ServerSetup (Superadmin)
                </a>

                <div class="navbar-dropdown">
                    <a class="navbar-item" href="/admin/editSmtp` + query + `">
                        EmailServer Setup (Superadmin)
                    </a>
					<a class="navbar-item" href="/admin/testemail` + query + `">
                        Send TestMail (Superadmin) 
                    </a>
					<a class="navbar-item" href="/admin/errorlog` + query + `">
                        System Error Log (Superamdin)
                    </a>
                </div>
			</div>

            <div class="navbar-item has-dropdown is-hoverable">
                <a class="navbar-link">
                    Setup 
                </a>

                <div class="navbar-dropdown">
                    <a class="navbar-item" href="/client/options` + query + `">
                        Personal Data
                    </a>
                    <a class="navbar-item" href="/client/changepw` + query + `">
                        Change Password
                    </a>
                    <hr class="navbar-divider">
                    <a class="navbar-item" href="/client/logOut` + query + `">
                        Log Out
                    </a>
                </div>
            </div>
        </div>
    </div>



<script>
    document.addEventListener('DOMContentLoaded', () => {

        // Get all "navbar-burger" elements
        const $navbarBurgers = Array.prototype.slice.call(document.querySelectorAll('.navbar-burger'), 0);

        // Check if there are any navbar burgers
        if ($navbarBurgers.length > 0) {

            // Add a click event on each of them
            $navbarBurgers.forEach( el => {
                el.addEventListener('click', () => {

                    // Get the target from the "data-target" attribute
                    const target = el.dataset.target;
                    const $target = document.getElementById(target);

                    // Toggle the "is-active" class on both the "navbar-burger" and the "navbar-menu"
                    el.classList.toggle('is-active');
                    $target.classList.toggle('is-active');

                });
            });
        }

    });
</script>`

	return template.HTML(erg)
}
