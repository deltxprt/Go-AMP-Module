package data

// SendGetInstance is the request to get a single instance
type SendGetInstance struct {
	InstanceId string `json:"instanceId"`
}

// ReceiveGetInstance return a single instance
type ReceiveGetInstance struct {
	InstanceID        string `json:"InstanceID"`
	TargetID          string `json:"TargetID"`
	InstanceName      string `json:"InstanceName"`
	FriendlyName      string `json:"FriendlyName"`
	Module            string `json:"Module"`
	ModuleDisplayName string `json:"ModuleDisplayName"`
	InstalledVersion  struct {
		Major         int `json:"Major"`
		Minor         int `json:"Minor"`
		Build         int `json:"Build"`
		Revision      int `json:"Revision"`
		MajorRevision int `json:"MajorRevision"`
		MinorRevision int `json:"MinorRevision"`
	} `json:"InstalledVersion"`
	IsHTTPS               bool          `json:"IsHTTPS"`
	IP                    string        `json:"IP"`
	Port                  int           `json:"Port"`
	Daemon                bool          `json:"Daemon"`
	DaemonAutostart       bool          `json:"DaemonAutostart"`
	ExcludeFromFirewall   bool          `json:"ExcludeFromFirewall"`
	Running               bool          `json:"Running"`
	AppState              int           `json:"AppState"`
	Tags                  []interface{} `json:"Tags"`
	DiskUsageMB           int           `json:"DiskUsageMB"`
	ReleaseStream         int           `json:"ReleaseStream"`
	ManagementMode        int           `json:"ManagementMode"`
	Suspended             bool          `json:"Suspended"`
	IsContainerInstance   bool          `json:"IsContainerInstance"`
	ContainerMemoryMB     int           `json:"ContainerMemoryMB"`
	ContainerMemoryPolicy int           `json:"ContainerMemoryPolicy"`
	ContainerCPUs         float64       `json:"ContainerCPUs"`
	Metrics               struct {
		CPUUsage struct {
			RawValue int    `json:"RawValue"`
			MaxValue int    `json:"MaxValue"`
			Percent  int    `json:"Percent"`
			Units    string `json:"Units"`
			Color    string `json:"Color"`
			Color2   string `json:"Color2"`
			Color3   string `json:"Color3"`
		} `json:"CPU Usage"`
		MemoryUsage struct {
			RawValue int    `json:"RawValue"`
			MaxValue int    `json:"MaxValue"`
			Percent  int    `json:"Percent"`
			Units    string `json:"Units"`
			Color    string `json:"Color"`
			Color3   string `json:"Color3"`
		} `json:"Memory Usage"`
		ActiveUsers struct {
			RawValue int    `json:"RawValue"`
			MaxValue int    `json:"MaxValue"`
			Percent  int    `json:"Percent"`
			Units    string `json:"Units"`
			Color    string `json:"Color"`
			Color3   string `json:"Color3"`
		} `json:"Active Users"`
	} `json:"Metrics"`
	ApplicationEndpoints []struct {
		DisplayName string `json:"DisplayName"`
		Endpoint    string `json:"Endpoint"`
		Uri         string `json:"Uri"`
	} `json:"ApplicationEndpoints"`
	DeploymentArgs struct {
		GenericModuleAppApplicationPort1     string `json:"GenericModule.App.ApplicationPort1"`
		GenericModuleAppApplicationPort2     string `json:"GenericModule.App.ApplicationPort2"`
		GenericModuleAppApplicationPort3     string `json:"GenericModule.App.ApplicationPort3"`
		GenericModuleAppRemoteAdminPort      string `json:"GenericModule.App.RemoteAdminPort"`
		GenericModuleMetaAuthor              string `json:"GenericModule.Meta.Author"`
		GenericModuleMetaConfigRoot          string `json:"GenericModule.Meta.ConfigRoot"`
		GenericModuleMetaConfigManifest      string `json:"GenericModule.Meta.ConfigManifest"`
		GenericModuleMetaDescription         string `json:"GenericModule.Meta.Description"`
		GenericModuleMetaDisplayImageSource  string `json:"GenericModule.Meta.DisplayImageSource"`
		GenericModuleMetaDisplayName         string `json:"GenericModule.Meta.DisplayName"`
		GenericModuleMetaEndpointURIFormat   string `json:"GenericModule.Meta.EndpointURIFormat"`
		GenericModuleMetaOS                  string `json:"GenericModule.Meta.OS"`
		GenericModuleMetaURL                 string `json:"GenericModule.Meta.URL"`
		GenericModuleMetaMinAMPVersion       string `json:"GenericModule.Meta.MinAMPVersion"`
		GenericModuleMetaOriginalSource      string `json:"GenericModule.Meta.OriginalSource"`
		FileManagerPluginSFTPSFTPIPBinding   string `json:"FileManagerPlugin.SFTP.SFTPIPBinding"`
		GenericModuleAppApplicationIPBinding string `json:"GenericModule.App.ApplicationIPBinding"`
		FileManagerPluginSFTPSFTPPortNumber  string `json:"FileManagerPlugin.SFTP.SFTPPortNumber"`
	} `json:"DeploymentArgs"`
	DisplayImageSource string `json:"DisplayImageSource"`
}

// ReceiveGetInstances return the Hosts and under it the Instances as "AvailableInstances"
type ReceiveGetInstances struct {
	Result []struct {
		Id           int    `json:"Id"`
		InstanceId   string `json:"InstanceId"`
		FriendlyName string `json:"FriendlyName"`
		Disabled     bool   `json:"Disabled"`
		IsRemote     bool   `json:"IsRemote"`
		Platform     struct {
			CPUInfo struct {
				Sockets      int    `json:"Sockets"`
				Cores        int    `json:"Cores"`
				Threads      int    `json:"Threads"`
				Vendor       string `json:"Vendor"`
				ModelName    string `json:"ModelName"`
				TotalCores   int    `json:"TotalCores"`
				TotalThreads int    `json:"TotalThreads"`
			} `json:"CPUInfo"`
			InstalledRAMMB int    `json:"InstalledRAMMB"`
			OS             int    `json:"OS"`
			PlatformName   string `json:"PlatformName"`
			SystemType     int    `json:"SystemType"`
			Virtualization int    `json:"Virtualization"`
		} `json:"Platform"`
		Datastores []struct {
			Id           int    `json:"Id"`
			FriendlyName string `json:"FriendlyName"`
		} `json:"Datastores"`
		CreatesInContainers bool   `json:"CreatesInContainers"`
		State               int    `json:"State"`
		StateReason         string `json:"StateReason"`
		CanCreate           bool   `json:"CanCreate"`
		LastUpdated         string `json:"LastUpdated"`
		AvailableInstances  []struct {
			InstanceID       string `json:"InstanceID"`
			TargetID         string `json:"TargetID"`
			InstanceName     string `json:"InstanceName"`
			FriendlyName     string `json:"FriendlyName"`
			Module           string `json:"Module"`
			InstalledVersion struct {
				Major         int `json:"Major"`
				Minor         int `json:"Minor"`
				Build         int `json:"Build"`
				Revision      int `json:"Revision"`
				MajorRevision int `json:"MajorRevision"`
				MinorRevision int `json:"MinorRevision"`
			} `json:"InstalledVersion"`
			IsHTTPS               bool          `json:"IsHTTPS"`
			IP                    string        `json:"IP"`
			Port                  int           `json:"Port"`
			Daemon                bool          `json:"Daemon"`
			DaemonAutostart       bool          `json:"DaemonAutostart"`
			ExcludeFromFirewall   bool          `json:"ExcludeFromFirewall"`
			Running               bool          `json:"Running"`
			AppState              int           `json:"AppState"`
			Tags                  []interface{} `json:"Tags"`
			DiskUsageMB           int           `json:"DiskUsageMB"`
			ReleaseStream         int           `json:"ReleaseStream"`
			ManagementMode        int           `json:"ManagementMode"`
			Suspended             bool          `json:"Suspended"`
			IsContainerInstance   bool          `json:"IsContainerInstance"`
			ContainerMemoryMB     int           `json:"ContainerMemoryMB"`
			ContainerMemoryPolicy int           `json:"ContainerMemoryPolicy"`
			ContainerCPUs         float64       `json:"ContainerCPUs"`
			Metrics               struct {
				CPUUsage struct {
					RawValue int    `json:"RawValue"`
					MaxValue int    `json:"MaxValue"`
					Percent  int    `json:"Percent"`
					Units    string `json:"Units"`
					Color    string `json:"Color"`
					Color2   string `json:"Color2"`
					Color3   string `json:"Color3"`
				} `json:"CPU Usage"`
				MemoryUsage struct {
					RawValue int    `json:"RawValue"`
					MaxValue int    `json:"MaxValue"`
					Percent  int    `json:"Percent"`
					Units    string `json:"Units"`
					Color    string `json:"Color"`
					Color3   string `json:"Color3"`
				} `json:"Memory Usage"`
				ActiveUsers struct {
					RawValue int    `json:"RawValue"`
					MaxValue int    `json:"MaxValue"`
					Percent  int    `json:"Percent"`
					Units    string `json:"Units"`
					Color    string `json:"Color"`
					Color3   string `json:"Color3"`
				} `json:"Active Users"`
			} `json:"Metrics"`
			ApplicationEndpoints []struct {
				DisplayName string `json:"DisplayName"`
				Endpoint    string `json:"Endpoint"`
				Uri         string `json:"Uri"`
			} `json:"ApplicationEndpoints"`
			DeploymentArgs struct {
				FileManagerPluginSFTPSFTPIPBinding      string `json:"FileManagerPlugin.SFTP.SFTPIPBinding"`
				FileManagerPluginSFTPSFTPPortNumber     string `json:"FileManagerPlugin.SFTP.SFTPPortNumber"`
				ADSModuleNetworkMetricsServerPort       string `json:"ADSModule.Network.MetricsServerPort,omitempty"`
				ADSModuleDefaultsDefaultAuthServerURL   string `json:"ADSModule.Defaults.DefaultAuthServerURL,omitempty"`
				CoreLoginUseAuthServer                  string `json:"Core.Login.UseAuthServer,omitempty"`
				CoreWebserverUsingReverseProxy          string `json:"Core.Webserver.UsingReverseProxy,omitempty"`
				CoreLoginAuthServerURL                  string `json:"Core.Login.AuthServerURL,omitempty"`
				CoreSecurityEnablePassthruAuth          string `json:"Core.Security.EnablePassthruAuth,omitempty"`
				ADSModuleNetworkDefaultIPBinding        string `json:"ADSModule.Network.DefaultIPBinding,omitempty"`
				MinecraftModuleMinecraftServerIPBinding string `json:"MinecraftModule.Minecraft.ServerIPBinding,omitempty"`
				MinecraftModuleMinecraftPortNumber      string `json:"MinecraftModule.Minecraft.PortNumber,omitempty"`
				GenericModuleAppApplicationPort1        string `json:"GenericModule.App.ApplicationPort1,omitempty"`
				GenericModuleAppApplicationPort2        string `json:"GenericModule.App.ApplicationPort2,omitempty"`
				GenericModuleAppApplicationPort3        string `json:"GenericModule.App.ApplicationPort3,omitempty"`
				GenericModuleAppRemoteAdminPort         string `json:"GenericModule.App.RemoteAdminPort,omitempty"`
				GenericModuleMetaAuthor                 string `json:"GenericModule.Meta.Author,omitempty"`
				GenericModuleMetaConfigRoot             string `json:"GenericModule.Meta.ConfigRoot,omitempty"`
				GenericModuleMetaConfigManifest         string `json:"GenericModule.Meta.ConfigManifest,omitempty"`
				GenericModuleMetaDescription            string `json:"GenericModule.Meta.Description,omitempty"`
				GenericModuleMetaDisplayImageSource     string `json:"GenericModule.Meta.DisplayImageSource,omitempty"`
				GenericModuleMetaDisplayName            string `json:"GenericModule.Meta.DisplayName,omitempty"`
				GenericModuleMetaEndpointURIFormat      string `json:"GenericModule.Meta.EndpointURIFormat,omitempty"`
				GenericModuleMetaOS                     string `json:"GenericModule.Meta.OS,omitempty"`
				GenericModuleMetaURL                    string `json:"GenericModule.Meta.URL,omitempty"`
				GenericModuleMetaMinAMPVersion          string `json:"GenericModule.Meta.MinAMPVersion,omitempty"`
				GenericModuleMetaOriginalSource         string `json:"GenericModule.Meta.OriginalSource,omitempty"`
				GenericModuleAppApplicationIPBinding    string `json:"GenericModule.App.ApplicationIPBinding,omitempty"`
			} `json:"DeploymentArgs"`
			DisplayImageSource string `json:"DisplayImageSource,omitempty"`
			ModuleDisplayName  string `json:"ModuleDisplayName,omitempty"`
		} `json:"AvailableInstances"`
		AvailableIPs []string      `json:"AvailableIPs"`
		Tags         []interface{} `json:"Tags"`
		TagsList     string        `json:"TagsList"`
		URL          string        `json:"URL"`
	} `json:"result"`
}

// SendStartStopInstance return TaskActionResult
// It's the same for Start/Stop/Restart
type SendStartStopInstance struct {
	InstanceName string `json:"InstanceName"`
}
