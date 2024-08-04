<?php

namespace App\Http\Controllers\Pages;

use App\Http\Controllers\Controller;
use App\Models\Product;
use Illuminate\Http\RedirectResponse;
use Illuminate\Http\Request;
use Illuminate\View\View;

class ProductController extends Controller
{
    public function list(Request $request): RedirectResponse | View {
        $page = $request->input('page');
        if($page !== null) {
            if(!is_int($page)) {
                return redirect('/products');
            }

        } else {
            $page = 1;
        }

        $products = Product::select('name')->get();

        return view('products.list', ['products' => $products]);
    }

    public function import(): string {
        return view('products.import');
    }

    public function item(string $id): string {
        return view('products.item');
    }

}
