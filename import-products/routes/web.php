<?php

use App\Http\Controllers\Pages\ProductController;
use Illuminate\Support\Facades\Route;

Route::get('/', [ProductController::class, 'list']);

Route::controller(ProductController::class)
    ->prefix('products')
    ->group(function () {
        Route::get('/', 'list');
        Route::get('/import', 'import');
        Route::get('/{id}', 'item');
    }
);
