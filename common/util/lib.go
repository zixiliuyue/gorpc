package util

import (
	// "context"
	"fmt"
	"gorpc/common"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"sync/atomic"
	// "github.com/emicklei/go-restful"
	// "github.com/gin-gonic/gin"
	"bytes"
	"github.com/rs/xid"
	"io"
	"io/ioutil"
	"net/url"
)

// GetDailAddress returns the address for net.Dail
func GetDailAddress(URL string) (string, error) {
	uri, err := url.Parse(URL)
	if err != nil {
		return "", err
	}
	var port = uri.Port()
	if uri.Port() == "" {
		port = "80"
	}
	return uri.Hostname() + ":" + port, err
}

func PeekRequest(req *http.Request) (content []byte, err error) {
	buf := &bytes.Buffer{}
	content, err = ioutil.ReadAll(io.TeeReader(req.Body, buf))
	if err != nil {
		return
	}
	req.Body = NewCloserWrapper(buf, req.Body.Close)
	return
}

func NewCloserWrapper(r io.Reader, closeFunc func() error) io.ReadCloser {
	return &closerWrapper{Reader: r, closeFunc: closeFunc}
}

type closerWrapper struct {
	io.Reader
	closeFunc func() error
}

func (c *closerWrapper) Close() error { return c.closeFunc() }

func InStrArr(arr []string, key string) bool {
	for _, a := range arr {
		if key == a {
			return true
		}
	}
	return false
}

// IsNil returns whether value is nil value, including map[string]interface{}{nil}, *Struct{nil}
func IsNil(value interface{}) bool {
	rflValue := reflect.ValueOf(value)
	if rflValue.IsValid() {
		return rflValue.IsNil()
	}
	return true
}

type AtomicBool int32

func NewBool(yes bool) *AtomicBool {
	var n = AtomicBool(0)
	if yes {
		n = AtomicBool(1)
	}
	return &n
}

func (b *AtomicBool) SetIfNotSet() bool {
	return atomic.CompareAndSwapInt32((*int32)(b), 0, 1)
}

func (b *AtomicBool) Set() {
	atomic.StoreInt32((*int32)(b), 1)
}

func (b *AtomicBool) UnSet() {
	atomic.StoreInt32((*int32)(b), 0)
}

func (b *AtomicBool) IsSet() bool {
	return atomic.LoadInt32((*int32)(b)) == 1
}

func (b *AtomicBool) SetTo(yes bool) {
	if yes {
		atomic.StoreInt32((*int32)(b), 1)
	} else {
		atomic.StoreInt32((*int32)(b), 0)
	}
}

type Int64Slice []int64

func (p Int64Slice) Len() int           { return len(p) }
func (p Int64Slice) Less(i, j int) bool { return p[i] < p[j] }
func (p Int64Slice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func GenerateRID() string {
	unused := "0000"
	id := xid.New()
	return fmt.Sprintf("cc%s%s", unused, id.String())
}

// Int64Join []int64 to string
func Int64Join(data []int64, separator string) string {
	var ret string
	for _, item := range data {
		ret += strconv.FormatInt(item, 10) + separator
	}
	return strings.Trim(ret, separator)
}

// BuildMongoField build mongodb sub item field key
func BuildMongoField(key ...string) string {
	return strings.Join(key, ".")
}

// BuildMongoSyncItemField build mongodb sub item synchronize field key
func BuildMongoSyncItemField(key string) string {
	return BuildMongoField(common.MetadataField, common.MetaDataSynchronizeField, key)
}
