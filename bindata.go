package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _templates_homepage_html_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x90\x41\x6e\xc4\x20\x0c\x45\xf7\x3d\x85\x85\x66\x59\xc5\xea\xb6\x72\x38\x41\xd5\x4d\x4f\x40\x13\xb7\xa0\x7a\x68\x05\x9e\xc5\x28\xca\xdd\x4b\xf0\x20\x65\xc5\xe7\xcb\x7e\xff\x03\x45\xbd\x8a\x7f\x02\xa0\xc8\x61\x3d\x44\x93\x9a\x54\xd8\x13\xda\xd9\x3d\x20\x49\xf9\x07\x62\xe1\xaf\xd9\x61\xd5\xa0\x69\xc1\x20\x32\x2d\xb5\x3a\x28\x2c\xb3\xab\x7a\x17\xae\x91\x59\x5d\x07\xe2\x20\xd2\xe7\xef\x7a\x7f\xa0\x0f\x8f\xcb\x83\xd9\xae\x2f\x9e\xc2\xa0\x3a\xff\xb1\x94\xf4\xa7\xf0\x16\x6e\x79\x89\x5c\x08\x43\x6b\xd1\x66\x6c\x17\xcf\xcb\x74\x13\x13\xdb\x06\x25\xe4\x6f\x86\x4b\x0e\x57\x7e\x86\x4b\x35\xc8\xeb\x0c\x93\xf1\x2a\xec\xfb\x08\x94\x74\x0a\xb4\xc9\x8a\xdb\xd6\x77\xf7\xdd\xf9\x26\xcd\x9d\xde\xbb\x73\x54\x18\x31\x9c\xd7\x41\x22\xb4\x78\x42\x7b\x5b\xeb\xd6\xff\xf1\x3f\x00\x00\xff\xff\x70\x3c\x63\x22\x4f\x01\x00\x00")

func templates_homepage_html_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_templates_homepage_html_tmpl,
		"templates/homepage.html.tmpl",
	)
}

func templates_homepage_html_tmpl() (*asset, error) {
	bytes, err := templates_homepage_html_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/homepage.html.tmpl", size: 335, mode: os.FileMode(436), modTime: time.Unix(1422159409, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templates_script_html_tmpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x84\x91\x3d\x6e\xc3\x30\x0c\x85\xf7\x9c\x82\xd0\x92\xcd\x42\xb3\x56\xd6\x52\x74\x2b\xba\xf4\x04\x8c\xcc\xc2\x4e\x15\xdb\x10\x69\x23\x46\x90\xbb\x57\x3f\x76\xfa\x83\x02\x9d\x24\x3d\x3e\x7e\xe0\x13\x4d\x2b\x67\x6f\x77\x00\xa6\x25\x6c\xd2\x25\x5e\xa5\x13\x4f\xd6\xe8\x72\x66\x0d\x8c\xef\xfa\x0f\x68\x03\xbd\xd7\x4a\xb3\xa0\x74\x4e\xa3\xf7\x95\x63\x56\x10\xc8\xd7\x8a\x65\xf1\xc4\x2d\x91\x28\xbb\x4b\x44\xbd\x21\xcd\x71\x68\x96\x95\x9d\x34\x0a\x2b\x34\x3e\x1f\xac\xc1\x0d\xab\xec\x9b\x0b\xdd\x28\xf0\x82\x53\xef\x5a\x0a\x46\x63\x1c\x23\x7a\xee\xf6\x83\xbd\x5e\x2b\xce\xae\xea\x15\xcf\x74\xbb\xc5\xfa\x61\x65\xeb\xef\x70\x73\x9c\x44\x86\x1e\x9c\x47\xe6\x5a\x95\x97\xb2\xcf\x17\x72\x93\x90\xd1\x45\xf8\xd3\xeb\x3c\x61\x50\xf6\x29\x1d\xbf\x8c\x4d\x37\xdf\x5d\x43\xcf\x83\x27\x18\x26\x19\xa7\x94\xb9\xcc\x10\x1d\x39\xb3\xde\x42\x9b\x32\x2e\xc8\x32\x52\xad\x84\x2e\xa2\x4f\x38\x63\x51\xd7\xb6\x19\x03\x14\x21\x85\x82\x1a\xf6\x31\x66\x9f\xf3\xed\x1f\x33\xad\x54\xff\xe7\x01\x07\xf7\x73\x43\x27\x56\xf6\xab\x3f\x7e\x52\xde\xf8\x67\x00\x00\x00\xff\xff\x03\xde\x73\x61\xf9\x01\x00\x00")

func templates_script_html_tmpl_bytes() ([]byte, error) {
	return bindata_read(
		_templates_script_html_tmpl,
		"templates/script.html.tmpl",
	)
}

func templates_script_html_tmpl() (*asset, error) {
	bytes, err := templates_script_html_tmpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "templates/script.html.tmpl", size: 505, mode: os.FileMode(436), modTime: time.Unix(1422161545, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _static_all_css = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x94\x56\xdf\x8f\xa3\x36\x10\xfe\x57\xa8\x56\x27\xdd\x9e\x20\x4b\xb2\xf9\xb1\x0b\xea\x43\xaf\xd5\x4a\x95\xda\xb7\xbe\xad\xf2\x60\xf0\x00\x6e\x8c\x8d\xb0\xc9\x26\x47\xf9\xdf\x3b\xb6\x81\x90\x0b\x39\xe9\x94\x87\x90\xc1\x9e\x1f\xdf\x7c\xdf\x4c\x9e\xbe\xfc\xe2\x09\x59\x97\x84\xb3\x6f\xb0\x48\x95\xf2\x8e\xcf\x8b\x70\xb1\xf2\xfe\xf3\xfe\xfe\xf3\x1f\xef\x2f\x96\x82\x50\x80\xbf\x72\xa6\x17\x4c\x3e\x8d\x67\xbd\x2f\x4f\x85\x2e\x79\x9b\x49\xa1\x83\x8c\x94\x8c\x9f\x23\x45\x84\x0a\x14\xd4\x2c\x8b\x83\x52\x05\x1a\x4e\x3a\x50\x78\x36\x20\xf4\xdf\x46\xe9\x68\x19\x86\x9f\xe2\xe0\x03\x92\x03\xd3\xf3\x6f\xbb\x44\xd2\x73\x5b\x92\x3a\x67\x22\x0a\x3b\x52\x6b\x96\x72\xf0\x89\x62\x14\x7c\x0a\x9a\x30\xae\xfc\x8c\xe5\x29\xa9\x34\x93\xc2\x3c\x36\x35\xf8\x99\x94\x1a\x6a\xbf\x00\x42\xcd\x57\x5e\xcb\xa6\xf2\x4b\xc2\x84\x5f\x82\x68\x7c\x41\x8e\xbe\x82\xd4\xde\x50\x4d\x89\xee\xcf\x2d\x65\xaa\xe2\xe4\x1c\x25\x5c\xa6\x87\x8e\x34\x94\x49\x3f\x25\xe2\x48\x94\x5f\xd5\x32\xaf\x41\x29\xff\x88\x51\xe5\x78\x92\x09\xce\x04\x04\xf6\x42\x7c\x04\x93\x1a\xe1\x01\x82\x91\x8b\x28\x21\x0a\xcc\x5b\xe7\x28\x12\x52\x7f\x7e\x4f\x11\x99\x5a\x72\xb5\x7f\x1c\x5d\x08\x29\x20\x2e\x80\xe5\x85\xc6\xea\xde\x0b\x46\x29\x88\xbd\xaf\xa1\xc4\xd7\x1a\xae\xce\x75\xa4\x4d\x48\x7a\x30\xb5\x08\x1a\xa4\x92\xcb\x3a\xd2\x35\x22\x5c\x91\x1a\x84\xee\x48\x44\xb0\xa2\x23\x82\x13\x15\x12\xd3\x69\x65\xa3\x4d\x0a\x06\xb6\x24\xa9\xdf\x35\xd3\x1c\xf6\x6d\x22\x6b\xc4\x24\x48\xa4\xd6\xb2\x8c\x96\xd5\xc9\xa3\xf8\x08\xb4\x4b\x7c\x85\xe9\x89\xdc\x75\xf0\xc3\x25\xb5\x0b\xc3\x8e\x66\xc2\xd9\x94\x3e\x73\x88\x98\xc6\x12\xd3\xae\x58\xf6\x46\xec\x58\xb4\x82\xb2\x43\x14\x0f\x93\x0c\xa3\x87\x2c\x0b\x63\x97\xe6\x43\x88\x6e\x14\xf2\x84\x4f\xee\xbc\x60\x77\x55\x83\x51\x9b\x6a\x62\xdd\x6d\x3e\xc5\x16\xd6\x01\x95\xb8\x92\x8a\x99\x4e\x45\x35\x20\x26\x58\xe0\x5d\xac\x8d\x27\x2d\xab\x28\x58\x6c\x30\x1f\xf4\xdd\xf6\x55\x06\x8b\x95\xb1\xb0\x32\xef\xcb\x47\x4c\xd4\x31\xb7\x6d\x89\x6a\xe4\xca\x63\x6b\x10\xcb\xb8\xfc\x88\x5c\x0f\x3a\x47\xa4\x81\x79\x4b\x28\xbd\x75\x58\x9d\xba\xa2\x6e\x83\x52\x7e\x43\xf8\x4e\x26\x5f\x26\xf2\xc8\xb4\x15\xf1\x37\xa6\xf8\x8e\x79\xec\x70\x85\x2e\xc7\x48\xa4\xd1\xb2\x4b\x25\x12\xf9\x90\x50\x24\x19\xf8\x8a\x94\xd5\x95\x80\x4a\x29\x24\xf6\x37\x05\x7f\x7c\x8a\x2f\x58\x61\x56\x5d\xd2\x60\x85\xc2\x67\xa2\x6a\xb4\x2f\x2b\xed\xa8\x8e\x80\x20\xbd\x7d\x23\x29\x24\x07\x69\x5d\x1b\x98\x28\x50\x8b\xda\x7a\x18\x7f\x8c\xda\x72\x9e\x2e\xe9\x1d\x99\x62\x09\x87\x21\x82\x73\xd9\x5a\x95\x5a\xda\x65\xa8\x7c\x47\xcc\xfe\x84\x91\xbf\x67\x13\x79\xd7\xe7\x0a\x7e\x75\xe6\xbd\x3f\x31\xa1\x8a\x40\x5f\x59\xb0\x4b\x25\xd3\xfb\x76\x98\x02\xa4\xaa\x80\xa0\xfb\x14\x22\x77\x3f\x4e\x9b\x5a\x61\xf2\x95\x64\x08\x68\xdd\x07\x7b\x47\x65\x10\xcc\x8e\xee\xa7\x61\x47\x63\xdb\x5f\xa2\x90\x91\x86\xeb\xfe\x52\x14\xd9\xde\x65\x32\x6d\x54\xc0\x84\xc0\xd1\x60\xef\xdd\xda\x47\x9a\xc4\x15\xa1\xd4\xb4\x33\xec\xec\xd1\x76\xca\x4d\x37\xfb\xba\x49\x35\x69\x01\xe9\x01\x3b\x7e\x5d\x34\xc1\x11\x60\x84\x37\x72\x63\xd4\xe0\xe9\x7b\xff\xee\x86\x68\xca\x04\xea\x3d\xe6\xd5\xa3\x62\x93\x0a\x54\xc5\x44\x30\x6d\xf8\x9d\xd3\x28\xfc\xeb\xd3\x6d\x9f\xb0\x65\xdc\x14\x7c\x84\x3a\x2d\x66\xc1\x37\x7d\xce\x18\x70\x1a\xff\x88\xef\xc3\xc5\x9f\x92\xc3\x4c\x06\x97\xdc\x9d\x21\x48\x4d\x12\x7c\xa6\xd8\x7b\x17\x28\xa4\xb2\x26\x66\x4e\xcc\x55\x63\x69\x6a\xcb\x41\xfe\x0d\xcd\x35\xb3\x4f\x49\xce\xa8\xa7\x18\x47\xd6\x8f\x52\xf0\x56\xd5\xa5\x31\x8b\x67\x1c\x1d\xde\x62\xbb\xb2\x5f\x3b\x33\x47\x38\xe4\x20\xe8\x1c\x47\x46\xc1\x5d\x8b\x7c\xd0\xe5\xcd\x68\xd5\x86\xae\xc3\x48\x46\x91\x72\x52\x29\x88\x86\x87\xb8\x7f\x61\x74\xdf\xfb\xa7\xbe\x2e\xda\x4b\x3c\xbb\x1d\x6f\xb6\xc2\x43\x46\xcc\xa7\x73\xeb\x6f\xe6\x7d\x18\xbe\x6e\x5f\x5e\xc6\xbc\x97\x38\xd9\xbc\xd0\x4c\xf4\x11\x01\xf3\x31\xe6\x4b\x6d\xf1\xdc\xda\xc3\x3b\xde\x30\x5d\x1e\x60\x99\x6d\x32\x88\xed\x84\xb8\xb4\xc3\x61\x5f\xac\x86\x53\x6f\x6f\xf3\xae\xe2\x1f\xc4\xee\x1a\x8e\xc2\x53\xfd\x06\x0a\x0c\x11\xdc\xea\x1c\x4e\xac\xcc\x6c\xe6\x6c\xa6\xd4\xf5\x66\xf3\xdb\x76\xdd\x3b\x1f\xb6\xde\x1a\x9d\x97\xe4\x14\x7c\x30\xaa\x0b\xd3\x08\x13\xec\xfb\x35\x83\xfe\xfa\x45\x7a\xeb\x75\xf7\xf2\x1a\xbe\xfe\x8e\x27\xb0\xfa\xb9\x7a\x87\xbd\x97\x65\x59\x7c\xf5\xc7\xe2\x0a\xf3\x6e\x81\xaa\x40\x06\xc2\xa0\x4f\xdc\x69\xe9\x67\xf3\xb7\xc7\x0b\xbc\xe5\x16\x59\xf8\x18\xdf\xc6\xde\x86\xbb\x3f\x5e\xbe\x4e\x23\xdc\xf0\xd9\xec\xdb\x01\xcf\x2b\x24\xed\x8f\x81\x9b\xc1\xd9\xb1\xb3\x9f\x11\xb7\x91\x56\xcb\xd7\xed\xdb\xf3\x4c\x24\x5b\xe2\x64\xfb\xaf\x27\x21\x76\x16\xdb\x4b\x68\x87\xf1\xd2\x76\xa8\x1f\xc3\xf7\x50\xdd\xae\xbf\x6e\xde\xb6\xdd\xff\x01\x00\x00\xff\xff\x3d\x9a\x46\x16\x83\x0a\x00\x00")

func static_all_css_bytes() ([]byte, error) {
	return bindata_read(
		_static_all_css,
		"static/all.css",
	)
}

func static_all_css() (*asset, error) {
	bytes, err := static_all_css_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "static/all.css", size: 2691, mode: os.FileMode(436), modTime: time.Unix(1422159336, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _static_all_js = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\x7b\xfb\x52\xdb\xc8\xf2\xf0\xff\x79\x0a\x5b\x75\x3e\x4a\xb3\x16\x02\x67\xf7\x9c\xb3\x6b\x31\x71\x11\x02\x09\xb9\x41\x02\x09\x21\xc6\x67\x6b\x24\x8d\x6c\xc7\xb2\xe5\x48\x32\x86\xc4\x7e\xf7\xaf\x7b\x6e\x1a\xd9\x26\x9b\xdd\xfa\x55\x51\xa0\xb9\xf5\xf4\xf4\xbd\x7b\x86\x5b\x96\x37\x86\x5e\x4a\xcb\xe1\xa8\x08\x1e\x25\xf3\x69\x54\x8e\xb2\x69\x63\xe2\x32\xf2\x1d\xc7\x42\x5a\xde\xcf\x78\x96\x34\x58\x30\x4a\x5c\x27\x0b\xbf\xf0\xa8\x74\x28\x0d\x09\x34\x61\x0e\xfe\x6e\x8c\xa6\x45\xc9\xa6\x11\x4e\x3b\xcc\x73\x76\x4f\x72\x5e\xce\xf3\xa9\xc3\xb0\xe1\x04\xeb\x73\xce\x04\x10\x35\xa9\x11\x06\xb8\x4f\x44\x65\xaf\x3f\xcb\xb3\x32\xc3\x3d\xfd\x32\xbb\x28\xf3\xd1\x74\xe0\x47\x2c\x4d\x61\x2f\x81\x40\x4f\x62\xd0\xb8\x1a\x4d\xe3\x6c\xd1\x07\x4c\x22\xbd\x9b\xc2\xad\x36\x4d\xa0\x23\x66\x2d\x97\xce\x74\x3e\x09\x79\x0e\x0d\x7d\x24\x3f\xe5\xd3\x41\x39\xdc\xd9\x71\xe6\xd3\x98\x27\xa3\x29\x8f\x9d\x66\x35\x5a\xcc\xd2\x51\xc4\x1f\x1a\x05\x44\x67\x3c\x2f\xef\x4f\x8b\x63\x00\xcc\x73\x16\xa6\x30\xb7\xb9\x7d\xc0\x75\x24\x30\x87\x6c\xd2\xc6\x20\x7b\xa2\xa8\xaf\xf1\xdd\xba\x2d\x12\xe3\xff\x08\x25\x04\x55\x21\xa4\x79\xef\xac\x78\x5a\xf0\x86\xea\x9d\xce\x61\x4e\x20\x7a\x10\x55\x33\x89\xd2\x47\x61\x0d\x0d\x5a\x47\x71\x9d\x29\x9a\xd9\x2b\x23\x62\x53\x14\x9f\xf5\xbd\x29\x45\xc9\x5b\xa1\x44\x30\x46\x9d\x28\xcd\x8a\x79\xce\xff\x9c\x8f\xe2\x3f\x9d\x96\xdb\x3e\xfe\xe3\x97\x37\xac\x1c\xfa\x39\x03\xee\x4f\x5c\xf2\xe4\xc9\x93\x7d\xe2\x85\x8c\xee\x07\x06\x6e\xc4\x5c\xe6\x85\x5e\xa4\x81\x2b\x84\x7c\x36\x9b\xa5\xf7\x2e\xf3\x43\x90\x1c\x8f\xe5\x03\xa0\xc3\xb4\x2c\x48\x85\x50\x6c\x16\xc2\x49\x9b\x8c\x94\xc3\x3c\x5b\x34\x8e\xf3\x3c\xcb\x5d\x21\x7c\x8f\x0f\xcc\x32\x25\x39\x52\x47\x62\x2a\xc4\xcc\x12\xdd\x02\x79\xad\xe4\x56\x2f\xf1\x1e\x13\x4d\x05\xbd\xa5\x2b\xd7\x47\x3f\xb5\x9e\x04\xeb\xb3\xe6\xd3\x62\x38\x4a\x4a\x75\xb2\xc8\x8b\xcd\x06\x4c\xf5\xe1\x69\x56\xab\xcd\x5d\x37\xa6\x59\xf4\x58\x55\x56\x60\xa6\x09\x32\xa3\x5a\x36\xad\xed\x91\x90\x3b\x3b\xbb\xed\xe6\x43\x83\x46\x83\x5d\xe2\x43\x93\xdf\x9d\x81\x04\x4d\x59\x39\xba\xe5\x8d\x28\x8b\x41\x19\xba\x11\xeb\xc4\x4c\x63\x3d\x53\xe8\xa0\xcc\xd9\x18\x21\x8d\x38\xa3\xcf\x58\xc9\xfd\x69\xb6\x58\x2e\x37\x4e\xd2\x9a\xf2\x45\x03\xc7\x57\x95\x1c\x7c\x45\xe4\xc9\xf7\x4a\x2e\x60\xf6\x2a\xaa\x50\xa4\x61\xf5\x1d\x30\xff\xf5\x5a\xbb\x9a\x87\xb0\x23\xbb\xc7\x8f\x32\x30\x66\xf9\x3c\x2a\xb3\x9c\x32\x18\x01\x01\x34\x28\x31\x2f\xf2\x12\xd8\x16\x84\x06\xf1\x1e\x48\xde\xba\xeb\x92\xb3\xfb\x98\x78\x63\xfa\x38\x18\x6f\xc8\x54\x30\x6e\xb5\xc8\xa0\x37\xde\x7d\xdc\xa7\x66\xac\x37\xee\x1b\x15\xaa\x10\xe9\x45\x7d\x2d\xd6\xde\x00\x38\x57\x1d\x3e\x57\xb6\x59\x48\x2f\x08\xd2\x0c\x96\xf2\x8b\x92\x45\xe3\xcb\x9c\x45\x9c\x3c\xd0\xef\xa2\x1b\xf0\x72\x22\x14\x5e\xd9\x7f\xa5\x00\x7e\x81\xb3\x02\x50\x7a\x31\x49\x36\xc1\x0d\xac\x98\xee\x99\xf0\xa2\x60\x03\x4e\x15\xcf\x19\x21\xab\xaf\x6e\xee\x89\xf5\x20\x99\x16\xfd\xa6\x6c\xc2\xa9\x73\x34\x2f\xca\x6c\x22\x86\x9d\x0a\xf3\x84\x29\xbe\x29\x02\x46\x54\x1a\xe2\xd2\x75\xfe\x5f\xe1\x10\x2f\xa6\x8e\xe3\xf1\x9f\xd3\xb8\x36\x1c\xc4\xd8\xf8\xf6\x41\xa4\x29\x4c\xe2\x16\x8d\x7c\xa1\x3c\x2e\x69\x71\xfd\xa5\x29\x1c\xb7\x22\xff\x4b\x36\x9a\xca\x1d\x85\xf8\x0d\x98\x3a\x96\xed\x9f\xf2\xd1\xa4\x5b\xf1\xdd\xd2\x2a\x1c\x71\xc9\xaa\xb3\x75\x30\xe7\xb3\x14\x49\xbd\xf7\xbf\xde\x4d\x71\x73\xc7\xf6\xfb\xad\xa5\xf9\xfa\xd7\xde\xc0\x73\x60\xcf\x8a\x1e\xa5\x24\x87\x5e\x7e\x10\x76\x77\xdb\x1d\xf6\x24\xec\xb6\x3b\xfb\xd6\xb4\xb9\x9c\x16\x6a\xa3\x80\xfe\x32\x97\xf4\x10\x3c\x4d\x98\xad\x5c\x21\x21\x41\xa8\x8f\x0d\x5c\x9a\x23\xcf\xe7\x1b\x1c\x3a\x2c\x0a\xf0\x1a\x00\x7e\x9d\x49\x43\xc5\x24\x69\x22\x51\x41\xe6\xae\x73\xc2\x46\x29\x48\x13\xd8\x69\xd6\x75\x3a\x0d\xa7\xc5\x3a\x70\x14\xef\x27\x39\x45\x82\x95\x08\x04\x6e\xd7\x59\xeb\x8d\x18\xbd\xd5\xf6\xc3\xa2\x77\xcd\xc6\x9b\x09\x0a\xb0\x18\xb4\x19\x20\x67\x47\x14\x8f\x0f\xde\xb5\xbb\xdf\xd9\x7f\x12\x75\x85\x33\x99\xb0\x3b\x77\xdf\xd3\xc1\x40\x2b\x22\x9d\x48\xb8\xe5\x42\x30\xdc\x72\x6c\xda\xa7\xad\x0f\x80\x4e\xb4\x21\x22\x52\x00\xba\xcc\x18\x3b\xdc\xb3\xb3\xdb\x0e\x50\x96\x83\xe8\x40\x6f\x11\x44\xa0\xe1\xb0\x43\x04\x41\x51\x03\xd4\x87\x81\x1e\x53\x0c\xa9\xd4\x61\x22\x25\x88\xbb\x6d\x49\x92\xbb\x80\x75\x84\x32\x7e\x61\x34\x05\xd6\xdc\x8e\x06\x0c\x4c\x0f\x22\xf9\x45\x85\x69\x63\x46\xbf\x30\x7f\x0e\xec\x3a\x1c\x00\x41\x71\x68\x0c\x43\x77\x74\xcc\x82\x30\xe7\x6c\xdc\x60\xab\xd5\x1d\xa8\x8e\x84\x98\xa2\x6f\xad\x8c\x18\xf0\xb3\x38\x5b\x4c\xcf\x55\x94\xd0\x18\x15\xe7\x9a\xfa\x67\x49\x63\x5b\xf0\xd0\x28\xb3\xd7\x19\x90\x9a\x4b\xad\x68\x68\x4b\xdf\xb8\x65\xe9\x1c\x16\x39\x5a\x67\x1b\x0e\xa9\xc4\x66\xb2\xae\xdb\x5e\x0c\xaa\xdc\x0e\xf8\xa6\x11\xe4\x40\xa2\xef\xb1\x65\x00\x79\x5f\x90\x51\xd0\x2c\x26\x82\x64\x31\xfc\x0a\x34\xac\x04\x63\x80\x83\xd4\xd0\x38\x01\x00\x11\x4d\x59\x2f\xe9\x7b\x1b\xa1\x65\xfd\xbc\x52\x66\x62\x60\x16\x98\x32\x03\x19\xed\x29\x02\x9e\x32\x8a\x3e\xee\xae\x72\x61\x67\xb0\x8a\x39\x64\xb9\x5c\xef\x3f\x7f\x0f\xf2\x7e\xbf\x3e\xfd\x32\x1f\xc5\x70\x84\x2d\x0b\xde\x5c\x9c\x1e\xc3\x8a\x6f\xeb\x2b\x9e\xf3\x68\x9c\x39\x04\x9d\x2b\x85\x7e\xa4\xf5\x82\xe7\x47\xac\xe0\xb6\x27\x5d\xf0\x70\x3c\x2a\x71\x5a\xd3\xfd\x9b\x5b\x82\x56\xaa\x3d\xff\x02\x76\xc5\xbb\x8c\xa9\x68\x05\x65\x30\xce\x22\xc1\x16\x13\x6f\x80\xd0\xeb\xbe\x37\xe0\xd9\x3b\xb7\xd9\x28\x6e\xec\xaf\x1e\xe1\x82\x99\xe5\x1d\x35\x08\xb0\xe1\x21\x0a\xe9\x14\xc4\x3f\xf5\x91\x0d\x5a\xbb\x04\x7c\xd1\xe1\xdf\xf2\xbc\x80\x45\x1e\xda\xcf\x2e\x6c\xdf\x61\xc1\xb7\x6e\x48\xf7\xf2\xdb\x9b\x8e\xdb\xfb\xdf\x0d\x09\xfa\x2d\xe2\xde\x90\x65\x40\xf6\x3a\xf7\x38\x72\x13\xba\xdd\x0e\x1e\x71\x99\xdf\x92\x5e\xa7\xd1\xdf\x98\x76\x08\x2c\x86\x89\x57\x3c\x7c\x35\x2a\x6f\xf6\xdc\x9b\x8b\x16\xd9\x23\xc2\xad\x31\x0a\x3f\xa1\xcf\xef\x78\xe4\xde\x11\xd8\xb1\xd7\xee\xa3\x05\xd3\x87\xbc\x17\x4b\x91\x0e\x5e\xf8\x64\xc6\xf2\x82\x9f\xa4\x19\x43\x4b\x4b\xba\xca\xe7\x85\x80\xe3\x0a\xc6\xbf\x32\xfa\x7d\x65\xa5\x53\x4f\x4d\x3a\x85\x87\x6e\x02\x98\xaf\xac\xc7\xfa\x04\x0c\x36\x8a\xad\x71\x75\x03\xe6\x2a\x48\x33\x80\xaa\x75\xc8\x17\x6e\xaf\x1a\x5b\x1b\xe2\xd4\x98\x31\xed\xe1\xbc\x58\x07\xa8\x1e\xea\xc5\x3e\x18\x97\x9d\x9d\xe4\x80\x0b\xad\xf8\x2e\xe3\x92\x08\x34\x03\x92\x0c\x07\xc2\x90\x58\x7f\x2e\xe8\x7b\x3e\x38\xbe\x9b\xb9\x8e\x7b\x73\x13\xff\x02\x54\xbb\x79\xf6\x0b\x71\x3c\x67\x00\xfb\xbc\x64\x3f\x18\x0d\xe2\x4c\xc0\x7d\x46\x17\x92\x82\x03\x90\xbd\x1e\x80\x14\x3f\x7d\xef\x98\xbe\x64\x72\x60\x5c\x1b\x40\x7a\x00\x7a\xcf\x7a\xfb\x7d\xe3\xa9\xa1\x7d\x5c\xb5\x89\xb0\x5f\x01\x24\xa2\x72\x62\x5b\x0f\x80\x0d\x17\x5c\x38\x9d\x96\x2e\x76\x7b\x6d\xc8\x05\xc4\xda\xad\x53\x8e\xd5\x14\xd8\x5e\x41\x7a\xac\xa7\xc9\x55\xa6\x29\x66\xe0\xb0\x87\x9d\x64\xb5\x18\x8e\x20\x5d\x42\x22\x92\x95\xe2\x1c\xdd\x3f\xa0\xe1\xca\x24\x35\x42\xce\x73\x5b\x33\xbc\x82\xd1\x1c\x84\xfb\xbe\x8b\x02\xb3\x5c\xba\xce\xd1\xc5\x45\xfb\x28\x9b\xcc\x18\xa6\xcf\x20\xdd\x91\xf8\x46\x6d\xe9\x1a\x1c\x67\x0c\x31\xec\xfc\x9b\x28\x05\x0a\x9a\xdf\x40\xbb\xef\x97\x4b\x90\x3d\xf8\xf9\xe3\x80\x16\x6c\xb9\x84\xbe\xa7\xae\xd3\xf6\xff\xf0\xdb\x40\xf8\x7b\xd1\xfa\xc3\xd6\xd5\xa3\x2a\x7b\x37\xf8\x48\xd7\xa7\x64\xd9\x8d\x40\xd0\xbf\xce\x79\x7e\x7f\xc1\x53\x8e\x2e\xe0\x10\xb3\xca\xb5\xbe\xee\x5a\x1b\xc5\xad\xc5\x48\xa7\x04\x33\x4e\x80\x41\x70\x2c\x84\x69\x25\x0c\x62\x44\xee\xec\x09\xe3\x1e\x54\x18\x20\xa7\x7f\x62\x53\x70\x87\xa6\x32\xb0\x31\xdb\x75\x64\x74\xe1\xab\xe0\x42\xa4\x65\x0c\x61\x0c\x78\x79\x9c\x72\xe1\x2a\x9e\xde\x1f\xa5\xac\x28\xde\x42\x08\x23\x71\x49\xe8\x43\xe3\x22\x4a\x52\xd9\xd1\x6a\x63\xda\x25\x1b\x88\x49\xce\x2f\x8e\xdc\x88\x7c\xe7\xa8\xd7\xc2\x0f\xd1\x18\x14\x2b\xa4\x09\xfa\xa0\xa8\xd2\xaa\xd0\x8f\x34\x70\x6f\x2c\xdc\x30\xb5\xb3\x5b\x15\x31\x0c\xa4\xfe\x92\x31\xca\xd1\x88\xb9\xaa\xed\xee\xdd\x14\xad\x3d\xe2\x01\x56\x63\x30\x36\xbc\x17\xb7\x5a\x7d\x94\x3a\x1d\xc5\xd2\x58\xa3\xcb\x4d\x56\x27\xbd\xd4\x9c\x51\x25\x27\x42\x48\xbc\x5b\x46\xe1\xbb\xa9\x24\xa3\x79\xb8\x5c\xc2\xe7\xbf\x1f\xff\x0e\x0d\x23\x3e\x21\xfa\x07\x29\x3e\xbf\xe3\x27\x1a\x63\x5c\xe0\xff\x1b\x5b\x87\xa2\x61\x96\x34\xf5\xac\x2d\xf2\x76\x02\x76\x7d\xbf\x49\x17\x3a\x15\xe8\x31\x06\xf6\xc4\x7c\xd2\x56\x2b\x04\x9b\x15\x88\x2c\xe1\x29\x95\x7f\x64\xeb\x83\x6c\x7d\x10\x41\xf6\x02\xf3\xf8\x13\xcb\x4d\x3f\xa5\xcd\x76\xb5\xcb\x73\x1d\x73\xc2\x02\x91\x98\x31\x09\x23\x9a\xe7\x39\x30\xec\x12\x62\x05\x5e\x4a\x78\xa5\xfc\x0e\xe5\x84\x98\x27\x6c\x9e\x96\xe7\x39\xbf\x85\x79\x3c\x96\x73\x06\x08\x5c\x7c\xbd\xa4\xcd\xfd\xd5\x73\x3b\x44\x2d\xb3\x19\x06\x07\x0c\xe2\x2c\xd8\xd9\x76\x5f\x7a\x29\xc4\xde\xf6\x8a\x99\x84\xfd\x4c\xee\xb4\xb1\x60\x03\x83\xe6\xbe\xd9\xba\x6d\x45\xf1\x2f\x50\xc6\x5e\xf4\x20\x6a\xea\x5b\xa2\xc9\x56\xb2\xcb\x06\x6b\x2d\x3a\x95\x74\x79\x6e\x45\xfb\xe8\x8f\x11\x2f\xe1\xbe\xc4\x46\x39\x4f\x21\x3d\x8e\x6d\x22\x3d\x4c\x37\x61\x29\xe4\x9c\x21\xcb\x8f\xc0\x44\xc9\xe1\x31\xbf\xaf\x1a\xe1\xbc\x2c\x81\x36\x32\x19\x8c\x72\xce\xa7\xd7\x76\xe3\x93\xda\x24\x1d\xc1\x1e\xd7\x76\x43\x8d\x64\x49\x52\x70\x3d\x22\x1b\x9f\xa8\x22\xcb\x84\x97\xec\x15\xbf\x57\xf0\x30\x4d\x31\x2d\x96\x56\xdf\x51\x99\xa7\xd8\xd0\x8c\x0c\xa9\x4e\x4d\x4b\x2e\x0f\x21\xb5\x56\x0d\x32\x55\x69\xb4\x24\x48\xfc\x09\xec\xb3\x33\xf5\xb1\x5c\x42\xde\x99\x47\xca\x1a\x6c\x93\x34\x59\xb8\x84\x18\xb5\x4e\x5c\xdc\x33\x16\xc9\xf7\x37\x69\x14\x38\x86\xef\x65\x7e\xff\xfd\x85\x1b\xfb\x53\xa0\x9f\xb0\x4c\x01\x47\x21\xd0\xb1\x79\xc4\xca\x68\xe8\x26\xc0\x57\x8e\x12\xc1\x41\x7b\x62\x71\x02\x88\x42\x31\x17\x77\x26\x19\x04\xf7\xd9\xad\x28\x5f\x3e\x8a\xba\xb8\x6b\x92\x43\xf6\x2c\xd1\xeb\xa8\xf1\x39\xfa\x96\x08\xb4\x10\xc7\xcb\x4c\x8d\x6e\x15\x81\x38\x50\x51\xb1\x2c\xe0\xa9\x58\x78\xc4\x8b\xee\x03\xfd\x52\xb2\xbe\x2b\x4e\x75\xbe\x43\xfe\x90\x8c\x06\x73\x91\x0e\x74\x9a\xfb\x1e\x37\xc9\x01\xb6\x60\x8b\x8e\xd8\xf5\x99\x57\xe8\xcf\xcf\x2b\x4f\x31\xfd\xa7\x57\x9f\x54\xab\xff\xb5\x5a\x91\x8e\x5b\x93\x16\xb9\x01\xc4\x5b\x9b\x02\x05\x26\x49\x1d\x5b\x0b\x9d\xf4\xa8\x4d\x0a\x94\x51\x5d\x5d\xf3\xd5\x61\x3e\x28\x3b\xff\x14\xd4\x64\x76\x63\xc5\xb5\x59\x71\xad\x56\x5c\x07\x35\x91\x67\xfa\x6b\xb9\xdc\x0f\x6a\x9a\xa1\x47\xae\xcd\x88\xd2\x1f\xa6\x3e\x82\x9a\x86\x31\xfd\x65\xa6\x1b\x4d\x7c\xc4\xcc\x37\x86\x15\x30\x0f\x8c\x4f\x51\x20\xe3\xbb\x66\x59\x67\x5f\x9f\x5e\xa9\x08\xd3\x5f\x81\xad\x45\x4c\x7d\x04\x75\x45\x63\xe6\xb3\xae\x8e\x4c\x7f\x05\x96\x9e\x31\xf9\x37\x60\x1b\x56\x6e\x67\x47\x4c\xab\xdb\x46\x17\x24\xfa\xab\x7b\xea\x3d\x27\xc1\x90\x9e\x5a\x25\xb7\xe1\x8f\xcc\xee\xa9\xff\x7a\x7d\xb8\x32\x77\xea\xa8\xe1\xfa\x8c\xee\xf6\x6e\x48\x24\xd4\x40\x84\x17\x12\xe9\xd3\x79\x08\x62\x27\x4c\xfa\xa3\xe1\x0f\x2c\x39\xe2\x50\x1f\xb5\x51\x90\x79\x8d\x04\x2c\xcc\xce\xda\x5c\xb2\xde\xe1\x12\x53\x56\x47\x0b\x82\x96\xfe\x23\x26\xcf\xa0\xff\xe0\xbd\x09\x1a\x0c\x31\xa4\xf8\xb6\x5c\xb6\xdb\x8f\x0f\x2a\xc1\xd8\xd9\x69\x3f\xfe\xf5\x49\xd5\x26\xe6\x0b\x32\x3b\x65\x4f\xc0\x2b\x40\x12\x3b\xf4\x9f\xd1\xcd\xda\x2f\x44\x04\x46\xbe\x15\x39\x94\x5a\x75\xeb\x4d\x4d\xac\x94\xdd\xf3\xfc\x13\x82\xfb\x4c\xed\xc2\xd6\x4f\x9a\x8b\x45\x3e\x2a\xb7\x2b\xfb\xba\x29\x10\x25\x04\xc8\xa3\x56\x04\x37\x3b\xf9\x3b\xb8\x5f\xd7\x71\xbf\xae\xe1\x7e\x2d\xf8\xfb\xaf\xbf\x8d\xfc\xf5\x3f\x45\x1e\x45\xe2\x65\x75\x89\x91\x8e\x8a\x92\x4f\x71\xa2\xbc\xcb\xf8\x4f\xfd\x2e\x63\x09\xd9\xcb\x5d\xed\x26\xe3\x5e\x5d\x48\x60\x20\xad\x9c\x58\xa2\x83\x9e\xc2\x72\xd2\xe0\xa6\x74\xa8\x23\xbc\x5a\x24\xbf\x33\xda\x6c\xc6\xf2\x73\x46\xb9\xb1\x2f\x10\x8d\xdd\x29\x20\x63\x49\xbd\x29\xba\x1c\xb3\xeb\x2b\xa4\x0b\x08\x13\xba\x27\x70\x32\x72\x1f\xa6\x37\x14\x4e\x51\x7f\xce\xc4\x87\x15\x8a\xbc\x36\xde\x16\x67\xa9\x5d\x44\x2e\x2c\xb1\xa3\xfb\xab\xd7\x56\xd4\xc4\xe2\x98\xd6\x4b\x74\xf2\xa8\x32\x68\x67\xd6\xfd\x41\xa0\x54\x0b\xcb\x39\x01\xa4\x40\xae\xd5\xa6\xbd\xbe\x74\x00\x25\x84\xe1\x81\x0c\xc3\xdf\x08\x70\x08\x2c\xd8\x6d\x1f\x0c\xba\x90\x71\xb3\xde\xa0\xef\x45\xb0\x34\x14\x07\x26\xe0\x4c\x42\x51\xdb\x07\x32\x87\x9e\x46\xda\x4b\x3c\xa0\x1a\xac\xf3\x70\x5a\xe4\xc1\x21\xe7\x05\x6a\x12\xa9\xae\xb0\x02\xfb\x0c\x39\x9f\x80\x5b\xde\x38\x06\x90\xb0\x7e\x00\x91\xf7\xe3\x35\x68\x43\x22\xae\xef\xdd\x20\x76\x11\x41\x82\x3e\x0f\xeb\x43\x56\xf1\xc6\xe5\x0a\x8c\x29\x05\x1e\x84\x5d\xf7\x15\xa4\x04\x61\x9f\x78\xb7\xea\x6a\x52\xda\x1f\x9c\xda\x16\x99\x6f\x55\xec\x76\x63\x48\x9a\x4a\xde\x30\x40\x15\x85\x76\x77\x89\xd7\x84\x0c\x13\xa3\xce\x2a\x6d\x7b\x2b\x23\x48\x19\x1d\x85\x32\x26\x32\xf5\x49\x9f\xe9\x5b\x2e\xf8\x84\x6c\xc7\xe3\x98\xb1\xc4\xb0\xc0\x4b\x02\x37\xc1\x0c\x86\x93\x9d\x9d\x3a\x4e\x40\x41\xac\xbe\x27\x80\xc9\x2b\x20\x1e\x22\x27\x57\x6f\x60\xa8\x80\x32\x81\x1b\xf8\x06\x9b\xb6\x8b\xad\x64\xd5\x67\xb2\xa8\xdb\x57\x34\xdc\x6d\x07\x98\x83\x70\xc5\x7f\x5c\x60\x51\x90\x77\x59\x8f\xf7\x3b\x6b\x32\x5b\x4d\x35\x75\x49\x0e\x5a\xc8\xab\x72\x6d\xab\x55\xc9\x24\xd6\x21\x91\x95\x89\x0f\x19\x5a\x02\xea\x21\x2a\x2b\xa0\x6b\xa0\x6c\x11\x7e\xcd\x28\x8d\xc9\x5a\x92\xa6\x0b\xb9\x67\x96\x31\x98\x3c\x68\x04\xce\x41\x5f\xbc\x6f\x35\x4b\xf0\xce\xd6\x0e\xac\x4d\xcb\xdb\x63\xbc\x2d\x05\xc9\xac\x97\x40\x43\xbb\x02\x2a\x16\x62\x09\x54\xae\x35\x1e\x27\xa2\x87\xcc\x8d\x20\xd3\xc4\xea\xf3\xcb\x3e\x38\x8f\x11\x6a\xa3\x6b\xea\x58\xb0\x00\xfc\x90\x59\x23\xae\x44\xc3\xda\x95\xa8\x73\x3a\x05\x63\x07\x56\x58\x38\xb5\x06\x32\xcc\x91\xea\x97\xa0\xe1\xf1\x06\xf4\x3d\xe6\x30\x03\x54\xd6\xde\x59\x9f\x0e\x84\xb2\xa1\x85\x20\x41\x44\x07\x62\xbb\xd0\xda\x26\xf2\x0b\x11\xf5\x3e\xc5\x02\x1b\x34\x68\xec\xc5\xd2\x8c\xc0\xdf\x44\xa8\x22\x2c\x39\xc6\xcd\x5e\x0b\x4b\xca\xf3\xee\x66\x17\x68\x77\x25\x16\x00\x38\x21\x10\xa9\xb1\xb2\x64\xd1\x50\xcc\x73\x8f\x58\x6d\x0a\xcc\x21\x40\x6c\x20\x95\x7d\xf1\xf9\xb4\x2a\x76\x3e\x03\x02\xd2\x39\xab\x2e\x1e\xd6\x2f\x96\x01\x1c\xda\x8d\xd0\x4f\xea\x57\x0e\xe2\xba\x61\xcb\x14\x21\x3e\x91\xa9\xf5\xaf\xb6\x5c\x8d\x1f\xb3\x7f\xca\x6f\xb9\x72\x9d\xe1\x6b\xdc\xc6\x1b\x0a\x65\xb7\x6a\x0c\xc7\x05\x1d\x26\x2a\xa1\xc8\x3a\x22\xea\x9d\xcc\x5f\x48\x36\x09\xab\x08\x7d\x27\x40\x40\x62\xd1\xea\x84\xa9\xbb\x46\xfd\xbe\xa2\x7a\x90\x00\xdb\x89\xe7\x07\x63\x5d\x8d\x12\xde\x43\x14\x82\x76\x76\x42\x94\xbb\xb7\x40\x9a\x11\xd6\x3b\xcc\x85\x63\xa4\x92\x32\x0f\xed\x4d\x11\x84\x0a\xd1\x3a\xdf\xb7\xf6\xe2\x0d\x38\x08\x49\x46\x3a\x21\x38\x74\xc3\x71\xac\x0e\x59\x4d\x14\x80\x08\xb9\x1e\x7c\x63\xbb\xbb\x01\x68\xc2\x7b\xa4\x69\xd7\x7d\x0b\xeb\x99\xb0\x54\x91\x0f\x8b\xdc\xc8\x78\x3a\x20\x28\xc8\xaf\xc8\xc6\x48\x07\xdd\xe3\x6a\x55\xf1\xea\x88\xd9\xb7\x7b\x68\x31\xcf\xbb\xe7\x60\x70\x3b\xf8\x8b\x3a\xd9\xd4\x69\xb1\x6a\xf6\x73\x56\x59\x1b\x69\x69\xc0\xcb\x62\x98\x27\x29\x8e\x94\x11\x86\xd2\x16\xd1\xbe\x60\x78\x88\xa5\xa5\x6c\x0a\x81\x1d\xc8\x35\x1a\x07\x56\x71\x9e\xe9\xfa\x53\x42\x43\x74\x20\x89\x32\x49\x60\x90\xa4\xa1\x02\x33\xfd\x82\xb9\x09\xca\x3a\xa7\xf8\x24\xa4\x0d\x91\x53\x42\x56\xc6\x4e\x55\xec\x7c\xc1\x6c\x67\x00\x01\x80\x60\xc4\x4c\xa5\xc3\xe0\xf4\xa7\x42\x04\xaa\xea\x44\xa4\x2d\x7e\x4d\x2a\x9e\x29\x30\x22\x84\x1d\x6b\x5f\x27\xce\xda\x9c\x2b\x8f\x12\x49\xa7\x08\x2e\x87\x10\x48\x91\x23\xda\x73\x16\xe2\x59\x8f\xe3\x39\xc2\xa8\x38\xd5\xf5\x4e\x4c\x53\x8f\x43\xce\x6c\xae\x6d\x03\x24\x16\xf2\xa4\x49\x63\x30\xca\x24\x16\x7f\xa4\x1c\xa9\x00\xc5\x24\xd7\x34\x5e\x61\x54\x14\x09\x0b\x74\x0a\x0e\x53\x86\xec\xb1\x22\x7e\xd3\xdd\x7f\xc2\xab\x84\x4b\x87\x96\xdc\x8e\xca\x09\x78\x9e\x8e\xa2\x31\xb8\x6c\x59\xb5\x36\x8b\x44\xb4\xce\xad\x00\x7c\x2d\xb1\x1f\x80\x6f\xc1\x04\x03\x96\x25\x7a\x03\xba\xb6\x01\xaf\x27\x01\xfb\x80\x72\x4f\x9e\x1f\x8c\x5f\xbd\xf4\x10\x24\x41\x42\xc1\xd5\x30\xec\x7a\x8b\xfb\x73\x19\xa6\x24\xc4\xba\x0f\x53\x7a\x34\x36\x61\x01\xa0\xd5\x8c\xfc\xc1\xce\x0e\xb8\xec\x71\x30\x06\x7f\xfb\x7d\x0d\x30\xe5\xf8\xd0\x40\xd4\xe2\x28\x88\x2a\xb6\x30\x20\x82\x30\x17\x6f\x20\x20\xa1\x5b\xac\x10\xfe\x18\xc4\x4f\x42\x1a\x1f\x70\xfb\xd9\xc2\x3a\xb8\x47\x02\x42\x0d\x56\xdb\x82\xa5\xef\xdb\xf5\x87\x90\x3d\xc9\x23\x19\x91\x11\xeb\x79\xce\x7b\x11\x97\x52\xf4\x25\x81\xa5\x6e\xe6\x55\xd9\xeb\x2e\x93\xae\x1d\xd1\x3f\x65\xd4\xf9\xf3\x4f\xed\x6c\x85\x30\x15\x7f\x26\xd3\x87\x1f\x11\x55\xbe\xf6\x50\xdb\x33\x4c\x10\xcc\xa5\x54\xc0\x7a\xa7\xa2\xa8\x29\xfe\x56\x71\x89\x75\x35\xef\x0f\x01\x60\x2a\x6d\x12\x9a\xc8\xaa\x7e\x87\x4b\xac\x88\xe3\x02\x9c\xcb\xc9\x46\xfa\x3a\x52\xfe\xd1\xea\x7a\x23\x9f\xe5\x89\xef\x53\x19\x6b\x7f\x75\x2f\xbc\x13\x12\x5c\x58\x8f\x40\x5e\xf6\x51\x8e\x2f\xea\x71\x75\xcd\x32\x6e\x09\xa3\xde\xa9\x42\xa1\x6a\xaf\x6a\xeb\xb7\x18\xd7\x2d\x20\xc0\xd7\xac\xc1\x78\x64\x03\x89\x47\xc5\x0c\x85\x5f\x80\xa9\x25\x5d\xea\x7e\x40\x06\x72\xa7\x22\xd2\x54\x26\x0e\x04\x1e\xd5\x34\xf2\x4f\x49\x28\x25\x3a\x22\x76\xd1\xee\x8d\x30\x46\x08\x1e\xec\xd1\x03\x97\xf4\x4c\xd0\x11\x8b\xc5\x91\x95\x61\xdb\x92\xf2\x9c\xb0\xcd\x6a\x5f\x54\xb9\x20\xac\x2c\x6b\x28\x78\x2b\x1c\x88\x7b\x6b\xae\x1e\x22\x01\xb1\x21\xf0\x45\x4b\x4d\xaa\x47\x3e\xa1\xa5\x62\x4c\xa9\xd8\x20\x18\x80\x8a\xa1\x1e\xae\xd5\x0d\x31\x11\xe1\xf4\x12\xed\x31\xea\x16\x03\xa7\x8a\xe5\x13\x8c\x94\x36\x67\x47\x1b\x53\x3d\x39\xd5\x74\xb7\x65\xb7\xbc\x0c\x91\x48\x0d\x50\x41\x05\x1e\x83\xca\x4b\x0c\x40\x41\xff\x0a\x1b\x05\xab\xba\x65\xb0\xb2\x84\xcb\x8a\xf5\xe8\xa4\x20\x60\xec\x99\xe8\x41\x06\xc5\x61\x65\xe7\x6d\x77\x15\x54\x71\x35\xd2\xae\x1e\xb5\xb4\x5a\x89\xb9\x3b\xc1\xe4\x0e\xe0\x00\xda\xcd\x01\xba\xad\x81\x70\x62\xea\x85\x03\x84\x8c\x09\x98\x92\x01\xfa\xa2\x81\xf0\x45\x03\xf4\x45\x90\xbb\x40\x04\x31\x10\xf5\x54\xf4\x6a\x63\xa9\x58\x0b\x40\x13\x0e\x62\x1e\xbf\x81\xd3\x03\x93\x1e\xfb\x2f\x57\x17\x7f\x91\x60\xa8\x05\x52\x27\x21\xf8\x31\x97\xae\x46\xcc\x0d\x45\x3e\x6c\xe4\xe7\x39\x2f\x78\x69\xf5\xae\x3e\x58\xbb\x1d\x49\xb7\x84\xa7\x79\x85\x1e\xfc\x43\x4d\xe5\x60\xe1\xd6\x64\x78\xe3\xd9\x2a\xc8\xfe\x2b\xf0\xf8\x41\xbc\x5c\x72\x88\x53\xa5\x79\x48\x75\xb6\x7d\x41\xc3\xc0\x4e\xfe\x8e\x56\xb5\x7d\x5e\xd5\x34\x51\x2f\xb5\x53\x78\x66\x06\x5e\x68\x98\xcf\xa5\xf6\xcd\xa9\x9a\x2f\x8f\x2a\x0d\x92\x59\xf8\xd1\xba\x9e\x11\xcf\x87\xd4\x62\x51\x11\xa1\xe1\xea\xe3\x96\xa7\xbd\x5b\x6a\x3b\x66\xbd\x4c\x99\xde\x48\x45\xfc\xe8\x3a\x17\xc7\x1f\x8f\xdf\x1f\x3b\x5e\xfb\xf8\x57\xe2\xbd\x35\xdd\xa7\x6f\x4f\xce\x1c\xef\xf7\x7d\x48\x9c\xce\x4c\xe7\xd1\xd9\xdb\x93\xd3\xe7\x8e\xf7\x5f\xe8\x0e\x86\xf4\x35\xab\xd5\x19\x41\xec\xb1\x04\xff\xd0\xe6\x2f\x56\x72\xce\xb9\x70\xaf\x0f\xcd\xca\x71\xd6\x5f\x52\xf3\x5c\x3b\x13\x48\x26\x8d\x2f\xf1\xe5\xbd\x84\x9f\xeb\x1e\x9c\x05\xad\x60\xc8\x5c\xe7\x7d\x96\x95\x8d\x34\x1b\x0c\xb8\x78\x96\xd3\x98\x66\x8d\x14\xfc\x57\xda\x00\xf9\xf0\xab\xc7\x08\xf2\xe2\x75\xe8\xc3\x4c\xba\xfe\xca\x49\x00\x17\x64\x7f\x42\xcf\xa5\x91\x26\xb2\x2d\x8c\x03\x3a\x2d\x11\xda\x87\x98\xfe\x48\x9a\xa1\x24\x57\xa9\x80\x24\x03\xa4\x04\x98\x0b\x80\xd8\x82\xc7\x8e\xa8\x03\x5b\x75\x20\x98\xf5\x2f\xbc\x54\xbc\x80\xcc\xf0\x71\xb1\x6b\xbe\xfd\x72\x34\xc1\xa7\x84\x93\x59\x77\x4b\x1f\x18\xf2\x4e\xd5\x3d\x61\xf9\xf8\x12\x86\xd2\xd1\x94\xe3\x03\x90\x6d\xfd\xb0\x82\xc0\x4e\x93\xe2\x2a\x1f\x95\x58\x9b\x4b\x46\x29\xcf\xdf\xc0\x0c\x5c\xb1\xa5\xdb\x15\x68\x0a\x8f\x19\x05\x68\xa3\x22\x7d\x41\x23\xaf\x9e\x9f\x13\x2b\xbf\x07\x33\xa4\xae\xd8\xf1\xda\xf7\xb9\x08\x1d\xd1\x3e\xba\x31\x11\x0e\xc8\xf0\x1f\xeb\xd5\xc0\xe9\xd1\x34\xc9\x6a\x84\xd6\xec\xce\x06\xee\x5b\x86\x3e\x50\x15\xfc\xde\x61\xb9\xcb\xbb\x92\xaa\x5e\x19\xcf\xf7\x42\x10\xae\xc0\x6e\x5f\x49\x3f\x0f\xac\x76\x88\xf7\x8e\xf5\x1c\xa7\x4f\xaf\xbc\x2b\xff\x95\x7b\x86\xa9\x71\xed\x95\xc8\x3b\xf3\x4a\x44\xad\x61\xda\x1d\x82\x10\xb1\xa2\x3c\xd5\xef\x75\xe4\x13\x11\x08\xd1\xe7\x21\x78\x44\x37\x6a\xb5\x91\x16\xb8\xa9\xee\xc2\x18\x0e\x93\xec\x39\xa0\x00\xbf\x01\x49\xd1\xea\xc5\x7d\x30\x18\x90\x4f\x01\xb5\xc4\x6e\xf6\xb3\x86\x4a\x8e\x3f\xc9\x23\x63\x46\x27\x48\x01\xf1\x99\x24\x9f\x6d\x0e\xaf\xe5\xa4\x8b\x8d\xd0\xe6\xad\x75\x95\x02\x51\x9a\xbe\xfc\x3c\xa4\xe1\x72\x79\xa1\xcc\xc4\x57\x69\x55\x0e\xd5\x5b\x52\x7c\x95\x78\xed\x5d\xa0\xf6\x5e\xd7\x94\x57\xc5\xf6\x43\x9f\xeb\x8f\x2f\x9a\x93\x43\xff\x16\x63\xf3\xa1\x3f\xa1\xd8\x18\x56\xe6\xf6\x92\xd9\x50\xbc\x0f\x0c\x49\xe3\x0c\xb2\x0c\x5c\x08\xa8\xd5\x15\x0f\x2f\xb2\x68\xcc\xf1\xc1\xd3\x25\xf3\x63\xfa\x81\x59\xd1\x9a\x9d\xe0\xc9\x97\x36\xa3\x29\x44\x91\xbf\xca\x28\x72\x96\x2d\xdc\xc7\x98\x38\xfe\xe7\xf8\x37\xb2\xaa\x63\x0b\xfa\x99\xcd\xf8\x74\x4d\x6e\x64\xd2\x22\x8e\x39\x14\xc2\x9f\x72\x96\xa3\xd4\x67\xf3\x52\x9e\x7e\xa8\xe8\x36\xb4\x8a\xbc\x68\x4f\xe5\xe8\x17\xc8\x97\xba\xee\x27\xd9\x88\x3d\x7c\x8e\x36\x15\x6f\xf0\x86\xbc\x61\x4e\xd2\x00\xcc\x9d\x96\x5c\xd9\x72\x1a\x8b\x51\x39\x6c\x08\xc4\xa2\x2c\xd5\x03\x5f\x94\xa6\xcb\xbc\xc8\x2c\x95\x80\x65\x92\x04\x73\x48\xe7\xef\xec\xf5\x43\x90\xe6\x4a\xce\xcf\xa6\x82\x30\x33\x39\x70\xed\x59\xc2\x82\x83\x18\xb8\x73\x3d\x7a\xb9\x3e\xaa\x1f\x1a\xab\xf1\x4f\xeb\xe3\x1c\xeb\x53\x7a\xf4\xa3\x1c\x15\xb5\x7f\x09\xd6\x32\xe8\xff\x8c\x15\x68\x16\x2b\x92\x1c\x01\xd0\x0d\x92\xa0\x3e\x4a\xef\x87\x21\x8f\xc2\x4c\x6c\xef\x56\x14\x12\xd7\xbb\x78\xef\xc5\xa7\xf1\xa6\x23\x89\x44\x3f\x56\x1b\x60\xca\xb5\x8d\x75\xb5\xb7\xc5\x02\x20\x27\x8f\x6b\x9c\x50\x8f\x1f\xec\x30\xdc\x75\x62\xfd\x20\x60\xa2\xef\xda\x37\x14\x0f\x29\x75\x59\xc3\xa7\xda\xf0\xf2\x07\x32\x26\x8e\x17\xfb\xce\xf6\x8d\x99\xee\x8f\xcc\xcd\xbc\x24\x10\x79\x08\xba\x84\x07\x0e\x30\x9f\x80\x3d\xb9\x37\x24\x95\x8a\xaf\xc4\x53\xeb\xbe\xf5\x60\x5d\x42\xc3\x27\x7d\xa1\x30\xcf\x6f\xd8\x3a\x64\x40\x0d\x3c\xce\x94\x47\x25\xc0\x9f\x4f\xf9\xdd\x4c\x7c\xa6\xf7\xe2\xd5\xb2\x1f\xb3\x92\x69\xf3\x66\xf0\x7c\x6b\x93\xe1\x82\xc3\xfa\xb8\x80\xb5\xe5\x28\x6d\x00\x80\xb2\x91\x73\x05\xb2\xc1\xca\x92\x4f\x66\x25\xc2\x12\x16\x22\x49\x33\xf0\x3b\x92\xd6\x7b\x18\xbd\xe0\x6d\x86\x92\x4e\x64\x9b\x38\x89\x57\x57\xba\x40\x1d\xe3\x6b\xa0\x13\x51\x6c\x8a\xd2\x1c\xfe\x53\x86\xcc\x8d\xab\xec\x66\x67\x67\xcb\x23\xa3\x5a\x42\x4a\xc4\x42\xbb\x47\x57\xdd\x1a\x5b\x8b\xba\xa9\xca\xfb\x1a\xfa\x05\x2e\xb0\xef\x91\xd2\x84\xc7\xed\xdf\xfe\xfb\xdb\xef\xbf\xfe\xe7\xb7\xff\xca\xd7\xe8\x29\x48\x6a\xa9\xb5\x06\x4c\xdc\x72\xa9\x6f\x9f\x27\x10\xa6\x6e\x17\xb2\x95\xd6\x8e\x36\x8a\xf7\xa7\x4d\xf1\xaf\x8b\x8f\x88\xeb\xd0\x9b\x21\x6f\x88\x50\x89\x8f\xb5\x35\x78\xdb\x82\x63\xc1\xc3\x32\x70\x08\x59\x00\x1e\xb2\x91\x45\x22\x05\x8a\x05\xb7\x0d\xa3\x1f\xd8\xf5\x8a\x89\xff\x64\xa8\xfc\xc1\x47\xe1\x0f\xec\xd7\x36\x4e\x64\x94\x4a\x59\x26\x06\xde\xeb\x23\xc3\x3b\x6e\xb3\xec\x6a\x73\x59\x68\x54\x05\x10\x17\x6b\xae\xea\x6b\x3e\x9b\x2a\x34\x62\xa2\xde\x1a\x30\x98\x1a\x60\x5d\x1f\xb6\x95\xe2\x72\xa5\xce\xa0\xed\x60\x7e\xff\x9d\x09\xc9\x72\x9d\x45\xd1\xd9\xdb\x73\x5a\x29\x10\x21\x92\xb7\xe6\xc3\xac\x28\x31\xbc\x6e\x39\x9d\xdf\xf7\x7f\xdf\xdf\x5b\xe0\xbf\x38\x98\xeb\x62\x1d\x88\x21\xcd\xf0\x7f\x3b\xe8\x91\xeb\xc8\xf7\x09\x80\xab\xdc\x34\x1d\x45\x63\xb5\xf1\x59\x7d\x63\x31\x5b\x98\xd0\x6d\x93\xcf\x6b\x93\x57\x9f\xfd\x77\xb6\x4d\xfb\xec\xbf\x87\xc0\x04\x7e\x8b\xa3\x7e\x06\x82\x7f\xb6\xf2\x85\xb3\x8d\x17\x55\x8c\x49\xf3\x98\xe2\xbb\x8a\xd1\x4c\x04\xf4\x68\xba\xec\x55\xe7\x9b\xef\x8e\x01\x41\x10\xd2\xd9\xbc\x94\x2f\xfa\x9c\x12\x54\xf7\x28\x9b\x96\x28\xe0\x78\xd9\x85\x95\x81\xaa\x8b\x3a\xd5\xbf\xbe\xfd\x8a\x37\x58\xf8\x7a\xe7\x12\x40\x13\x29\x6c\xf6\x38\xf3\x93\x51\x5e\x94\x47\xc3\x51\x1a\xef\xec\x88\xd9\x55\x47\xb5\x50\xdc\x31\x05\x32\x7a\x13\x43\xcd\xda\xc4\x80\x30\x55\x7d\x11\x4d\xd7\x9a\x48\x82\x1a\x44\xb5\xbf\x78\x1d\x64\x2e\xae\x20\x92\xfb\x21\xb8\x10\x81\xb0\x19\x88\x46\x2c\x3b\xdc\x3f\xec\x63\x41\x6c\x06\x82\xb3\x00\xb5\x7f\xa6\x1e\x6d\x62\xd9\x57\x3f\xe0\x24\x7e\x94\x73\x56\xf2\x4b\x20\x10\x16\x1d\x31\x80\xc5\xe0\xd8\x26\xf9\xd5\x66\xa1\xa7\x46\xf3\x10\xa2\x47\x00\xff\xe2\xf2\xcd\xeb\x96\x78\x4e\x22\xf4\x05\xff\x17\x25\xca\xb3\x34\xbd\xcc\x66\x54\x7f\xbf\xe0\xa3\xc1\xb0\x44\xf0\xef\x20\xc9\x7d\xf4\xff\x03\x00\x00\xff\xff\x55\x5b\xf1\xf7\xa9\x3a\x00\x00")

func static_all_js_bytes() ([]byte, error) {
	return bindata_read(
		_static_all_js,
		"static/all.js",
	)
}

func static_all_js() (*asset, error) {
	bytes, err := static_all_js_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "static/all.js", size: 15017, mode: os.FileMode(436), modTime: time.Unix(1422161500, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/homepage.html.tmpl": templates_homepage_html_tmpl,
	"templates/script.html.tmpl":   templates_script_html_tmpl,
	"static/all.css":               static_all_css,
	"static/all.js":                static_all_js,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func     func() (*asset, error)
	Children map[string]*_bintree_t
}

var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"static": &_bintree_t{nil, map[string]*_bintree_t{
		"all.css": &_bintree_t{static_all_css, map[string]*_bintree_t{}},
		"all.js":  &_bintree_t{static_all_js, map[string]*_bintree_t{}},
	}},
	"templates": &_bintree_t{nil, map[string]*_bintree_t{
		"homepage.html.tmpl": &_bintree_t{templates_homepage_html_tmpl, map[string]*_bintree_t{}},
		"script.html.tmpl":   &_bintree_t{templates_script_html_tmpl, map[string]*_bintree_t{}},
	}},
}}

// Restore an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	if err != nil { // File
		return RestoreAsset(dir, name)
	} else { // Dir
		for _, child := range children {
			err = RestoreAssets(dir, path.Join(name, child))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
