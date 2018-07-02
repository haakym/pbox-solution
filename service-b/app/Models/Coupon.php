<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Model;

class Coupon extends Model
{
    /**
     * Get coupons by constraints and
     * remove from DB if any matches
     */
    public static function redeem($filters)
    {
        $coupons = self::where('brand', $filters['brand'])
            ->where('value', $filters['value'])
            ->limit($filters['limit'])
            ->get();

        if (count($coupons)) {
            $toRemove = $coupons->pluck('id');
            
            self::whereIn('id', $toRemove)->delete();
        }
        
        return $coupons;
    }
}
