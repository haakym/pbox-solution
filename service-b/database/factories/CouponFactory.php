<?php

use Faker\Generator as Faker;

$factory->define(App\Models\Coupon::class, function (Faker $faker) {
    return [
        'brand' => $faker->word,
        'value' => $faker->numberBetween(1, 10),
    ];
});
