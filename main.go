package main

import (
	"context"
	"crypto/aes"
	"database/sql"

	// "crypto/cipher"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io/fs"

	"log"
	"os"
	// "os/user"
	"strconv"
	// "strings"
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

	err = WriteCSV(db)
    if err != nil {
        log.Fatal(err)
    }


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

func upload_to_minio(endpoint, accessKey, secretKey, bucketName, filename string, filelogname string /*, trdNewKey int64, skmNewKey int64*/) {

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: true,
	})
	permissions := fs.FileMode(0644)
	if err != nil {
		fmt.Println("Failed to initialize Minio client:", err)
		os.WriteFile("log_"+filelogname+"_error.txt", []byte("Failed to initialize Minio client:"+err.Error()), permissions)
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		os.WriteFile("log_"+filelogname+"_error.txt", []byte("Failed to open file:"+err.Error()), permissions)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Failed to get file info:", err)
		os.WriteFile("log_"+filelogname+"_error.txt", []byte("Failed to get file info:"+err.Error()), permissions)
	}

	_, err = minioClient.PutObject(context.Background(), bucketName, "pending_inventory_data/"+filename, file, fileInfo.Size(), minio.PutObjectOptions{})
	if err != nil {
		fmt.Println("Failed to upload file to Minio:", err)
		os.WriteFile("log_"+filelogname+"_error.txt", []byte("Failed to upload file to Minio:"+err.Error()), permissions)
	}
	os.WriteFile("log_"+filelogname+".txt", []byte(fmt.Sprintf("Successfully uploaded "+filename)), permissions)

	logFile, err := os.Open("log_" + filelogname + ".txt")
	if err != nil {
		fmt.Println("Failed to open log file: ", err)
		os.WriteFile("log_"+filelogname+"_error.txt", []byte("Failed to open log file:"+err.Error()), permissions)
	}
	defer logFile.Close()

	logFileInfo, err := logFile.Stat()
	if err != nil {
		fmt.Println("Failed to get log file info: ", err)
		os.WriteFile("log_"+filelogname+"_error.txt", []byte("Failed to get log file info:"+err.Error()), permissions)
	}

	_, err = minioClient.PutObject(context.Background(), bucketName, "log_inventory/sucess_bplus_log/"+"log_"+filelogname+".txt", logFile, logFileInfo.Size(), minio.PutObjectOptions{})
	if err != nil {
		fmt.Println("Failed to upload log to Minio: ", err)
		os.WriteFile("log_"+filelogname+"_error.txt", []byte("Failed to upload log to Minio:"+err.Error()), permissions)
	}

	_, err = minioClient.PutObject(context.Background(), bucketName, "pending_inventory_data/"+filename, file, fileInfo.Size(), minio.PutObjectOptions{})
	if err != nil {
		fmt.Println("Failed to upload file to Minio:", err)
		os.WriteFile("log_"+filelogname+"_error.txt", []byte("Failed to upload file to Minio:"+err.Error()), permissions)
	}

	os.WriteFile("log_"+filelogname+".txt", []byte(fmt.Sprintf("Successfully uploaded "+filename)), permissions)
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
			os.WriteFile("log_"+filelogname+"_error.txt", []byte("Failed to upload log to Minio:"+err.Error()), permissions)
		}
		fmt.Println("Successfully uploaded log file to Minio")
		os.Remove("log_" + filelogname + "_error.txt")
	}
	file, err = os.Open("TrdAndSkmData.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// reader := csv.NewReader(file)
	// records, err := reader.ReadAll()
	// if err != nil {
	// 	panic(err)
	// }

	// records[1][0] = strconv.Itoa(int(trdNewKey))
	// records[1][1] = strconv.Itoa(int(skmNewKey))

	// file, err = os.Create("TrdAndSkmData.csv")
	// if err != nil {
	// 	panic(err)
	// }
	// defer file.Close()
	// writer := csv.NewWriter(file)
	// for _, record := range records {
	// 	err := writer.Write(record)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	// writer.Flush()
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


func WriteCSV(db *sql.DB) (error) {
    defer db.Close()

	file, err := os.Create("data-test.csv")
    if err != nil {
		return err
    }
    defer file.Close()
	
	writer := csv.NewWriter(file)

    rows, err := db.Query(`
	SELECT TOP 10 * FROM DOCINFO DC
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
	`)
    if err != nil {
        return err
    }
    defer rows.Close()

    columns, err := rows.Columns()
    if err != nil {
        return err
    }

	// เขียนแถวแรกเป็นคอลัมน์
	err = writer.WriteAll([][]string{columns})
	if err != nil {
		panic(err)
	}

    var data [][]string
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
				if row[i] == "<nil>"{
					row[i] = ""
				}
            case int:
                row[i] = strconv.Itoa(v.(int))
            case float64:
                row[i] = strconv.FormatFloat(v.(float64), 'f', -1, 'g')
            case time.Time:
                row[i] = v.(time.Time).Format("2006-01-02 15:04:05")
				if row[i] == "<nil>"{
					row[i] = ""
				}
            default:
                row[i] = fmt.Sprintf("%v", v)
				if row[i] == "<nil>"{
					row[i] = ""
				}
            }
        }

        data = append(data, row)
    }

    err = writer.WriteAll(data)
    if err != nil {
        return err
    }

    writer.Flush()
	fmt.Println("ข้อมูลถูก Export ลงไฟล์ CSV เรียบร้อย")

    return nil
}

// func writeCSV(data [][]string, fileName string) error {
//     file, err := os.Create("data-test.csv")
//     if err != nil {
//         return err
//     }
//     defer file.Close()

//     writer := csv.NewWriter(file)
//     err = writer.WriteAll(data)
//     if err != nil {
//         return err
//     }

//     writer.Flush()
//     return nil
// }