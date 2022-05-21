package protocol

type Protocol struct {
	// CPU
	CPUUtilization float64 `json:"CPUUtilization"`
	CPULoad        float64 `json:"CPULoad"`
	// Memory
	FreeMemory uint64 `json:"FreeMemory"`
	UsedMemory uint64 `json:"UsedMemory"`
	// Net
	NetByteSend uint64 `json:"NetByteSend"`
	NetByteRec  uint64 `json:"NetByteRec"`
}
