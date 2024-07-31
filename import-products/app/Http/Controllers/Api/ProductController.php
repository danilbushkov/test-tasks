<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use Illuminate\Http\Request;

class ProductController extends Controller
{
    public function import(Request $request): string {
        $file = $request->file();

        return response()->json([
            'file' => $file,
        ]);
    }
}
