<?php

namespace App\Http\Controllers\Api;

use App\Http\Controllers\Controller;
use App\Imports\ProductsImport;
use Illuminate\Http\JsonResponse;
use Illuminate\Http\Request;
use Illuminate\Http\Response;
use Maatwebsite\Excel\Facades\Excel;
use Throwable;

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

        try {
            Excel::import(new ProductsImport, $file);
        } catch (Throwable $e) {
            return response()->json([
                'error' => [
                    'message' => 'Ошибка при считывании файла. Возможно, заданы не все необходимые поля',
                ]
            ], 422);
        }

        return response(null, 204);
    }
}
