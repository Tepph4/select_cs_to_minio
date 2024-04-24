package main

import (
	"context"
	"crypto/aes"
	"database/sql"

	// "crypto/cipher"
	"encoding/csv"
	"encoding/hex"
	"fmt"

	"log"
	"os"
	_"os/user"
	"strconv"
	_"strings"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {

	user := "admin"
	password := "newadmin"
	server := "LAPTOP-IIKGOJND"
	database := "AgentThaiBev"

	// สร้าง connection string
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;", server, user, password, database)

	// เชื่อมต่อฐานข้อมูล
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		fmt.Println("Error Can't Connect database ", err.Error())
	}

	queryScript := generateQuery()
	for tables, querysql := range queryScript {
		recordPerFile := 1000
		if tables == "TRANSTKD"{
			recordPerFile = 500
		}
		fmt.Printf("Saving Data %s to CSV...\n", tables)
		err = WriteCSV(db, recordPerFile, tables, querysql)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer db.Close()
	// cipher key
	// key := "thisis32bitlongpassphraseimusing"

	// // plaintext
	// pt := "This is a secret"

	// c := EncryptAES([]byte(key), pt)

	// // ciphertext
	// fmt.Println(c)

}

func EncryptAES(key []byte, plaintext string) string {

	c, err := aes.NewCipher(key)
	CheckError(err)

	out := make([]byte, len(plaintext))

	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}

func PKCS7Unpadding(plaintext []byte) ([]byte, error) {
	length := len(plaintext)
	padding := int(plaintext[length-1])
	if padding > length || padding > aes.BlockSize {
		return nil, fmt.Errorf("invalid padding size")
	}

	return plaintext[:length-padding], nil
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
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

func upload_to_minio(endpoint, accessKey, secretKey, bucketName, filename string, filelogname string, trdNewKey int64, skmNewKey int64) {

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: true,
	})
	if err != nil {
		fmt.Println("Failed to initialize Minio client:", err)
		writeToFile("log_"+filelogname+"_error.txt", "Failed to initialize Minio client:"+err.Error())
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		writeToFile("log_"+filelogname+"_error.txt", "Failed to open file:"+err.Error())
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Failed to get file info:", err)
		writeToFile("log_"+filelogname+"_error.txt", "Failed to get file info:"+err.Error())
	}

	_, err = minioClient.PutObject(context.Background(), bucketName, "pending_inventory_data/"+filename, file, fileInfo.Size(), minio.PutObjectOptions{})
	if err != nil {
		fmt.Println("Failed to upload file to Minio:", err)
		writeToFile("log_"+filelogname+"_error.txt", "Failed to upload file to Minio:"+err.Error())
	}
	writeToFile("log_"+filelogname+".txt", fmt.Sprintf("Successfully uploaded "+filename))

	logFile, err := os.Open("log_" + filelogname + ".txt")
	if err != nil {
		fmt.Println("Failed to open log file: ", err)
		writeToFile("log_"+filelogname+"_error.txt", "Failed to open log file:"+err.Error())
	}
	defer logFile.Close()

	logFileInfo, err := logFile.Stat()
	if err != nil {
		fmt.Println("Failed to get log file info: ", err)
		writeToFile("log_"+filelogname+"_error.txt", "Failed to get log file info:"+err.Error())
	}

	_, err = minioClient.PutObject(context.Background(), bucketName, "log_inventory/sucess_bplus_log/"+"log_"+filelogname+".txt", logFile, logFileInfo.Size(), minio.PutObjectOptions{})
	if err != nil {
		fmt.Println("Failed to upload log to Minio: ", err)
		writeToFile("log_"+filelogname+"_error.txt", "Failed to upload log to Minio:"+err.Error())
	}

	_, err = minioClient.PutObject(context.Background(), bucketName, "pending_inventory_data/"+filename, file, fileInfo.Size(), minio.PutObjectOptions{})
	if err != nil {
		fmt.Println("Failed to upload file to Minio:", err)
		writeToFile("log_"+filelogname+"_error.txt", "Failed to upload file to Minio:"+err.Error())
	}

	writeToFile("log_"+filelogname+".txt", fmt.Sprintf("Successfully uploaded "+filename))
	os.Remove("log" + filelogname + ".txt")
	os.Remove("log" + filename + ".txt")
	errorlogFile, err := os.Open("log_" + filelogname + "_error.txt")
	if err != nil {
		fmt.Println("Successfully uploaded data to Minio")
	}
	if err == nil {
		defer logFile.Close()

		errorlogFileInfo, err := logFile.Stat()
		if err != nil {
			fmt.Println("Error", err)
		}
		_, err = minioClient.PutObject(context.Background(), bucketName, "log_inventory/failed_bplus_log/"+"log_"+filename+".txt", errorlogFile, errorlogFileInfo.Size(), minio.PutObjectOptions{})
		if err != nil {
			fmt.Println("Failed to upload log to Minio: ", err)
			writeToFile("log_"+filelogname+"_error.txt", "Failed to upload log to Minio:"+err.Error())
		}
		fmt.Println("Successfully uploaded log file to Minio")
		os.Remove("log_" + filelogname + "_error.txt")
	}
	file, err = os.Open("TrdAndSkmData.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	records[1][0] = strconv.Itoa(int(trdNewKey))
	records[1][1] = strconv.Itoa(int(skmNewKey))

	file, err = os.Create("TrdAndSkmData.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	for _, record := range records {
		err := writer.Write(record)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
}

func GetLastTrdAndSkmKey() (int64, int64, error) {
	fmt.Println("Start")
	TrdAndSkmData, err := os.Open("TrdAndSkmData.csv")
	if err != nil {
		return 0, 0, err
	}
	defer TrdAndSkmData.Close()

	reader := csv.NewReader(TrdAndSkmData)

	_, err = reader.Read()
	if err != nil {
		return 0, 0, err
	}
	record, err := reader.Read()
	if err != nil {
		return 0, 0, err
	}
	trdKeyMin, err := strconv.ParseInt(record[0], 10, 64)
	if err != nil {
		return 0, 0, err
	}

	skmKeyMin, err := strconv.ParseInt(record[1], 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return trdKeyMin, skmKeyMin, nil
}


func WriteCSV(db *sql.DB, batchSize int, table string, querysql []string) error {

	totalRecord, err := getTotalOfRecords(db, querysql[1])
	if err != nil {
		return err
	}
	if totalRecord == 0 {
		log.Fatal("No records found in the table")
	}

	var data [][]string
	numFile := 1
	for offset := 0; offset < totalRecord; offset += batchSize {

		dirName := fmt.Sprintf("CSV-%s", table) 
		_, err := os.Stat(dirName)
		if err != nil && os.IsNotExist(err) {
			err := os.Mkdir(dirName, 0o755)
			if err != nil {
				return fmt.Errorf("สร้างโฟลเดอร์ %s ล้มเหลว: %v", dirName, err)
			}
		}

		fileName := fmt.Sprintf("%s/%s-data-%d.csv",dirName, table, numFile)

		// Create Log File
		// epoch := time.Now().Unix()
		// dcName, err := os.Hostname()
		// if err != nil {
		// 	return fmt.Errorf("can't fetch current hostname: %v", err)
		// }
		// currentUser, err := user.Current()
		// if err != nil {
		// 	return fmt.Errorf("can't fetch current user: %v", err)
		// }
		// if err != nil {
		// 		return fmt.Errorf("can't fetch host name: %v", err)
		// }
		// fileLogName := fmt.Sprintf("%s/%s%s%d", dirName, dcName, currentUser.Username, epoch)
		// newFilename := strings.Replace(fileName, `\`, `_`, -1)
		// newFileLogname := strings.Replace(fileLogName, `\`, `_`, -1)
		// fmt.Println("newFilename",newFilename)
		// fmt.Println("Start Query")
		// writeToFile("log"+newFileLogname+".txt", "Start Query")
		// writeToFile("log"+newFileLogname+".txt", "Finish Query")	
		/////////////////////////////////////////////////////////////////////

		file, err := os.Create(fileName)
		if err != nil {
			return err
		}
		numFile++
		defer file.Close()

		writer := csv.NewWriter(file)

		allRows, err := db.Query(querysql[0])
		if err != nil {
			return err
		}
		defer allRows.Close()

		columns, err := allRows.Columns()
		if err != nil {
			return err
		}

		// เขียนแถวแรกเป็นคอลัมน์
		err = writer.WriteAll([][]string{columns})
		if err != nil {
			panic(err)
		}

		// Fetch data in batches
		stmt, err := db.Prepare(querysql[2])
		if err != nil {
			return err
		}
		defer stmt.Close()

		// Execute the statement with parameters
		rows, err := stmt.Query("ใบขายสด%", "สินค้ารางวัล", offset, batchSize)
		if err != nil {
			return err
		}
		defer rows.Close()

		for rows.Next() {
			scanValues := make([]interface{}, len(columns))
			scanArgs := make([]interface{}, len(scanValues))
			for i, _ := range scanValues {
				scanArgs[i] = &scanValues[i]
			}

			err := rows.Scan(scanArgs...)
			if err != nil {
				return err
			}

			row := make([]string, len(scanValues))
			for i, v := range scanValues {
				switch v.(type) {
				case []byte:
					row[i] = string(v.([]byte))
					if row[i] == "<nil>" {
						row[i] = ""
					}
				case int:
					row[i] = strconv.Itoa(v.(int))
				case float64:
					row[i] = strconv.FormatFloat(v.(float64), 'f', -1, 'g')
				case time.Time:
					row[i] = v.(time.Time).Format("2006-01-02 15:04:05")
					if row[i] == "<nil>" {
						row[i] = ""
					}
				default:
					row[i] = fmt.Sprintf("%v", v)
					if row[i] == "<nil>" {
						row[i] = ""
					}
				}
			}

			data = append(data, row)
		}

		// Write batch data to CSV
		err = writer.WriteAll(data)
		if err != nil {
			return err
		}
		writer.Flush()

		data = nil // Reset data slice for next batch
	}
	numFile = 0
	fmt.Printf("ข้อมูล %s ถูก Export ลงไฟล์ CSV เรียบร้อย\n", table)

	return nil
}

func getTotalOfRecords(db *sql.DB, query string) (int, error) {
	var result int
	row := db.QueryRow(query)
	err := row.Scan(&result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func generateQuery() map[string][]string {
	// สร้าง map เพื่อเก็บหลาย Slice
	query := map[string][]string{
		"DOCINFO": {`
			SELECT TOP 1 * FROM DOCINFO 
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')	
			`, `
			SELECT COUNT(*) AS result FROM DOCINFO 
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
			`, `
			SELECT * FROM DOCINFO 
			INNER JOIN DOCTYPE ON DI_DT = DT_KEY
			WHERE (DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE @p1 OR DT_THAIDESC LIKE @p2)
			ORDER BY DI_CRE_DATE
			OFFSET @p3 ROWS  
			FETCH NEXT @p4 ROWS ONLY
		`},
		"CASHBOOK": {`
			SELECT TOP 1 * FROM CASHBOOK 
			INNER JOIN DOCINFO ON CASHB_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
			`, `
			SELECT COUNT(*) FROM CASHBOOK 
			INNER JOIN DOCINFO ON CASHB_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
			`, `
			SELECT * FROM CASHBOOK 
			INNER JOIN DOCINFO ON CASHB_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE (DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE @p1 OR DT_THAIDESC LIKE @p2)
			ORDER BY DI_CRE_DATE
			OFFSET @p3 ROWS
			FETCH NEXT @p4 ROWS ONLY
		`},
		"TRANPAYH": {`
			SELECT TOP 1 * FROM TRANPAYH
			INNER JOIN DOCINFO ON TPH_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
			`, `
			SELECT COUNT(*) FROM TRANPAYH
			INNER JOIN DOCINFO ON TPH_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
			`, `
			SELECT * FROM TRANPAYH
			INNER JOIN DOCINFO ON TPH_DI = DI_KEY
			INNER JOIN DOCTYPE ON DI_DT = DT_KEY
			WHERE (DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE @p1 OR DT_THAIDESC LIKE @p2)
			ORDER BY DI_CRE_DATE
			OFFSET @p3 ROWS
			FETCH NEXT @p4 ROWS ONLY
		`},
		"TRANPAYD": {`
			SELECT TOP 1 * FROM TRANPAYD
			INNER JOIN TRANPAYH ON TPD_TPH = TPH_KEY
			INNER JOIN DOCINFO ON TPH_DI = DI_KEY
			INNER JOIN DOCTYPE ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
			`, `
			SELECT COUNT(*) FROM TRANPAYD
			INNER JOIN TRANPAYH ON TPD_TPH = TPH_KEY
			INNER JOIN DOCINFO ON TPH_DI = DI_KEY
			INNER JOIN DOCTYPE ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
		`, `
			SELECT * FROM TRANPAYD
			INNER JOIN TRANPAYH ON TPD_TPH = TPH_KEY
			INNER JOIN DOCINFO ON TPH_DI = DI_KEY
			INNER JOIN DOCTYPE ON DI_DT = DT_KEY
			WHERE (DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE @p1 OR DT_THAIDESC LIKE @p2)
			ORDER BY DI_CRE_DATE
			OFFSET @p3 ROWS
			FETCH NEXT @p4 ROWS ONLY
		`},
		"ARPAYMENT": {`
			SELECT TOP 1 * FROM ARPAYMENT
			INNER JOIN DOCINFO ON ARP_DI = DI_KEY
			INNER JOIN DOCTYPE ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
			`, `
			SELECT COUNT(*) FROM ARPAYMENT
			INNER JOIN DOCINFO ON ARP_DI = DI_KEY
			INNER JOIN DOCTYPE ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
		`, `
			SELECT * FROM ARPAYMENT
			INNER JOIN DOCINFO ON ARP_DI = DI_KEY
			INNER JOIN DOCTYPE ON DI_DT = DT_KEY
			WHERE (DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE @p1 OR DT_THAIDESC LIKE @p2)
			ORDER BY DI_CRE_DATE
			OFFSET @p3 ROWS
			FETCH NEXT @p4 ROWS ONLY
		`},
		"TRANSTKJ": {`
			SELECT TOP 1 * FROM TRANSTKJ
			INNER JOIN DOCINFO ON TRJ_DI = DI_KEY
			INNER JOIN DOCTYPE ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
			`, `
			SELECT COUNT(*) FROM TRANSTKJ
			INNER JOIN DOCINFO ON TRJ_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
		`, `
			SELECT * FROM TRANSTKJ
			INNER JOIN DOCINFO ON TRJ_DI = DI_KEY
			INNER JOIN DOCTYPE ON DI_DT = DT_KEY
			WHERE (DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE @p1 OR DT_THAIDESC LIKE @p2)
			ORDER BY DI_CRE_DATE
			OFFSET @p3 ROWS
			FETCH NEXT @p4 ROWS ONLY
		`},
		"TRANSTKH": {`
			SELECT TOP 1 * FROM TRANSTKH
			INNER JOIN DOCINFO ON TRH_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
			`, `
			SELECT COUNT(*) FROM TRANSTKH
			INNER JOIN DOCINFO ON TRH_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
		`, `
			SELECT * FROM TRANSTKH
			INNER JOIN DOCINFO ON TRH_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE (DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE @p1 OR DT_THAIDESC LIKE @p2)
			ORDER BY DI_CRE_DATE
			OFFSET @p3 ROWS
			FETCH NEXT @p4 ROWS ONLY
		`},
		"TRANSTKD": {`
			SELECT TOP 1 * FROM TRANSTKD
			INNER JOIN TRANSTKH ON TRD_TRH = TRH_KEY
			INNER JOIN DOCINFO ON TRH_DI = DI_KEY
			INNER JOIN DOCTYPE ON DI_DT = DT_KEY
			WHERE (DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
			`, `
			SELECT COUNT(*) FROM TRANSTKD
			INNER JOIN TRANSTKH ON TRD_TRH = TRH_KEY
			INNER JOIN DOCINFO ON TRH_DI = DI_KEY
			INNER JOIN DOCTYPE ON DI_DT = DT_KEY
			WHERE (DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
		`, `
			SELECT * FROM TRANSTKD
			INNER JOIN TRANSTKH ON TRD_TRH = TRH_KEY
			INNER JOIN DOCINFO ON TRH_DI = DI_KEY
			INNER JOIN DOCTYPE ON DI_DT = DT_KEY
			WHERE (DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE @p1 OR DT_THAIDESC LIKE @p2)
			ORDER BY DI_CRE_DATE
			OFFSET @p3 ROWS
			FETCH NEXT @p4 ROWS ONLY
		`},
		"SKUMOVE": {`
			SELECT TOP 1 * FROM SKUMOVE
			INNER JOIN DOCINFO ON SKM_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
			`, `
			SELECT COUNT(*) FROM SKUMOVE
			INNER JOIN DOCINFO ON SKM_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
		`, `
			SELECT * FROM SKUMOVE
			INNER JOIN DOCINFO ON SKM_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE (DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE @p1 OR DT_THAIDESC LIKE @p2)
			ORDER BY DI_CRE_DATE
			OFFSET @p3 ROWS
			FETCH NEXT @p4 ROWS ONLY
		`},
		"SLDETAIL": {`
			SELECT TOP 1 * FROM SLDETAIL
			INNER JOIN DOCINFO ON SLD_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
			`, `
			SELECT COUNT(*) FROM SLDETAIL
			INNER JOIN DOCINFO ON SLD_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
		`, `
			SELECT * FROM SLDETAIL
			INNER JOIN DOCINFO ON SLD_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE (DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE @p1 OR DT_THAIDESC LIKE @p2)
			ORDER BY DI_CRE_DATE
			OFFSET @p3 ROWS
			FETCH NEXT @p4 ROWS ONLY
		`},
		"VATTABLE": {`
			SELECT TOP 1 * FROM VATTABLE
			INNER JOIN DOCINFO ON VAT_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
			`, `
			SELECT COUNT(*) FROM VATTABLE
			INNER JOIN DOCINFO ON VAT_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
		`, `
			SELECT * FROM VATTABLE
			INNER JOIN DOCINFO ON VAT_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE (DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE @p1 OR DT_THAIDESC LIKE @p2)
			ORDER BY DI_CRE_DATE
			OFFSET @p3 ROWS
			FETCH NEXT @p4 ROWS ONLY
		`},
		"ARDETAIL": {`
			SELECT TOP 1 * FROM ARDETAIL
			INNER JOIN DOCINFO ON ARD_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
			`, `
			SELECT COUNT(*) FROM ARDETAIL
			INNER JOIN DOCINFO ON ARD_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE(DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE 'ใบขายสด%' OR DT_THAIDESC LIKE 'สินค้ารางวัล')
		`, `
			SELECT * FROM ARDETAIL
			INNER JOIN DOCINFO ON ARD_DI = DI_KEY
			INNER JOIN DOCTYPE d ON DI_DT = DT_KEY
			WHERE (DI_CRE_DATE BETWEEN '2021-01-01' AND '2024-03-10') AND (DT_THAIDESC LIKE @p1 OR DT_THAIDESC LIKE @p2)
			ORDER BY DI_CRE_DATE
			OFFSET @p3 ROWS
			FETCH NEXT @p4 ROWS ONLY
		`},
		// เพิ่ม Slice เพิ่มเติมตามต้องการ
	}
	return query
}
