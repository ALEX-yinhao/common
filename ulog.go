package common

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"errors"
)

var Log = logrus.New()

func InitLog(path string, prefix string, suffix string, size int) error {

	file, err := os.OpenFile(path + prefix + "_" + time.Now().Format("20060102150405") + suffix, os.O_WRONLY | os.O_APPEND | os.O_CREATE | os.O_SYNC, 0600)
	if err != nil {
		Log.Fatal("log  init failed : ", err)
		return err
	}

	info, err := file.Stat()
	if err != nil {
		Log.Fatal(err)
		return err
	}
	fileWriter := logFileWriter{file, info.Size(), path, prefix, suffix, int64(size)}
	Log.SetOutput(&fileWriter)
	return nil
}

type logFileWriter struct {
	file    *os.File
	size    int64

	path    string
	prefix  string
	suffix  string
	maxsize int64
}

func (p *logFileWriter) Write(data []byte) (n int, err error) {
	if p == nil {
		return 0, errors.New("logFileWriter is nil")
	}
	if p.file == nil {
		return 0, errors.New("file not opened")
	}
	n, e := p.file.Write(data)
	p.size += int64(n)

	if p.size > 1024 * 1024 * p.maxsize {
		p.file.Close()
		p.file, _ = os.OpenFile(p.path + p.prefix + "_" + time.Now().Format("20060102150405") + p.suffix, os.O_WRONLY | os.O_APPEND | os.O_CREATE | os.O_SYNC, 0600)
		p.size = 0
	}
	return n, e
}