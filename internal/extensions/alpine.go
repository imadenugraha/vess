package extensions

// GetAlpineSupport returns Alpine-specific installation information for an extension
func GetAlpineSupport(extName string) *OSSupport {
	alpineSupport := map[string]*OSSupport{
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
			BuildDeps:   []string{"postgresql-dev"},
			RuntimeDeps: []string{"postgresql-libs"},
			InstallCmd:  "docker-php-ext-install pdo_pgsql",
			PECLInstall: false,
		},
		"pgsql": {
			BuildDeps:   []string{"postgresql-dev"},
			RuntimeDeps: []string{"postgresql-libs"},
			InstallCmd:  "docker-php-ext-install pgsql",
			PECLInstall: false,
		},
		"gd": {
			BuildDeps:   []string{"freetype-dev", "libjpeg-turbo-dev", "libpng-dev", "libwebp-dev"},
			RuntimeDeps: []string{"freetype", "libjpeg-turbo", "libpng", "libwebp"},
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
			RuntimeDeps: []string{"libzip"},
			InstallCmd:  "docker-php-ext-install zip",
			PECLInstall: false,
		},
		"intl": {
			BuildDeps:   []string{"icu-dev"},
			RuntimeDeps: []string{"icu-libs"},
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
			BuildDeps:   []string{"libxslt-dev"},
			RuntimeDeps: []string{"libxslt"},
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
			BuildDeps:   []string{"imagemagick-dev"},
			RuntimeDeps: []string{"imagemagick"},
			InstallCmd:  "pecl install imagick && docker-php-ext-enable imagick",
			PECLInstall: true,
		},
		"memcached": {
			BuildDeps:   []string{"libmemcached-dev", "zlib-dev"},
			RuntimeDeps: []string{"libmemcached-libs"},
			InstallCmd:  "pecl install memcached && docker-php-ext-enable memcached",
			PECLInstall: true,
		},
		"mongodb": {
			BuildDeps:   []string{"openssl-dev"},
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

	return alpineSupport[extName]
}

// GetAlpineBuildDeps returns all build dependencies for Alpine
func GetAlpineBuildDeps(extensions []string) []string {
	depsMap := make(map[string]bool)

	for _, ext := range extensions {
		support := GetAlpineSupport(ext)
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

// GetAlpineRuntimeDeps returns all runtime dependencies for Alpine
func GetAlpineRuntimeDeps(extensions []string) []string {
	depsMap := make(map[string]bool)

	for _, ext := range extensions {
		support := GetAlpineSupport(ext)
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
