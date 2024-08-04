<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Imports\ProductsImport;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Http\Response;
use Maatwebsite\Excel\Facades\Excel;

class ProductController extends Controller
{
    public function import(Request $request): Response|JsonResponse {
        if(!$request->hasFile('file')){
            return response()->json([
                'error' => [
                    'message' => 'Файл не найден',
                ],
            ], 422);
        }
        $file = $request->file('file');
        if (!$file->isValid()
                || !in_array($file->getClientOriginalExtension(), ['xls', 'xlsx'])) {
            return response()->json([
                'error' => [
                    'message' => 'Файл не xlsx|xls формата',
                ],
            ], 422);
        }
        echo 'test';
        Excel::import(new ProductsImport, $file);


        return response(null, 204);
    }
}
