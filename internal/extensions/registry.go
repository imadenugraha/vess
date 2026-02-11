package extensions

var registry = map[string]*Extension{
	// Core extensions
	"mysqli": {
		Name:        "mysqli",
		Description: "MySQL Improved Extension",
		PHPVersions: []string{"7.4", "8.0", "8.1", "8.2", "8.3"},
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("mysqli"),
			"ubuntu": GetUbuntuSupport("mysqli"),
		},
		Conflicts: []string{},
	},
	"pdo_mysql": {
		Name:        "pdo_mysql",
		Description: "MySQL PDO Driver",
		PHPVersions: []string{"7.4", "8.0", "8.1", "8.2", "8.3"},
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("pdo_mysql"),
			"ubuntu": GetUbuntuSupport("pdo_mysql"),
		},
		Conflicts: []string{},
	},
	"pdo_pgsql": {
		Name:        "pdo_pgsql",
		Description: "PostgreSQL PDO Driver",
		PHPVersions: []string{"7.4", "8.0", "8.1", "8.2", "8.3"},
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("pdo_pgsql"),
			"ubuntu": GetUbuntuSupport("pdo_pgsql"),
		},
		Conflicts: []string{},
	},
	"pgsql": {
		Name:        "pgsql",
		Description: "PostgreSQL Extension",
		PHPVersions: []string{"7.4", "8.0", "8.1", "8.2", "8.3"},
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("pgsql"),
			"ubuntu": GetUbuntuSupport("pgsql"),
		},
		Conflicts: []string{},
	},
	"gd": {
		Name:        "gd",
		Description: "GD Graphics Library",
		PHPVersions: []string{"7.4", "8.0", "8.1", "8.2", "8.3"},
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("gd"),
			"ubuntu": GetUbuntuSupport("gd"),
		},
		Conflicts:     []string{},
		ConfigureArgs: []string{"--with-freetype", "--with-jpeg"},
	},
	"opcache": {
		Name:        "opcache",
		Description: "Zend OPcache",
		PHPVersions: []string{"7.4", "8.0", "8.1", "8.2", "8.3"},
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("opcache"),
			"ubuntu": GetUbuntuSupport("opcache"),
		},
		Conflicts: []string{},
	},
	"zip": {
		Name:        "zip",
		Description: "Zip Archive Extension",
		PHPVersions: []string{"7.4", "8.0", "8.1", "8.2", "8.3"},
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("zip"),
			"ubuntu": GetUbuntuSupport("zip"),
		},
		Conflicts: []string{},
	},
	"intl": {
		Name:        "intl",
		Description: "Internationalization Extension",
		PHPVersions: []string{"7.4", "8.0", "8.1", "8.2", "8.3"},
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("intl"),
			"ubuntu": GetUbuntuSupport("intl"),
		},
		Conflicts: []string{},
	},
	"bcmath": {
		Name:        "bcmath",
		Description: "BC Math Extension",
		PHPVersions: []string{"7.4", "8.0", "8.1", "8.2", "8.3"},
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("bcmath"),
			"ubuntu": GetUbuntuSupport("bcmath"),
		},
		Conflicts: []string{},
	},
	"exif": {
		Name:        "exif",
		Description: "EXIF Extension",
		PHPVersions: []string{"7.4", "8.0", "8.1", "8.2", "8.3"},
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("exif"),
			"ubuntu": GetUbuntuSupport("exif"),
		},
		Conflicts: []string{},
	},
	"pcntl": {
		Name:        "pcntl",
		Description: "Process Control Extension",
		PHPVersions: []string{"7.4", "8.0", "8.1", "8.2", "8.3"},
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("pcntl"),
			"ubuntu": GetUbuntuSupport("pcntl"),
		},
		Conflicts: []string{},
	},
	"soap": {
		Name:        "soap",
		Description: "SOAP Extension",
		PHPVersions: []string{"7.4", "8.0", "8.1", "8.2", "8.3"},
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("soap"),
			"ubuntu": GetUbuntuSupport("soap"),
		},
		Conflicts: []string{},
	},
	"sockets": {
		Name:        "sockets",
		Description: "Sockets Extension",
		PHPVersions: []string{"7.4", "8.0", "8.1", "8.2", "8.3"},
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("sockets"),
			"ubuntu": GetUbuntuSupport("sockets"),
		},
		Conflicts: []string{},
	},
	"xmlrpc": {
		Name:        "xmlrpc",
		Description: "XML-RPC Extension",
		PHPVersions: []string{"7.4", "8.0"}, // Removed in PHP 8.1+
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("xmlrpc"),
			"ubuntu": GetUbuntuSupport("xmlrpc"),
		},
		Conflicts: []string{},
	},
	"xsl": {
		Name:        "xsl",
		Description: "XSL Extension",
		PHPVersions: []string{"7.4", "8.0", "8.1", "8.2", "8.3"},
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("xsl"),
			"ubuntu": GetUbuntuSupport("xsl"),
		},
		Conflicts: []string{},
	},
	// PECL extensions
	"redis": {
		Name:        "redis",
		Description: "Redis Extension (PECL)",
		PHPVersions: []string{"7.4", "8.0", "8.1", "8.2", "8.3"},
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("redis"),
			"ubuntu": GetUbuntuSupport("redis"),
		},
		Conflicts: []string{},
	},
	"imagick": {
		Name:        "imagick",
		Description: "ImageMagick Extension (PECL)",
		PHPVersions: []string{"7.4", "8.0", "8.1", "8.2", "8.3"},
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("imagick"),
			"ubuntu": GetUbuntuSupport("imagick"),
		},
		Conflicts: []string{},
	},
	"memcached": {
		Name:        "memcached",
		Description: "Memcached Extension (PECL)",
		PHPVersions: []string{"7.4", "8.0", "8.1", "8.2", "8.3"},
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("memcached"),
			"ubuntu": GetUbuntuSupport("memcached"),
		},
		Conflicts: []string{},
	},
	"mongodb": {


		Name:        "mongodb",
		Description: "MongoDB Extension (PECL)",
		PHPVersions: []string{"7.4", "8.0", "8.1", "8.2", "8.3"},
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("mongodb"),
			"ubuntu": GetUbuntuSupport("mongodb"),
		},
		Conflicts: []string{},
	},
	"xdebug": {
		Name:        "xdebug",
		Description: "Xdebug Debugging Extension (PECL)",
		PHPVersions: []string{"7.4", "8.0", "8.1", "8.2", "8.3"},
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("xdebug"),
			"ubuntu": GetUbuntuSupport("xdebug"),
		},
		Conflicts: []string{},
	},
	"apcu": {
		Name:        "apcu",
		Description: "APCu Cache Extension (PECL)",
		PHPVersions: []string{"7.4", "8.0", "8.1", "8.2", "8.3"},
		OSSupport: map[string]*OSSupport{
			"alpine": GetAlpineSupport("apcu"),
			"ubuntu": GetUbuntuSupport("apcu"),
		},
		Conflicts: []string{},
	},
}

// GetRegistry returns the complete extension registry
func GetRegistry() map[string]*Extension {
	return registry
}

// GetExtension returns a specific extension by name
func GetExtension(name string) (*Extension, bool) {
	ext, exists := registry[name]
	return ext, exists
}

// GetAllExtensionNames returns all available extension names
func GetAllExtensionNames() []string {
	names := make([]string, 0, len(registry))
	for name := range registry {
		names = append(names, name)
	}
	return names
}

// SupportsVersion checks if an extension supports a PHP version
func SupportsVersion(extName, phpVersion string) bool {
	ext, exists := GetExtension(extName)
	if !exists {
		return false
	}
	
	for _, v := range ext.PHPVersions {
		if v == phpVersion {
			return true
		}
	}
	return false
}

// SupportsOS checks if an extension supports an OS
func SupportsOS(extName, osType string) bool {
	ext, exists := GetExtension(extName)
	if !exists {
		return false
	}
	
	_, supported := ext.OSSupport[osType]
	return supported
}
