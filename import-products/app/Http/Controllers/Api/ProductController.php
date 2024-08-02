<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Imports\ProductsImport;
use Illuminate\Http\Request;
use Maatwebsite\Excel\Facades\Excel;

class ProductController extends Controller
{
    public function import(Request $request): string {
        //$file = $request->file();
        //Excel::import(new ProductsImport, './tests/import_example.xlsx');


        return response()->json([
            'file' => '',
        ]);
    }
}
