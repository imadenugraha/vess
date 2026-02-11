package extensions

// GetUbuntuSupport returns Ubuntu-specific installation information for an extension
func GetUbuntuSupport(extName string) *OSSupport {
	ubuntuSupport := map[string]*OSSupport{
		"mysqli": {
			BuildDeps:   []string{},
			RuntimeDeps: []string{},
			InstallCmd:  "docker-php-ext-install mysqli",
			PECLInstall: false,
		},
		"pdo_mysql": {
			BuildDeps:   []string{},
			RuntimeDeps: []string{},
			InstallCmd:  "docker-php-ext-install pdo_mysql",
			PECLInstall: false,
		},
		"pdo_pgsql": {
			BuildDeps:   []string{"libpq-dev"},
			RuntimeDeps: []string{"libpq5"},
			InstallCmd:  "docker-php-ext-install pdo_pgsql",
			PECLInstall: false,
		},
		"pgsql": {
			BuildDeps:   []string{"libpq-dev"},
			RuntimeDeps: []string{"libpq5"},
			InstallCmd:  "docker-php-ext-install pgsql",
			PECLInstall: false,
		},
		"gd": {
			BuildDeps:   []string{"libfreetype6-dev", "libjpeg62-turbo-dev", "libpng-dev", "libwebp-dev"},
			RuntimeDeps: []string{"libfreetype6", "libjpeg62-turbo", "libpng16-16", "libwebp7"},
			InstallCmd:  "docker-php-ext-configure gd --with-freetype --with-jpeg --with-webp && docker-php-ext-install gd",
			PECLInstall: false,
		},
		"opcache": {
			BuildDeps:   []string{},
			RuntimeDeps: []string{},
			InstallCmd:  "docker-php-ext-install opcache",
			PECLInstall: false,
		},
		"zip": {
			BuildDeps:   []string{"libzip-dev"},
			RuntimeDeps: []string{"libzip4"},
			InstallCmd:  "docker-php-ext-install zip",
			PECLInstall: false,
		},
		"intl": {
			BuildDeps:   []string{"libicu-dev"},
			RuntimeDeps: []string{"libicu70"},
			InstallCmd:  "docker-php-ext-install intl",
			PECLInstall: false,
		},
		"bcmath": {
			BuildDeps:   []string{},
			RuntimeDeps: []string{},
			InstallCmd:  "docker-php-ext-install bcmath",
			PECLInstall: false,
		},
		"exif": {
			BuildDeps:   []string{},
			RuntimeDeps: []string{},
			InstallCmd:  "docker-php-ext-install exif",
			PECLInstall: false,
		},
		"pcntl": {
			BuildDeps:   []string{},
			RuntimeDeps: []string{},
			InstallCmd:  "docker-php-ext-install pcntl",
			PECLInstall: false,
		},
		"soap": {
			BuildDeps:   []string{"libxml2-dev"},
			RuntimeDeps: []string{"libxml2"},
			InstallCmd:  "docker-php-ext-install soap",
			PECLInstall: false,
		},
		"sockets": {
			BuildDeps:   []string{},
			RuntimeDeps: []string{},
			InstallCmd:  "docker-php-ext-install sockets",
			PECLInstall: false,
		},
		"xmlrpc": {
			BuildDeps:   []string{"libxml2-dev"},
			RuntimeDeps: []string{"libxml2"},
			InstallCmd:  "docker-php-ext-install xmlrpc",
			PECLInstall: false,
		},
		"xsl": {
			BuildDeps:   []string{"libxslt1-dev"},
			RuntimeDeps: []string{"libxslt1.1"},
			InstallCmd:  "docker-php-ext-install xsl",
			PECLInstall: false,
		},
		// PECL extensions
		"redis": {
			BuildDeps:   []string{},
			RuntimeDeps: []string{},
			InstallCmd:  "pecl install redis && docker-php-ext-enable redis",
			PECLInstall: true,
		},
		"imagick": {
			BuildDeps:   []string{"libmagickwand-dev"},
			RuntimeDeps: []string{"libmagickwand-6.q16-6"},
			InstallCmd:  "pecl install imagick && docker-php-ext-enable imagick",
			PECLInstall: true,
		},
		"memcached": {
			BuildDeps:   []string{"libmemcached-dev", "zlib1g-dev"},
			RuntimeDeps: []string{"libmemcached11"},
			InstallCmd:  "pecl install memcached && docker-php-ext-enable memcached",
			PECLInstall: true,
		},
		"mongodb": {
			BuildDeps:   []string{"libssl-dev"},
			RuntimeDeps: []string{},
			InstallCmd:  "pecl install mongodb && docker-php-ext-enable mongodb",
			PECLInstall: true,
		},
		"xdebug": {
			BuildDeps:   []string{},
			RuntimeDeps: []string{},
			InstallCmd:  "pecl install xdebug && docker-php-ext-enable xdebug",
			PECLInstall: true,
		},
		"apcu": {
			BuildDeps:   []string{},
			RuntimeDeps: []string{},
			InstallCmd:  "pecl install apcu && docker-php-ext-enable apcu",
			PECLInstall: true,
		},
	}

	return ubuntuSupport[extName]
}

// GetUbuntuBuildDeps returns all build dependencies for Ubuntu
func GetUbuntuBuildDeps(extensions []string) []string {
	depsMap := make(map[string]bool)
	
	for _, ext := range extensions {
		support := GetUbuntuSupport(ext)
		if support != nil {
			for _, dep := range support.BuildDeps {
				depsMap[dep] = true
			}
		}
	}
	
	deps := make([]string, 0, len(depsMap))
	for dep := range depsMap {
		deps = append(deps, dep)
	}
	return deps
}

// GetUbuntuRuntimeDeps returns all runtime dependencies for Ubuntu
func GetUbuntuRuntimeDeps(extensions []string) []string {
	depsMap := make(map[string]bool)
	
	for _, ext := range extensions {
		support := GetUbuntuSupport(ext)
		if support != nil {
			for _, dep := range support.RuntimeDeps {
				depsMap[dep] = true
			}
		}
	}
	
	deps := make([]string, 0, len(depsMap))
	for dep := range depsMap {
		deps = append(deps, dep)
	}
	return deps
}
