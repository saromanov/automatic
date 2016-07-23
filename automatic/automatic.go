package automatic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/buger/jsonparser"
	"github.com/fatih/color"
)

var (
	Deploy = "deploy"
	Test   = "test"
)

type Automatic struct {
	workConfig []byte
}

// Load config provides load config data
// config data should be at the current dir
func (auto *Automatic) LoadConfig(path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	auto.workConfig = file
	return nil
}

func (auto *Automatic) Process(name string) {
	jsonparser.ArrayEach(auto.workConfig, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		fmt.Println(dataType)
		if err != nil {
			//return
		}

		parse(value)

		valuePrint, _, _, errPrint := jsonparser.Get(value, "parallel")
		if errPrint != nil {
			return
		}
		var parallelList []interface{}
		err = json.Unmarshal(valuePrint, &parallelList)
		if err != nil {
			return
		}
		failChan := make(chan int, 1)
		var wg sync.WaitGroup
		wg.Add(len(parallelList))
		for _, item := range parallelList {
			dataMap := item.(map[string]interface{})
			execPath, ok := dataMap["exec"]
			if ok {
				go func(command string) {
					defer wg.Done()
					_, err := ExecCommand(command)
					if err != nil {
						color.Red(fmt.Sprintf("Falled to execute command: %s", command))
						failChan <- 1
						return
					}

					failChan<- 2

					//fmt.Println(out)
				}(execPath.(string))
			}
		}

		switch <-failChan {
		case 1:
			color.Red("Stop")
			os.Exit(1)
		case 2:
			return
		} 
	}, "deploy")
}

// parse provides parsing of the config file
func parse(value []byte) {
	valueData, errPath := jsonparser.GetString(value, "script", "path")
	if errPath == nil {
		// Execution of the script
		fmt.Println("EXECUTE PATH: ", valueData)
	}

	valuePrint, errPrint := jsonparser.GetString(value, "print")
	if errPrint == nil {
		color.Green(valuePrint)
	}

	valueExec, errExec := jsonparser.GetString(value, "exec")
	if errExec == nil {
		var wg sync.WaitGroup
		wg.Add(1)

		go func(command string) {
			defer wg.Done()
			_, err := ExecCommand(command)
			if err != nil {
				color.Red(fmt.Sprintf("Falled to execute command: %s", command))
				return
			}

			//fmt.Println(out)
		}(valueExec)

		wg.Wait()
	}
}
