package webvars

//Webvars Model
type Webvars struct {
	Host           string
	PoweredBy      string
	WebServer      string
	IsLinux        bool
	IsWindows      bool
	IsPlesk        bool
	IsCpanel       bool
	IsDirectAdmin  bool
	IsWebSitePanel bool
	IsMaestroPanel bool
}

const (
	WEB          = 80
	RDP          = 3389
	SSH          = 22
	MSSQL        = 1433
	MYSQL        = 3306
	WINS         = 42
	CPANEL       = 2082
	DIRECTADMIN  = 2222
	MAESTROPANEL = 9715
	PLESK        = 8443
	PLESK_HA     = 8087 //Hosting Accelerator Control Panel
	WSPANEL      = 9001
	DOCKER       = 2375
	POSTGRESQL   = 5432
	POWERSHELL   = 5985
	VCENTER      = 5988
	VMWARE_UI    = 8333 //VMware Server Management User Interface
	REDIS        = 6379
	HYPERV       = 6600
	HyperVM_HTTP = 8887
	ZIMBRA       = 7047
	TOMCAT       = 8080
	PUPPET       = 8140
	IIS_MANAGER  = 8172
	COULD_FUSION = 8500
	WSUS         = 8530 //Windows Server Update Services
	HADOOP       = 9000
	CASSANDRA    = 9042
	MONGODB      = 27017
	WINRM        = 47001
)

func GetServerVariables(host string) (w Webvars, err error) {

	w = Webvars{}
	w.Host = host
	w.IsLinux = detectPorts(host, SSH)
	w.IsWindows = detectPorts(host, RDP, MSSQL, POWERSHELL, IIS_MANAGER, WINRM)
	w.IsCpanel = detectPorts(host, CPANEL)
	w.IsDirectAdmin = detectPorts(host, DIRECTADMIN)
	w.IsMaestroPanel = detectPorts(host, MAESTROPANEL)
	w.IsPlesk = detectPorts(host, PLESK, PLESK_HA)

	if detectPorts(host, WEB) {
		w.PoweredBy = getHeader(host, "X-Powered-By")
		w.WebServer = getHeader(host, "Server")
	}

	return w, err
}
