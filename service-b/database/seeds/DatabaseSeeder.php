<?php

use Illuminate\Database\Seeder;
use App\Models\Coupon;

class DatabaseSeeder extends Seeder
{
    /**
     * Seed the application's database.
     *
     * @return void
     */
    public function run()
    {
        factory(Coupon::class, 10)->create();
    }
}
