package functions

import (
	"database/sql"
	"fmt"
	"os"
	"os/user"
	"strings"
	"time"
)

func formatNullableFloat(f *float64) string {
	if f == nil {
		return ""
	}
	return fmt.Sprintf("%f", *f)
}

func writeToFile(filename, data string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to open file ", filename, err)
	}
	defer f.Close()

	if _, err := f.WriteString(data + "\n"); err != nil {
		fmt.Println("Failed to write to file ", filename, err)
	}
}

func generate_csv(connectionString string) () {
	// connect database
	conn, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		fmt.Println("Error creating connection pool: ", err.Error())
	}
	defer conn.Close()
	err = conn.Ping()
	if err != nil {
		fmt.Println("Error establishing connection: ", err.Error())
	}
	query := `use AgentThaiBev 
			SELECT top 10 * FROM DOCINFO DC
			INNER JOIN DOCTYPE DT ON DC.DI_DT = DT.DT_KEY
			LEFT  JOIN BANKSTATEMENT BS ON BS.BSTM_DI = DI_KEY
			INNER JOIN CASHBOOK CB ON CB.CASHB_DI = DC.DI_KEY
			INNER JOIN TRANPAYH TH ON TH.TPH_DI = DC.DI_KEY
			INNER JOIN TRANPAYD TD ON TD.TPD_TPH = TH.TPH_KEY
			INNER JOIN ARPAYMENT ARP ON ARP.ARP_DI = DC.DI_KEY
			INNER JOIN TRANSTKJ TJ ON TJ.TRJ_DI = DC.DI_KEY
			INNER JOIN TRANSTKH TRH ON TRH.TRH_DI = DC.DI_KEY
			INNER JOIN TRANSTKD TRD ON TRD.TRD_TRH = TRH.TRH_KEY
			INNER JOIN SKUMOVE SK ON SK.SKM_DI = DC.DI_KEY
			INNER JOIN SLDETAIL SL ON SL.SLD_DI = DC.DI_KEY
			INNER JOIN VATTABLE VT ON VT.VAT_DI = DC.DI_KEY
			INNER JOIN ARDETAIL ARD ON ARD.ARD_DI = DC.DI_KEY
			WHERE(DC.DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-31') AND (DT.DT_THAIDESC LIKE 'ใบขายสด%' OR DT.DT_THAIDESC LIKE 'สินค้ารางวัล')
			`
		dcName, err := os.Hostname()
		if err != nil {
				fmt.Println("can't fetch host name:", err)
				return 
		}
		currentUser, err := user.Current()
		if err != nil {
			fmt.Println("can't fetch current user:", err)
			return 
		}
		epoch := time.Now().Unix()
		// fileName := fmt.Sprintf("%s_%s_%d.csv", dcName, currentUser.Username, epoch)
		fileLogName := fmt.Sprintf("%s_%s_%d", dcName, currentUser.Username, epoch)
		// newFilename := strings.Replace(fileName, `\`, `_`, -1)
		newFileLogname := strings.Replace(fileLogName, `\`, `_`, -1)
		
		fmt.Println("Start Query")
		writeToFile("log_"+newFileLogname+".txt", "Start Query")
		rows, err := conn.Query(query)
		if err != nil {
			fmt.Println(err)
			writeToFile("log_"+newFileLogname+"_error.txt", "Failed to query: "+err.Error())
			return /* "", "", err */
		}
		fmt.Println("Finish Query")
		writeToFile("log_"+newFileLogname+".txt", "Finish Query")
		defer rows.Close()
}
