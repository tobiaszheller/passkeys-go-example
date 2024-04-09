package server

// aaguidMapping contains mapping 'aaguid' to 'name'. Some rnties are missing, taken from
// https://github.com/passkeydeveloper/passkey-authenticator-aaguids/blob/main/aaguid.json
var aaguidMapping = map[string]string{
	"ea9b8d66-4d01-1d21-3ce4-b6b48cb575d4": "Google Password Manager",
	"adce0002-35bc-c60a-648b-0b25f1f05503": "Chrome on Mac",
	"08987058-cadc-4b81-b6e1-30de50dcbe96": "Windows Hello",
	"9ddd1817-af5a-4672-a2b9-3e3dd95000a9": "Windows Hello",
	"6028b017-b1d4-4c02-b4b3-afcdafc96bb2": "Windows Hello",
	"dd4ec289-e01d-41c9-bb89-70fa845d4bf2": "iCloud Keychain (Managed)",
	"fbfc3007-154e-4ecc-8c0b-6e020557d7bd": "iCloud Keychain",
}
