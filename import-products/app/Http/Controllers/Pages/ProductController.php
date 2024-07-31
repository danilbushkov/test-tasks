<?php

namespace App\Http\Controllers\Pages;

use App\Http\Controllers\Controller;
use Illuminate\Http\RedirectResponse;
use Illuminate\Http\Request;

class ProductController extends Controller
{
    public function list(Request $request): RedirectResponse | string {
        $page = $request->input('page');
        if($page !== null) {
            if(!is_int($page)) {
                return redirect('/products');
            }

        } else {
            $page = 1;
        }

        return view('products.list');
    }

    public function import(): string {
        return 'Import';
    }

    public function item(string $id): string {
        return $id;
    }

}
