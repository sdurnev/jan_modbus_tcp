package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"github.com/goburrow/modbus"
	"math"
	"time"
)

//type a map.s.int32

/*
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!! VERSION !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
*/
const version = "0.01.3"

/*
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!! VERSION !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
*/

func main() {
	paramName := [...]string{
		"19000_G_ULN[0]",
		"19002_G_ULN[1]",
		"19004_G_ULN[2]",
		"19006_G_ULL[0]",
		"19008_G_ULL[1]",
		"19010_G_ULL[2]",
		"19012_G_ILN[0]",
		"19014_G_ILN[1]",
		"19016_G_ILN[2]",
		"19018_G_I_SUM3",
		"19020_G_PLN[0]",
		"19022_G_PLN[1]",
		"19024_G_PLN[2]",
		"19026_G_P_SUM3",
		"19028_G_SLN[0]",
		"19030_G_SLN[1]",
		"19032_G_SLN[2]",
		"19034_G_S_SUM3",
		"19036_G_QLN[0]",
		"19038_G_QLN[1]",
		"19040_G_QLN[2]",
		"19042_G_Q_SUM3",
		"19044_G_COS_PHI[0]",
		"19046_G_COS_PHI[1]",
		"19048_G_COS_PHI[2]",
		"19050_G_FREQ",
		"19052_G_PHASE_SEQ",
		"19054_G_WH[0]",
		"19056_G_WH[1]",
		"19058_G_WH[2]",
		"19060_G_WH_SUML13",
		"19062_G_WH_V[0]",
		"19064_G_WH_V[1]",
		"19066_G_WH_V[2]",
		"19068_G_WH_V_HT_SUML13",
		"19070_G_WH_Z[0]",
		"19072_G_WH_Z[1]",
		"19074_G_WH_Z[2]",
		"19076_G_WH_Z_SUML13",
		"19078_G_WH_S[0]",
		"19080_G_WH_S[1]",
		"19082_G_WH_S[2]",
		"19084_G_WH_S_SUML13",
		"19086_G_QH[0]",
		"19088_G_QH[1]",
		"19090_G_QH[2]",
		"19092_G_QH_SUML13",
		"19094_G_IQH[0]",
		"19096_G_IQH[1]",
		"19098_G_IQH[2]",
		"19100_G_IQH_SUML13",
		"19102_G_CQH[0]",
		"19104_G_CQH[1]",
		"19106_G_CQH[2]",
		"19108_G_CQH_SUML13",
		"19110_G_THD_ULN[0]",
		"19112_G_THD_ULN[1]",
		"19114_G_THD_ULN[2]",
		"19116_G_THD_ILN[0]",
		"19118_G_THD_ILN[1]",
		"19120_G_THD_ILN[2]",
	}
	var data []float32

	addressIP := flag.String("ip", "localhost", "a string")
	tcpPort := flag.String("port", "502", "a string")
	slaveID := flag.Int("id", 1, "an int")
	regQuantity := flag.Uint("q", 61, "an uint")
	flag.Parse()
	serverParam := fmt.Sprint(*addressIP, ":", *tcpPort)
	s := byte(*slaveID)

	//	fmt.Println(serverParam)

	handler := modbus.NewTCPClientHandler(serverParam)
	handler.SlaveId = s
	handler.Timeout = 2 * time.Second
	// Connect manually so that multiple requests are handled in one session
	err := handler.Connect()
	defer handler.Close()
	client := modbus.NewClient(handler)

	results, err := client.ReadInputRegisters(19000, uint16(*regQuantity)*2)
	if err != nil {
		fmt.Printf("{\"status\":\"error\", \"error\":\"%s\"}", err)
		//fmt.Printf("%s\n", err)
	}

	//fmt.Println(len(results))
	//fmt.Println(results)
	i := 0
	for i < len(results) {
		a := Float32frombytes(results[i : i+4])
		if math.IsNaN(float64(a)) {
			data = append(data, 0)
		} else {
			data = append(data, a)
		}
		i += 4
	}

	for l := 0; l < len(data); l++ {
		if l == 0 {
			fmt.Printf("{ \"%s\": ", paramName[l])
		} else {
			fmt.Printf(", \"%s\": ", paramName[l])
		}
		fmt.Print(data[l])
	}
	if len(results) != 0 {
		fmt.Printf(", \"version\": \"%s\"}", version)
	}
}

func Float32frombytes(bytes []byte) float32 {
	bits := binary.BigEndian.Uint32(bytes)
	float := math.Float32frombits(bits)
	return float
}
