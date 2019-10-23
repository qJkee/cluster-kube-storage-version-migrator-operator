// Code generated by go-bindata.
// sources:
// bindata/kube-storage-version-migrator/deployment.yaml
// bindata/kube-storage-version-migrator/namespace.yaml
// bindata/kube-storage-version-migrator/roles.yaml
// bindata/kube-storage-version-migrator/serviceaccount.yaml
// DO NOT EDIT!

package assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _kubeStorageVersionMigratorDeploymentYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: migrator
  namespace: openshift-kube-storage-version-migrator
  labels:
    app: migrator
spec:
  replicas: 1
  selector:
    matchLabels:
      app: migrator
  template:
    metadata:
      labels:
        app: migrator
    spec:
      serviceAccountName: kube-storage-version-migrator-sa
      containers:
      - name: migrator
        image: quay.io/sanchezl/storage-version-migration-migrator:v0.1
        command:
          - /migrator
          - '--alsologtostderr'
        args:
          - '--v=5'
        terminationMessagePolicy: FallbackToLogsOnError
`)

func kubeStorageVersionMigratorDeploymentYamlBytes() ([]byte, error) {
	return _kubeStorageVersionMigratorDeploymentYaml, nil
}

func kubeStorageVersionMigratorDeploymentYaml() (*asset, error) {
	bytes, err := kubeStorageVersionMigratorDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kube-storage-version-migrator/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _kubeStorageVersionMigratorNamespaceYaml = []byte(`apiVersion: v1
kind: Namespace
metadata:
  name: openshift-kube-storage-version-migrator
  labels:
    openshift.io/run-level: "1"
    openshift.io/cluster-monitoring: "true"`)

func kubeStorageVersionMigratorNamespaceYamlBytes() ([]byte, error) {
	return _kubeStorageVersionMigratorNamespaceYaml, nil
}

func kubeStorageVersionMigratorNamespaceYaml() (*asset, error) {
	bytes, err := kubeStorageVersionMigratorNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kube-storage-version-migrator/namespace.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _kubeStorageVersionMigratorRolesYaml = []byte(`kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: storage-version-migration-migrator
rules:
- apiGroups: ["*"]
  resources: ["*"]
  verbs: ["get", "list", "update"]
- apiGroups: ["migration.k8s.io"]
  resources: ["storageversionmigrations"]
  verbs: ["watch"]
---
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: storage-version-migration-crd-creator
rules:
- apiGroups: ["apiextensions.k8s.io"]
  resources: ["customresourcedefinitions"]
  verbs: ["create", "delete", "get"]
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: storage-version-migration-migrator
subjects:
- kind: ServiceAccount
  name: kube-storage-version-migrator-sa
  namespace: openshift-kube-storage-version-migrator
roleRef:
  kind: ClusterRole
  name: storage-version-migration-migrator
  apiGroup: rbac.authorization.k8s.io
---
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: storage-version-migration-crd-creator
subjects:
- kind: ServiceAccount
  name: kube-storage-version-migrator-sa
  namespace: openshift-kube-storage-version-migrator
roleRef:
  kind: ClusterRole
  name: storage-version-migration-crd-creator
  apiGroup: rbac.authorization.k8s.io
`)

func kubeStorageVersionMigratorRolesYamlBytes() ([]byte, error) {
	return _kubeStorageVersionMigratorRolesYaml, nil
}

func kubeStorageVersionMigratorRolesYaml() (*asset, error) {
	bytes, err := kubeStorageVersionMigratorRolesYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kube-storage-version-migrator/roles.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _kubeStorageVersionMigratorServiceaccountYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: openshift-apiserver
  name: kube-storage-version-migrator-sa
`)

func kubeStorageVersionMigratorServiceaccountYamlBytes() ([]byte, error) {
	return _kubeStorageVersionMigratorServiceaccountYaml, nil
}

func kubeStorageVersionMigratorServiceaccountYaml() (*asset, error) {
	bytes, err := kubeStorageVersionMigratorServiceaccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "kube-storage-version-migrator/serviceaccount.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
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

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
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
	"kube-storage-version-migrator/deployment.yaml":     kubeStorageVersionMigratorDeploymentYaml,
	"kube-storage-version-migrator/namespace.yaml":      kubeStorageVersionMigratorNamespaceYaml,
	"kube-storage-version-migrator/roles.yaml":          kubeStorageVersionMigratorRolesYaml,
	"kube-storage-version-migrator/serviceaccount.yaml": kubeStorageVersionMigratorServiceaccountYaml,
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
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"kube-storage-version-migrator": {nil, map[string]*bintree{
		"deployment.yaml":     {kubeStorageVersionMigratorDeploymentYaml, map[string]*bintree{}},
		"namespace.yaml":      {kubeStorageVersionMigratorNamespaceYaml, map[string]*bintree{}},
		"roles.yaml":          {kubeStorageVersionMigratorRolesYaml, map[string]*bintree{}},
		"serviceaccount.yaml": {kubeStorageVersionMigratorServiceaccountYaml, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
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

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}