<?php

use App\Http\Controllers\Api\ProductController;
use Illuminate\Support\Facades\Route;

Route::post('/products/import', [ProductController::class, 'import']);
