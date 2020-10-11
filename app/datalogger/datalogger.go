package datalogger

import (
	"encoding/json"
	"github.com/tkrajina/gpxgo/gpx"
	"os"
)

type DataLogger struct {
	DestinationFilePath string
	destinationFile     *os.File
}

func (dl *DataLogger) LogDataPoint(dp *DataPoint) error {
	//Ensure we have a file handle
	dl.ensureFileWriter()

	//Serial Data to JSON
	jsonBytes, err := json.Marshal(dp)
	if err != nil {
		return err
	}

	//Write to file
	_, err = dl.destinationFile.Write(jsonBytes)
	if err != nil {
		return err
	}
	_, err = dl.destinationFile.WriteString("\n")
	if err != nil {
		return err
	}
	err = dl.destinationFile.Sync()
	if err != nil {
		return err
	}

	return nil
}

func (dl *DataLogger) ensureFileWriter() error {
	//Stop if we have a handle open
	if dl.destinationFile != nil {
		return nil
	}

	//Open the file
	f, err := os.Open(dl.DestinationFilePath)
	if err != nil {
		return err
	}

	//Save the file handle
	dl.destinationFile = f

	return nil
}
