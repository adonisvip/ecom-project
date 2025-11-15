# Database Migrations

This directory contains migration files for the ecom-auth service database.

## Database

- **Database Name**: `ecom_auth`

## Migration Approach

This project uses **GORM AutoMigrate** for database migrations. GORM automatically creates and updates database tables based on the model definitions in `models/models.go`.

## Running Migrations

### Using GORM AutoMigrate (Recommended)

```bash
# From the ecom-auth directory
go run cmd/migrate/main.go
```

This will:
- Connect to the database using environment variables
- Automatically create/update the `users` table based on the `models.Users` struct
- Add missing columns and indexes
- Preserve existing data

### Environment Variables

Make sure these environment variables are set (or use docker-compose):

```bash
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_DATABASE=ecom_auth
```

### Using Docker

```bash
# From the ecom-project directory
docker exec -it ecom-auth go run cmd/migrate/main.go
```

## SQL Migration Files (Legacy)

The `001_create_users_table.sql` file is kept for reference or manual migration if needed. However, GORM AutoMigrate is the preferred method.

### Manual SQL Migration (Alternative)

If you prefer to run SQL migrations manually:

```bash
docker exec -i ecom-postgres psql -U postgres -d ecom_auth < ecom-auth/migrations/001_create_users_table.sql
```

## How GORM AutoMigrate Works

GORM AutoMigrate will:
- ✅ Create tables if they don't exist
- ✅ Add missing columns
- ✅ Add missing indexes
- ✅ Add missing foreign keys
- ⚠️ **Won't** delete unused columns (for safety)
- ⚠️ **Won't** change column types if they exist

## Model Structure

The `models.Users` struct defines the table structure:
- `id` - Primary key (auto-increment)
- `username` - Unique, not null
- `password` - Not null
- `email` - Unique, not null
- `fullname` - Optional
- `phone` - Optional
- `token` - Optional
- `refresh_token` - Optional
- `created_at` - Auto-managed timestamp
- `updated_at` - Auto-managed timestamp

## Notes

- Always backup your database before running migrations in production
- Test migrations in a development environment first
- GORM AutoMigrate is idempotent - safe to run multiple times
- To add new models, add them to the `AutoMigrate()` call in `cmd/migrate/main.go`

