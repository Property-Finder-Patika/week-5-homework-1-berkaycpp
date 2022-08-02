package main

import "fmt"

/*****************************************************************************/
// base license manager class that will derive both real and proxy managers
type licenseManager interface {
	// real manager methods
	checkAvailableLicense() bool
	pickLicense()
	dropLicense()
}

/*****************************************************************************/

/*****************************************************************************/
// proxy class with real manager and attributes listed if there are
type ProxyLicenseManager struct {
	realLicenseManager *RealLicenseManager
	version            int
}

// constructor of the proxy class
func newProxyLicenseManager(_realLicenseManager *RealLicenseManager) *ProxyLicenseManager {
	return &ProxyLicenseManager{
		realLicenseManager: _realLicenseManager,
		version:            _realLicenseManager.version,
	}
}

// overriding the functions
func (lm *ProxyLicenseManager) checkAvailableLicense() bool {
	fmt.Println("Checking availability of the licenses by proxy...")
	if !lm.realLicenseManager.checkAvailableLicense() {
		fmt.Println("No available license!")
		return false
	} else {
		fmt.Println("Available license exists.")
		return true
	}
}

func (lm *ProxyLicenseManager) pickLicense() {
	fmt.Println("Checking availability of the licenses by proxy...")
	if !lm.realLicenseManager.checkAvailableLicense() {
		fmt.Println("No available license!")
		return
	} else {
		lm.realLicenseManager.pickLicense()
		fmt.Println("The license is successfully reserved.")
	}
}

func (lm *ProxyLicenseManager) dropLicense() {
	lm.realLicenseManager.dropLicense()
	fmt.Println("The license in use is dropped.")
}

/*****************************************************************************/

/*****************************************************************************/
// real license manager class
// these are the real license manager functions, only signatures are shown to eliminate errors
type RealLicenseManager struct {
	version int
}

// overriding the functions
func (lm *RealLicenseManager) checkAvailableLicense() bool {
	return false
}

func (lm *RealLicenseManager) pickLicense() {

}

func (lm *RealLicenseManager) dropLicense() {

}

/*****************************************************************************/

/*****************************************************************************/
// the main flow
func main() {
	var appLicenseManager = RealLicenseManager{1001}
	appLicenseManagerProxy := newProxyLicenseManager(&appLicenseManager)

	appLicenseManagerProxy.checkAvailableLicense()
}

/*****************************************************************************/
