# vess - PHP Dockerfile Generator

🐳 A powerful CLI tool that generates OS-specific PHP Dockerfiles, builds Docker images, and exports comprehensive PHP extension metadata to JSON.

## Features

- 🎯 **Generate Dockerfiles** from simple environment files
- 🏗️ **Build Docker images** directly from generated Dockerfiles
- 📊 **Export extension metadata** to JSON with complete dependency information
- 🐧 **Multi-OS support**: Alpine and Ubuntu base images
- 🔢 **PHP 7.4 to 8.3** support
- 📦 **20+ PHP extensions** with automatic dependency resolution
- 🚀 **Multi-stage builds** for optimized image sizes
- ✨ **Zero configuration** - just specify extensions in an env file

## Installation

### From Source

```bash
git clone https://github.com/imadenugraha/vess.git
cd vess
go build -o vess
sudo mv vess /usr/local/bin/
```

### Quick Build

```bash
go build -o vess
```

## Usage

### Generate Dockerfile

Create a Dockerfile from an env file specifying PHP extensions:

```bash
vess generate --os alpine --php-version 8.3 --env-file app.env --output Dockerfile
```

**Short form:**

```bash
vess generate -o alpine -p 8.3 -e app.env -f Dockerfile
```

### Build Docker Image

Build a Docker image from a generated Dockerfile:

```bash
vess build --dockerfile Dockerfile --tag my-php:8.3-alpine
```

**With no cache:**

```bash
vess build -d Dockerfile -t my-php:8.3 --no-cache
```

### Export Extension Metadata

Export all available PHP extension metadata to JSON:

```bash
vess export --os alpine --php-version 8.3 --output extensions.json
```

### Complete Workflow

Generate Dockerfile, build image, and verify:

```bash
# Generate Dockerfile
vess generate -o alpine -p 8.3 -t fpm -e examples/basic.env -f Dockerfile

# Build image
vess build -d Dockerfile -t my-app:latest

# Verify installed extensions
docker run --rm my-app:latest php -m
```

## Environment File Format

Create an `.env` file with PHP extensions:

```env
# Basic PHP extensions
PHP_EXTENSIONS=mysqli,pdo_mysql,opcache,zip,bcmath

# You can add comments
# PHP_EXTENSIONS=gd,intl,redis,imagick
```

See [examples/](examples/) for more configuration samples.

## Supported Extensions

### Core Extensions (bundled with PHP)

- `mysqli` - MySQL Improved Extension
- `pdo_mysql` - MySQL PDO Driver
- `pdo_pgsql` - PostgreSQL PDO Driver
- `pgsql` - PostgreSQL Extension
- `gd` - GD Graphics Library
- `opcache` - Zend OPcache
- `zip` - Zip Archive Extension
- `intl` - Internationalization
- `bcmath` - BC Math Extension
- `exif` - EXIF Extension
- `pcntl` - Process Control
- `soap` - SOAP Extension
- `sockets` - Sockets Extension
- `xsl` - XSL Extension

### PECL Extensions (from PECL repository)

- `redis` - Redis Extension
- `imagick` - ImageMagick Extension
- `memcached` - Memcached Extension
- `mongodb` - MongoDB Driver
- `xdebug` - Xdebug Debugger
- `apcu` - APCu Cache

## Supported PHP Versions

- PHP 7.4
- PHP 8.0
- PHP 8.1
- PHP 8.2
- PHP 8.3

## Base Image Types

vess supports three PHP base image types:

### CLI (Command Line Interface)

- **Use for**: Queue workers, cron jobs, scheduled tasks, Laravel Artisan commands
- **Characteristics**: No web server, runs PHP scripts from command line
- **Container behavior**: Starts interactive PHP shell by default
- **Example**: `vess generate -o alpine -p 8.3 -t cli -e examples/cli-worker.env`

### FPM (FastCGI Process Manager)

- **Use for**: Production web applications with Nginx, microservices, modern PHP apps
- **Characteristics**: PHP-FPM daemon listening on port 9000
- **Container behavior**: Requires separate web server (Nginx/Apache) as reverse proxy
- **Example**: `vess generate -o alpine -p 8.3 -t fpm -e examples/basic.env`
- **Default**: FPM is the default image type

### Apache

- **Use for**: Traditional all-in-one deployments, simple hosting, legacy applications
- **Characteristics**: Apache web server with mod_php built-in, listens on port 80
- **Container behavior**: Self-contained web server, no external proxy needed
- **Example**: `vess generate -o ubuntu -p 8.3 -t apache -e examples/apache-simple.env`
- **⚠️ Limitation**: Only available with Ubuntu/Debian (not Alpine)

## Supported Operating Systems

- **Alpine Linux** - Lightweight, optimal for production (`--os alpine`)
  - Supports: CLI, FPM
  - ⚠️ Does not support: Apache
- **Ubuntu/Debian** - Full-featured, better for development (`--os ubuntu`)
  - Supports: CLI, FPM, Apache

### Compatibility Matrix

| Image Type | Alpine | Ubuntu/Debian |
| ---------- | ------ | ------------- |
| CLI | ✅ Yes | ✅ Yes |
| FPM | ✅ Yes | ✅ Yes |
| Apache | ❌ No | ✅ Yes |

## Examples

### Basic Web Application (FPM)

```bash
# Create env file
echo "PHP_EXTENSIONS=mysqli,pdo_mysql,opcache,zip" > app.env

# Generate FPM image for use with Nginx
vess generate -o alpine -p 8.3 -t fpm -e app.env -f Dockerfile
vess build -d Dockerfile -t my-web-app:latest
```

### Apache All-in-One Application

```bash
# Generate Apache image (Ubuntu only)
vess generate -o ubuntu -p 8.3 -t apache -e examples/apache-simple.env -f Dockerfile.apache
vess build -d Dockerfile.apache -t my-apache-app:latest

# Run and test
docker run -d -p 8080:80 my-apache-app:latest
curl http://localhost:8080
```

### CLI Worker for Queue Processing

```bash
# Generate CLI image for Laravel queue worker
vess generate -o alpine -p 8.3 -t cli -e examples/cli-worker.env -f Dockerfile.worker
vess build -d Dockerfile.worker -t my-worker:latest

# Run worker
docker run -d my-worker:latest php artisan queue:work
```

### Laravel Application (Multiple Containers)

```bash
# FPM container for web requests
vess generate -o alpine -p 8.3 -t fpm -e examples/laravel.env -f Dockerfile.web
vess build -d Dockerfile.web -t laravel-web:8.3

# CLI container for queue workers
vess generate -o alpine -p 8.3 -t cli -e examples/cli-worker.env -f Dockerfile.worker  
vess build -d Dockerfile.worker -t laravel-worker:8.3
```

### Development Environment with Xdebug

```bash
# Create dev env file
echo "PHP_EXTENSIONS=mysqli,pdo_mysql,redis,xdebug,opcache" > dev.env

# Generate Ubuntu-based image (better for dev)
vess generate -o ubuntu -p 8.3 -e dev.env -f Dockerfile.dev
vess build -d Dockerfile.dev -t php-dev:8.3
```

## Global Flags

- `--os, -o` - Operating system: `alpine` or `ubuntu` (default: `alpine`)
- `--php-version, -p` - PHP version: `7.4`, `8.0`, `8.1`, `8.2`, `8.3` (default: `8.3`)
- `--verbose, -v` - Enable verbose output

## Command Reference

### `vess generate`

Generates a Dockerfile from an env file.

**Flags:**

- `--env-file, -e` - Path to env file (required)
- `--output, -f` - Output Dockerfile path (default: `Dockerfile`)
- `--type, -t` - PHP base image type: `cli`, `fpm`, `apache` (default: `fpm`)

### `vess build`

Builds a Docker image from a Dockerfile.

**Flags:**

- `--dockerfile, -d` - Path to Dockerfile (default: `Dockerfile`)
- `--tag, -t` - Image tag (required)
- `--no-cache` - Build without cache

### `vess export`

Exports PHP extension metadata to JSON.

**Flags:**

- `--output` - Output JSON file path (default: `extensions.json`)

## Output Structure

### Generated Dockerfile

Dockerfiles use multi-stage builds for optimal image size:

1. **Builder stage**: Installs build dependencies and compiles extensions
2. **Final stage**: Contains only runtime dependencies and compiled extensions

### JSON Export

The exported JSON contains:

```json
{
  "os": "alpine",
  "php_version": "8.3",
  "extensions": [
    {
      "name": "mysqli",
      "description": "MySQL Improved Extension",
      "build_dependencies": [],
      "runtime_dependencies": [],
      "install_command": "docker-php-ext-install mysqli",
      "pecl_install": false,
      "supported_php_versions": ["7.4", "8.0", "8.1", "8.2", "8.3"],
      "conflicts": []
    }
  ]
}
```

## Troubleshooting

### Docker daemon not running

```bash
Error: Docker error: failed to ping Docker daemon
```

**Solution**: Start Docker daemon:

```bash
sudo systemctl start docker
```

### Extension not found

```bash
Error: unknown extension: foo
```

**Solution**: Check available extensions:

```bash
vess export --output extensions.json
cat extensions.json | grep '"name"'
```

### Build failures

If the build fails, try:

```bash
# Build with verbose output
vess generate -v -o alpine -p 8.3 -e app.env

# Build without cache
vess build -d Dockerfile -t my-app:latest --no-cache
```

## Contributing

Contributions are welcome! To add new extensions:

1. Update `internal/extensions/registry.go`
2. Add OS-specific installation info in `alpine.go` and `ubuntu.go`
3. Test the generation and build process
4. Submit a pull request

## License

MIT License

## Author

Built with ❤️ for the PHP community
