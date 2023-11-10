package option

type Option struct {
	Strict     bool
	CPUCore    bool
	CPUCoreNum int
	CPUArch    bool
	DiskAvail  bool
	DiskOfDir  []string
}
