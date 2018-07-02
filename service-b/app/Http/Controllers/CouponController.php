<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Models\Coupon;

class CouponController extends Controller
{
    public function index()
    {
        return Coupon::all();
    }

    public function getCoupons()
    {
        $requestParams = request()->only(['brand', 'limit', 'value']);

        $coupons = Coupon::redeem($requestParams);
        
        return response()->json($coupons);
    }
}
