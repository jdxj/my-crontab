package model

type IpClassify struct {
	Id         int     `json:"id"`
	Name       string  `json:"name"`
	Internet   bool    `json:"internet"`
	Ipv6Size   int     `json:"ipv6Size"`
	SetupPrice float64 `json:"setupPrice"`
	Descr      string  `json:"descr"`
}

type Group struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type Prices struct {
	ANNUALLY     float64 `json:"ANNUALLY"`
	QUARTERLY    float64 `json:"QUARTERLY,omitempty"`
	MONTHLY      float64 `json:"MONTHLY,omitempty"`
	SEMIANNUALLY float64 `json:"SEMI_ANNUALLY,omitempty"`
}

type Product struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	CpuCores int    `json:"cpuCores"`
	// MB
	Memory int `json:"memory"`
	// GB
	Disk int `json:"disk"`
	// MB/s
	DiskRead int `json:"diskRead"`
	// MB/s
	DiskWrite int `json:"diskWrite"`
	Mbps      int `json:"mbps"`
	// GB
	Traffic        int        `json:"traffic"`
	IpClassify     IpClassify `json:"ipClassify"`
	Hots           int        `json:"hots"`
	CpuFrequency   string     `json:"cpuFrequency"`
	Ssd            bool       `json:"ssd"`
	Raid           string     `json:"raid"`
	Ipv6           bool       `json:"ipv6"`
	SpecialPrice   bool       `json:"specialPrice"`
	VerifiedIdCard bool       `json:"verifiedIdCard"`
	Group          Group      `json:"group"`
	// important
	SoldOut bool   `json:"soldOut"`
	Prices  Prices `json:"prices"`
}

type NotifyInput struct {
	Text string
}
