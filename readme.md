# Pbox Take Home Test Solution

## Folder Structure

- service-a: go
- service-b: php (Laravel)

## To Run

### 1. Install dependencies etc. for service-b

```bash
cd service-b
composer install
copy .env.example .env
php artisan key:generate
```

### 2. Create dummy coupon data for service-b

```bash
cd service-b
php artisan db:seed
```

### 3. Run containers

For first time:

```bash
docker-composer up --build
```

and thereafter:

```bash
docker-composer up
```

## Web Routes

- service-a: http://localhost:8000/coupons
- service-b:
    - http://localhost:3000/api/coupons
    - http://localhost:3000/api/get-coupons
